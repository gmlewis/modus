## Detailed Analysis: MoonBit Array[Int16] Memory Representation in WASM

After thorough analysis of the WAT file, MoonBit source code, and all Array[Int16] operations, I can provide a comprehensive explanation of how MoonBit represents `Array[Int16]` in WASM linear memory.

### Key Discovery: String-Based Implementation with Sentinel Option Encoding

MoonBit `Array[Int16]` employs a **unique hybrid strategy** that differs from all other array types analyzed:

1. **Array[Int16]**: **Reuses String infrastructure** with 16-bit signed integer interpretation
2. **Array[Int16?]**: **32-bit integer storage** with sentinel-based None encoding

### Memory Layout

#### Outer Array Wrapper Structure (Same as other Arrays)
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (constant 1573120 = 0x180100)
0x08    4     Array length (number of elements)
0x0C    4     Pointer to inner FixedArray
```

#### Inner Array Structures - TWO DISTINCT VARIANTS

**For Array[Int16] (non-optional):**
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (kind=1, elem_size_shift=1, length)
0x08    N*2   Int16 data (N elements * 2 bytes each, signed 16-bit)
0x08+N*2 Pad  Zero padding to align to 2-byte boundaries
```

**For Array[Int16?] (optional):**
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (kind=1, elem_size_shift=2, length)
0x08    N*4   Extended integer data (N elements * 4 bytes each)
```

### Option Encoding Strategy (for Array[Int16?])

**Sentinel-Based Encoding:**
- **Some(int16)**: Direct storage as 32-bit signed integer
- **None**: `32768` (0x8000) - Just outside Int16 positive range

**Why This Works:**
- Int16 signed range: -32,768 to 32,767
- None value: 32,768 (one beyond maximum positive value)
- Efficient single-value encoding without separate objects

### Detailed Implementation Analysis

#### Array[Int16] Implementation

**Creation Function:**
- Uses `moonbit.int16_array_make` → `moonbit.unsafe_make_string`
- **Memory allocation**: `((length + 1) & ~1) * 2` bytes (2-byte aligned)
- **Header calculation**: `(1 << 30) | (1 << 28) | length`
- **elem_size_shift=1**: Indicates 2 bytes per element (2^1 = 2)

**Element Storage:**
```wasm
local.get $array_ptr
local.get $index
i32.const 1
i32.shl                       ; index * 2 (2 bytes per element)
i32.add
local.get $value
i32.store16 offset=8 align=1  ; Store 16-bit value at offset 8+(index*2)
```

**Element Access:**
- Uses `moonbit.int16_array_item_s`
- Access: `i32.load16_s offset=8+(index*2)` (sign-extended to 32-bit)

#### Array[Int16?] Implementation

**Creation Function:**
- Uses `moonbit.i32_array_make(size, -1)` then overwrites with `32768` for None
- **Memory allocation**: `size * 4` bytes (standard 32-bit integer array)
- **Header calculation**: `(1 << 30) | (2 << 28) | length`
- **elem_size_shift=2**: Indicates 4 bytes per element

**Element Storage:**
```wasm
; For Some(value):
local.get $array_ptr
local.get $index
i32.const 4
i32.mul
i32.add
local.get $int16_value         ; Direct storage as 32-bit int
i32.store offset=8 align=1

; For None:
local.get $array_ptr
local.get $index
i32.const 4
i32.mul
i32.add
i32.const 32768               ; Store None sentinel
i32.store offset=8 align=1
```

**Element Access:**
- Uses `moonbit.array_item` (same as other 32-bit integer arrays)
- Direct `i32.load offset=8+(index*4)`

### Memory Examples

#### Array[Int16] Examples

**Empty Array `[]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 0, ptr_to_empty]
Empty inner:   [stored at address 19648] (separate from other int arrays)
```

**Single value `[1]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x50000001, 0x0001, padding]
               ; 16-bit value 1, padded to 2-byte alignment
```

**Min/Max values `[Int16::min_value, Int16::max_value]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 2, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x50000002, 0x8000, 0x7FFF]
               ; -32768, 32767 (no padding needed, exactly 4 bytes)
```

**Four values `[1, 2, 3, 4]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 4, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x50000004, 0x0001, 0x0002, 0x0003, 0x0004]
               ; Four 16-bit values, total 8 bytes (4-byte aligned)
```

#### Array[Int16?] Examples

**Single None `[None]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000001, 32768, 0, 0, 0]
               ; None = 32768 stored as 32-bit integer
```

**Single Some `[Some(42)]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000001, 42, 0, 0, 0]
               ; 42 stored as 32-bit integer
```

**Mixed extremes `[None, Some(Int16::min_value), Some(0), Some(Int16::max_value)]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 4, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000004,
                32768,        ; None
                -32768,       ; Some(Int16::min_value)
                0,            ; Some(0)
                32767]        ; Some(Int16::max_value)
```

### Performance and Memory Characteristics

#### Array[Int16] Advantages
1. **Memory efficient**: 2 bytes per element (half the size of 32-bit integers)
2. **String infrastructure reuse**: Leverages optimized 16-bit element handling
3. **Alignment optimization**: 2-byte boundary alignment for efficiency
4. **Direct WASM operations**: Native i32.store16/i32.load16_s operations

#### Array[Int16?] Trade-offs
1. **Memory overhead**: 4 bytes per element (2x non-optional storage)
2. **Simple sentinel**: No separate object allocation overhead
3. **Range restriction**: None value uses one possible integer value
4. **Consistent access**: Same patterns as other 32-bit integer option arrays

### Comparison with Other Array Types

| Array Type | Non-Optional Storage | Optional Storage | None Encoding |
|------------|---------------------|------------------|---------------|
| Array[Bool] | 4 bytes/element | 4 bytes/element | -1 |
| Array[Byte] | 1 byte/element | 4 bytes/element | -1 |
| Array[Char] | 4 bytes/element | 4 bytes/element | -1 |
| Array[Int] | 4 bytes/element | 8 bytes/element | 2^32 (64-bit) |
| Array[Int16] | **2 bytes/element** | **4 bytes/element** | **32768** |
| Array[Float] | 4 bytes/element | 4+12 bytes/element | Separate objects |
| Array[Double] | 8 bytes/element | 4+16 bytes/element | Separate objects |

### Type Header Analysis

| Array Type | Kind | Elem Size Shift | Element Size | Infrastructure Reuse |
|------------|------|-----------------|--------------|---------------------|
| Array[Int16] | 1 | 1 | 2 bytes | **String implementation** |
| Array[Int16?] | 1 | 2 | 4 bytes | Integer array implementation |

### Unique Design Aspects

#### String Infrastructure Reuse
- **Memory allocation**: Same pattern as String ((len+1)&~1)*2
- **Element access**: Same 16-bit load/store operations as String
- **Alignment**: 2-byte boundary alignment like String
- **Interpretation**: Signed integers instead of Unicode characters

#### Sentinel Value Selection
- **Strategic choice**: 32768 = 2^15 (just outside Int16 range)
- **Minimal overhead**: Single sentinel value vs separate Option objects
- **Type safety**: Clear distinction from valid Int16 values
- **Range efficiency**: Uses full 32-bit space for encoding

### Empty Array Specialization

- **Array[Int16]**: Address 19648 (distinct from other integer arrays)
- **Array[Int16?]**: Would share with other 32-bit integer option arrays
- **Separation rationale**: Different element sizes require different empty array structures

### Verification Across All Operations

This hybrid implementation strategy is consistent across:
- ✅ **Array creation**: Specialized functions for optimal storage
- ✅ **Element access**: Type-appropriate load/store operations
- ✅ **Memory allocation**: Efficient for both compact and optional variants
- ✅ **Option handling**: Clean sentinel-based None encoding
- ✅ **Type safety**: Runtime distinction via value ranges
- ✅ **Infrastructure reuse**: Leverages existing String optimization

### Key Design Insights

1. **Smart infrastructure reuse**: Repurposes String implementation for 16-bit integers
2. **Optimal memory usage**: 2-byte storage when possible, 4-byte when needed
3. **Sentinel strategy**: Elegant None encoding without object overhead
4. **Type-specific optimization**: Different approach than other integer types
5. **Range-aware design**: Uses value space efficiently for encoding

### Use Case Optimization

**Array[Int16] is optimal for:**
- Embedded systems and IoT applications
- Audio sample processing (16-bit audio)
- Graphics coordinates and dimensions
- Memory-constrained numerical computing
- Large datasets with moderate value ranges

**Array[Int16?] is optimal for:**
- Sparse datasets with missing 16-bit values
- Configuration arrays with optional parameters
- Sensor data with occasional missing readings
- Database-style nullable 16-bit columns

This implementation demonstrates MoonBit's sophisticated approach to memory optimization, showing how different data types can benefit from different underlying infrastructures while maintaining consistent high-level semantics and optimal performance characteristics for their specific use cases.
