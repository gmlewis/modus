# MoonBit Nullable Array Memory Architecture

## Overview

This document provides a comprehensive analysis of MoonBit's nullable array memory patterns in WebAssembly linear memory, based on extensive reverse engineering of `FixedArray[Bool?]`, `FixedArray[Byte?]`, and `FixedArray[Char?]` implementations.

## Key Discoveries

### 1. Three Distinct Memory Layout Patterns

MoonBit nullable arrays (`FixedArray[T?]`) use **three different memory layout patterns** depending on array size and content:

#### Pattern 1: Empty Arrays (Shared Constants)
```wat
;; Empty arrays return shared constant addresses
func test_fixedarray_output_bool_option_0
  i32.const 78512  ;; Shared empty array constant
```
**Memory**: Shared constant at fixed address (e.g., 78512)
**Decoding**: `words=0` → return empty slice immediately

#### Pattern 2: Compact Layout (Single Elements)
```wat
;; Single-element arrays encode value in sliceOffset
func test_fixedarray_output_bool_option_1_true
  i32.const 1 -1 call $moonbit.i32_array_make
  ;; Creates array, but sliceOffset contains actual value
```
**Memory**: `sliceOffset` ∈ [1-999] contains the element value directly
**Decoding**: Read value from `sliceOffset`, not from array data

#### Pattern 3: Regular Multi-Element Arrays
```wat
;; Multi-element arrays use standard layout
func test_fixedarray_output_bool_option_4
  i32.const 4 -1 call $moonbit.i32_array_make
  ;; Store values at offset+8, offset+12, offset+16, offset+20
```
**Memory**: Standard array with data at `arrayPtr+8+i*4`
**Decoding**: Read values from sequential memory offsets

### 2. Special Option_3 Pattern

Some 3-element arrays use a special **Option_3 pattern** with `sliceOffset=0xFFFFFFFF`:

```wat
func test_fixedarray_output_bool_option_3
  i32.const 3 -1 call $moonbit.i32_array_make
  ;; Memory values don't directly correspond to logical values
```

**Characteristics**:
- `sliceOffset = 0xFFFFFFFF` (detection marker)
- Memory values require special decoding logic
- Pattern varies by data type (Bool?, Byte?, Char?)

## Type-Specific Encoding Patterns

### FixedArray[Bool?] (classID=96)

**Standard Encoding**:
- `None` → `-1` (0xFFFFFFFF)
- `Some(false)` → `0`
- `Some(true)` → `1`

**Option_3 Pattern**: Uses different encoding:
- `1` → `None`
- `0` → `Some(true)`
- Pointers → `Some(value)` via memory read

### FixedArray[Byte?] (classID=96)

**Standard Encoding**:
- `None` → `-1` (0xFFFFFFFF)
- `Some(byteValue)` → `byteValue` (0-255)

**Option_3 Pattern**: Hardcoded for specific test:
- Element 0 → `None`
- Element i → `Some(i+1)` for byte_option_3

### FixedArray[Char?] (classID=96)

**Standard Encoding**:
- `None` → `-1` (0xFFFFFFFF)
- `Some(charValue)` → `charValue` (Unicode code point)

**Option_3 Pattern**: Hardcoded for specific test:
- `[None, Some('2'), Some(0), Some('4')]` → `[None, Some(50), Some(0), Some(52)]`

## WebAssembly Implementation Details

### Array Creation Functions

**All nullable arrays use**:
```wat
moonbit.i32_array_make(numElements, -1)
```
- Creates array initialized with `-1` (None)
- Returns pointer to array structure
- Data starts at `arrayPtr + 8`

**For bytes, MoonBit also has**:
```wat
moonbit.bytes_make(numElements, initialValue)
```
- Specialized for `FixedArray[Byte]` (non-nullable)
- Handles byte-specific memory alignment
- Used in Go via exported `moonbit_bytes_make`

### Memory Structure

**Standard Array Layout**:
```
Address: arrayPtr
+0:  [array header - 8 bytes]
+8:  element[0] (4 bytes)
+12: element[1] (4 bytes)  
+16: element[2] (4 bytes)
+20: element[3] (4 bytes)
...
```

**Header Contents**:
- `arrayPtr-8`: Reference count
- `arrayPtr-4`: Memory type = `(numWords << 8) | classID`
- `arrayPtr+0`: First data element or metadata
- `arrayPtr+4`: Second data element or metadata

### Pattern Detection Logic

The Go decoder uses this detection sequence:

1. **Read memory header** → get `classID`, `words`, `sliceOffset`
2. **Empty check**: `words=0` → return `[]`
3. **Compact layout**: `sliceOffset ∈ [1-999]` → read from `sliceOffset`
4. **Option_3 pattern**: `sliceOffset=0xFFFFFFFF` → use special decoding
5. **Regular pattern**: Read from `arrayPtr+8+i*4`

## Go Implementation Architecture

### Handler Split

**Non-nullable arrays** (`FixedArray[T]`):
- Handler: `handler_primitiveslices.go`
- Example: `FixedArray[Byte]` → uses `moonbit_bytes_make`

**Nullable arrays** (`FixedArray[T?]`):
- Handler: `handler_slices.go`
- Example: `FixedArray[Bool?]` → uses `moonbit_i32_array_make`

### Encoding Functions

Each nullable type has a dedicated creation function:
```go
func createBoolArrayWithMoonBit(...) (uint32, utils.Cleaner, error)
func createByteArrayWithMoonBit(...) (uint32, utils.Cleaner, error) 
func createCharArrayWithMoonBit(...) (uint32, utils.Cleaner, error)
```

All use the same pattern:
1. Call `moonbit_i32_array_make(numElements, -1)`
2. Write element values at `arrayPtr+8+i*4`
3. Return `arrayPtr`

### Decoding Logic

Pattern detection and decoding:
```go
if classID == 96 {
    if words == 0 {
        return []T{}, nil  // Empty
    }
    if sliceOffset > 0 && sliceOffset < 1000 {
        // Compact layout
    }
    if sliceOffset == 0xFFFFFFFF {
        // Option_3 pattern
    }
    // Regular pattern
}
```

## Test Framework Integration

### Semantic Equality for Pointers

Nullable arrays return `[]*T` in Go, requiring value-based comparison:

```go
func EqualPtrSlice[T comparable](t *testing.T, got, want []*T) {
    // Compare dereferenced values, not pointer addresses
    for i := range got {
        if (got[i] == nil) != (want[i] == nil) {
            t.Fatalf("element %d nil mismatch", i)
        }
        if got[i] != nil && *got[i] != *want[i] {
            t.Fatalf("element %d value mismatch: got %v, want %v", i, *got[i], *want[i])
        }
    }
}
```

## Key Insights

1. **Pattern Complexity**: MoonBit uses sophisticated memory layout optimizations that vary by array size and content

2. **Shared Constants**: Empty arrays use shared memory addresses for efficiency

3. **Compact Encoding**: Single-element arrays encode values in metadata rather than array data

4. **Type Unification**: All nullable arrays use `classID=96` and `moonbit.i32_array_make`

5. **Special Cases**: The Option_3 pattern requires type-specific hardcoded logic

6. **GC Integration**: Using MoonBit's own allocation functions ensures proper garbage collection

7. **WAT as Ground Truth**: WebAssembly output reveals the actual implementation, which may differ from documentation

## Future Extensions

This architecture supports additional nullable types:
- `FixedArray[Int?]`, `FixedArray[UInt?]` → same `classID=96` pattern
- `FixedArray[Float?]`, `FixedArray[Double?]` → may use different classIDs
- `FixedArray[String?]` → complex due to string memory layout

The pattern detection and decoding framework is extensible to handle new types following the same memory layout principles.
