## FixedArray[Int16] and FixedArray[Int16?] Memory Representation Analysis

Based on comprehensive examination of the WAT implementation, here's the detailed analysis of how MoonBit represents `FixedArray[Int16]` and `FixedArray[Int16?]` in WebAssembly linear memory.

### FixedArray[Int16] (Non-Optional) - Type 243 (String Infrastructure)

**Memory Layout:**
FixedArray[Int16] **reuses the String infrastructure** with 16-bit element storage:

```
FixedArray[Int16] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (2 bytes)       │ (2 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 243        Int16 Value     Int16 Value...
```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.int16_array_make` → `moonbit.unsafe_make_string`
2. **Array Header**: Type 243 (String) - same infrastructure as Array[UInt16] and String
3. **Element Storage**: 2-byte elements using `i32.store16`, with 2-byte alignment
4. **Memory Allocation**: `((size + 1) & -2) << 1` bytes (padded to even count, then doubled for 2-byte elements)
5. **Element Access**: `moonbit.int16_array_item_s` with 2-byte stride and `i32.load16_s` (signed load)
6. **Value Range**: -32768 to 32767 (standard 16-bit signed integer range)

**Example Memory Layout** (FixedArray [1, 2, 3]):
```
Offset: 0x0000BFA0 (Array wrapper - if needed)
[1 0 0 0] [0 2 0 0] [C0 BE 0 0] [3 0 0 0]
 RefCount   Type 0    Inner Ptr   Length

Offset: 0x0000BEC0 (String/FixedArray data)
[1 0 0 0] [243 3 0 0] [1 0] [2 0] [3 0] [xx xx]
 RefCount  Type+Len    El[0] El[1] El[2] Padding
```

### FixedArray[Int16?] (Optional) - Type 241 with Sentinel Encoding

**Memory Layout:**
FixedArray[Int16?] **switches to 4-byte integer storage** to accommodate sentinel values:

```
FixedArray[Int16?] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 241        32-bit Value    32-bit Value...
```

**Sentinel Encoding Strategy:**

1. **None Representation**: `32768` (2^15, just outside Int16's valid range)
2. **Some(value) Representation**: `value` stored directly as 32-bit integer
3. **Sentinel Rationale**: 32768 is one above Int16.max_value (32767), providing clean None encoding

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.i32_array_make` with `-1` as default value
2. **Array Header**: Type 241 (FixedArray[Int]) - shares type with other integer arrays
3. **Element Storage**: 4-byte elements with `size * 4` allocation
4. **Element Access**: Standard `moonbit.array_item` with 4-byte stride and `i32.load`
5. **Memory Overhead**: 2x storage overhead compared to non-optional variant

**Example Memory Layout** (FixedArray [Some(1), None, Some(-32768)]):
```
FixedArray Data (20 bytes):
[1 0 0 0] [241 3 0 0] [1 0 0 0] [0 128 0 0] [0 128 255 255]
 RefCount  Type+Len    Some(1)   None (32768) Some(-32768)
```

### Critical Design Insights

**FixedArray[Int16] vs FixedArray[Int16?]:**

1. **Infrastructure Reuse**: Non-optional reuses String (Type 243), optional uses integer array (Type 241)
2. **Memory Efficiency Trade-off**:
   - FixedArray[Int16]: 2 bytes per element (optimal density)
   - FixedArray[Int16?]: 4 bytes per element (100% overhead for optionality)
3. **Sentinel Strategy**: Uses 32768 (2^15), exploiting the gap just above Int16's maximum value
4. **Access Patterns**: Different load instructions (i32.load16_s vs i32.load)

**Mathematical Elegance:**

The 32768 sentinel works perfectly because:
- Int16 range: -32768 to 32767 (-2^15 to 2^15-1)
- Sentinel: 32768 = 2^15 (exactly one above maximum valid value)
- Storage: 32-bit space accommodates both Int16 values and the sentinel

**Comparison with Related Types:**

| Type | Non-Optional | Optional | Sentinel | Infrastructure |
|------|-------------|----------|----------|----------------|
| **Int16** | 2 bytes (String) | 4 bytes | 32768 | String → Integer |
| **UInt16** | 2 bytes (String) | 4 bytes | -1 | String → Integer |
| **Int** | 4 bytes | 8 bytes | 2^32 | Integer → Int64 |
| **Char** | 4 bytes | 4 bytes | -1 | Integer (unified) |

**Performance Characteristics:**

- **FixedArray[Int16]**: Optimal memory density with String infrastructure optimizations
- **FixedArray[Int16?]**: 2x memory overhead but maintains fast direct array access
- **Padding Behavior**: String infrastructure handles odd-length arrays with automatic padding
- **Signed vs Unsigned Load**: Uses `i32.load16_s` for correct sign extension of negative values

**Design Trade-offs:**

1. **Memory vs Infrastructure Reuse**: Chooses proven String infrastructure for 2-byte storage optimization
2. **Sentinel Availability**: Exploits mathematical gap at 2^15 for efficient None encoding
3. **Type System Switching**: Changes infrastructure when sentinel encoding is needed

**Edge Case Handling:**

- **Negative Values**: Correctly handled with signed 16-bit loads and sign extension
- **Boundary Values**: Int16.min_value (-32768) and Int16.max_value (32767) both supported
- **Sentinel Collision**: Impossible since 32768 > 32767 (max valid Int16)

This analysis reveals MoonBit's sophisticated approach to 16-bit integer optimization. The non-optional variant leverages the mature String infrastructure for optimal 2-byte storage, while the optional variant switches to integer arrays with a mathematically perfect sentinel value. The 32768 sentinel represents an elegant solution - using the natural gap at 2^15 to enable efficient None encoding without compromising the full range of valid Int16 values.

The analysis of FixedArray[Int16] and FixedArray[Int16?] is now complete! This reveals another excellent example of MoonBit's optimization strategy:

**Key Takeaways:**
1. **FixedArray[Int16] reuses String infrastructure** (Type 243) for optimal 2-byte storage density
2. **FixedArray[Int16?] switches to integer arrays** (Type 241) with 32768 (2^15) as the None sentinel
3. **Perfect mathematical sentinel** - 32768 is exactly one above Int16's maximum value (32767)
4. **100% memory overhead for optionality** but maintains direct access without indirection
5. **Infrastructure switching** - leverages different optimized systems for different requirements

This design showcases MoonBit's pragmatic optimization approach:

- **Non-optional optimization**: Reuses mature String infrastructure for 2-byte elements
- **Mathematical precision**: 32768 sentinel exploits the natural gap at 2^15
- **Type system flexibility**: Switches infrastructure when different capabilities are needed
- **Consistent patterns**: Same sentinel strategy as other 16-bit types but with signed values

The comparison with UInt16 is particularly interesting:
- **UInt16**: Uses -1 sentinel (outside 0-65535 range)
- **Int16**: Uses 32768 sentinel (outside -32768 to 32767 range)

Both achieve the same goal (efficient None encoding) but use different mathematical gaps appropriate to their respective value ranges, demonstrating MoonBit's systematic approach to finding optimal sentinel values for each type's specific constraints.
