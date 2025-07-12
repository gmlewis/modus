## Array[UInt] and Array[UInt?] Memory Representation Analysis

Based on comprehensive examination of both MoonBit source code comments and WAT implementation, here's the detailed analysis of how MoonBit represents `Array[UInt]` and `Array[UInt?]` in WebAssembly linear memory.

### Array[UInt] (Non-Optional) - Type 241

**Memory Layout:**
Array[UInt] uses **direct 32-bit storage** with compact element packing:

```
Array[UInt] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 241        UInt Value      UInt Value...
```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.i32_array_make` with parameters:
   - Size: number of elements
   - Default value: 0 (for initialization)
2. **Array Header**: Type 241 (FixedArray[UInt]) - same type identifier as Array[Int]
3. **Element Storage**: Direct 4-byte unsigned integers, allocated with `size * 4` bytes
4. **Element Access**: `moonbit.array_item` with 4-byte stride and `i32.load`
5. **Value Range**: 0 to 4294967295 (0x00000000 to 0xFFFFFFFF)

**Example Memory Layout** (Array [1, 2, 3]):
```
Offset: 0x0000BFA0 (Array wrapper)
[1 0 0 0] [0 2 0 0] [C0 BE 0 0] [3 0 0 0]
 RefCount   Type 0    Inner Ptr   Length

Offset: 0x0000BEC0 (FixedArray data)
[1 0 0 0] [241 3 0 0] [1 0 0 0] [2 0 0 0] [3 0 0 0]
 RefCount  Type+Len    Element[0] Element[1] Element[2]
```

### Array[UInt?] (Optional) - Type 241 with In-Band Encoding

**Memory Layout:**
Array[UInt?] uses **in-band encoding** with 64-bit elements to accommodate sentinel values:

```
Array[UInt?] Object (16+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (8 bytes)       │ (8 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 241       64-bit Value     64-bit Value...
```

**In-Band Encoding Strategy:**

1. **None Representation**: `0x0000000100000000` (4294967296 = 2^32)
2. **Some(value) Representation**: `0x00000000XXXXXXXX` (value zero-extended to 64 bits)
3. **Sentinel Rationale**: 2^32 is outside UInt's valid range (0 to 2^32-1), providing clean None encoding

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.int64_array_make` instead of `i32_array_make`
2. **Array Header**: Same Type 241 as non-optional variant
3. **Element Storage**: 8-byte elements with `size * 8` allocation
4. **Element Access**: `moonbit.int64_array_item` with 8-byte stride and `i64.load`
5. **Value Encoding**: UInt values extended to i64 using `i64.extend_i32_s`

**Example Memory Layout** (Array [Some(11), None, Some(33)]):
```
Offset: 0x0000BF20 (Array wrapper)
[1 0 0 0] [0 2 0 0] [C0 BE 0 0] [3 0 0 0]
 RefCount   Type 0    Inner Ptr   Length

Offset: 0x0000BEC0 (FixedArray data)
[1 0 0 0] [241 3 0 0] [11 0 0 0 0 0 0 0] [0 0 0 0 1 0 0 0] [33 0 0 0 0 0 0 0]
 RefCount  Type+Len    Some(11)           None              Some(33)
```

### Critical Design Insights

**Array[UInt] vs Array[UInt?]:**

1. **Storage Efficiency Trade-off**:
   - Array[UInt]: 4 bytes per element
   - Array[UInt?]: 8 bytes per element (100% overhead for optionality)
2. **Unified Type System**: Both use Type 241, demonstrating MoonBit's type consolidation
3. **Sentinel Strategy**: Uses 2^32 sentinel, exploiting the gap between UInt32 max and Int64 space

**Comparison with Other Types:**

- **vs Array[Int?]**: Uses identical in-band encoding strategy with 2^32 sentinel
- **vs Array[String?]**: UInt? uses in-band encoding while String? uses null pointer encoding
- **vs Array[Int64?]**: UInt? uses efficient in-band encoding while Int64? requires separate Option objects

**Performance Characteristics:**

- **Array[UInt]**: Optimal memory density and cache efficiency
- **Array[UInt?]**: 2x memory overhead but maintains direct element access
- **Sentinel Detection**: Simple comparison with 2^32 constant for None checking
- **Value Extraction**: Direct truncation from i64 to i32 for Some values

**Encoding Details:**
```
UInt Value Range: 0x00000000 to 0xFFFFFFFF
None Sentinel:    0x0000000100000000 (2^32)
```

This analysis reveals MoonBit's sophisticated approach to optional unsigned integers. Unlike Array[Int64?] which requires expensive Option objects due to the lack of available sentinel values, Array[UInt?] leverages the fact that UInt only uses 32 bits of a 64-bit space, allowing the 33rd bit position to serve as a clean None sentinel. This demonstrates the compiler's ability to choose optimal representations based on the mathematical properties of each type.

The analysis of Array[UInt] and Array[UInt?] is now complete! This reveals another excellent example of MoonBit's optimization strategy:

**Key Takeaways:**
1. **Array[UInt] uses direct 32-bit storage** with Type 241 for maximum memory efficiency
2. **Array[UInt?] employs the same in-band encoding as Array[Int?]** using 2^32 as the None sentinel
3. **Both share Type 241** but use different underlying WAT functions (i32_array_make vs int64_array_make)
4. **Memory overhead for optionality is exactly 100%** (4→8 bytes per element) but avoids separate object allocations
5. **The 2^32 sentinel works perfectly** because UInt's range (0 to 2^32-1) leaves 2^32 unused in 64-bit space

This demonstrates MoonBit's consistent approach to in-band encoding for types where mathematical sentinel values are available, contrasting with the reference-based approaches needed for types like Int64 where no unused bit patterns exist.
