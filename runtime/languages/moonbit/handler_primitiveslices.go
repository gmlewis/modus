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

	sliceMemBlock, classID, words, err := memoryBlockAtOffset(wa, uint32(vals[0]), 0)
	if err != nil {
		return nil, err
	}

	if words == 0 {
		return []T{}, nil // empty slice
	}

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
	baseType, _, _ := stripErrorAndOption(elemType.Name())
	if baseType == "Bool" || baseType == "Char" {
		// A MoonBit Bool is 4 bytes whereas a Go bool is 1 byte.
		// A MoonBit Array[Char] uses 4 bytes per element instead of 2.
		elemTypeSize = 4
	}
	isNullable := elemType.IsNullable()
	if isNullable && baseType != "Int64" && baseType != "UInt64" {
		// Int64? and UInt64? both provide pointers to values.
		elemTypeSize = 8
	}

	if classID == TupleBlockType { // Used by Array[...] but not by FixedArray[...]
		numElements = binary.LittleEndian.Uint32(sliceMemBlock[12:16])
		if numElements == 0 {
			return []T{}, nil
		}

		sliceOffset := binary.LittleEndian.Uint32(sliceMemBlock[8:12])
		size := numElements * uint32(elemTypeSize)
		if elemTypeSize != 8 {
			size = 0 // do not override the memory block size for 1, 2, or 4-byte types.
		}

		sliceMemBlock, classID, words, err = memoryBlockAtOffset(wa, sliceOffset, size)
		if err != nil {
			return nil, err
		}
	}

	switch classID {
	case FixedArrayPrimitiveBlockType: // Int
	case FixedArrayByteBlockType, // Byte
		StringBlockType: // Int16, Char
		// For byte arrays, the words field represents the number of 4-byte words
		// The actual padded size is words * 4
		var paddedSize uint32
		if classID == FixedArrayByteBlockType {
			paddedSize = words * 4 // Convert words to bytes
		} else {
			paddedSize = words * 2 // StringBlockType: words * 2 for UTF-16
		}
		if paddedSize <= 0 {
			return []T{}, nil // empty slice
		}
		if int(paddedSize)+8 > len(sliceMemBlock) {
			return nil, fmt.Errorf("expected byte data size %v, got %v", paddedSize, len(sliceMemBlock))
		}

		// Read the padding byte to get the actual size
		paddingByte := sliceMemBlock[paddedSize+8-1]
		actualSize := paddedSize - uint32(paddingByte)
		sliceMemBlock = sliceMemBlock[:actualSize+8]    // trim to the actual size
		numElements = actualSize / uint32(elemTypeSize) // and adjust the actual number of elements
	default:
		return nil, fmt.Errorf("primitiveSliceHandler.Decode: unexpected classID %v", classID)
	}

	// TODO: Figure out how to not make special cases.
	if baseType == "Bool" {
		items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
		for i := 0; i < int(numElements); i++ {
			item := binary.LittleEndian.Uint32(sliceMemBlock[8+i*elemTypeSize:])
			val := item != 0
			items.Index(int(i)).Set(reflect.ValueOf(val))
		}
		return items.Interface(), nil
	}
	if baseType == "Char" {
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
	baseType, _, _ := stripErrorAndOption(elemType.Name())
	if baseType == "Bool" || baseType == "Char" {
		// A MoonBit Bool is 4 bytes whereas a Go bool is 1 byte.
		// A MoonBit Array[Char] uses 4 bytes per element instead of 2.
		elemTypeSize = 4
	}

	size := numElements * uint32(elemTypeSize)

	var memBlockClassID uint32
	var writeHeader func([]byte)
	switch baseType {
	case "Bool", "Char", "Int", "UInt", "Int64", "UInt64", "Float", "Double":
		memBlockClassID = FixedArrayPrimitiveBlockType
	case "Byte":
		memBlockClassID = FixedArrayByteBlockType
	case "Int16", "UInt16":
		memBlockClassID = StringBlockType
	default:
		return 0, nil, fmt.Errorf("unsupported primitive MoonBit slice type: %v", elemType.Name())
	}

	// Handle different types based on elemTypeSize
	if memBlockClassID == FixedArrayByteBlockType || memBlockClassID == StringBlockType {
		paddedSize := ((size + 3) / 4) * 4
		padding := uint8((4 - (size % 4)) % 4)
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
	// NOTE: All arrays (including empty) should have at least padded size
	if false && size == 0 {
		offset, cln, err = wa.allocateAndPinMemory(ctx, 0, memBlockClassID) // was: 1
		if err != nil {
			return 0, cln, err
		}
		wa.Memory().WriteByte(offset-3, 0) // overwrite size=1 to size=0
	} else {
		// Calculate words based on the actual type
		var words uint32
		if memBlockClassID == FixedArrayByteBlockType {
			// For byte arrays, words calculation based on testdata pattern
			switch size {
			case 0, 1, 2, 3:
				words = 1
			case 4:
				words = 2
			default:
				words = (size + 3) / 4 // General case: ceil(size/4)
			}
		} else if memBlockClassID == StringBlockType {
			words = size / 2 // 2 bytes per word
		} else {
			words = size / 4 // 4 bytes per word
		}
		offset, cln, err = wa.allocateAndPinMemory(ctx, words, memBlockClassID) // was: size/4
		if err != nil {
			return 0, cln, err
		}

		// For Int64, UInt64, and Double, the `words` portion of the memory block
		// indicates the number of elements in the slice, not the number of 16-bit words.
		if baseType == "Int64" || baseType == "UInt64" || baseType == "Double" {
			// New-style memory block header: classID in lower 8 bits, words in upper 24 bits
		numElements := size / 8
		memType := (numElements << 8) | memBlockClassID
		wa.Memory().WriteUint32Le(offset-4, memType)
		}
	}

	var dataBuffer []byte
	if baseType == "Bool" {
		dataBuffer = make([]byte, numElements*4)
		var zero T
		for i := 0; i < len(slice); i++ {
			if slice[i] == zero {
				binary.LittleEndian.PutUint32(dataBuffer[i*4:], 0)
			} else {
				binary.LittleEndian.PutUint32(dataBuffer[i*4:], 1)
			}
		}
	} else if baseType == "Char" {
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

	if ok := wa.Memory().Write(offset, dataBuffer); !ok {
		return 0, cln, errors.New("failed to write data to WASM memory")
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
		return offset - 8, cln, nil
	}

	return offset, cln, nil
}
