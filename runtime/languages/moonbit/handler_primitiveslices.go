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

	typ, hasError, hasOption := stripErrorAndOption(ti.ListElementType().Name())
	gmlPrintf("GML: handler_primitiveslices.go: NewPrimitiveSliceHandler('%v'): '%v', hasError=%v, hasOption=%v", ti.Name(), typ, hasError, hasOption)

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
	gmlPrintf("GML: handler_primitiveslices.go: primitiveSliceHandler[%T].Read(offset: %v)", []T{}, debugShowOffset(offset))
	return h.Decode(ctx, wa, []uint64{uint64(offset)})

	// if offset == 0 {
	// 	return nil, nil
	// }

	// data, size, _, err := wa.(*wasmAdapter).readSliceHeader(offset)
	// if err != nil {
	// 	return nil, err
	// }

	// return h.doReadSlice(wa, data, size)
}

func (h *primitiveSliceHandler[T]) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	ptr, cln, err := h.doWriteSlice(ctx, wa.(*wasmAdapter), obj)
	if err != nil {
		return cln, err
	}

	if ok := utils.CopyMemory(wa.Memory(), ptr, offset, 12); !ok {
		return cln, errors.New("failed to copy slice header")
	}

	return cln, nil
}

func (h *primitiveSliceHandler[T]) Decode(ctx context.Context, wasmAdapter langsupport.WasmAdapter, vals []uint64) (any, error) {
	wa, ok := wasmAdapter.(wasmMemoryReader)
	if !ok {
		return nil, fmt.Errorf("expected a wasmMemoryReader, got %T", wasmAdapter)
	}

	gmlPrintf("GML: handler_primitiveslices.go: primitiveSliceHandler.Decode(len(vals): %v)", len(vals))

	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 value when decoding a primitive slice but got %v: %+v", len(vals), vals)
	}

	if vals[0] == 0 {
		return nil, nil
	}

	memBlock, _, err := memoryBlockAtOffset(wa, uint32(vals[0]), 0, true)
	if err != nil {
		return nil, err
	}

	numElements := binary.LittleEndian.Uint32(memBlock[12:16])
	if numElements == 0 {
		// For debugging:
		sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
		_, _, _ = memoryBlockAtOffset(wa, sliceOffset, 0, true)

		return []T{}, nil
	}

	sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
	elemTypeSize := h.converter.TypeSize()
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
	size := numElements * uint32(elemTypeSize)
	gmlPrintf("GML: handler_primitiveslices.go: primitiveSliceHandler.Decode: sliceOffset=%v, numElements=%v, size=%v", debugShowOffset(sliceOffset), numElements, size)

	// For reverse engineering:
	_, _, _ = memoryBlockAtOffset(wa, sliceOffset, 0, true)

	sliceMemBlock, _, err := memoryBlockAtOffset(wa, sliceOffset, size, true)
	if err != nil {
		return nil, err
	}

	// gmlPrintf("GML: handler_primitiveslices.go: primitiveSliceHandler.Decode: (sliceOffset: %v, numElements: %v, size: %v), sliceMemBlock=%+v", sliceOffset, numElements, size, sliceMemBlock)

	if elemType.Name() == "Bool" {
		items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
		for i := 0; i < int(numElements); i++ {
			// if isNullable {
			// 	value32 := binary.LittleEndian.Uint32(memBlock[8+i*elemTypeSize:])
			// 	value := uint64(value32)
			// 	item, err := h.elementHandler.Decode(ctx, wa, []uint64{value})
			// 	if err != nil {
			// 		return nil, err
			// 	}
			// 	if !utils.HasNil(item) {
			// 		items.Index(int(i)).Set(reflect.ValueOf(item))
			// 	}
			// 	continue
			// }
			item := binary.LittleEndian.Uint32(sliceMemBlock[8+i*elemTypeSize:])
			val := item != 0
			items.Index(int(i)).Set(reflect.ValueOf(val))
		}
		return items.Interface(), nil
	}

	// TODO: Figure out how to not make special cases.
	if elemType.Name() == "Char" {
		items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
		for i := 0; i < int(numElements); i++ {
			val := int16(binary.LittleEndian.Uint32(sliceMemBlock[8+i*elemTypeSize:]))
			items.Index(int(i)).Set(reflect.ValueOf(val))
		}
		return items.Interface(), nil
	}
	if elemType.Name() == "Int16" {
		items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
		for i := 0; i < int(numElements); i++ {
			val := int16(binary.LittleEndian.Uint16(sliceMemBlock[8+i*elemTypeSize:]))
			items.Index(int(i)).Set(reflect.ValueOf(val))
		}
		return items.Interface(), nil
	}

	items := h.converter.BytesToSlice(sliceMemBlock[8:])
	return items, nil
}

func (h *primitiveSliceHandler[T]) Encode(ctx context.Context, wasmAdapter langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	wa, ok := wasmAdapter.(wasmMemoryWriter)
	if !ok {
		return nil, nil, fmt.Errorf("expected a wasmMemoryWriter, got %T", wasmAdapter)
	}

	gmlPrintf("GML: handler_primitiveslices.go: primitiveSliceHandler.Encode: obj=%T", obj)
	ptr, cln, err := h.doWriteSlice(ctx, wa, obj)
	if err != nil {
		return nil, cln, err
	}

	return []uint64{uint64(ptr)}, cln, nil
}

/*
func (h *primitiveSliceHandler[T]) doReadSlice(wa langsupport.WasmAdapter, data, size uint32) (any, error) {
	if data == 0 {
		// nil slice
		return nil, nil
	}

	if size == 0 {
		// empty slice
		return []T{}, nil
	}

	bufferSize := size * uint32(h.converter.TypeSize())
	buf, ok := wa.Memory().Read(data, bufferSize)
	if !ok {
		return nil, errors.New("failed to read data from WASM memory")
	}

	items := h.converter.BytesToSlice(buf)
	return items, nil
}
*/

func (h *primitiveSliceHandler[T]) doWriteSlice(ctx context.Context, wa wasmMemoryWriter, obj any) (uint32, utils.Cleaner, error) {
	gmlPrintf("GML: handler_primitiveslices.go: doWriteSlice: T: %T", []T{})

	if utils.HasNil(obj) {
		return 0, nil, nil
	}

	slice, ok := utils.ConvertToSliceOf[T](obj)
	if !ok {
		return 0, nil, fmt.Errorf("expected a %T, got %T", []T{}, obj)
	}

	numElements := uint32(len(slice))
	elementSize := h.converter.TypeSize()
	elemType := h.typeInfo.ListElementType()
	if elemType.Name() == "Bool" {
		// A MoonBit Bool is 4 bytes whereas a Go bool is 1 byte.
		elementSize = 4
	}
	// gmlPrintf("GML: handler_primitiveslices.go: doWriteSlice: len(slice): %v, numElements: %v, elementSize: %v, slice: %+v", len(slice), numElements, elementSize, slice)

	var size uint32
	// var headerValue uint32
	var memBlockClassID uint32
	var writeHeader func([]byte)

	// Handle different types based on elementSize
	if elementSize == 1 {
		// Byte arrays: round up to nearest 4 bytes + padding byte
		paddedSize := ((numElements + 5) / 4) * 4
		// gmlPrintf("GML: handler_primitiveslices.go: doWriteSlice: numElements: %v, paddedSize: %v", numElements, paddedSize)
		size = numElements
		// headerValue = ((paddedSize / 4) << 8) | 246 // 246 is the byte array header type
		memBlockClassID = 246
		var zero T
		for i := numElements; i < paddedSize; i++ {
			// gmlPrintf("GML: handler_primitiveslices.go: doWriteSlice: ADDING PADDING BYTE #%v of %v", i+1-numElements, paddedSize-numElements)
			slice = append(slice, zero) // add the padding bytes
		}

		writeHeader = func(mem []byte) {
			// // Write the data bytes
			// for i, b := range slice {
			// 	mem[i] = b
			// }
			// Write padding byte at the end
			padding := uint8(3 - (numElements % 4))
			mem[paddedSize-1] = padding
		}
	} else {
		// Int arrays: 4 bytes per element + header
		size = numElements * uint32(elementSize)
		// headerValue = (numElements << 8) | 241 // 241 is the int array header type
		memBlockClassID = ArrayBlockType

		// writeHeader = func(mem []byte) {
		// 	for i, val := range slice {
		// 		binary.LittleEndian.PutUint32(mem[i*4:(i+1)*4], uint32(val.(int32)))
		// 	}
		// }
	}

	// Allocate memory
	var offset uint32
	var cln utils.Cleaner
	var err error
	if size == 0 {
		offset, cln, err = wa.allocateAndPinMemory(ctx, 1, memBlockClassID) // cannot allocate 0 bytes
		if err != nil {
			return 0, cln, err
		}
		wa.Memory().WriteByte(offset-3, 0) // overwrite size=1 to size=0
	} else {
		offset, cln, err = wa.allocateAndPinMemory(ctx, size, memBlockClassID)
		if err != nil {
			return 0, cln, err
		}
		// For both Int64 and UInt64, the `words` portion of the memory block
		// indicates the number of elements in the slice, not the number of 16-bit words.
		if elemType.Name() == "Int64" || elemType.Name() == "UInt64" {
			memType := ((size / 8) << 8) | memBlockClassID
			wa.Memory().WriteUint32Le(offset-4, memType)
		}
	}

	// Write header
	// header := make([]byte, 8)
	// binary.LittleEndian.PutUint32(header[4:8], headerValue) // Header at offset 4
	// if ok := wa.Memory().Write(offset, header); !ok {
	// 	return 0, cln, errors.New("failed to write header to WASM memory")
	// }

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
	} else {
		// Allocate data buffer and write using the appropriate function
		dataBuffer = h.converter.SliceToBytes(slice)
	}
	// gmlPrintf("GML: handler_primitiveslices.go: doWriteSlice: BEFORE WRITING HEADER: dataBuffer: %+v", dataBuffer)
	if writeHeader != nil {
		writeHeader(dataBuffer)
	}
	// gmlPrintf("GML: handler_primitiveslices.go: doWriteSlice: AFTER WRITING HEADER: dataBuffer: %+v", dataBuffer)

	if ok := wa.Memory().Write(offset, dataBuffer); !ok {
		return 0, cln, errors.New("failed to write data to WASM memory")
	}

	// For debugging:
	_, _, _ = memoryBlockAtOffset(wa, offset-8, 0, true)

	// Finally, write the slice memory block.
	slicePtr, sliceCln, err := wa.allocateAndPinMemory(ctx, 8, 0)
	innerCln := utils.NewCleanerN(1)
	innerCln.AddCleaner(sliceCln)
	if err != nil {
		return 0, cln, err
	}
	wa.Memory().WriteUint32Le(slicePtr, offset-8)
	wa.Memory().WriteUint32Le(slicePtr+4, numElements)

	// For debugging purposes:
	_, _, _ = memoryBlockAtOffset(wa, slicePtr-8, 0, true)

	return slicePtr - 8, cln, nil
}
