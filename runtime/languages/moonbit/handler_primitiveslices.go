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
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/langsupport/primitives"
	"github.com/gmlewis/modus/runtime/utils"
)

func (p *planner) NewPrimitiveSliceHandler(ti langsupport.TypeInfo) (h langsupport.TypeHandler, err error) {
	defer func() {
		if err == nil {
			p.typeHandlers[ti.Name()] = h
		}
	}()

	typeDef, err := p.metadata.GetTypeDefinition(ti.Name())
	if err != nil {
		return nil, fmt.Errorf("planner.NewPrimitiveSliceHandler: p.metadata.GetTypeDefinition('%v'): %w", ti.Name(), err)
	}

	typ, _, _ := stripErrorAndOption(ti.ListElementType().Name())

	switch typ {
	case "Bool":
		return newPrimitiveSliceHandler[bool](ti, typeDef), nil
		// https://docs.moonbitlang.com/en/latest/language/fundamentals.html#number
	case "Int16": // 16-bit signed integer, e.g. `(42 : Int16)`
		return newPrimitiveSliceHandler[int16](ti, typeDef), nil
	case "Int": // 32-bit signed integer, e.g. `42`
		return newPrimitiveSliceHandler[int32](ti, typeDef), nil
	case "Int64": // 64-bit signed integer, e.g. `1000L`
		return newPrimitiveSliceHandler[int64](ti, typeDef), nil
	case "UInt16": // 16-bit unsigned integer, e.g. `(14 : UInt16)`
		return newPrimitiveSliceHandler[uint16](ti, typeDef), nil
	case "UInt": // 32-bit unsigned integer, e.g. `14U`
		return newPrimitiveSliceHandler[uint32](ti, typeDef), nil
	case "UInt64": // 64-bit unsigned integer, e.g. `14UL`
		return newPrimitiveSliceHandler[uint64](ti, typeDef), nil
	case "Double": // 64-bit floating point, defined by IEEE754, e.g. `3.14`
		return newPrimitiveSliceHandler[float64](ti, typeDef), nil
	case "Float": // 32-bit floating point, defined by IEEE754, e.g. `(3.14 : Float)`
		return newPrimitiveSliceHandler[float32](ti, typeDef), nil
	case "Char": // represents a Unicode code point, e.g. `'a'`, `'\x41'`, `'\u{30}'`, `'\u03B1'`,
		return newPrimitiveSliceHandler[int16](ti, typeDef), nil
	case "Byte": // either a single ASCII character, e.g. `b'a'`, `b'\xff'`
		return newPrimitiveSliceHandler[uint8](ti, typeDef), nil
	// case "BigInt": // represents numeric values larger than other types, e.g. `10000000000000000000000N`
	// case "String": // holds a sequence of UTF-16 code units, e.g. `"Hello, World!"`
	case "@time.Duration":
		return newPrimitiveSliceHandler[time.Duration](ti, typeDef), nil

	default:
		return nil, fmt.Errorf("unsupported primitive MoonBit slice type: %s", ti.Name())
	}
}

func newPrimitiveSliceHandler[T primitive](ti langsupport.TypeInfo, typeDef *metadata.TypeDefinition) *primitiveSliceHandler[T] {
	return &primitiveSliceHandler[T]{
		*NewTypeHandler(ti),
		typeDef,
		primitives.NewPrimitiveTypeConverter[T](),
	}
}

type primitiveSliceHandler[T primitive] struct {
	typeHandler
	typeDef   *metadata.TypeDefinition
	converter primitives.TypeConverter[T]
}

func (h *primitiveSliceHandler[T]) Read(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {
	return h.Decode(ctx, wa, []uint64{uint64(offset)})
}

func (h *primitiveSliceHandler[T]) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	ptr, cln, err := h.doWriteSlice(ctx, wa.(*wasmAdapter), obj)
	if err != nil {
		return cln, err
	}

	if ok := wa.Memory().WriteUint32Le(offset, ptr); !ok {
		return cln, errors.New("failed to write struct pointer to memory")
	}

	return cln, nil
}

func (h *primitiveSliceHandler[T]) Decode(ctx context.Context, wasmAdapter langsupport.WasmAdapter, vals []uint64) (any, error) {
	wa, ok := wasmAdapter.(wasmMemoryReader)
	if !ok {
		return nil, fmt.Errorf("expected a wasmMemoryReader, got %T", wasmAdapter)
	}

	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 value when decoding a primitive slice but got %v: %+v", len(vals), vals)
	}

	if vals[0] == 0 {
		return nil, nil
	}

	// First read to get the header and determine the classID
	headerBlock, classID, words, err := memoryBlockAtOffset(wa, uint32(vals[0]), 0)
	if err != nil {
		return nil, err
	}

	// For new classIDs, calculate the correct size and re-read
	var sliceMemBlock []byte
	if classID == 96 || classID == 64 || classID == 112 {
		// Calculate the correct size for these classIDs
		elemTypeSize := h.converter.TypeSize()
		elemType := h.typeInfo.ListElementType()
		if elemType.Name() == "Bool" || elemType.Name() == "Char" {
			elemTypeSize = 4
		}
		dataSize := words * uint32(elemTypeSize)
		totalSize := dataSize // override with correct size
		sliceMemBlock, classID, words, err = memoryBlockAtOffset(wa, uint32(vals[0]), totalSize)
		if err != nil {
			return nil, err
		}
	} else {
		sliceMemBlock = headerBlock
	}

	// In current MoonBit version, outer Array wrapper may have words=0 even for non-empty arrays
	// Don't return early here, let the inner array logic determine if it's truly empty
	numElements := words
	elemTypeSize := h.converter.TypeSize()
	if classID == FixedArrayPrimitiveBlockType && elemTypeSize == 8 {
		// For Int64 and UInt64, the `words` portion of the memory block
		// indicates the number of elements in the slice, not the number of 16-bit words.
		size := numElements * uint32(elemTypeSize)
		sliceMemBlock, _, _, err = memoryBlockAtOffset(wa, uint32(vals[0]), size)
		if err != nil {
			return nil, err
		}
	}

	elemType := h.typeInfo.ListElementType()
	if elemType.Name() == "Bool" || elemType.Name() == "Char" {
		// A MoonBit Bool is 4 bytes whereas a Go bool is 1 byte.
		// A MoonBit Array[Char] uses 4 bytes per element instead of 2.
		elemTypeSize = 4
	}
	isNullable := elemType.IsNullable()
	if isNullable && elemType.Name() != "Int64?" && elemType.Name() != "UInt64?" {
		// Int64? and UInt64? both provide pointers to values.
		elemTypeSize = 8
	}

	if classID == TupleBlockType { // Used by Array[...] but not by FixedArray[...]
		numElements = binary.LittleEndian.Uint32(sliceMemBlock[12:16])
		// In older MoonBit versions, empty arrays had numElements == 0
		if numElements == 0 {
			return []T{}, nil
		}
		// For current MoonBit version, numElements might be a pointer for empty arrays

		sliceOffset := binary.LittleEndian.Uint32(sliceMemBlock[8:12])
		size := numElements * uint32(elemTypeSize)
		if elemTypeSize != 8 {
			size = 0 // do not override the memory block size for 1, 2, or 4-byte types.
		}

		sliceMemBlock, classID, words, err = memoryBlockAtOffset(wa, sliceOffset, size)
		if err != nil {
			return nil, err
		}
		// Don't return early based on words=0, let the switch statement handle it
	}

	switch classID {
	case FixedArrayPrimitiveBlockType, // Int
		96,  // FixedArray[UInt] in current MoonBit version
		64,  // FixedArray[Byte] in current MoonBit version
		112: // FixedArray[Double/Int64] in current MoonBit version
		// For classID 96 (FixedArray[UInt]), trust the words field as element count
		if classID == 96 {
			// For FixedArray[UInt], numElements comes from words and is authoritative
			if numElements == 0 {
				return []T{}, nil
			}
		} else {
			// Fix for arrays where numElements is calculated incorrectly (for other classIDs)
			if numElements == 0 && len(sliceMemBlock) > 8 {
				// Calculate numElements from actual memory block data size
				dataSize := len(sliceMemBlock) - 8 // subtract header size
				if elemType.Name() == "Bool" || elemType.Name() == "Char" {
					elemTypeSize = 4
				}
				numElements = uint32(dataSize) / uint32(elemTypeSize)
			} else if numElements == 0 {
				return []T{}, nil
			}
		}
	case FixedArrayByteBlockType, // Byte
		StringBlockType: // Int16, Char
		// remainderOffset := words*4 + 7
		// remainder := uint32(3 - sliceMemBlock[remainderOffset]%4)
		// size := (words-1)*4 + remainder
		size := words
		if size <= 0 {
			return []T{}, nil // empty slice
		}
		if int(size)+8 > len(sliceMemBlock) {
			return nil, fmt.Errorf("expected byte data size %v, got %v", size, len(sliceMemBlock))
		}

		sliceMemBlock = sliceMemBlock[:size+8]    // trim to the actual size
		numElements = size / uint32(elemTypeSize) // and adjust the actual number of elements
	case TupleBlockType: // classID 0 - can occur for shared empty arrays
		// For TupleBlockType (classID 0), calculate element count from memory block
		// In current MoonBit, this case occurs for inner arrays of Some([...])
		if len(sliceMemBlock) <= 8 {
			return []T{}, nil // Only header, truly empty
		}
		// If we have data, treat it like a regular primitive array
		// Calculate numElements from available data
		dataSize := len(sliceMemBlock) - 8
		elemTypeSize := h.converter.TypeSize()
		if elemType.Name() == "Bool" || elemType.Name() == "Char" {
			elemTypeSize = 4
		}
		numElements = uint32(dataSize) / uint32(elemTypeSize)
	default:
		return nil, fmt.Errorf("primitiveSliceHandler.Decode: unexpected classID %v", classID)
	}

	// TODO: Figure out how to not make special cases.
	if elemType.Name() == "Bool" {
		items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
		for i := 0; i < int(numElements); i++ {
			item := binary.LittleEndian.Uint32(sliceMemBlock[8+i*elemTypeSize:])
			val := item != 0
			items.Index(int(i)).Set(reflect.ValueOf(val))
		}
		return items.Interface(), nil
	}
	if elemType.Name() == "Char" {
		items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
		for i := 0; i < int(numElements); i++ {
			val := int16(binary.LittleEndian.Uint32(sliceMemBlock[8+i*elemTypeSize:]))
			items.Index(int(i)).Set(reflect.ValueOf(val))
		}
		return items.Interface(), nil
	}

	items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
	for i := 0; i < int(numElements); i++ {
		var v uint64
		switch elemTypeSize {
		case 8:
			v = binary.LittleEndian.Uint64(sliceMemBlock[8+i*elemTypeSize:])
		case 4:
			v = uint64(binary.LittleEndian.Uint32(sliceMemBlock[8+i*elemTypeSize:]))
		case 2:
			v = uint64(binary.LittleEndian.Uint16(sliceMemBlock[8+i*elemTypeSize:]))
		case 1:
			v = uint64(sliceMemBlock[8+i*elemTypeSize])
		default:
			return nil, fmt.Errorf("unsupported element type size: %v", elemTypeSize)
		}
		val := h.converter.Decode(v)
		items.Index(int(i)).Set(reflect.ValueOf(val))
	}
	return items.Interface(), nil
}

func (h *primitiveSliceHandler[T]) Encode(ctx context.Context, wasmAdapter langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	wa, ok := wasmAdapter.(wasmMemoryWriter)
	if !ok {
		return nil, nil, fmt.Errorf("expected a wasmMemoryWriter, got %T", wasmAdapter)
	}

	ptr, cln, err := h.doWriteSlice(ctx, wa, obj)
	if err != nil {
		return nil, cln, err
	}

	return []uint64{uint64(ptr)}, cln, nil
}

func (h *primitiveSliceHandler[T]) doWriteSlice(ctx context.Context, wa wasmMemoryWriter, obj any) (uint32, utils.Cleaner, error) {
	if utils.HasNil(obj) {
		return 0, nil, nil
	}

	slice, ok := utils.ConvertToSliceOf[T](obj)
	if !ok {
		return 0, nil, fmt.Errorf("expected a %T, got %T", []T{}, obj)
	}

	numElements := uint32(len(slice))
	elemTypeSize := h.converter.TypeSize()
	elemType := h.typeInfo.ListElementType()
	if elemType.Name() == "Bool" || elemType.Name() == "Char" {
		// A MoonBit Bool is 4 bytes whereas a Go bool is 1 byte.
		// A MoonBit Array[Char] uses 4 bytes per element instead of 2.
		elemTypeSize = 4
	}

	size := numElements * uint32(elemTypeSize)
	var memBlockClassID uint32
	var writeHeader func([]byte)
	switch elemType.Name() {
	case "Bool", "Char", "Int", "Float":
		memBlockClassID = FixedArrayPrimitiveBlockType // 241
	case "UInt":
		memBlockClassID = 96 // FixedArray[UInt] in current MoonBit version
	case "Byte":
		memBlockClassID = 64 // FixedArray[Byte] in current MoonBit version
	case "Int64", "UInt64", "Double":
		memBlockClassID = 112 // FixedArray[Double/Int64] in current MoonBit version
	case "Int16", "UInt16":
		memBlockClassID = StringBlockType // 80
	default:
		return 0, nil, fmt.Errorf("unsupported primitive MoonBit slice type: %v", elemType.Name())
	}

	// Handle different types based on elemTypeSize
	if memBlockClassID == 64 || memBlockClassID == StringBlockType { // Byte (64) or Int16/UInt16 (80)
		paddedSize := ((size + 4) / 4) * 4
		padding := uint8(3 - (size % 4))
		if padding != 0 {
			writeHeader = func(mem []byte) {
				// Write padding byte at the end
				mem[paddedSize-1] = padding
			}
		}

		size = paddedSize
		var zero T
		for i := numElements; i < paddedSize; i++ {
			slice = append(slice, zero) // add the padding bytes
		}
	}

	// Allocate memory
	var offset uint32
	var cln utils.Cleaner
	var err error
	if size == 0 {
		// For empty arrays, manually create the exact structure MoonBit produces
		// MoonBit empty FixedArray[UInt]: [4294967295, 1610612736, 0, 0] (16 bytes)
		offset, cln, err = wa.allocateAndPinMemory(ctx, 4, memBlockClassID) // 4 words = 16 bytes
		if err != nil {
			return 0, cln, err
		}
		// Manually write the exact structure that MoonBit expects
		wa.Memory().WriteUint32Le(offset, 4294967295)   // Special marker for empty array
		wa.Memory().WriteUint32Le(offset+4, 1610612736) // (96 << 24) | 0
		wa.Memory().WriteUint32Le(offset+8, 0)          // padding
		wa.Memory().WriteUint32Le(offset+12, 0)         // padding
		// Fix the header to match: memType should indicate 4 words
		memType := (4 << 8) | memBlockClassID // 4 words, classID 96
		wa.Memory().WriteUint32Le(offset-4, memType)
	} else {
		// For non-empty arrays, use MoonBit's exported malloc function
		// This ensures GC compatibility by using MoonBit's own allocation
		// Cast to concrete adapter type to access GetFunction
		concreteWa, ok := wa.(*wasmAdapter)
		if !ok {
			// Fall back to manual allocation if we can't access GetFunction
			allocSize := size / 4
			if numElements == 1 {
				// For 1-element arrays, allocate space for [elem, header, elem] = 3 words
				allocSize = 3
			}

			offset, cln, err = wa.allocateAndPinMemory(ctx, allocSize, memBlockClassID)
			if err != nil {
				return 0, cln, err
			}

			// For Int64, UInt64, and Double, the `words` portion of the memory block
			// indicates the number of elements in the slice, not the number of 16-bit words.
			if elemType.Name() == "Int64" || elemType.Name() == "UInt64" || elemType.Name() == "Double" {
				memType := ((size / 8) << 8) | memBlockClassID
				wa.Memory().WriteUint32Le(offset-4, memType)
			}
		} else {
			// Use MoonBit's malloc function
			fnMalloc := concreteWa.GetFunction("malloc")
			if fnMalloc == nil {
				return 0, cln, fmt.Errorf("malloc function not found")
			}

			// Call malloc(size) - MoonBit's malloc handles the headers
			// The returned pointer points to the data area (after headers)
			res, err := fnMalloc.Call(ctx, uint64(size))
			if err != nil {
				return 0, cln, fmt.Errorf("failed to call malloc: %w", err)
			}

			offset = uint32(res[0])
			if offset == 0 {
				return 0, cln, fmt.Errorf("malloc returned null pointer")
			}

			// MoonBit's malloc returns data pointer, but we need to adjust the array header
			// for the specific array type we're creating
			// The array header is at offset-4
			arrayHeader := uint32(0x60000000) | numElements // (1<<30) | (2<<28) | numElements
			wa.Memory().WriteUint32Le(offset-4, arrayHeader)

			// Create a simple cleaner function
			cln = utils.NewCleanerN(0)
		}
	}

	var dataBuffer []byte
	if elemType.Name() == "Bool" {
		dataBuffer = make([]byte, numElements*4)
		var zero T
		for i := 0; i < len(slice); i++ {
			if slice[i] == zero {
				binary.LittleEndian.PutUint32(dataBuffer[i*4:], 0)
			} else {
				binary.LittleEndian.PutUint32(dataBuffer[i*4:], 1)
			}
		}
	} else if elemType.Name() == "Char" {
		dataBuffer = make([]byte, numElements*4)
		for i := 0; i < len(slice); i++ {
			val := reflect.ValueOf(slice[i])
			binary.LittleEndian.PutUint32(dataBuffer[i*4:], uint32(val.Int()))
		}
	} else {
		// Allocate data buffer and write using the appropriate function
		dataBuffer = h.converter.SliceToBytes(slice)
	}
	if writeHeader != nil {
		writeHeader(dataBuffer)
	}

	// Write the data buffer
	if size == 0 {
		// For empty arrays, data is already written above
	} else {
		// For non-empty arrays, write data starting at the correct offset
		// If we used malloc, offset already points to the data area
		// If we used manual allocation, we need to adjust
		concreteWa, usedMalloc := wa.(*wasmAdapter)
		if usedMalloc && concreteWa.GetFunction("malloc") != nil {
			// malloc returns pointer to data area, write directly
			if ok := wa.Memory().Write(offset, dataBuffer); !ok {
				return 0, cln, errors.New("failed to write data to WASM memory")
			}
		} else {
			// Manual allocation, offset points to start of data section
			if ok := wa.Memory().Write(offset, dataBuffer); !ok {
				return 0, cln, errors.New("failed to write data to WASM memory")
			}
		}
	}

	// For FixedArray, try minimal header addition for 1-element arrays
	if strings.HasPrefix(h.typeDef.Name, "FixedArray[") && numElements == 1 {
		// For 1-element arrays, try expanding to match MoonBit's [elem, header, elem] pattern
		// Allocate additional space and restructure as [first_elem, header, first_elem]
		header := (memBlockClassID << 24) | numElements // (96 << 24) | 1 = 1610612737

		// Read the current element
		firstElem, _ := wa.Memory().ReadUint32Le(offset)

		// Write the structured format: [elem, header, elem]
		wa.Memory().WriteUint32Le(offset, firstElem)   // first element
		wa.Memory().WriteUint32Le(offset+4, header)    // header
		wa.Memory().WriteUint32Le(offset+8, firstElem) // duplicate element

		// Update the allocation header to reflect 3 words instead of 1
		newMemType := (3 << 8) | memBlockClassID // 3 words, classID 96
		wa.Memory().WriteUint32Le(offset-4, newMemType)
	}

	if strings.HasPrefix(h.typeDef.Name, "Array[") {
		// Finally, write the slice memory block.
		slicePtr, sliceCln, err := wa.allocateAndPinMemory(ctx, 2, TupleBlockType) // was: 8
		innerCln := utils.NewCleanerN(1)
		innerCln.AddCleaner(sliceCln)
		if err != nil {
			return 0, cln, err
		}
		wa.Memory().WriteUint32Le(slicePtr, offset-8)
		wa.Memory().WriteUint32Le(slicePtr+4, numElements)

		return slicePtr - 8, cln, nil
	}

	if strings.HasPrefix(h.typeDef.Name, "FixedArray[") {
		// For FixedArray, adjust return value based on allocation method
		concreteWa, usedMalloc := wa.(*wasmAdapter)
		if usedMalloc && concreteWa.GetFunction("malloc") != nil && size > 0 {
			// malloc returns data pointer, but we need array pointer
			return offset - 8, cln, nil
		} else {
			// Manual allocation or empty array, return as-is
			return offset, cln, nil
		}
	}

	return offset, cln, nil
}
