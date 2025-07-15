## Detailed Analysis: MoonBit Array[Double] Memory Representation in WASM

After thorough analysis of the WAT file, MoonBit source code, and all Array[Double] operations, I can provide a comprehensive explanation of how MoonBit represents `Array[Double]` in WASM linear memory.

### Key Discovery: Specialized Storage with Option Indirection

MoonBit `Array[Double]` employs **two completely different storage strategies** depending on optionality:

1. **Array[Double]**: Direct 64-bit IEEE 754 storage for maximum efficiency
2. **Array[Double?]**: Reference-based storage with separate Option objects

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

**For Array[Double] (non-optional):**
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (kind=1, elem_size_shift=3, length)
0x08    N*8   Double data (N elements * 8 bytes each, IEEE 754)
```

**For Array[Double?] (optional):**
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (kind=2, elem_size_shift=2, length)
0x08    N*4   Pointer data (N pointers * 4 bytes each)
```

### Option Object Structure (for Array[Double?])

**None Object:**
- Pre-allocated shared object at address 10248
- No individual allocation per None value

**Some(Double) Object:**
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (1)
0x04    4     Type info (constant 2097153 = 0x00200001)
0x08    8     Double value (IEEE 754 64-bit)
```

### Detailed Implementation Analysis

#### Array[Double] Implementation

**Creation Function:**
- Uses `moonbit.float_array_make(size, 0.0)`
- **Memory allocation**: `size * 8` bytes (8 bytes per double)
- **Header calculation**: `(1 << 30) | (3 << 28) | length`
- **elem_size_shift=3**: Indicates 8 bytes per element (2^3 = 8)

**Element Storage:**
```wasm
local.get $array_ptr
local.get $index
i32.const 8
i32.mul
i32.add
local.get $value
f64.store offset=8 align=1    ; Store at offset 8+(index*8)
```

**Element Access:**
- Uses `moonbit.float_array_item`
- Direct `f64.load offset=8+(index*8)`
- Returns raw double value

#### Array[Double?] Implementation

**Creation Function:**
- Uses `moonbit.ref_array_make(size, None_address)`
- **Memory allocation**: `size * 4` bytes (4 bytes per pointer)
- **Header calculation**: `(2 << 30) | (2 << 28) | length`
- **kind=2**: Reference array type
- **elem_size_shift=2**: Indicates 4 bytes per element (2^2 = 4)

**Element Storage:**
```wasm
; For Some(value):
; 1. Allocate 16-byte Some object
i32.const 16
call moonbit.gc.malloc
local.tee $some_ptr
i32.const 2097153
i32.store offset=4 align=1    ; Some object header
local.get $some_ptr
local.get $double_value
f64.store offset=8 align=1    ; Store double in Some object

; 2. Store pointer in array
local.get $array_ptr
local.get $index
i32.const 4
i32.mul
i32.add
local.get $some_ptr
i32.store offset=8 align=1    ; Store pointer at offset 8+(index*4)
```

**Element Access:**
- Uses `moonbit.array_item` (same as integer arrays)
- `i32.load offset=8+(index*4)` to get pointer
- Dereference pointer to access Option object
- Load double from Some object at offset 8

### Memory Examples

#### Array[Double] Examples

**Empty Array `[]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 0, ptr_to_empty]
Empty inner:   [stored at address 40480]
```

**Two doubles `[1.0, 2.0]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 2, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x70000002, 1.0_as_f64, 2.0_as_f64]
Binary layout: [1, 0, 0, 0, 0x70000002, 0x3FF0000000000000, 0x4000000000000000]
```

#### Array[Double?] Examples

**Single None `[None]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0xA0000001, 10248]
None object:   [pre-allocated shared object at 10248]
```

**Single Some `[Some(1.0)]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0xA0000001, ptr_to_some]
Some object:   [1, 0, 0, 0, 0x00200001, 1.0_as_f64]
```

**Mixed options `[Some(1.0), None, Some(3.0)]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 3, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0xA0000003, ptr1, 10248, ptr3]
Some object 1: [1, 0, 0, 0, 0x00200001, 1.0_as_f64]
Some object 3: [1, 0, 0, 0, 0x00200001, 3.0_as_f64]
```

### Performance and Memory Characteristics

#### Array[Double] Advantages
1. **Direct storage**: No pointer indirection for element access
2. **Memory efficient**: 8 bytes per element (minimum possible for IEEE 754)
3. **Cache friendly**: Contiguous memory layout
4. **WASM optimized**: Direct f64.load/f64.store operations

#### Array[Double?] Trade-offs
1. **Memory overhead**: 4 bytes per element + 16 bytes per Some object
2. **Allocation overhead**: Separate allocation for each Some value
3. **Indirection cost**: Two memory accesses per element (array → option → value)
4. **Reference counting**: Additional GC overhead for Option objects

### Type Header Decoding

| Array Type | Kind | Elem Size Shift | Element Size | Storage Type |
|------------|------|-----------------|--------------|--------------|
| Array[Double] | 1 | 3 | 8 bytes | Direct IEEE 754 |
| Array[Double?] | 2 | 2 | 4 bytes | Pointers to Option |

### Empty Array Specialization

Different array types have different pre-allocated empty arrays:
- **Array[Double]**: Address 40480 (FixedArray[Double])
- **Array[Bool]/Array[Char]**: Address 19632 (FixedArray[Int])
- **Array[Byte]**: Address 45368 (FixedArray[Byte])

This type-specific approach prevents runtime type confusion and enables specialized optimizations.

### Verification Across All Operations

This dual storage strategy is consistent across:
- ✅ **Array creation**: Different creation functions optimize for each case
- ✅ **Element access**: Specialized access functions for direct vs indirect storage
- ✅ **Memory allocation**: Optimal allocation strategies for each storage type
- ✅ **Option handling**: Proper None sharing and Some object management
- ✅ **Type safety**: Runtime type checking via different header encodings
- ✅ **IEEE 754 compliance**: Standard floating-point representation

### Comparison with Other Array Types

| Array Type | Non-Optional Storage | Optional Storage | Memory Efficiency |
|------------|---------------------|------------------|-------------------|
| Array[Bool] | 4 bytes/element | 4 bytes/element | Same |
| Array[Byte] | 1 byte/element | 4 bytes/element | Different |
| Array[Char] | 4 bytes/element | 4 bytes/element | Same |
| Array[Double] | 8 bytes/element | 4+16 bytes/element | Very Different |

### Key Design Insights

1. **Floating-point specialization**: Direct IEEE 754 storage when possible
2. **Option cost awareness**: Clear separation between efficient and flexible storage
3. **Memory layout optimization**: Type-specific empty arrays and headers
4. **WASM alignment**: Leverages native f64.load/f64.store operations
5. **GC integration**: Proper reference counting for complex Option structures

This sophisticated dual approach demonstrates MoonBit's commitment to performance optimization while maintaining full type safety and Option semantics, making `Array[Double]` highly efficient for numerical computing while still supporting optional values when needed.
