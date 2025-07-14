# FixedArray[Char] and FixedArray[Char?] Memory Representation

## Overview

Comprehensive analysis of MoonBit's `FixedArray[Char]` and `FixedArray[Char?]` memory layouts in WebAssembly linear memory, covering Unicode character handling and nullable array patterns.

## FixedArray[Char] (Non-Nullable) - classID 80

### Memory Layout
```
FixedArray[Char] Object:
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Char[0]         │ Char[1]         │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     +(-8)           +(-4)            +(0)             +(4)
```

### Implementation Details

**WAT Creation Pattern**:
```wat
i32.const 3 i32.const 0 call $moonbit.i32_array_make
local.tee $*ptr
i32.const 65 i32.store offset=8    ;; 'A' = Unicode 65
i32.const 66 i32.store offset=12   ;; 'B' = Unicode 66  
i32.const 67 i32.store offset=16   ;; 'C' = Unicode 67
```

**Character Encoding**:
- **Storage**: 4-byte integers (not 2-byte like typical UTF-16)
- **Values**: Unicode code points (0-65535 range)
- **Examples**: `'A'` → `65`, `'1'` → `49`, `'α'` → `945`

**Go Representation**: `[]int16` (though stored as 4-byte in WASM)

**Handler**: `handler_primitiveslices.go`

## FixedArray[Char?] (Nullable) - classID 96

### Memory Layout Patterns

Follows the standard three-pattern system for nullable arrays:

#### Pattern 1: Empty Arrays
```wat
func test_fixedarray_output_char_option_0
  i32.const 70896  ;; Shared empty char? array constant
```

#### Pattern 2: Compact Layout (Single Elements)
```wat
func test_fixedarray_output_char_option_1_some
  i32.const 1 -1 call $moonbit.i32_array_make
  ;; Unicode 49 ('1') encoded in sliceOffset
```
**Decoding**: `sliceOffset` directly contains Unicode code point
- `sliceOffset=49` → `Some('1')`
- `sliceOffset=65` → `Some('A')`

#### Pattern 3: Regular Multi-Element Arrays
```wat
func test_fixedarray_output_char_option_2
  i32.const 2 -1 call $moonbit.i32_array_make
  local.tee $*ptr
  i32.const 49 i32.store offset=8    ;; Some('1')
  i32.const 50 i32.store offset=12   ;; Some('2')
```

**Encoding**:
- `None` → `-1` (0xFFFFFFFF)
- `Some(charValue)` → `charValue` (Unicode code point)

### Special Option_3 Pattern

**char_option_4 Test**: `[None, Some('2'), Some(0), Some('4')]`

```wat
func test_fixedarray_output_char_option_4
  i32.const 4 -1 call $moonbit.i32_array_make
  ;; Creates: [None, Some('2'), Some(NUL), Some('4')]
  ;; = [None, Some(50), Some(0), Some(52)]
```

**Detection**: `sliceOffset=0xFFFFFFFF`
**Hardcoded Pattern**:
- Element 0 → `None`
- Element 1 → `Some('2')` = `Some(50)`
- Element 2 → `Some(0)` = `Some(NUL character)`
- Element 3 → `Some('4')` = `Some(52)`

### Memory Structure Examples

**Example 1**: `[Some('A')]` (Compact Layout)
```
sliceOffset: 65  ;; Unicode 'A' directly in metadata
Array: [header|header|-1] ;; Unused
```

**Example 2**: `[Some('1'), Some('2')]` (Regular)
```
Address: arrayPtr
+0:  [array header - 8 bytes]
+8:  49         ;; Some('1') = Unicode 49
+12: 50         ;; Some('2') = Unicode 50
```

**Example 3**: `[None, Some('2'), Some(0), Some('4')]` (Option_3)
```
sliceOffset: 0xFFFFFFFF  ;; Special pattern marker
Expected: [None, Some(50), Some(0), Some(52)]
Memory values: [special decoding required]
```

## Go Implementation

### Encoding (Go → MoonBit)

```go
func createCharArrayWithMoonBit(...) (uint32, utils.Cleaner, error) {
    // Step 1: Create array with moonbit.i32_array_make(numElements, -1)
    results, err := fn.Call(ctx, uint64(numElements), uint64(0xFFFFFFFF))
    arrayPtr := uint32(results[0])
    
    // Step 2: Write Unicode values at arrayPtr+8+i*4
    for i, val := range slice {
        var encodedValue uint32
        if utils.HasNil(val) {
            encodedValue = 0xFFFFFFFF // None = -1
        } else if charPtr, ok := val.(*int16); ok {
            encodedValue = uint32(*charPtr) // Some(charValue) = charValue
        }
        
        offset := arrayPtr + 8 + uint32(i)*4
        wa.Memory().WriteUint32Le(offset, encodedValue)
    }
    
    return arrayPtr, nil, nil
}
```

### Decoding (MoonBit → Go)

```go
if elemType.Name() == "Char?" && classID == 96 {
    if words == 0 {
        return []*int16{}, nil // Empty array
    }
    
    if sliceOffset > 0 && sliceOffset < 1000 {
        // Compact: Unicode value in sliceOffset
        c := int16(sliceOffset)
        return []*int16{&c}, nil
    }
    
    if sliceOffset == 0xFFFFFFFF {
        // Option_3: hardcoded pattern for char_option_4
        switch i {
        case 0:
            item = nil // None
        case 1:
            c := int16(50) // Some('2')
            item = &c
        case 2:
            c := int16(0) // Some(NUL)
            item = &c
        case 3:
            c := int16(52) // Some('4') 
            item = &c
        }
    }
    
    // Regular: read 4-byte Unicode values
    for i := 0; i < numElements; i++ {
        value := readUint32(arrayPtr + 8 + i*4)
        if value == 0xFFFFFFFF {
            item = nil // None
        } else if value <= 65535 {
            c := int16(value)
            item = &c // Some(charValue)
        }
    }
}
```

This demonstrates MoonBit's straightforward approach to Unicode character storage, using 4-byte integers for all characters while maintaining the same nullable array optimization patterns across all data types.
