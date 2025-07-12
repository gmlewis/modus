## Detailed Analysis: MoonBit Array[Byte] Memory Representation in WASM

After thorough analysis of the WAT file, MoonBit source code, and all Array[Byte] operations, I can provide a comprehensive explanation of how MoonBit represents `Array[Byte]` in WASM linear memory.

### Key Discovery: Optimized Byte Storage with Dual Encoding

MoonBit `Array[Byte]` uses the same **double-indirection structure** as `Array[Bool]`, but with **memory-optimized byte storage** for the non-optional case and **4-byte integer storage** for the optional case.

### Memory Layout

#### Outer Array Wrapper Structure (Same as Array[Bool])
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (constant 1573120 = 0x180100)
0x08    4     Array length (number of elements)
0x0C    4     Pointer to inner FixedArray
```

#### Inner FixedArray Structure - TWO VARIANTS

**For Array[Byte] (non-optional):**
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (kind=1, elem_size_shift=0, length)
0x08    N     Byte data (N bytes, 4-byte aligned allocation)
```

**For Array[Byte?] (optional):**
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (kind=1, elem_size_shift=2, length)
0x08    N*4   Integer data (N elements * 4 bytes each)
```

### Detailed Field Analysis

#### Inner Array Type Headers

**Non-optional Array[Byte]:**
- **Header calculation**: `(1 << 30) | (0 << 28) | length`
- **elem_size_shift=0**: Indicates 1 byte per element (2^0 = 1)
- **Memory allocation**: `(length + 3) & ~3` bytes (4-byte aligned)
- **Element access**: `base + 8 + index` using `i32.load8_u`

**Optional Array[Byte?]:**
- **Header calculation**: `(1 << 30) | (2 << 28) | length`
- **elem_size_shift=2**: Indicates 4 bytes per element (2^2 = 4)
- **Memory allocation**: `length * 4` bytes
- **Element access**: `base + 8 + (index * 4)` using `i32.load`
- **Value encoding**: `Some(n)` = n, `None` = -1 (0xFFFFFFFF)

### Memory Examples from Analysis

#### Array[Byte] Examples

**Empty Array `[]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 0, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x40000000] (length=0)
```

**Single byte `[1]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x40000001, 1, pad, pad, pad] (4-byte aligned)
```

**Four bytes `[1, 2, 3, 4]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 4, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x40000004, 1, 2, 3, 4] (exactly 4 bytes)
```

#### Array[Byte?] Examples

**Single option `[Some(1)]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000001, 1, 0, 0, 0] (4 bytes per element)
```

**Mixed options `[Some(1), None]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 2, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000002, 1, 0, 0, 0, 255, 255, 255, 255]
```

### Element Access Patterns

#### Array[Byte] Access
Using `moonbit.bytes_item`:
```wasm
local.get $arr          ; inner array pointer
local.get $index        ; element index
i32.add                 ; arr + index (no multiplication!)
i32.load8_u offset=8    ; load 1 byte from offset 8
```

#### Array[Byte?] Access
Using `moonbit.array_item`:
```wasm
local.get $arr          ; inner array pointer
local.get $index        ; element index
i32.const 4             ; 4 bytes per element
i32.mul                 ; index * 4
i32.add                 ; arr + (index * 4)
i32.load offset=8       ; load 4 bytes from offset 8
```

### Key Implementation Functions

**Array[Byte] Creation:**
- Uses `moonbit.bytes_make(size, 0)`
- Stores elements with `i32.store8 offset=8+index`
- Compact 1-byte storage

**Array[Byte?] Creation:**
- Uses `moonbit.i32_array_make(size, -1)`
- Stores elements with `i32.store offset=8+(index*4)`
- 4-byte storage to accommodate -1 for `None`

### Memory Optimizations

#### Empty Array Sharing
- All empty `Array[Byte]` instances share the same pre-allocated inner array at address 45368
- Different from `Array[Bool]` empty array (address 19632) due to different type headers

#### 4-Byte Alignment
- Non-optional byte arrays are padded to 4-byte boundaries
- Memory allocation: `(size + 3) & ~3` ensures WASM alignment requirements

#### Type Discrimination
- `elem_size_shift` in header distinguishes between byte and integer storage
- Same `moonbit.array_length` function works for both variants
- Different access functions (`bytes_item` vs `array_item`) based on storage type

### Verification Across All Operations

This dual encoding is consistent across:
- ✅ **Array creation**: `bytes_make` vs `i32_array_make` based on optionality
- ✅ **Element access**: Different access functions for different storage types
- ✅ **Memory allocation**: Optimized for space (bytes) vs compatibility (integers)
- ✅ **Option handling**: -1 encoding requires 4-byte integers
- ✅ **Empty arrays**: Type-specific shared empty arrays
- ✅ **Length encoding**: Same 28-bit encoding for both variants

### Performance Implications

1. **Space efficiency**: `Array[Byte]` uses 1/4 the memory of `Array[Byte?]`
2. **Access speed**: Direct byte access vs integer access with range checking
3. **Alignment**: All allocations respect WASM 4-byte alignment requirements
4. **Type safety**: Runtime type discrimination prevents incorrect access patterns

This sophisticated dual encoding allows MoonBit to provide memory-efficient byte arrays while maintaining full Option type semantics when needed, demonstrating the compiler's optimization capabilities within WASM constraints.
