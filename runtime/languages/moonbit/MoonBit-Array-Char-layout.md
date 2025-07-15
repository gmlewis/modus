## Detailed Analysis: MoonBit Array[Char] Memory Representation in WASM

After thorough analysis of the WAT file, MoonBit source code, and all Array[Char] operations, I can provide a comprehensive explanation of how MoonBit represents `Array[Char]` in WASM linear memory.

### Key Discovery: Unified Integer Storage for All Char Operations

MoonBit `Array[Char]` uses the **same double-indirection structure** as `Array[Bool]`, but with **4-byte integer storage** for both optional and non-optional cases. This is different from `Array[Byte]` which optimizes storage for the non-optional case.

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

#### Inner FixedArray Structure (Unified for both Array[Char] and Array[Char?])
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (kind=1, elem_size_shift=2, length)
0x08    N*4   Character data (N elements * 4 bytes each)
```

### Character Encoding and Value Ranges

#### MoonBit Char Type
- **Value range**: -32768 to 32767 (16-bit signed integer range)
- **Storage**: 32-bit signed integers in arrays
- **Unicode coverage**: Most of Basic Multilingual Plane (U+0000 to U+7FFF)
- **Encoding**: Direct Unicode code point values

#### Value Examples
```
Character '1' (U+0031) → 49
Character '2' (U+0032) → 50
Character '3' (U+0033) → 51
Character '4' (U+0034) → 52
Minimum char value → -32768
Maximum char value → 32767
```

### Memory Examples from Analysis

#### Array[Char] Examples

**Empty Array `[]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 0, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000000] (shared with other Int arrays)
```

**Single character `['1']`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000001, 49, 0, 0, 0]
```

**Two characters `['1', '2']`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 2, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000002, 49, 0, 0, 0, 50, 0, 0, 0]
```

#### Array[Char?] Examples

**Single None `[None]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000001, 255, 255, 255, 255] (None = -1)
```

**Single Some `[Some('1')]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000001, 49, 0, 0, 0] (Same as non-optional)
```

**Mixed options `[Some('1'), None, Some('3')]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 3, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000003, 49, 0, 0, 0, 255, 255, 255, 255, 51, 0, 0, 0]
```

### Element Access Pattern

Both `Array[Char]` and `Array[Char?]` use the same access pattern via `moonbit.array_item`:

```wasm
local.get $arr          ; inner array pointer
local.get $index        ; element index
i32.const 4             ; 4 bytes per element
i32.mul                 ; index * 4
i32.add                 ; arr + (index * 4)
i32.load offset=8       ; load 4 bytes from offset 8
```

### Key Implementation Details

#### Array Creation
**Both Array[Char] and Array[Char?]:**
- Use `moonbit.i32_array_make(size, initial_value)`
- `Array[Char]`: initial_value = 0
- `Array[Char?]`: initial_value = -1
- Store characters with `i32.store offset=8+(index*4)`

#### Option Encoding
- **Some(char)**: Unicode code point value (-32768 to 32767)
- **None**: -1 (0xFFFFFFFF)
- No special encoding needed since -1 is outside the valid char range

#### Type Headers
- **Header calculation**: `(1 << 30) | (2 << 28) | length`
- **elem_size_shift=2**: Indicates 4 bytes per element
- **Same encoding as Array[Bool] and Array[Int]**

### Memory Optimizations and Design Decisions

#### Shared Empty Arrays
- Empty `Array[Char]` shares the same pre-allocated inner array as `Array[Bool]` (address 19632)
- Both use type 241 (FixedArray[Int]) with identical headers

#### Unified Storage Strategy
Unlike `Array[Byte]` which uses different storage for optional vs non-optional:
- **Array[Char]** and **Array[Char?]** use identical storage formats
- **Rationale**: Char values can be negative, requiring 32-bit storage anyway
- **Benefit**: Simplified implementation and consistent access patterns

#### Unicode Limitations
- **Coverage**: Basic Multilingual Plane (U+0000 to U+7FFF)
- **Missing**: Supplementary planes (U+8000+) and surrogate pairs
- **Trade-off**: Memory efficiency vs full Unicode support

### Verification Across All Operations

This unified integer storage is consistent across:
- ✅ **Array creation**: Same `i32_array_make` for both variants
- ✅ **Element access**: Identical `array_item` function for both
- ✅ **Option handling**: -1 encoding naturally fits in 32-bit signed integers
- ✅ **Character encoding**: Direct Unicode code point storage
- ✅ **Memory sharing**: Empty arrays shared with other integer array types
- ✅ **Type safety**: Runtime type checking via header information

### Comparison with Other Array Types

| Array Type | Inner Storage | Element Size | Option Encoding | Empty Array Sharing |
|------------|---------------|--------------|-----------------|-------------------|
| Array[Bool] | FixedArray[Int] | 4 bytes | 0/1 → -1 | Yes (19632) |
| Array[Byte] | FixedArray[Byte] | 1 byte | N/A | No (45368) |
| Array[Byte?] | FixedArray[Int] | 4 bytes | value → -1 | No |
| Array[Char] | FixedArray[Int] | 4 bytes | N/A | Yes (19632) |
| Array[Char?] | FixedArray[Int] | 4 bytes | value → -1 | Yes (19632) |

### Performance Implications

1. **Consistent access**: No conditional logic needed between Array[Char] and Array[Char?]
2. **Memory overhead**: 4 bytes per character (vs 1-2 bytes for pure UTF-8/UTF-16)
3. **Unicode support**: Limited to 16-bit signed range but covers most common use cases
4. **Type safety**: Clear distinction between characters and arbitrary integers
5. **Option efficiency**: Natural -1 encoding without additional wrapper structures

This design demonstrates MoonBit's pragmatic approach to Unicode handling, prioritizing implementation simplicity and performance over complete Unicode coverage while maintaining type safety and memory efficiency within reasonable bounds.
