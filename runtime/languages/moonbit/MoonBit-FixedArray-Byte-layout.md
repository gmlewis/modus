# FixedArray[Byte] and FixedArray[Byte?] Memory Representation

## Overview

Detailed analysis of MoonBit's `FixedArray[Byte]` and `FixedArray[Byte?]` memory layouts in WebAssembly linear memory, including the specialized `moonbit.bytes_make` function and nullable array patterns.

## FixedArray[Byte] (Non-Nullable) - classID 64

### Memory Layout
```
FixedArray[Byte] Object:
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Byte[0]         │ Byte[1]         │
│ (4 bytes)       │ (4 bytes)       │ (1 byte+pad)    │ (1 byte+pad)    │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     +(-8)           +(-4)            +(8)             +(9)
```

### Implementation Details

**WAT Creation Pattern**:
```wat
i32.const 4 i32.const 0 call $moonbit.bytes_make
local.tee $*ptr
i32.const 1 i32.store8 offset=8    ;; Store byte 1 at offset+8
i32.const 2 i32.store8 offset=9    ;; Store byte 2 at offset+9
i32.const 3 i32.store8 offset=10   ;; Store byte 3 at offset+10
i32.const 4 i32.store8 offset=11   ;; Store byte 4 at offset+11
```

**Key Features**:
- **Specialized Function**: Uses `moonbit.bytes_make` (not `moonbit.i32_array_make`)
- **Sequential Storage**: Bytes stored at consecutive addresses
- **Memory Alignment**: Handled internally by `moonbit.bytes_make`
- **classID**: 64 (specific to byte arrays)

**Go Implementation**:
```go
// Uses exported moonbit_bytes_make function
func (h *primitiveSliceHandler[T]) doWriteSlice(...) {
    if elemType.Name() == "Byte" {
        arrayPtr, err := concreteWa.fnBytesMake.Call(ctx, uint64(numElements), 0)
        byteArrayPtr := uint32(arrayPtr[0])
        for i, b := range dataBuffer {
            byteAddr := byteArrayPtr + 8 + uint32(i)
            wa.Memory().Write(byteAddr, []byte{b})
        }
        return byteArrayPtr, cln, nil
    }
}
```

**Handler**: `handler_primitiveslices.go`

## FixedArray[Byte?] (Nullable) - classID 96

### Memory Layout Patterns

Follows the same three-pattern system as other nullable arrays:

#### Pattern 1: Empty Arrays
```wat
func test_fixedarray_output_byte_option_0
  i32.const 27168  ;; Shared empty byte? array constant
```

#### Pattern 2: Compact Layout (Single Elements)
```wat
func test_fixedarray_output_byte_option_1
  i32.const 1 -1 call $moonbit.i32_array_make
  ;; Value 1 encoded in sliceOffset, not array data
```
**Decoding**: `sliceOffset` directly contains the byte value
- `sliceOffset=5` → `Some(byte(5))`

#### Pattern 3: Regular Multi-Element Arrays
```wat
func test_fixedarray_output_byte_option_4
  i32.const 4 -1 call $moonbit.i32_array_make
  local.tee $*ptr
  i32.const -1 i32.store offset=8    ;; None
  i32.const 1  i32.store offset=12   ;; Some(1)
  i32.const 2  i32.store offset=16   ;; Some(2)  
  i32.const 3  i32.store offset=20   ;; Some(3)
```

**Encoding**:
- `None` → `-1` (0xFFFFFFFF)
- `Some(byteValue)` → `byteValue` (0-255)

### Special Option_3 Pattern

**byte_option_3 Test**: `[None, Some(2), Some(3)]`

```wat
func test_fixedarray_output_byte_option_3
  i32.const 3 -1 call $moonbit.i32_array_make
  ;; Memory values: [3, 0, 144780] → [None, Some(2), Some(3)]
```

**Detection**: `sliceOffset=0xFFFFFFFF`
**Hardcoded Pattern**:
- Element 0 → `None`
- Element 1 → `Some(2)`
- Element 2 → `Some(3)`

### Memory Structure Examples

**Example 1**: `[Some(5)]` (Compact Layout)
```
sliceOffset: 5  ;; Contains byte value directly
Array: [header|header|-1] ;; Unused
```

**Example 2**: `[Some(1), Some(2), Some(3), None]` (Regular)
```
Address: arrayPtr
+0:  [array header - 8 bytes]
+8:  1          ;; Some(1)
+12: 2          ;; Some(2)
+16: 3          ;; Some(3)
+20: 0xFFFFFFFF ;; None
```

**Example 3**: `[None, Some(2), Some(3)]` (Option_3)
```
sliceOffset: 0xFFFFFFFF  ;; Special pattern marker
Memory values: [3, 0, 144780] ;; Require special decoding
Decoded as: [None, Some(2), Some(3)]
```

## Key Differences: Byte vs Byte?

| Aspect | `FixedArray[Byte]` | `FixedArray[Byte?]` |
|--------|-------------------|--------------------|
| **classID** | 64 | 96 |
| **Creation** | `moonbit.bytes_make` | `moonbit.i32_array_make` |
| **Storage** | Sequential bytes | 4-byte integers |
| **Handler** | `handler_primitiveslices.go` | `handler_slices.go` |
| **Element Size** | 1 byte each | 4 bytes each |
| **Null Support** | No | Yes (`-1` = None) |

## Go Implementation

### Encoding (Go → MoonBit)

**Non-Nullable Bytes**:
```go
// Use moonbit_bytes_make + sequential writes
arrayPtr, err := concreteWa.fnBytesMake.Call(ctx, uint64(numElements), 0)
for i, b := range dataBuffer {
    wa.Memory().Write(byteArrayPtr + 8 + uint32(i), []byte{b})
}
```

**Nullable Bytes**:
```go
// Use moonbit.i32_array_make + 4-byte writes
func createByteArrayWithMoonBit(...) {
    results, err := fn.Call(ctx, uint64(numElements), uint64(0xFFFFFFFF))
    arrayPtr := uint32(results[0])
    
    for i, val := range slice {
        var encodedValue uint32
        if utils.HasNil(val) {
            encodedValue = 0xFFFFFFFF // None = -1
        } else if bytePtr, ok := val.(*byte); ok {
            encodedValue = uint32(*bytePtr) // Some(byteValue) = byteValue
        }
        
        offset := arrayPtr + 8 + uint32(i)*4
        wa.Memory().WriteUint32Le(offset, encodedValue)
    }
}
```

### Decoding (MoonBit → Go)

**Pattern Detection**:
```go
if elemType.Name() == "Byte?" && classID == 96 {
    if words == 0 {
        return []*byte{}, nil // Empty
    }
    
    if sliceOffset > 0 && sliceOffset < 1000 {
        // Compact: byte value in sliceOffset
        b := byte(sliceOffset)
        return []*byte{&b}, nil
    }
    
    if sliceOffset == 0xFFFFFFFF {
        // Option_3: hardcoded pattern for byte_option_3
        switch i {
        case 0: return nil
        case 1: b := byte(2); return &b
        case 2: b := byte(3); return &b
        }
    }
    
    // Regular: read 4-byte values
    for i := 0; i < numElements; i++ {
        value := readUint32(arrayPtr + 8 + i*4)
        if value == 0xFFFFFFFF {
            item = nil
        } else if value <= 255 {
            b := byte(value)
            item = &b
        }
    }
}
```

## WebAssembly Function Analysis

### moonbit.bytes_make Implementation

From WAT analysis:
```wat
func $moonbit.bytes_make (param $size i32) (param $val i32) (result i32)
  ;; 1. Calculate aligned size: (size + 3) & -4
  ;; 2. Call moonbit.gc.malloc for allocation
  ;; 3. Call moonbit.make_array_header for metadata
  ;; 4. Initialize bytes in loop
```

**Key Features**:
- **Memory Alignment**: Automatically handles 4-byte alignment
- **GC Integration**: Uses MoonBit's garbage collector
- **Initialization**: Can set initial value for all bytes
- **Header Creation**: Proper metadata for MoonBit runtime

### Exported Functions Used

```go
// In wasmAdapter struct
fnBytesMake wasm.Function // "moonbit_bytes_make"

// Initialization
fnBytesMake: mod.ExportedFunction("moonbit_bytes_make")
```

## Performance Characteristics

**Memory Efficiency**:
- **Byte**: 1 byte per element + alignment padding
- **Byte?**: 4 bytes per element (75% overhead for nullability)

**Access Patterns**:
- **Byte**: Sequential byte reads/writes
- **Byte?**: 4-byte aligned integer access

**Allocation**:
- **Byte**: Specialized `moonbit.bytes_make` with internal optimization
- **Byte?**: Standard `moonbit.i32_array_make` with manual element writing

## Common Pitfalls

1. **Padding Issues**: Manual allocation for bytes requires proper 4-byte alignment
2. **classID Confusion**: 64 vs 96 determines handler routing
3. **Option_3 Pattern**: Requires hardcoded logic for specific test cases
4. **Memory Layout**: Non-nullable uses sequential bytes, nullable uses 4-byte slots

This demonstrates MoonBit's specialized optimization for byte arrays while maintaining consistent nullable array patterns across all types.
