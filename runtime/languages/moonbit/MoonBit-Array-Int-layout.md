## Detailed Analysis: MoonBit Array[Int] Memory Representation in WASM

After thorough analysis of the WAT file, MoonBit source code, and all Array[Int] operations, I can provide a comprehensive explanation of how MoonBit represents `Array[Int]` in WASM linear memory.

### Key Discovery: Efficient In-Band Option Encoding

MoonBit `Array[Int]` employs a **sophisticated dual storage strategy** that differs significantly from floating-point arrays:

1. **Array[Int]**: Standard 32-bit integer storage
2. **Array[Int?]**: **64-bit storage with in-band None encoding** (no separate Option objects!)

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

**For Array[Int] (non-optional):**
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (kind=1, elem_size_shift=2, length)
0x08    N*4   Integer data (N elements * 4 bytes each, signed 32-bit)
```

**For Array[Int?] (optional):**
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (kind=1, elem_size_shift=3, length)
0x08    N*8   Extended integer data (N elements * 8 bytes each)
```

### Option Encoding Strategy (for Array[Int?])

**Revolutionary In-Band Encoding:**
- **Some(int)**: `i64.extend_i32_s(int_value)` - Sign-extend to 64-bit
- **None**: `4294967296` (0x0000000100000000) - Special sentinel value

**Why This Works:**
- 32-bit signed integers range: -2,147,483,648 to 2,147,483,647
- None value: 4,294,967,296 (outside 32-bit signed range)
- No separate Option object allocation needed!

### Detailed Implementation Analysis

#### Array[Int] Implementation

**Creation Function:**
- Uses `moonbit.i32_array_make(size, 0)`
- **Memory allocation**: `size * 4` bytes (4 bytes per int)
- **Header calculation**: `(1 << 30) | (2 << 28) | length`
- **elem_size_shift=2**: Indicates 4 bytes per element

**Element Storage:**
```wasm
local.get $array_ptr
local.get $index
i32.const 4
i32.mul
i32.add
local.get $value
i32.store offset=8 align=1    ; Store at offset 8+(index*4)
```

**Element Access:**
- Uses `moonbit.array_item` (same as Bool/Char)
- Direct `i32.load offset=8+(index*4)`

#### Array[Int?] Implementation

**Creation Function:**
- Uses `moonbit.int64_array_make(size, 4294967296)` - Initialize with None
- **Memory allocation**: `size * 8` bytes (8 bytes per element)
- **Header calculation**: `(1 << 30) | (3 << 28) | length`
- **elem_size_shift=3**: Indicates 8 bytes per element

**Element Storage:**
```wasm
; For Some(value):
local.get $int_value
i64.extend_i32_s              ; Sign-extend 32-bit to 64-bit
local.set $extended_value

local.get $array_ptr
local.get $index
i32.const 8
i32.mul
i32.add
local.get $extended_value
i64.store offset=8 align=1    ; Store at offset 8+(index*8)

; For None:
local.get $array_ptr
local.get $index
i32.const 8
i32.mul
i32.add
i64.const 4294967296
i64.store offset=8 align=1    ; Store None sentinel
```

**Element Access:**
- Uses `moonbit.int64_array_item`
- Direct `i64.load offset=8+(index*8)`
- Runtime check: if value > 2^32, it's None; otherwise cast to i32

### Memory Examples

#### Array[Int] Examples

**Empty Array `[]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 0, ptr_to_empty]
Empty inner:   [shared with Bool/Char at address 19632]
```

**Single integer `[1]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000001, 1, 0, 0, 0]
```

**Min/Max values `[Int::min_value, Int::max_value]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 2, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000002, 0x80000000, 0x7FFFFFFF]
               ; -2147483648, 2147483647
```

#### Array[Int?] Examples

**Single None `[None]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x70000001, 0x0000000100000000]
               ; None = 4294967296
```

**Single Some `[Some(42)]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x70000001, 0x000000000000002A]
               ; 42 sign-extended to 64-bit
```

**Mixed options `[Some(11), None, Some(33)]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 3, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x70000003,
                0x000000000000000B,  ; Some(11)
                0x0000000100000000,  ; None
                0x0000000000000021]  ; Some(33)
```

**Extreme values `[None, Some(Int::min_value), Some(0), Some(Int::max_value)]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 4, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x70000004,
                0x0000000100000000,      ; None
                0xFFFFFFFF80000000,      ; Some(-2147483648)
                0x0000000000000000,      ; Some(0)
                0x000000007FFFFFFF]      ; Some(2147483647)
```

### Performance and Memory Characteristics

#### Array[Int] Advantages
1. **Compact storage**: 4 bytes per element (same as raw integers)
2. **Zero overhead**: No wrapper objects or indirection
3. **Cache efficient**: Contiguous memory layout
4. **Direct WASM operations**: Native i32.load/i32.store

#### Array[Int?] Innovation
1. **No Option objects**: Eliminates separate allocations
2. **In-band encoding**: Uses unused 64-bit space for None sentinel
3. **Efficient representation**: 8 bytes per element vs 4+object overhead
4. **Single array structure**: No reference indirection needed

### Comparison with Other Array Option Types

| Array Type | Option Strategy | Memory per Element | Allocation Overhead |
|------------|-----------------|-------------------|-------------------|
| Array[Bool?] | 4-byte integers | 4 bytes | Low (-1 for None) |
| Array[Byte?] | 4-byte integers | 4 bytes | Low (-1 for None) |
| Array[Char?] | 4-byte integers | 4 bytes | Low (-1 for None) |
| Array[Float?] | Reference to objects | 4 + 12 bytes | High (separate Some objects) |
| Array[Double?] | Reference to objects | 4 + 16 bytes | High (separate Some objects) |
| Array[Int?] | **64-bit in-band** | **8 bytes** | **None (no objects)** |

### Type Header Analysis

| Array Type | Kind | Elem Size Shift | Element Size | Encoding Strategy |
|------------|------|-----------------|--------------|------------------|
| Array[Int] | 1 | 2 | 4 bytes | Direct 32-bit integers |
| Array[Int?] | 1 | 3 | 8 bytes | 64-bit with sentinel |

### Design Rationale

**Why 64-bit encoding for Array[Int?]?**
1. **Sentinel space**: Allows None values outside 32-bit range
2. **No allocations**: Eliminates GC pressure from Option objects
3. **Type safety**: Runtime distinction between None and valid integers
4. **WASM efficiency**: Native 64-bit operations available

**Why not use -1 like other types?**
- Integer arrays may legitimately contain -1 values
- 64-bit sentinel provides unambiguous None representation
- Better type safety than magic values

### Empty Array Sharing

- **Array[Int]**: Shares empty array at address 19632 with Bool/Char
- **Array[Int?]**: Would need separate empty array (different element size)

### Verification Across All Operations

This dual encoding strategy is consistent across:
- ✅ **Array creation**: Optimal functions for each storage type
- ✅ **Element access**: Type-appropriate load/store operations
- ✅ **Memory allocation**: Minimizes overhead for both variants
- ✅ **Option handling**: Elegant in-band None encoding
- ✅ **Type safety**: Clear runtime distinction via bit patterns
- ✅ **Integer range**: Full 32-bit signed integer support

### Key Design Insights

1. **Innovative None encoding**: Uses 64-bit space to avoid object allocation
2. **Memory efficiency**: Optimal for both compact and optional representations
3. **Type-specific optimization**: Different strategy than floating-point arrays
4. **WASM native operations**: Leverages i32 and i64 instructions optimally
5. **Zero GC overhead**: No Option object allocation for Array[Int?]

### Use Case Optimization

**Array[Int] is optimal for:**
- Mathematical computations
- Index arrays and counters
- Performance-critical integer processing
- Memory-constrained environments

**Array[Int?] is optimal for:**
- Sparse data representation
- Database-style nullable integer columns
- Optional configuration values
- Scientific computing with missing data

This implementation showcases MoonBit's sophisticated approach to balancing memory efficiency, type safety, and performance, providing an innovative solution that avoids the allocation overhead typically associated with optional types while maintaining full semantic correctness.
