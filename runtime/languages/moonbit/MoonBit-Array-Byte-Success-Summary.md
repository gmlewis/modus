# Array[Byte] Success Summary: Key Insights for Future Sessions

## Problem Overview

**Issue**: All Array[Byte] unit tests were failing (`TestArrayOutput_byte_[0-4]`) while other primitive types like Array[Int], Array[Int16], Array[Float] were working correctly.

**Root Cause**: Array[Byte] was being forced to use fixed array infrastructure instead of MoonBit's native array creation system.

## Solution Summary

**Three-step approach using MoonBit's native functions**:
1. Create Bytes object using `moonbit_bytes_make`
2. Write byte data as 1-byte values to Bytes object
3. Convert Bytes to Array[Byte] using `fnBytes2Array` (`moonbit_bytes_to_array`)

**Result**: 100% success - all Array[Byte] tests now pass with no regression in other types.

## Critical Technical Discoveries

### 1. Size-Mismatched Types Need Special Handling

**Key insight**: Types where Go and MoonBit have different size representations require specialized approaches.

**Size-mismatched types**:
- `Bool`: Go bool (1 byte) vs MoonBit Bool (4 bytes)
- `Byte`: Go uint8 (1 byte) vs MoonBit Byte (context-dependent)
- `Char`: Go rune (4 bytes) vs MoonBit Char (context-dependent)

**Size-matched types** (work with standard infrastructure):
- `Int16`, `Int64`, `Float`, `Double`: Direct size correspondence

### 2. MoonBit's Native Functions Are Always Better

**Failing approach**: Manual memory allocation and wrapper creation
```go
// This causes GC unreachable errors
wrapperPtr, err := wa.allocateAndPinMemory(ctx, 16, 0)
wa.Memory().WriteUint32Le(wrapperPtr+4, 1573120)
```

**Successful approach**: Use MoonBit's own functions
```go
// This works perfectly
fn := wasmAdapter.GetFunction("moonbit_bytes_make")
results, err := fn.Call(ctx, uint64(numElements), uint64(0))
```

**Available MoonBit functions** (from `adapter.go`):
- `moonbit_bytes_make` + `moonbit_bytes_to_array` (for Array[Byte])
- `moonbit_i32_array_make` (for Bool?, Byte?, Char?)
- `moonbit_ref_array_make` (for reference types)
- `moonbit_int16_array_make` (for Int16)
- `moonbit_int64_array_make` (for Int64, Int?)
- `moonbit_float_array_make` (for Float)
- `moonbit_float32_array_make` (for Double)

### 3. WAT Analysis is Essential

**WAT reveals the truth**: What MoonBit actually generates vs. what Go code assumes

**Example from Array[Byte] WAT analysis**:
```wat
(call $moonbit.bytes_make
  (i32.const 4)  ;; numElements
  (i32.const 0)) ;; default value

;; ... write byte data ...

(call $moonbit.bytes_to_array
  (local.get $bytes_ptr))
```

**This revealed**:
- MoonBit uses `bytes_make` + `bytes_to_array` conversion
- Not direct array creation
- Conversion function was already exported but not used

### 4. Memory Layout Complexity

**Array[Byte] created by `fnBytes2Array` has complex structure**:

```
Array[Byte] Wrapper (returned by fnBytes2Array):
├── [0x00] Reference count
├── [0x04] Type info (1573120 = Array type)
├── [0x08] Array length
└── [0x0C] Pointer to Bytes object
    └── Bytes Object:
        ├── [0x00] Reference count
        ├── [0x04] Type info (0x4000000X pattern)
        └── [0x08] Raw byte data (1 byte per element)
```

**Key decode insight**: Must detect Bytes objects by type info pattern `(typeInfo & 0xFF000000) == 0x40000000` and read from offset +8.

### 5. Infrastructure Categorization

**Don't force types into wrong infrastructure**:

**Old (wrong) approach**:
```go
// Force Array[Byte] to use fixed array infrastructure
if !isFixedArray && (elemType.Name() == "Bool" || elemType.Name() == "Byte") {
    isFixedArray = true
}
```

**Correct approach**:
```go
// Let Array[Byte] use dynamic array path with native functions
if !isFixedArray && elemType.Name() == "Bool" {
    isFixedArray = true  // Only Bool still needs this
}
```

## Implementation Code Patterns

### Encoding Pattern (Create Array[Byte])

```go
func (h *primitiveSliceHandler[T]) createByteDataArray(ctx context.Context, wa wasmMemoryWriter, wasmAdapter *wasmAdapter, slice []T, numElements uint32) (uint32, error) {
    // Step 1: Create Bytes object
    fn := wasmAdapter.GetFunction("moonbit_bytes_make")
    results, err := fn.Call(ctx, uint64(numElements), uint64(0))
    if err != nil {
        return 0, fmt.Errorf("failed to call moonbit_bytes_make: %w", err)
    }
    bytesPtr := uint32(results[0])

    // Step 2: Write byte data
    for i, val := range slice {
        if byteVal, ok := any(val).(byte); ok {
            offset := bytesPtr + 8 + uint32(i) // 8 = header size
            wa.Memory().Write(offset, []byte{byteVal})
        }
    }

    // Step 3: Convert to Array[Byte]
    arrayResults, err := wasmAdapter.fnBytes2Array.Call(ctx, uint64(bytesPtr))
    if err != nil {
        return 0, fmt.Errorf("failed to call fnBytes2Array: %w", err)
    }
    arrayPtr := uint32(arrayResults[0])
    
    return arrayPtr, nil
}
```

### Decoding Pattern (Read Array[Byte])

```go
// In Decode function - detect Array[Byte] with Bytes object
if typeInfo == 1573120 { // Array wrapper
    arrayLength = binary.LittleEndian.Uint32(lengthBytes)
    dataPointer := binary.LittleEndian.Uint32(dataPointerBytes)
    
    if elemType.Name() == "Byte" {
        // Check if pointing to Bytes object
        bytesTypeInfo := binary.LittleEndian.Uint32(bytesTypeBytes)
        if (bytesTypeInfo & 0xFF000000) == 0x40000000 {
            // Bytes object - read from data section
            offset = dataPointer + 8
            isBytesData = true
        }
    }
}

// Special handling for Bytes data
if isBytesData && elemType.Name() == "Byte" {
    dataBytes, ok := wa.Memory().Read(offset, arrayLength)
    if !ok {
        return nil, fmt.Errorf("failed to read byte data")
    }
    
    items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(arrayLength), int(arrayLength))
    for i := uint32(0); i < arrayLength; i++ {
        val := h.converter.Decode(uint64(dataBytes[i]))
        items.Index(int(i)).Set(reflect.ValueOf(val))
    }
    return items.Interface(), nil
}
```

## Debugging Methodology

### 1. Start with WAT Analysis

```bash
# Find the failing test function
grep -A 50 "test_array_output_byte_2" build/testdata.wat

# Look for MoonBit function calls
grep -E "moonbit\.[a-z_]+" build/testdata.wat | grep -i byte
```

### 2. Check Available Functions

```bash
# See what functions are exported
grep -n "moonbit_.*_make" adapter.go
grep -n "moonbit_.*_to_" adapter.go
```

### 3. Add Debug Output

```go
// Debug memory structure
if debugBytes, ok := wa.Memory().Read(offset, 32); ok {
    fmt.Printf("DEBUG: First 32 bytes: %v\n", debugBytes)
    typeInfo := binary.LittleEndian.Uint32(debugBytes[4:8])
    length := binary.LittleEndian.Uint32(debugBytes[8:12])
    dataPtr := binary.LittleEndian.Uint32(debugBytes[12:16])
    fmt.Printf("DEBUG: typeInfo=%d, length=%d, dataPtr=%d\n", typeInfo, length, dataPtr)
}
```

### 4. Test Incrementally

```bash
# Test one size at a time
go test -run TestArrayOutput_byte_0 -v
go test -run TestArrayOutput_byte_1 -v
# ... continue until all pass
```

## Key Lessons for Future Work

### 1. Check for Native Functions First

**Before implementing any manual memory allocation**, check if MoonBit provides specialized functions.

**Pattern**: `moonbit_[type]_make` often exists for type-specific creation.

### 2. Size-Mismatched Types Are Special

**Don't assume all primitive types work the same way**. Types with size mismatches between Go and MoonBit often need completely different approaches.

### 3. Use MoonBit's Type System

**Work with MoonBit's classifications, not against them**. If MoonBit treats a type specially (like Byte), follow that pattern.

### 4. WAT is Always Right

**When Go assumptions conflict with WAT output, WAT wins**. The WAT shows what MoonBit actually generates.

### 5. Test Coverage Matters

**Test all sizes (0-4) to catch edge cases**. Empty arrays, single elements, and larger arrays often have different behaviors.

## Success Metrics

**Before Fix**:
- 0/5 Array[Byte] tests passing
- "wasm error: unreachable" for sizes 2-4
- "expected [1 2], got [1 13]" type corruption

**After Fix**:
- 5/5 Array[Byte] tests passing
- No regressions in other types
- Clean, maintainable code
- Proper MoonBit integration

## Next Steps for Related Types

### Array[Bool] (Still Failing)

**Likely approach**: Similar to Array[Byte] but may need different functions.

**Check for**:
- `moonbit_bool_array_make` or similar
- Boolean-specific conversion functions
- WAT analysis for Bool array creation

### Array[Char] (May Need Similar Fix)

**Likely approach**: Check if Char also has size mismatch issues.

**Pattern**: May need `moonbit_char_array_make` + conversion function.

## File Locations

**Key files modified**:
- `/app/runtime/languages/moonbit/handler_primitiveslices.go` - Main implementation
- `/app/runtime/languages/moonbit/adapter.go` - Function exports (already had fnBytes2Array)

**Test files**:
- `/app/runtime/languages/moonbit/tests/arrays_byte_test.go` - Test definitions
- `/app/runtime/languages/moonbit/testdata/arrays_byte.mbt` - MoonBit test code

**Documentation updated**:
- `MoonBit-Array-Byte-layout.md` - Memory layout analysis
- `MoonBit-FixedArray-Debugging-Guide.md` - Debugging techniques
- `MoonBit-Go-Unit-Test-Debugging-Guide.md` - Systematic debugging approach

This success case demonstrates that even complex failing tests can be resolved by understanding and working with MoonBit's native type system rather than trying to force types into incompatible infrastructure.