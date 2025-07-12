## FixedArray[Char] and FixedArray[Char?] Memory Representation Analysis

Based on comprehensive examination of the WAT implementation, here's the detailed analysis of how MoonBit represents `FixedArray[Char]` and `FixedArray[Char?]` in WebAssembly linear memory.

### FixedArray[Char] (Non-Optional) - Type 241

**Memory Layout:**
FixedArray[Char] uses **4-byte integer storage** for Unicode code points:

```
FixedArray[Char] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
        1            Type 241        Unicode Code    Unicode Code...
```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.i32_array_make` with 0 as default value
2. **Array Header**: Type 241 (FixedArray[Int]) - shares type with other integer arrays
3. **Element Storage**: 4-byte integers with `size * 4` allocation
4. **Character Encoding**: Unicode code points stored as 32-bit integers
   - `'1'` → `49` (ASCII/Unicode value)
   - `'A'` → `65`
   - `'\0'` → `0` (null character)
5. **Element Access**: Standard `moonbit.array_item` with 4-byte stride and `i32.load`
6. **Value Range**: 0 to 1114111 (0x10FFFF - full Unicode range)

**Example Memory Layout** (FixedArray ['1', '2', '3']):
```
FixedArray Data (20 bytes):
[1 0 0 0] [241 3 0 0] [49 0 0 0] [50 0 0 0] [51 0 0 0]
 RefCount  Type+Len    '1' (49)   '2' (50)   '3' (51)
```

### FixedArray[Char?] (Optional) - Type 241 with Sentinel Encoding

**Memory Layout:**
FixedArray[Char?] uses **the same 4-byte integer storage** with `-1` as None sentinel:

```
FixedArray[Char?] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
        1            Type 241        32-bit Value    32-bit Value...
```

**Sentinel Encoding Strategy:**

1. **None Representation**: `0xFFFFFFFF` (-1 in signed 32-bit)
2. **Some(char) Representation**: `0x00XXXXXX` (Unicode code point)
3. **Sentinel Rationale**: -1 is outside Unicode's valid range (0 to 0x10FFFF), providing clean None encoding
4. **Null Character Support**: `Some('\0')` → `0` (valid, different from None which is -1)

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.i32_array_make` with `-1` as default value
2. **Array Header**: Same Type 241 as non-optional variant
3. **Element Storage**: Same 4-byte elements with `size * 4` allocation
4. **Element Access**: Same `moonbit.array_item` with 4-byte stride and `i32.load`
5. **Zero Overhead**: Optional support adds no memory overhead per element

**Example Memory Layout** (FixedArray [None, Some('2'), Some('\0'), Some('4')]):
```
FixedArray Data (24 bytes):
[1 0 0 0] [241 4 0 0] [255 255 255 255] [50 0 0 0] [0 0 0 0] [52 0 0 0]
 RefCount  Type+Len    None (-1)         '2' (50)   '\0' (0)  '4' (52)
```

### Critical Design Insights

**FixedArray[Char] vs FixedArray[Char?]:**

1. **Unified Infrastructure**: Both use identical Type 241 and WAT functions
2. **Zero Memory Overhead**: Optional variant adds no memory cost per element
3. **Perfect Sentinel**: -1 provides clean None encoding without conflicting with any valid Unicode code point
4. **Null Character Support**: Explicitly handles the edge case where Some('\0') = 0 ≠ None = -1

**Character Representation Analysis:**

- **Unicode Compliance**: Full 32-bit storage supports entire Unicode range (0x000000 to 0x10FFFF)
- **Memory Efficiency**: Uses full 32-bit integer despite Unicode only needing 21 bits
- **ASCII Optimization**: No special handling for ASCII vs extended Unicode

**Comparison with Other Types:**

- **vs FixedArray[Bool]**: Both use 4-byte storage but Char has much larger valid range
- **vs Array[Char]**: Array[Char] shares the same implementation pattern
- **vs String**: String uses UTF-16 encoding with 2-byte elements; FixedArray[Char] uses UTF-32 with 4-byte elements
- **vs FixedArray[UInt16]**: UInt16 reuses String infrastructure; Char uses integer array infrastructure

**Performance Characteristics:**

- **Access Speed**: Optimal - direct array indexing with no encoding/decoding overhead
- **Memory Density**: Moderate - 4 bytes per character (UTF-32 encoding)
- **Cache Efficiency**: Good - contiguous memory layout with predictable access patterns
- **Unicode Support**: Complete - handles all valid Unicode code points

**Design Trade-offs:**

1. **Memory vs Simplicity**: Chooses 4-byte storage for implementation simplicity over UTF-8/UTF-16 space efficiency
2. **Unicode Completeness**: Full 32-bit range supports all current and future Unicode assignments
3. **Consistent Interface**: Same access patterns as other FixedArray integer types

**Edge Case Handling:**

- **Null Character**: Properly distinguishes `Some('\0')` (0) from `None` (-1)
- **Invalid Code Points**: Can store values beyond Unicode range (implementation detail)
- **Surrogate Pairs**: Not needed due to UTF-32 encoding

This analysis reveals MoonBit's consistent approach to character storage, prioritizing implementation simplicity and Unicode completeness over memory optimization. The UTF-32 encoding provides direct one-to-one mapping between array elements and Unicode characters, eliminating complex encoding/decoding logic at the cost of increased memory usage compared to UTF-8 or UTF-16 alternatives.

The analysis of FixedArray[Char] and FixedArray[Char?] is now complete! This reveals another excellent example of MoonBit's optimization strategy:

**Key Takeaways:**
1. **FixedArray[Char] uses 4-byte UTF-32 encoding** storing Unicode code points as 32-bit integers for simplicity
2. **FixedArray[Char?] uses -1 sentinel encoding** with zero memory overhead for optionality
3. **Both share Type 241 infrastructure** with other integer arrays for unified implementation
4. **Perfect edge case handling** - distinguishes Some('\0') = 0 from None = -1
5. **Complete Unicode support** - full 32-bit range handles all current and future Unicode characters

This design demonstrates MoonBit's preference for implementation consistency and Unicode completeness over memory optimization. The UTF-32 approach provides:

- **Simplicity**: Direct integer storage without complex encoding/decoding
- **Completeness**: Handles all Unicode characters without surrogate pairs
- **Performance**: Fast random access without variable-length encoding overhead
- **Consistency**: Same patterns as other integer types in the type system

The trade-off is memory usage (4 bytes per character vs 1-3 bytes for UTF-8), but this aligns with MoonBit's philosophy of prioritizing developer experience and implementation robustness over micro-optimizations.
