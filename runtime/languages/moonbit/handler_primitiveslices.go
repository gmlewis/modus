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
	"math"
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

	// Check if this is a wrapper structure (contains type info 1573120)
	offset := uint32(vals[0])
	isBytesData := false
	var arrayLength uint32

	// Try to read the type info at offset 4
	typeInfoBytes, ok := wa.Memory().Read(offset+4, 4)
	if ok {
		typeInfo := binary.LittleEndian.Uint32(typeInfoBytes)
		if typeInfo == 1573120 { // Array type wrapper
			// Read the length and data pointer from the wrapper
			lengthBytes, ok := wa.Memory().Read(offset+8, 4)
			if ok {
				arrayLength = binary.LittleEndian.Uint32(lengthBytes)
			}
			dataPointerBytes, ok := wa.Memory().Read(offset+12, 4)
			if ok {
				dataPointer := binary.LittleEndian.Uint32(dataPointerBytes)
				// Check if data pointer is pointing to a Bytes object (for Array[Byte] from fnBytes2Array)
				if elemType := h.typeInfo.ListElementType(); elemType.Name() == "Byte" {
					// Check if this is a Bytes object by reading its type info
					if bytesTypeBytes, ok := wa.Memory().Read(dataPointer+4, 4); ok {
						bytesTypeInfo := binary.LittleEndian.Uint32(bytesTypeBytes)
						// Check if this is a Bytes object (has pattern 0x4000000X where X is length)
						if (bytesTypeInfo & 0xFF000000) == 0x40000000 { // Bytes type info pattern
							// This is a Bytes object, adjust offset to actual data and use special handling
							offset = dataPointer + 8 // Skip Bytes header
							isBytesData = true
						} else {
							offset = dataPointer
						}
					} else {
						offset = dataPointer
					}
				} else {
					offset = dataPointer
				}
			}
		}
	}

	// Special handling for Array[Byte] data from fnBytes2Array
	if isBytesData && h.typeInfo.ListElementType().Name() == "Byte" {
		if arrayLength == 0 {
			return []T{}, nil
		}
		// Read raw byte data directly
		dataBytes, ok := wa.Memory().Read(offset, arrayLength)
		if !ok {
			return nil, fmt.Errorf("failed to read byte data at offset %d, length %d", offset, arrayLength)
		}
		// Convert to slice
		items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(arrayLength), int(arrayLength))
		for i := uint32(0); i < arrayLength; i++ {
			val := h.converter.Decode(uint64(dataBytes[i]))
			items.Index(int(i)).Set(reflect.ValueOf(val))
		}
		return items.Interface(), nil
	}

	// First read to get the header and determine the classID
	headerBlock, classID, words, err := memoryBlockAtOffset(wa, offset, 0)
	if err != nil {
		// Check if this is a dynamic Array[T] that needs fallback (e.g., from Map handler)
		isFixedArray := strings.HasPrefix(h.typeDef.Name, "FixedArray[")
		if !isFixedArray && strings.Contains(err.Error(), "invalid memory offset") {
			// This is likely a dynamic array created by moonbit.i32_array_make
			// Use direct memory reading approach as fallback
			return h.decodeDynamicPrimitiveArray(ctx, wa, offset)
		}
		return nil, err
	}

	// For new classIDs, calculate the correct size and re-read
	var sliceMemBlock []byte
	if classID == 96 || classID == 64 || classID == 112 {
		// Calculate the correct size for these classIDs
		elemTypeSize := h.converter.TypeSize()
		elemType := h.typeInfo.ListElementType()
		if elemType.Name() == "Bool" || elemType.Name() == "Char" || elemType.Name() == "Byte" {
			elemTypeSize = MoonBitBoolSize
		}
		dataSize := words * uint32(elemTypeSize)

		// For empty arrays (words=0), return empty slice immediately
		if words == 0 {
			return []T{}, nil
		}

		totalSize := dataSize // override with correct size
		sliceMemBlock, classID, words, err = memoryBlockAtOffset(wa, offset, totalSize)
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
		sliceMemBlock, _, _, err = memoryBlockAtOffset(wa, offset, size)
		if err != nil {
			return nil, err
		}
	}

	elemType := h.typeInfo.ListElementType()
	if elemType.Name() == "Bool" || elemType.Name() == "Char" || elemType.Name() == "Byte" {
		// A MoonBit Bool is 4 bytes whereas a Go bool is 1 byte.
		// A MoonBit Array[Char] uses 4 bytes per element instead of 2.
		elemTypeSize = MoonBitBoolSize
	}
	isNullable := elemType.IsNullable()
	if isNullable && elemType.Name() != "Int64?" && elemType.Name() != "UInt64?" {
		// Int64? and UInt64? both provide pointers to values.
		elemTypeSize = Int64Size
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
		BoolByteCharClassID, // FixedArray[UInt] in current MoonBit version
		64,                  // FixedArray[Byte] in current MoonBit version
		Int64DoubleClassID:  // FixedArray[Double/Int64] in current MoonBit version
		// For classID 96 (FixedArray[UInt]), trust the words field as element count
		if classID == BoolByteCharClassID {
			// For FixedArray[UInt], numElements comes from words and is authoritative
			if numElements == 0 {
				return []T{}, nil
			}
		} else {
			// Fix for arrays where numElements is calculated incorrectly (for other classIDs)
			if numElements == 0 && len(sliceMemBlock) > MemoryBlockHeaderSize {
				// Calculate numElements from actual memory block data size
				dataSize := len(sliceMemBlock) - MemoryBlockHeaderSize // subtract header size
				if elemType.Name() == "Bool" || elemType.Name() == "Char" || elemType.Name() == "Byte" {
					elemTypeSize = MoonBitBoolSize
				}
				numElements = uint32(dataSize) / uint32(elemTypeSize)
			} else if numElements == 0 {
				return []T{}, nil
			}
		}
	case StringBlockType: // FixedArray[Int16/UInt16]
		// For Int16/UInt16 arrays, handle similar to other primitive arrays
		if numElements == 0 && len(sliceMemBlock) > MemoryBlockHeaderSize {
			// Calculate numElements from actual memory block data size
			dataSize := len(sliceMemBlock) - MemoryBlockHeaderSize // subtract header size
			numElements = uint32(dataSize) / uint32(elemTypeSize)
		} else if numElements == 0 {
			return []T{}, nil
		}
	case FixedArrayByteBlockType: // Byte
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
		if len(sliceMemBlock) <= MemoryBlockHeaderSize {
			return []T{}, nil // Only header, truly empty
		}
		// If we have data, treat it like a regular primitive array
		// Calculate numElements from available data
		dataSize := len(sliceMemBlock) - MemoryBlockHeaderSize
		elemTypeSize := h.converter.TypeSize()
		if elemType.Name() == "Bool" || elemType.Name() == "Char" || elemType.Name() == "Byte" {
			elemTypeSize = MoonBitBoolSize
		}
		numElements = uint32(dataSize) / uint32(elemTypeSize)
	default:
		return nil, fmt.Errorf("primitiveSliceHandler.Decode: unexpected classID %v", classID)
	}

	// TODO: Figure out how to not make special cases.
	if elemType.Name() == "Bool" {
		for i := 0; i < 16 && i < len(sliceMemBlock); i += 4 {
		}

		items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
		for i := 0; i < int(numElements); i++ {
			offset := MemoryBlockHeaderSize + i*elemTypeSize
			item := binary.LittleEndian.Uint32(sliceMemBlock[offset:])
			val := item != 0
			items.Index(int(i)).Set(reflect.ValueOf(val))
		}
		return items.Interface(), nil
	}
	if elemType.Name() == "Char" {
		items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
		for i := 0; i < int(numElements); i++ {
			val := int16(binary.LittleEndian.Uint32(sliceMemBlock[MemoryBlockHeaderSize+i*elemTypeSize:]))
			items.Index(int(i)).Set(reflect.ValueOf(val))
		}
		return items.Interface(), nil
	}

	items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
	for i := 0; i < int(numElements); i++ {
		var v uint64
		switch elemTypeSize {
		case Int64Size:
			v = binary.LittleEndian.Uint64(sliceMemBlock[MemoryBlockHeaderSize+i*elemTypeSize:])
		case StandardPtrSize:
			v = uint64(binary.LittleEndian.Uint32(sliceMemBlock[MemoryBlockHeaderSize+i*elemTypeSize:]))
		case 2:
			v = uint64(binary.LittleEndian.Uint16(sliceMemBlock[MemoryBlockHeaderSize+i*elemTypeSize:]))
		case 1:
			v = uint64(sliceMemBlock[MemoryBlockHeaderSize+i*elemTypeSize])
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

	// Check if this is a dynamic Array[T] (not FixedArray[T])
	isFixedArray := strings.HasPrefix(h.typeDef.Name, "FixedArray[")

	// Both Array[Bool] and Array[Byte] now use dynamic array path with native MoonBit functions
	// Array[Bool] uses moonbit_i32_array_make, Array[Byte] uses fnBytes2Array

	if !isFixedArray {
		// For dynamic Array[T], use MoonBit's native array creation functions
		return h.createDynamicPrimitiveArray(ctx, wa, slice, numElements, elemType)
	}

	if elemType.Name() == "Bool" || elemType.Name() == "Char" || elemType.Name() == "Byte" {
		// A MoonBit Bool is 4 bytes whereas a Go bool is 1 byte.
		// A MoonBit Array[Char] uses 4 bytes per element instead of 2.
		elemTypeSize = MoonBitBoolSize
	}

	size := numElements * uint32(elemTypeSize)
	var memBlockClassID uint32
	var writeHeader func([]byte)
	switch elemType.Name() {
	case "Bool":
		// TEMP FIX: Use BoolByteCharClassID (96) for Bool to match WAT-generated arrays
		memBlockClassID = BoolByteCharClassID // 96
	case "Char", "Int", "Float":
		memBlockClassID = FixedArrayPrimitiveBlockType // 241
	case "UInt":
		memBlockClassID = BoolByteCharClassID // FixedArray[UInt] in current MoonBit version
	case "Byte":
		memBlockClassID = 64 // FixedArray[Byte] in current MoonBit version
	case "Int64", "UInt64", "Double":
		memBlockClassID = Int64DoubleClassID // FixedArray[Double/Int64] in current MoonBit version
	case "Int16", "UInt16":
		memBlockClassID = StringBlockType // 80
	default:
		return 0, nil, fmt.Errorf("unsupported primitive MoonBit slice type: %v", elemType.Name())
	}

	// Handle different types based on elemTypeSize
	// Note: Skip padding for Byte arrays when using moonbit_bytes_make
	if memBlockClassID == StringBlockType { // Int16/UInt16 (80) - Byte arrays (64) use moonbit_bytes_make
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
		wa.Memory().WriteUint32Le(offset, EmptyArrayMarker1)   // Special marker for empty array
		wa.Memory().WriteUint32Le(offset+4, EmptyArrayMarker2) // (96 << 24) | 0
		wa.Memory().WriteUint32Le(offset+8, 0)                 // padding
		wa.Memory().WriteUint32Le(offset+12, 0)                // padding
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
			// Use malloc + ptr2*_array approach for GC-compatible arrays
			// Step 1: Allocate raw memory for our data using MoonBit's malloc
			malloc_res, err := concreteWa.fnMalloc.Call(ctx, uint64(size))
			if err != nil {
				return 0, cln, fmt.Errorf("failed to call moonbit malloc: %w", err)
			}
			if len(malloc_res) == 0 || malloc_res[0] == 0 {
				return 0, cln, errors.New("moonbit malloc returned null pointer")
			}
			dataPtr := uint32(malloc_res[0])

			// Step 2: We'll write data later and then convert to proper array
			// For now, store the data pointer
			offset = dataPtr
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
	} else if elemType.Name() == "Byte" {
		// For Byte arrays, MoonBit expects 4-byte values in fixed array infrastructure
		dataBuffer = make([]byte, numElements*4)
		for i := 0; i < len(slice); i++ {
			val := reflect.ValueOf(slice[i])
			binary.LittleEndian.PutUint32(dataBuffer[i*4:], uint32(val.Uint()))
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
		// For non-empty arrays, write data and convert to proper MoonBit array
		concreteWa, ok := wa.(*wasmAdapter)
		if ok && concreteWa.fnMalloc != nil {
			// Step 1: Write our data to the allocated memory
			if ok := wa.Memory().Write(offset, dataBuffer); !ok {
				return 0, cln, errors.New("failed to write data to allocated memory")
			}

			// Step 2: Convert raw memory to proper MoonBit array using ptr2*_array
			var arrayPtr []uint64
			var err error
			switch elemType.Name() {
			case "UInt":
				arrayPtr, err = concreteWa.fnPtr2uintArray.Call(ctx, uint64(offset), uint64(numElements))
			case "Bool":
				// SPECIAL: Use moonbit_bytes_make for Bool arrays like the WAT functions do
				// Use the same function as WAT: moonbit.i32_array_make(numElements, 0)
				arrayPtr, err = concreteWa.fnMakeArrayInt.Call(ctx, uint64(numElements), 0)
				if err != nil {
					return 0, cln, fmt.Errorf("failed to call moonbit_bytes_make: %w", err)
				}
				if len(arrayPtr) > 0 && arrayPtr[0] != 0 {
					boolArrayPtr := uint32(arrayPtr[0])
					// Write individual bool values to the allocated array
					for i := uint32(0); i < numElements; i++ {
						value := binary.LittleEndian.Uint32(dataBuffer[i*4:])
						// Write each bool at offset+8+i*4 (data starts at offset 8)
						boolAddr := boolArrayPtr + MemoryBlockHeaderSize + i*4
						wa.Memory().WriteUint32Le(boolAddr, value)
					}
					// Update offset to point to the bool array
					offset = boolArrayPtr
					// Return early since the array is properly allocated and initialized
					// Let Array wrapper creation happen
					return offset, cln, nil
					// return offset, cln, nil
				}
				// If array creation failed, fall back to default behavior
			case "Int", "Char":
				arrayPtr, err = concreteWa.fnPtr2intArray.Call(ctx, uint64(offset), uint64(numElements))
			case "Float":
				arrayPtr, err = concreteWa.fnPtr2floatArray.Call(ctx, uint64(offset), uint64(numElements))
			case "Double":
				arrayPtr, err = concreteWa.fnPtr2doubleArray.Call(ctx, uint64(offset), uint64(numElements))
			case "Int64":
				arrayPtr, err = concreteWa.fnPtr2int64Array.Call(ctx, uint64(offset), uint64(numElements))
			case "UInt64":
				arrayPtr, err = concreteWa.fnPtr2uint64Array.Call(ctx, uint64(offset), uint64(numElements))
			case "Int16":
				// For Int16, use moonbit_int16_array_make
				arrayPtr, err = concreteWa.fnMakeArrayInt16.Call(ctx, uint64(numElements), 0)
				if err != nil {
					return 0, cln, fmt.Errorf("failed to call moonbit_int16_array_make: %w", err)
				}
				// Write data to the created array
				if len(arrayPtr) > 0 && arrayPtr[0] != 0 {
					int16ArrayPtr := uint32(arrayPtr[0])
					// Write individual int16 values to the allocated array
					for i := uint32(0); i < numElements; i++ {
						val := binary.LittleEndian.Uint16(dataBuffer[i*2:])
						// Write each int16 at offset+8+i*2 (data starts at offset 8)
						int16Addr := int16ArrayPtr + MemoryBlockHeaderSize + i*2
						wa.Memory().WriteUint16Le(int16Addr, val)
					}
					// Update offset to point to the int16 array
					offset = int16ArrayPtr
					// Return early since the array is properly allocated and initialized
					// Let Array wrapper creation happen
					return offset, cln, nil
					// return offset, cln, nil
				}
			case "UInt16":
				// For UInt16, use moonbit_int16_array_make (same as Int16)
				arrayPtr, err = concreteWa.fnMakeArrayInt16.Call(ctx, uint64(numElements), 0)
				if err != nil {
					return 0, cln, fmt.Errorf("failed to call moonbit_int16_array_make for UInt16: %w", err)
				}
				// Write data to the created array
				if len(arrayPtr) > 0 && arrayPtr[0] != 0 {
					uint16ArrayPtr := uint32(arrayPtr[0])
					// Write individual uint16 values to the allocated array
					for i := uint32(0); i < numElements; i++ {
						val := binary.LittleEndian.Uint16(dataBuffer[i*2:])
						// Write each uint16 at offset+8+i*2 (data starts at offset 8)
						uint16Addr := uint16ArrayPtr + MemoryBlockHeaderSize + i*2
						wa.Memory().WriteUint16Le(uint16Addr, val)
					}
					// Update offset to point to the uint16 array
					offset = uint16ArrayPtr
					// Return early since the array is properly allocated and initialized
					// Let Array wrapper creation happen
					return offset, cln, nil
					// return offset, cln, nil
				}
			case "Byte":
				// For Byte arrays, use the exported moonbit_bytes_make function
				arrayPtr, err = concreteWa.fnBytesMake.Call(ctx, uint64(numElements), 0)
				if err != nil {
					return 0, cln, fmt.Errorf("failed to call moonbit_bytes_make: %w", err)
				}
				// Write individual bytes to the allocated array
				if len(arrayPtr) > 0 && arrayPtr[0] != 0 {
					byteArrayPtr := uint32(arrayPtr[0])
					for i, b := range dataBuffer {
						// Write each byte at offset+8+i (data starts at offset 8)
						byteAddr := byteArrayPtr + MemoryBlockHeaderSize + uint32(i)
						if ok := wa.Memory().Write(byteAddr, []byte{b}); !ok {
							return 0, cln, fmt.Errorf("failed to write byte at address %d", byteAddr)
						}
					}
					// Update offset to point to the byte array
					offset = byteArrayPtr
					// Return early since the array is properly allocated and initialized
					// Let Array wrapper creation happen
					return offset, cln, nil
					// return offset, cln, nil
				}
			default:
				return 0, cln, fmt.Errorf("unsupported type for ptr2*_array conversion: %s", elemType.Name())
			}

			if err != nil {
				return 0, cln, fmt.Errorf("failed to convert to array using ptr2*_array for type %s: %w", elemType.Name(), err)
			}

			// Update offset to point to the proper array (if conversion was used)
			if len(arrayPtr) > 0 && arrayPtr[0] != 0 {
				offset = uint32(arrayPtr[0])
				// The array is now properly GC-managed, return early to skip manual headers
				// Let Array wrapper creation happen
				return offset, cln, nil
				// return offset, cln, nil
			}
		} else {
			// Fallback to manual allocation approach
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
		wa.Memory().WriteUint32Le(offset, firstElem)                       // first element
		wa.Memory().WriteUint32Le(offset+4, header)                        // header
		wa.Memory().WriteUint32Le(offset+MemoryBlockHeaderSize, firstElem) // duplicate element

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
		wa.Memory().WriteUint32Le(slicePtr, offset-MemoryBlockHeaderSize)
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
			// Let Array wrapper creation happen
			return offset, cln, nil
			// return offset, cln, nil
		}
	}

	return offset, cln, nil
}

// createDynamicPrimitiveArray creates dynamic Array[T] types for primitive types using MoonBit's native functions
// This handles the two-level structure: wrapper object + data array
func (h *primitiveSliceHandler[T]) createDynamicPrimitiveArray(ctx context.Context, wa wasmMemoryWriter, slice []T, numElements uint32, elemType langsupport.TypeInfo) (uint32, utils.Cleaner, error) {
	// Convert wasmMemoryWriter to WasmAdapter for function calls
	wasmAdapter, ok := wa.(*wasmAdapter)
	if !ok {
		return 0, nil, fmt.Errorf("expected *wasmAdapter, got %T", wa)
	}
	// Step 1: Create the data array using appropriate MoonBit function
	var offset uint32
	var err error

	switch elemType.Name() {
	case "Bool":
		// Array[Bool] → moonbit.i32_array_make
		offset, err = h.createBoolDataArray(ctx, wa, wasmAdapter, slice, numElements)
	case "String":
		// Array[String] → moonbit.ref_array_make
		offset, err = h.createStringDataArray(ctx, wa, wasmAdapter, slice, numElements)
	case "Int":
		// Array[Int] → moonbit.i32_array_make
		offset, err = h.createIntDataArray(ctx, wa, wasmAdapter, slice, numElements)
	case "Byte":
		// Array[Byte] → moonbit.i32_array_make (bytes are stored as i32)
		offset, err = h.createByteDataArray(ctx, wa, wasmAdapter, slice, numElements)
	case "Char":
		// Array[Char] → moonbit.i32_array_make (chars are stored as i32)
		offset, err = h.createCharDataArray(ctx, wa, wasmAdapter, slice, numElements)
	case "Int16":
		// Array[Int16] → moonbit.int16_array_make
		offset, err = h.createInt16DataArray(ctx, wa, wasmAdapter, slice, numElements)
	case "UInt16":
		// Array[UInt16] → moonbit.int16_array_make
		offset, err = h.createUInt16DataArray(ctx, wa, wasmAdapter, slice, numElements)
	case "Int64":
		// Array[Int64] → moonbit.int64_array_make
		offset, err = h.createInt64DataArray(ctx, wa, wasmAdapter, slice, numElements)
	case "UInt64":
		// Array[UInt64] → moonbit.int64_array_make
		offset, err = h.createUInt64DataArray(ctx, wa, wasmAdapter, slice, numElements)
	case "Float":
		// Array[Float] → moonbit.float32_array_make
		offset, err = h.createFloatDataArray(ctx, wa, wasmAdapter, slice, numElements)
	case "Double":
		// Array[Double] → moonbit.float_array_make
		offset, err = h.createDoubleDataArray(ctx, wa, wasmAdapter, slice, numElements)
	default:
		return 0, nil, fmt.Errorf("unsupported dynamic primitive array element type: %s", elemType.Name())
	}

	if err != nil {
		return 0, nil, fmt.Errorf("failed to create data array for %s: %w", elemType.Name(), err)
	}

	// Step 2: Create wrapper structure using cabi_realloc (MoonBit allocator)
	// Array[T] wrapper structure: [refCount, 1573120, length, dataPtr]
	// Array[Byte] doesn't need wrapper creation because fnBytes2Array already returns a complete Array[Byte] object
	if elemType.Name() == "Byte" {
		// fnBytes2Array already returns a complete Array[Byte] object, no wrapper needed
		// fnBytes2Array returned complete Array[Byte] object
		return offset, utils.NewCleaner(), nil
	}

	// Array[Bool] doesn't need wrapper creation because moonbit_i32_array_make returns a complete Array[Bool]
	if elemType.Name() == "Bool" {
		// moonbit_i32_array_make already returns a complete Array[Bool] object, no wrapper needed
		return offset, utils.NewCleaner(), nil
	}

	// Other types return data array directly
	return offset, utils.NewCleaner(), nil
}

// Helper functions for creating data arrays for different primitive types

func (h *primitiveSliceHandler[T]) createBoolDataArray(ctx context.Context, wa wasmMemoryWriter, wasmAdapter *wasmAdapter, slice []T, numElements uint32) (uint32, error) {
	// Use moonbit_i32_array_make - this should create the correct Bool array type info
	if wasmAdapter.fnMakeArrayInt == nil {
		return 0, fmt.Errorf("function moonbit_i32_array_make not found")
	}

	// Create Array[Bool] using i32 array function (correct for Bool type info)
	results, err := wasmAdapter.fnMakeArrayInt.Call(ctx, uint64(numElements), uint64(0))
	if err != nil {
		return 0, fmt.Errorf("failed to call moonbit_i32_array_make: %w", err)
	}
	if len(results) != 1 {
		return 0, fmt.Errorf("expected 1 result from moonbit_i32_array_make, got %d", len(results))
	}

	arrayPtr := uint32(results[0])

	// Update memory directly where true values should be (back to 32-bit approach)
	for i, val := range slice {
		if boolVal, ok := any(val).(bool); ok {
			offset := arrayPtr + 8 + uint32(i)*4 // 8 = header size, 4 bytes per bool
			var value uint32
			if boolVal {
				value = 1
			}
			wa.Memory().WriteUint32Le(offset, value)
		}
	}

	// Return the complete Array[Bool] pointer (no wrapper creation needed)
	return arrayPtr, nil
}

func (h *primitiveSliceHandler[T]) createStringDataArray(ctx context.Context, wa wasmMemoryWriter, wasmAdapter *wasmAdapter, slice []T, numElements uint32) (uint32, error) {
	return 0, fmt.Errorf("String arrays should not be handled by primitiveSliceHandler")
}

func (h *primitiveSliceHandler[T]) createIntDataArray(ctx context.Context, wa wasmMemoryWriter, wasmAdapter *wasmAdapter, slice []T, numElements uint32) (uint32, error) {
	fn := wasmAdapter.GetFunction("moonbit_bytes_make")
	if fn == nil {
		return 0, fmt.Errorf("function moonbit_bytes_make not found")
	}

	// Create array with initial value 0
	results, err := fn.Call(ctx, uint64(numElements), uint64(0))
	if err != nil {
		return 0, fmt.Errorf("failed to call moonbit_bytes_make: %w", err)
	}
	if len(results) != 1 {
		return 0, fmt.Errorf("expected 1 result from moonbit_bytes_make, got %d", len(results))
	}

	arrayPtr := uint32(results[0])

	// Write int values
	for i, val := range slice {
		if intVal, ok := any(val).(int32); ok {
			offset := arrayPtr + MemoryBlockHeaderSize + uint32(i)
			wa.Memory().WriteUint32Le(offset, uint32(intVal))
		}
	}

	return arrayPtr, nil
}

func (h *primitiveSliceHandler[T]) createByteDataArray(ctx context.Context, wa wasmMemoryWriter, wasmAdapter *wasmAdapter, slice []T, numElements uint32) (uint32, error) {
	// Step 1: Create Bytes object using moonbit_bytes_make
	fn := wasmAdapter.GetFunction("moonbit_bytes_make")
	if fn == nil {
		return 0, fmt.Errorf("function moonbit_bytes_make not found")
	}

	results, err := fn.Call(ctx, uint64(numElements), uint64(0))
	if err != nil {
		return 0, fmt.Errorf("failed to call moonbit_bytes_make: %w", err)
	}
	if len(results) != 1 {
		return 0, fmt.Errorf("expected 1 result from moonbit_bytes_make, got %d", len(results))
	}

	bytesPtr := uint32(results[0])

	// Step 2: Write byte data to Bytes object
	// Writing bytes to Bytes object
	for i, val := range slice {
		if byteVal, ok := any(val).(byte); ok {
			offset := bytesPtr + MemoryBlockHeaderSize + uint32(i)
			if !wa.Memory().Write(offset, []byte{byteVal}) {
				return 0, fmt.Errorf("failed to write byte at offset %d", offset)
			}
		}
	}

	arrayResults, err := wasmAdapter.fnBytes2Array.Call(ctx, uint64(bytesPtr))
	if err != nil {
		return 0, fmt.Errorf("failed to call fnBytes2Array: %w", err)
	}
	if len(arrayResults) != 1 {
		return 0, fmt.Errorf("expected 1 result from fnBytes2Array, got %d", len(arrayResults))
	}

	arrayPtr := uint32(arrayResults[0])

	return arrayPtr, nil
}

func (h *primitiveSliceHandler[T]) createCharDataArray(ctx context.Context, wa wasmMemoryWriter, wasmAdapter *wasmAdapter, slice []T, numElements uint32) (uint32, error) {
	fn := wasmAdapter.GetFunction("moonbit_bytes_make")
	if fn == nil {
		return 0, fmt.Errorf("function moonbit_bytes_make not found")
	}

	// Create array with initial value 0
	results, err := fn.Call(ctx, uint64(numElements), uint64(0))
	if err != nil {
		return 0, fmt.Errorf("failed to call moonbit_bytes_make: %w", err)
	}
	if len(results) != 1 {
		return 0, fmt.Errorf("expected 1 result from moonbit_bytes_make, got %d", len(results))
	}

	arrayPtr := uint32(results[0])

	// Write char values as uint32
	for i, val := range slice {
		if charVal, ok := any(val).(int16); ok {
			offset := arrayPtr + MemoryBlockHeaderSize + uint32(i)
			wa.Memory().WriteUint32Le(offset, uint32(charVal))
		}
	}

	return arrayPtr, nil
}

func (h *primitiveSliceHandler[T]) createInt16DataArray(ctx context.Context, wa wasmMemoryWriter, wasmAdapter *wasmAdapter, slice []T, numElements uint32) (uint32, error) {
	fn := wasmAdapter.GetFunction("moonbit_int16_array_make")
	if fn == nil {
		return 0, fmt.Errorf("function moonbit_int16_array_make not found")
	}

	// Create array with initial value 0
	results, err := fn.Call(ctx, uint64(numElements), uint64(0))
	if err != nil {
		return 0, fmt.Errorf("failed to call moonbit_int16_array_make: %w", err)
	}
	if len(results) != 1 {
		return 0, fmt.Errorf("expected 1 result from moonbit_int16_array_make, got %d", len(results))
	}

	arrayPtr := uint32(results[0])

	// Write int16 values
	for i, val := range slice {
		if int16Val, ok := any(val).(int16); ok {
			offset := arrayPtr + MemoryBlockHeaderSize + uint32(i)*2 // int16 = 2 bytes
			wa.Memory().WriteUint16Le(offset, uint16(int16Val))
		}
	}

	return arrayPtr, nil
}

func (h *primitiveSliceHandler[T]) createUInt16DataArray(ctx context.Context, wa wasmMemoryWriter, wasmAdapter *wasmAdapter, slice []T, numElements uint32) (uint32, error) {
	fn := wasmAdapter.GetFunction("moonbit_int16_array_make")
	if fn == nil {
		return 0, fmt.Errorf("function moonbit_int16_array_make not found")
	}

	// Create array with initial value 0
	results, err := fn.Call(ctx, uint64(numElements), uint64(0))
	if err != nil {
		return 0, fmt.Errorf("failed to call moonbit_int16_array_make: %w", err)
	}
	if len(results) != 1 {
		return 0, fmt.Errorf("expected 1 result from moonbit_int16_array_make, got %d", len(results))
	}

	arrayPtr := uint32(results[0])

	// Write uint16 values
	for i, val := range slice {
		if uint16Val, ok := any(val).(uint16); ok {
			offset := arrayPtr + MemoryBlockHeaderSize + uint32(i)*2 // uint16 = 2 bytes
			wa.Memory().WriteUint16Le(offset, uint16Val)
		}
	}

	return arrayPtr, nil
}

func (h *primitiveSliceHandler[T]) createInt64DataArray(ctx context.Context, wa wasmMemoryWriter, wasmAdapter *wasmAdapter, slice []T, numElements uint32) (uint32, error) {
	fn := wasmAdapter.GetFunction("moonbit_int64_array_make")
	if fn == nil {
		return 0, fmt.Errorf("function moonbit_int64_array_make not found")
	}

	// Create array with initial value 0
	results, err := fn.Call(ctx, uint64(numElements), uint64(0))
	if err != nil {
		return 0, fmt.Errorf("failed to call moonbit_int64_array_make: %w", err)
	}
	if len(results) != 1 {
		return 0, fmt.Errorf("expected 1 result from moonbit_int64_array_make, got %d", len(results))
	}

	arrayPtr := uint32(results[0])

	// Write int64 values
	for i, val := range slice {
		if int64Val, ok := any(val).(int64); ok {
			offset := arrayPtr + MemoryBlockHeaderSize + uint32(i)*8 // int64 = 8 bytes
			wa.Memory().WriteUint64Le(offset, uint64(int64Val))
		}
	}

	return arrayPtr, nil
}

func (h *primitiveSliceHandler[T]) createUInt64DataArray(ctx context.Context, wa wasmMemoryWriter, wasmAdapter *wasmAdapter, slice []T, numElements uint32) (uint32, error) {
	fn := wasmAdapter.GetFunction("moonbit_int64_array_make")
	if fn == nil {
		return 0, fmt.Errorf("function moonbit_int64_array_make not found")
	}

	// Create array with initial value 0
	results, err := fn.Call(ctx, uint64(numElements), uint64(0))
	if err != nil {
		return 0, fmt.Errorf("failed to call moonbit_int64_array_make: %w", err)
	}
	if len(results) != 1 {
		return 0, fmt.Errorf("expected 1 result from moonbit_int64_array_make, got %d", len(results))
	}

	arrayPtr := uint32(results[0])

	// Write uint64 values
	for i, val := range slice {
		if uint64Val, ok := any(val).(uint64); ok {
			offset := arrayPtr + MemoryBlockHeaderSize + uint32(i)*8 // uint64 = 8 bytes
			wa.Memory().WriteUint64Le(offset, uint64Val)
		}
	}

	return arrayPtr, nil
}

func (h *primitiveSliceHandler[T]) createFloatDataArray(ctx context.Context, wa wasmMemoryWriter, wasmAdapter *wasmAdapter, slice []T, numElements uint32) (uint32, error) {
	fn := wasmAdapter.GetFunction("moonbit_float32_array_make")
	if fn == nil {
		return 0, fmt.Errorf("function moonbit_float32_array_make not found")
	}

	// Create array with initial value 0.0
	results, err := fn.Call(ctx, uint64(numElements), math.Float64bits(0.0))
	if err != nil {
		return 0, fmt.Errorf("failed to call moonbit_float32_array_make: %w", err)
	}
	if len(results) != 1 {
		return 0, fmt.Errorf("expected 1 result from moonbit_float32_array_make, got %d", len(results))
	}

	arrayPtr := uint32(results[0])

	// Write float32 values
	for i, val := range slice {
		if float32Val, ok := any(val).(float32); ok {
			offset := arrayPtr + MemoryBlockHeaderSize + uint32(i)*4 // float32 = 4 bytes
			wa.Memory().WriteFloat32Le(offset, float32Val)
		}
	}

	return arrayPtr, nil
}

func (h *primitiveSliceHandler[T]) createDoubleDataArray(ctx context.Context, wa wasmMemoryWriter, wasmAdapter *wasmAdapter, slice []T, numElements uint32) (uint32, error) {
	fn := wasmAdapter.GetFunction("moonbit_float_array_make")
	if fn == nil {
		return 0, fmt.Errorf("function moonbit_float_array_make not found")
	}

	// Create array with initial value 0.0
	results, err := fn.Call(ctx, uint64(numElements), math.Float64bits(0.0))
	if err != nil {
		return 0, fmt.Errorf("failed to call moonbit_float_array_make: %w", err)
	}
	if len(results) != 1 {
		return 0, fmt.Errorf("expected 1 result from moonbit_float_array_make, got %d", len(results))
	}

	arrayPtr := uint32(results[0])

	// Write float64 values
	for i, val := range slice {
		if float64Val, ok := any(val).(float64); ok {
			offset := arrayPtr + MemoryBlockHeaderSize + uint32(i)*8 // float64 = 8 bytes
			wa.Memory().WriteFloat64Le(offset, float64Val)
		}
	}

	return arrayPtr, nil
}

// decodeDynamicPrimitiveArray reads dynamic Array[T] types created by moonbit.i32_array_make
// This is used as a fallback when memoryBlockAtOffset fails with "invalid memory offset"
func (h *primitiveSliceHandler[T]) decodeDynamicPrimitiveArray(ctx context.Context, wa wasmMemoryReader, offset uint32) (any, error) {
	// For dynamic arrays created by moonbit.i32_array_make, read structure directly
	// Structure: [length(4), classInfo(4), element0(4), element1(4), ...]

	// Read the array length at offset 0
	lengthBytes, ok := wa.Memory().Read(offset, 4)
	if !ok {
		return nil, fmt.Errorf("failed to read array length at offset %d", offset)
	}
	numElements := binary.LittleEndian.Uint32(lengthBytes)

	// Dynamic array decode logic

	if numElements == 0 {
		return []T{}, nil // empty array
	}

	elemType := h.typeInfo.ListElementType()
	elemTypeSize := h.converter.TypeSize()
	if elemType.Name() == "Bool" || elemType.Name() == "Char" || elemType.Name() == "Byte" {
		// A MoonBit Bool is 4 bytes whereas a Go bool is 1 byte.
		// A MoonBit Array[Char] uses 4 bytes per element instead of 2.
		elemTypeSize = MoonBitBoolSize
	}

	// Read the data elements starting at offset 8 (after length and classInfo)
	dataStartOffset := uint32(8)
	dataSize := numElements * uint32(elemTypeSize)
	dataBytes, ok := wa.Memory().Read(offset+dataStartOffset, dataSize)
	if !ok {
		return nil, fmt.Errorf("failed to read dynamic array data at offset %d, size %d", offset+dataStartOffset, dataSize)
	}

	// Create the result slice
	items := make([]T, numElements)

	// Read each element directly from the data
	for i := uint32(0); i < numElements; i++ {
		var item T
		offset := i * uint32(elemTypeSize) // Data starts at beginning of dataBytes

		switch elemType.Name() {
		case "Bool":
			value := binary.LittleEndian.Uint32(dataBytes[offset:])
			boolValue := value != 0
			item = any(boolValue).(T)
		case "Int":
			value := binary.LittleEndian.Uint32(dataBytes[offset:])
			intValue := int32(value)
			item = any(intValue).(T)
		case "Byte":
			value := binary.LittleEndian.Uint32(dataBytes[offset:]) // MoonBit stores bytes as 4-byte values
			byteValue := byte(value)
			item = any(byteValue).(T)
		case "Char":
			value := binary.LittleEndian.Uint32(dataBytes[offset:]) // MoonBit stores chars as 4-byte values
			charValue := int16(value)
			item = any(charValue).(T)
		case "Int16":
			value := binary.LittleEndian.Uint16(dataBytes[offset:])
			int16Value := int16(value)
			item = any(int16Value).(T)
		case "UInt16":
			value := binary.LittleEndian.Uint16(dataBytes[offset:])
			uint16Value := uint16(value)
			item = any(uint16Value).(T)
		case "Int64":
			value := binary.LittleEndian.Uint64(dataBytes[offset:])
			int64Value := int64(value)
			item = any(int64Value).(T)
		case "UInt64":
			value := binary.LittleEndian.Uint64(dataBytes[offset:])
			uint64Value := uint64(value)
			item = any(uint64Value).(T)
		case "Float":
			value := binary.LittleEndian.Uint32(dataBytes[offset:])
			floatValue := math.Float32frombits(value)
			item = any(floatValue).(T)
		case "Double":
			value := binary.LittleEndian.Uint64(dataBytes[offset:])
			doubleValue := math.Float64frombits(value)
			item = any(doubleValue).(T)
		default:
			return nil, fmt.Errorf("unsupported dynamic primitive array element type: %s", elemType.Name())
		}

		items[i] = item
	}

	return items, nil
}
