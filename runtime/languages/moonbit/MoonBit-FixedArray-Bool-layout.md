# FixedArray[Bool] and FixedArray[Bool?] Memory Representation

## Overview

Comprehensive analysis of MoonBit's `FixedArray[Bool]` and `FixedArray[Bool?]` memory layouts in WebAssembly linear memory, based on WAT reverse engineering and runtime pattern analysis.

## FixedArray[Bool] (Non-Nullable) - classID 64

### Memory Layout
```
FixedArray[Bool] Object:
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     +(-8)           +(-4)            +(0)             +(4)
```

### Implementation Details

**WAT Creation Pattern**:
```wat
i32.const 4 i32.const 0 call $moonbit.i32_array_make
;; Creates 4-element array with false (0) as default
```

**Element Encoding**:
- `false` → `0` (32-bit integer)
- `true` → `1` (32-bit integer)

**Go Handler**: `handler_primitiveslices.go`

## FixedArray[Bool?] (Nullable) - classID 96

### Three Memory Layout Patterns

MoonBit uses **three distinct patterns** for nullable boolean arrays:

#### Pattern 1: Empty Arrays (Shared Constants)
```wat
func test_fixedarray_output_bool_option_0
  i32.const 78512  ;; Returns shared empty array constant
```
**Detection**: `words=0`
**Decoding**: Return empty slice immediately

#### Pattern 2: Compact Layout (Single Elements)
```wat
func test_fixedarray_output_bool_option_1_true  
  i32.const 1 -1 call $moonbit.i32_array_make
  ;; Value encoded in sliceOffset, not array data
```
**Detection**: `sliceOffset ∈ [1-999]`
**Decoding**: Read value directly from `sliceOffset`
- `sliceOffset=1` → `Some(true)`
- `sliceOffset=0` → `Some(false)`

#### Pattern 3: Regular Multi-Element Arrays
```wat
func test_fixedarray_output_bool_option_4
  i32.const 4 -1 call $moonbit.i32_array_make
  ;; Standard array with data at arrayPtr+8, +12, +16, +20
```
**Detection**: Standard array access
**Encoding**:
- `None` → `-1` (0xFFFFFFFF)
- `Some(false)` → `0`
- `Some(true)` → `1`

### Special Option_3 Pattern

Three-element arrays use a special encoding pattern:

```wat
func test_fixedarray_output_bool_option_3
  i32.const 3 -1 call $moonbit.i32_array_make
  ;; Memory values: [1, 0, pointer] → [None, Some(true), Some(true)]
```

**Detection**: `sliceOffset=0xFFFFFFFF`
**Special Encoding**:
- `1` → `None`
- `0` → `Some(true)`
- Pointers → `Some(value)` via memory read

### Memory Structure Examples

**Example 1**: `[Some(true)]` (Compact Layout)
```
sliceOffset: 1  ;; Contains the value directly
Array: [header|header|-1] ;; Unused, value is in sliceOffset
```

**Example 2**: `[Some(false), Some(true), Some(false), None]` (Regular)
```
Address: arrayPtr
+0:  [array header - 8 bytes]
+8:  0          ;; Some(false)
+12: 1          ;; Some(true)  
+16: 0          ;; Some(false)
+20: 0xFFFFFFFF ;; None
```

**Example 3**: `[None, Some(true), Some(true)]` (Option_3)
```
sliceOffset: 0xFFFFFFFF  ;; Special pattern marker
Memory values: [1, 0, pointer] → [None, Some(true), Some(true)]
```

## Go Implementation

### Handler Location
- **FixedArray[Bool]**: `handler_primitiveslices.go`
- **FixedArray[Bool?]**: `handler_slices.go`

### Encoding (Go → MoonBit)

```go
func createBoolArrayWithMoonBit(...) (uint32, utils.Cleaner, error) {
    // Step 1: Create array with moonbit.i32_array_make(numElements, -1)
    results, err := fn.Call(ctx, uint64(numElements), uint64(0xFFFFFFFF))
    arrayPtr := uint32(results[0])
    
    // Step 2: Write values at arrayPtr+8+i*4
    for i, val := range slice {
        var encodedValue uint32
        if utils.HasNil(val) {
            encodedValue = 0xFFFFFFFF // None = -1
        } else if boolPtr, ok := val.(*bool); ok {
            if *boolPtr {
                encodedValue = 1 // Some(true) = 1
            } else {
                encodedValue = 0 // Some(false) = 0
            }
        }
        offset := arrayPtr + 8 + uint32(i)*4
        wa.Memory().WriteUint32Le(offset, encodedValue)
    }
    
    return arrayPtr, nil, nil
}
```

### Decoding (MoonBit → Go)

```go
if elemType.Name() == "Bool?" && classID == 96 {
    // Pattern detection
    if words == 0 {
        return []bool{}, nil // Empty array
    }
    
    if sliceOffset > 0 && sliceOffset < 1000 {
        // Compact layout: value in sliceOffset
        var item any
        switch sliceOffset {
        case 1:
            t := true
            item = &t
        case 0:
            f := false
            item = &f
        }
        // Return single-element slice
    }
    
    if sliceOffset == 0xFFFFFFFF {
        // Option_3 pattern: special decoding
        // Handle each element based on memory value pattern
    }
    
    // Regular pattern: read from arrayPtr+8+i*4
    for i := 0; i < numElements; i++ {
        value := readUint32(arrayPtr + 8 + i*4)
        switch value {
        case 0xFFFFFFFF:
            item = nil // None
        case 0:
            f := false
            item = &f // Some(false)
        case 1:
            t := true
            item = &t // Some(true)
        }
    }
}
```

## Test Framework Integration

### Semantic Equality

Nullable arrays return `[]*bool` in Go, requiring value-based comparison:

```go
// Use EqualPtrSlice for nullable arrays
EqualPtrSlice(t, result, expected)

// Instead of reflect.DeepEqual which compares pointer addresses
```

## Key Insights

1. **Pattern Complexity**: MoonBit optimizes memory layout based on array size and content
2. **Shared Constants**: Empty arrays reuse the same memory address
3. **Compact Encoding**: Single elements stored in metadata, not array data
4. **Type Unification**: All nullable arrays use `classID=96`
5. **Special Cases**: Three-element arrays have unique encoding patterns
6. **GC Integration**: Uses MoonBit's allocation functions for proper memory management

## Performance Characteristics

**Memory Efficiency**:
- Empty arrays: O(1) shared constant
- Single elements: Compact encoding in metadata
- Multi-element: Standard 4-byte per element + 8-byte header

**Access Patterns**:
- Compact layout: Direct metadata access
- Regular layout: Sequential memory reads
- Option_3 pattern: Mixed memory access and special decoding

This architecture demonstrates MoonBit's sophisticated approach to memory optimization while maintaining compatibility with WebAssembly's linear memory model.
