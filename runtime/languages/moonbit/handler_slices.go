/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package moonbit

import (
	"context"
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"
)

func (p *planner) NewSliceHandler(ctx context.Context, ti langsupport.TypeInfo) (langsupport.TypeHandler, error) {
	handler := &sliceHandler{
		typeHandler: *NewTypeHandler(ti),
	}
	p.AddHandler(handler)

	typeDef, err := p.metadata.GetTypeDefinition(ti.Name())
	if err != nil {
		return nil, fmt.Errorf("planner.NewSliceHandler: p.metadata.GetTypeDefinition('%v'): %w", ti.Name(), err)
	}
	handler.typeDef = typeDef

	elementHandler, err := p.GetHandler(ctx, ti.ListElementType().Name())
	if err != nil {
		return nil, fmt.Errorf("planner.NewSliceHandler: p.GetHandler('%v'): %w", ti.ListElementType().Name(), err)
	}
	handler.elementHandler = elementHandler

	// an empty slice (not nil)
	handler.emptyValue = reflect.MakeSlice(ti.ReflectedType(), 0, 0).Interface()

	return handler, nil
}

type sliceHandler struct {
	typeHandler
	typeDef        *metadata.TypeDefinition
	elementHandler langsupport.TypeHandler
	emptyValue     any
}

func (h *sliceHandler) Read(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {
	gmlPrintf("GML: handler_slices.go: sliceHandler.Read(offset: %v), type=%T", debugShowOffset(offset), h.emptyValue)
	return h.Decode(ctx, wa, []uint64{uint64(offset)})

	// if offset == 0 {
	// 	if h.typeInfo.IsPointer() {
	// 		return nil, nil
	// 	}
	// 	return h.emptyValue, nil
	// }

	// // First attempt to read slice directly, but if it fails, then treat the offset as a pointer to the slice
	// // and try again.
	// if v, err := h.Decode(ctx, wa, []uint64{uint64(offset)}); err == nil {
	// 	return v, nil
	// }

	// // Read pointer to slice
	// ptr, ok := wa.Memory().ReadUint32Le(offset)
	// if !ok {
	// 	gmlPrintf("GML: sliceHandler.Read: failed to read pointer to slice at offset %v", debugShowOffset(offset))
	// } else {
	// 	gmlPrintf("GML: sliceHandler.Read: ptr: %v", debugShowOffset(ptr))
	// }

	// return h.Decode(ctx, wa, []uint64{uint64(ptr)})
}

func (h *sliceHandler) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	ptr, cln, err := h.doWriteSlice(ctx, wa, obj)
	if err != nil {
		return cln, err
	}

	wa.Memory().WriteUint32Le(offset, ptr)

	return cln, nil
}

// Decode is always passed an address to a slice in memory.
func (h *sliceHandler) Decode(ctx context.Context, wasmAdapter langsupport.WasmAdapter, vals []uint64) (any, error) {
	wa, ok := wasmAdapter.(wasmMemoryReader)
	if !ok {
		return nil, fmt.Errorf("expected a wasmMemoryReader, got %T", wasmAdapter)
	}

	gmlPrintf("GML: handler_slices.go: sliceHandler.Decode(vals: %+v)", vals)

	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 value when decoding a slice but got %v: %+v", len(vals), vals)
	}

	if vals[0] == 0 {
		return nil, nil
	}

	memBlock, classID, words, err := memoryBlockAtOffset(wa, uint32(vals[0]), 0, true)
	if err != nil {
		return nil, err
	}

	if words == 0 {
		return h.emptyValue, nil // empty slice
	}

	numElements := uint32(words)
	elemType := h.typeInfo.ListElementType()
	elemTypeSize := uint32(4) // elemType.Size()
	isNullable := elemType.IsNullable()
	if isNullable && elemType.IsPrimitive() &&
		// elemType.Name() != "Byte?" && elemType.Name() != "Bool?" &&
		// elemType.Name() != "Char?" && elemType.Name() != "Int16?" &&
		// elemType.Name() != "Int64?" && elemType.Name() != "UInt64?" &&
		// elemType.Name() != "Float?" && elemType.Name() != "Double?" {
		(elemType.Name() == "Int?" || elemType.Name() == "UInt?" || elemType.Name() == "String?") {
		elemTypeSize = 8
	}
	if classID == FixedArrayPrimitiveBlockType || classID == PtrArrayBlockType { // && isNullable && elemType.IsPrimitive()
		memBlock, _, _, err = memoryBlockAtOffset(wa, uint32(vals[0]), words*elemTypeSize, true)
		if err != nil {
			return nil, err
		}
	} else {
		sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
		if sliceOffset == 0 {
			return nil, nil // nil slice
			// return h.emptyValue, nil // empty slice
		}

		// For debugging:
		memoryBlockAtOffset(wa, sliceOffset, 0, true) //nolint:errcheck

		if words == 1 {
			// sliceOffset is the pointer to the single-slice element.
			items := reflect.MakeSlice(h.typeInfo.ReflectedType(), 1, 1)
			item, err := h.elementHandler.Read(ctx, wasmAdapter, sliceOffset)
			if err != nil {
				return nil, err
			}
			if !utils.HasNil(item) {
				items.Index(0).Set(reflect.ValueOf(item))
			}
			return items.Interface(), nil
		}

		numElements = binary.LittleEndian.Uint32(memBlock[12:16])
		if numElements == 0 {
			return h.emptyValue, nil // empty slice
		}

		size := numElements * uint32(elemTypeSize)
		gmlPrintf("GML: handler_slices.go: sliceHandler.Decode: sliceOffset=%v, numElements=%v, size=%v", debugShowOffset(sliceOffset), numElements, size)

		memBlock, _, _, err = memoryBlockAtOffset(wa, sliceOffset, size, true)
		if err != nil {
			return nil, err
		}
	}

	// return h.doReadSlice(ctx, wa, sliceOffset, size)

	// elementSize := h.elementHandler.TypeInfo().Size()
	items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
	for i := uint32(0); i < numElements; i++ {
		// TODO: This is all quite a hack - figure out how to make this an elegant solution.
		if elemType.IsPrimitive() && isNullable {
			var value uint64
			// if elemType.Name() != "Bool?" && elemType.Name() != "Byte?" &&
			// 	elemType.Name() != "Char?" && elemType.Name() != "Int16?" &&
			// 	elemType.Name() != "Int64?" && elemType.Name() != "UInt64?" &&
			// 	elemType.Name() != "Float?" && elemType.Name() != "Double?" {
			if elemType.Name() == "Int?" || elemType.Name() == "UInt?" || elemType.Name() == "String?" {
				value = binary.LittleEndian.Uint64(memBlock[8+i*elemTypeSize:])
			} else {
				value32 := binary.LittleEndian.Uint32(memBlock[8+i*elemTypeSize:])
				value = uint64(value32)
			}
			item, err := h.elementHandler.Decode(ctx, wasmAdapter, []uint64{value})
			if err != nil {
				return nil, err
			}
			if !utils.HasNil(item) {
				items.Index(int(i)).Set(reflect.ValueOf(item))
			}
			continue
		}
		itemOffset := binary.LittleEndian.Uint32(memBlock[8+i*elemTypeSize:])
		item, err := h.elementHandler.Read(ctx, wasmAdapter, itemOffset)
		if err != nil {
			return nil, err
		}
		if !utils.HasNil(item) {
			items.Index(int(i)).Set(reflect.ValueOf(item))
		}
	}

	return items.Interface(), nil
}

func (h *sliceHandler) Encode(ctx context.Context, wasmAdapter langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	ptr, cln, err := h.doWriteSlice(ctx, wasmAdapter, obj)
	if err != nil {
		return nil, cln, err
	}

	return []uint64{uint64(ptr)}, cln, nil
}

/*
func (h *sliceHandler) doReadSlice(ctx context.Context, wa langsupport.WasmAdapter, data, size uint32) (any, error) {
	if data == 0 {
		// nil slice
		return nil, nil
	}

	if size == 0 {
		// empty slice
		return h.emptyValue, nil
	}

	elementSize := h.elementHandler.TypeInfo().Size()
	items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(size), int(size))
	for i := uint32(0); i < size; i++ {
		itemOffset := data + i*elementSize
		item, err := h.elementHandler.Read(ctx, wa, itemOffset)
		if err != nil {
			return nil, err
		}
		if !utils.HasNil(item) {
			items.Index(int(i)).Set(reflect.ValueOf(item))
		}
	}

	return items.Interface(), nil
}
*/

// type sliceWriter interface {
// 	doWriteSlice(ctx context.Context, wa langsupport.WasmAdapter, obj any) (ptr uint32, cln utils.Cleaner, err error)
// }

func (h *sliceHandler) doWriteSlice(ctx context.Context, wasmAdapter langsupport.WasmAdapter, obj any) (ptr uint32, cln utils.Cleaner, err error) {
	wa, ok := wasmAdapter.(wasmMemoryWriter)
	if !ok {
		return 0, nil, fmt.Errorf("expected a wasmMemoryWriter, got %T", wasmAdapter)
	}

	gmlPrintf("GML: handler_slices.go: sliceHandler.doWriteSlice(obj: %+v)", obj)
	if utils.HasNil(obj) {
		return 0, nil, nil
	}

	slice, err := utils.ConvertToSlice(obj)
	if err != nil {
		return 0, nil, err
	}

	numElements := uint32(len(slice))
	elemType := h.typeInfo.ListElementType()
	elemTypeSize := uint32(4) // elemType.Size()
	isNullable := elemType.IsNullable()
	if elemType.IsPrimitive() && isNullable &&
		// elemType.Name() != "Byte?" && elemType.Name() != "Bool?" &&
		// elemType.Name() != "Char?" && elemType.Name() != "Int16?" &&
		// elemType.Name() != "Int64?" && elemType.Name() != "UInt64?" &&
		// elemType.Name() != "Float?" && elemType.Name() != "Double?" {
		(elemType.Name() == "Int?" || elemType.Name() == "UInt?") {
		elemTypeSize = 8
	}
	size := numElements * uint32(elemTypeSize)
	memBlockClassID := uint32(PtrArrayBlockType)
	if elemType.Name() == "Byte?" || elemType.Name() == "Bool?" || elemType.Name() == "Char?" ||
		elemType.Name() == "Int?" || elemType.Name() == "UInt?" ||
		elemType.Name() == "Int16?" || elemType.Name() == "UInt16?" {
		memBlockClassID = uint32(FixedArrayPrimitiveBlockType)
	}

	// Allocate memory
	if size == 0 {
		ptr, cln, err = wa.allocateAndPinMemory(ctx, 1, memBlockClassID) // cannot allocate 0 bytes
		if err != nil {
			return 0, cln, err
		}
		wa.Memory().WriteByte(ptr-3, 0) // overwrite size=1 to size=0
	} else {
		ptr, cln, err = wa.allocateAndPinMemory(ctx, size, memBlockClassID)
		if err != nil {
			return 0, cln, err
		}

		// For `Int?`, `UInt?`, the `words` portion of the memory block
		// indicates the number of elements in the slice, not the number of 16-bit words.
		if elemType.Name() == "Int?" || elemType.Name() == "UInt?" {
			// NOT: elemType.Name() == "Int64?" || elemType.Name() == "UInt64?" || elemType.Name() == "Double?" {
			memType := ((size / 8) << 8) | memBlockClassID
			wa.Memory().WriteUint32Le(ptr-4, memType)
		}
	}

	// For debugging purposes:
	memoryBlockAtOffset(wa, ptr-8, size, true) //nolint:errcheck

	innerCln := utils.NewCleanerN(len(slice))

	defer func() {
		// unpin slice elements after the slice is written to memory
		if e := innerCln.Clean(); e != nil && err == nil {
			err = e
		}
	}()

	// elementTypeSize (above) is the skip size when writing the encoded value to memory.
	// elementSize (below) is the native size of the MoonBit value.
	elementSize := h.elementHandler.TypeInfo().Size() // BUG?!?  Int64?/UInt64? are 8 bytes, not 4
	if elemType.Name() == "Int64?" || elemType.Name() == "UInt64?" {
		elementSize = 8
	} else if elementSize > 4 {
		// Note that the element size itself may be larger than 4, but we are writing a pointer to it, so make it 4.
		elementSize = 4
	}

	for i, val := range slice {
		if elemType.IsPrimitive() && isNullable {
			if !utils.HasNil(val) {
				if elemType.Name() == "Int64?" || elemType.Name() == "UInt64?" || elemType.Name() == "Float?" || elemType.Name() == "Double?" {
					subBlockClassID := uint32(OptionBlockType) // TODO: rename
					valueBlock, c, err := wa.allocateAndPinMemory(ctx, elementSize, subBlockClassID)
					if err != nil {
						return 0, cln, err
					}
					innerCln.AddCleaner(c)
					if _, err := h.elementHandler.Write(ctx, wasmAdapter, valueBlock, val); err != nil {
						return 0, cln, err
					}
					wa.Memory().WriteUint32Le(ptr+uint32(i)*elemTypeSize, valueBlock-8)
					// For debugging:
					memoryBlockAtOffset(wa, valueBlock-8, 0, true) //nolint:errcheck
				} else {
					if _, err := h.elementHandler.Write(ctx, wasmAdapter, ptr+uint32(i)*elemTypeSize, val); err != nil {
						return 0, cln, err
					}
					switch elemType.Name() {
					case "Byte?", "Bool?", "Char?", "String?": // "UInt?" in MoonBit for some reason uses sign extending below.
						wa.Memory().Write(ptr+4+uint32(i)*elemTypeSize, []byte{0, 0, 0, 0})
					case "Int16?":
						highByte, ok := wa.Memory().ReadByte(ptr + 1 + uint32(i)*elemTypeSize)
						if ok && highByte&0x80 != 0 {
							wa.Memory().Write(ptr+2+uint32(i)*elemTypeSize, []byte{255, 255}) // Some(negative number)
						} else {
							wa.Memory().Write(ptr+2+uint32(i)*elemTypeSize, []byte{0, 0}) // Some(positive number)
						}
					case "UInt16?":
						wa.Memory().Write(ptr+2+uint32(i)*elemTypeSize, []byte{0, 0}) // Some(positive number)
						// case "Int64?", "UInt64?", "Float?", "Double?": // handled above
					case "Int?", "UInt?": // I don't know why uses sign extending for "UInt?", but seems to.
						highByte, ok := wa.Memory().ReadByte(ptr + 3 + uint32(i)*elemTypeSize)
						if ok && highByte&0x80 != 0 {
							wa.Memory().Write(ptr+4+uint32(i)*elemTypeSize, []byte{255, 255, 255, 255}) // Some(negative number)
						} else {
							wa.Memory().Write(ptr+4+uint32(i)*elemTypeSize, []byte{0, 0, 0, 0}) // Some(positive number)
						}
					default:
						return 0, cln, fmt.Errorf("unsupported nullable primitive type: %v", elemType.Name())
					}
				}
			} else if elemType.Name() == "Int64?" || elemType.Name() == "UInt64?" || elemType.Name() == "Float?" || elemType.Name() == "Double?" {
				noneBlock, c, err := wa.allocateAndPinMemory(ctx, 1, 0) // cannot allocate 0 bytes
				if err != nil {
					return 0, cln, err
				}
				innerCln.AddCleaner(c)
				wa.Memory().Write(noneBlock-8, []byte{255, 255, 255, 255, 0, 0, 0, 0}) // None in memBlock header
				wa.Memory().WriteUint32Le(ptr+uint32(i)*elemTypeSize, noneBlock-8)
			} else if elemType.Name() == "Byte?" || elemType.Name() == "Bool?" ||
				elemType.Name() == "Char?" || elemType.Name() == "Int16?" || elemType.Name() == "UInt16?" {
				wa.Memory().Write(ptr+uint32(i)*elemTypeSize, []byte{255, 255, 255, 255}) // None
			} else {
				wa.Memory().Write(ptr+uint32(i)*elemTypeSize, []byte{0, 0, 0, 0, 1, 0, 0, 0}) // None
			}
			continue
		}
		if !utils.HasNil(val) {
			c, err := h.elementHandler.Write(ctx, wasmAdapter, ptr+uint32(i)*elementSize, val)
			innerCln.AddCleaner(c)
			if err != nil {
				return 0, cln, err
			}
		} // TODO: else is it safe to assume the memory has been cleared out to zeros, or do we need to write them explicitly?
	}

	// For debugging purposes:
	memoryBlockAtOffset(wa, ptr-8, size, true) //nolint:errcheck

	if strings.HasPrefix(h.typeDef.Name, "FixedArray[") {
		return ptr - 8, cln, nil
	}

	// Finally, write the slice memory block.
	slicePtr, sliceCln, err := wa.allocateAndPinMemory(ctx, 8, TupleBlockType)
	innerCln.AddCleaner(sliceCln)
	if err != nil {
		return 0, cln, err
	}
	wa.Memory().WriteUint32Le(slicePtr, ptr-8)
	wa.Memory().WriteUint32Le(slicePtr+4, numElements)

	// For debugging purposes:
	memoryBlockAtOffset(wa, slicePtr-8, 0, true) //nolint:errcheck

	return slicePtr - 8, cln, nil
}
