## Array[UInt16] and Array[UInt16?] Memory Representation Analysis

Based on comprehensive examination of the WAT implementation, here's the detailed analysis of how MoonBit represents `Array[UInt16]` and `Array[UInt16?]` in WebAssembly linear memory.

### Array[UInt16] (Non-Optional) - Type 243 (String Infrastructure)

**Memory Layout:**
Array[UInt16] **reuses the String infrastructure** with 16-bit element storage:

```
Array[UInt16] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (2 bytes)       │ (2 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 243        UInt16 Value    UInt16 Value...
```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.int16_array_make` → `moonbit.unsafe_make_string`
2. **Array Header**: Type 243 (String) with `i32.const 1` and `i32.const 1` parameters
3. **Element Storage**: 2-byte elements using `i32.store16`, with 2-byte alignment
4. **Memory Allocation**: `((size + 1) & -2) << 1` bytes (padded to even count, then doubled for 2-byte elements)
5. **Element Access**: `moonbit.int16_array_item_u` with 2-byte stride and `i32.load16_u`
6. **Value Range**: 0 to 65535 (0x0000 to 0xFFFF)

**Example Memory Layout** (Array [1, 2, 3]):
```
Offset: 0x0000BFA0 (Array wrapper)
[1 0 0 0] [0 2 0 0] [C0 BE 0 0] [3 0 0 0]
 RefCount   Type 0    Inner Ptr   Length

Offset: 0x0000BEC0 (String/FixedArray data)
[1 0 0 0] [243 3 0 0] [1 0] [2 0] [3 0] [xx xx]
 RefCount  Type+Len    El[0] El[1] El[2] Padding
```

### Array[UInt16?] (Optional) - Type 241 with Sentinel Encoding

**Memory Layout:**
Array[UInt16?] uses **32-bit elements** with `-1` sentinel for None:

```
Array[UInt16?] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 241       32-bit Value     32-bit Value...
```

**Sentinel Encoding Strategy:**

1. **None Representation**: `0xFFFFFFFF` (-1 in signed 32-bit)
2. **Some(value) Representation**: `0x0000XXXX` (value zero-extended to 32 bits)
3. **Sentinel Rationale**: -1 is outside UInt16's valid range (0 to 65535), providing clean None encoding

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.i32_array_make` with `-1` as default value
2. **Array Header**: Type 241 (FixedArray[UInt]) - shares type with Array[UInt]
3. **Element Storage**: 4-byte elements with `size * 4` allocation
4. **Element Access**: `moonbit.array_item` with 4-byte stride and `i32.load`
5. **Value Storage**: UInt16 values stored in 32-bit slots with zero-extension

**Example Memory Layout** (Array [Some(11), None, Some(33)]):
```
Offset: 0x0000BF20 (Array wrapper)
[1 0 0 0] [0 2 0 0] [C0 BE 0 0] [3 0 0 0]
 RefCount   Type 0    Inner Ptr   Length

Offset: 0x0000BEC0 (FixedArray data)
[1 0 0 0] [241 3 0 0] [11 0 0 0] [255 255 255 255] [33 0 0 0]
 RefCount  Type+Len    Some(11)   None (-1)         Some(33)
```

### Critical Design Insights

**Array[UInt16] vs Array[UInt16?]:**

1. **Infrastructure Reuse**: Non-optional reuses String (Type 243), optional reuses UInt array (Type 241)
2. **Memory Efficiency Trade-off**:
   - Array[UInt16]: 2 bytes per element (optimal)
   - Array[UInt16?]: 4 bytes per element (100% overhead for optionality)
3. **Sentinel Strategy**: Uses -1 sentinel, exploiting the gap between UInt16 max (65535) and signed 32-bit representation

**Comparison with Other Types:**

- **vs Array[Int16]**: Both reuse String infrastructure for non-optional variants
- **vs Array[UInt]**: Optional variants share Type 241 but use different sentinels (-1 vs 2^32)
- **vs Array[String]**: Non-optional UInt16 uses String infrastructure, but for 16-bit elements rather than pointers

**Performance Characteristics:**

- **Array[UInt16]**: Optimal memory density with 2-byte elements and string-optimized access
- **Array[UInt16?]**: 2x memory overhead but maintains direct element access with simple sentinel checking
- **Padding Behavior**: String infrastructure handles odd-length arrays with automatic padding

**Encoding Details:**
```
UInt16 Value Range: 0x0000 to 0xFFFF (0 to 65535)
None Sentinel:      0xFFFFFFFF (-1 in signed 32-bit)
```

This analysis reveals MoonBit's sophisticated infrastructure reuse strategy. Array[UInt16] leverages the existing String implementation, which is already optimized for 16-bit UTF-16 character storage, while Array[UInt16?] switches to the 32-bit array infrastructure to accommodate the -1 sentinel value. This demonstrates the compiler's ability to choose optimal representations based on both the element type characteristics and the availability of suitable sentinel values for optional variants.

The analysis of Array[UInt16] and Array[UInt16?] is now complete! This reveals yet another fascinating optimization strategy:

**Key Takeaways:**
1. **Array[UInt16] cleverly reuses String infrastructure** (Type 243) since strings already handle 16-bit UTF-16 elements optimally
2. **Array[UInt16?] switches to 32-bit infrastructure** (Type 241) to accommodate the -1 sentinel value
3. **Memory overhead is exactly 100%** (2→4 bytes per element) for optionality, but avoids complex Option objects
4. **The -1 sentinel works perfectly** because it's outside UInt16's valid range (0-65535)
5. **Different types for different needs**: String infrastructure for optimal 16-bit storage vs UInt infrastructure for sentinel-based optionals

This demonstrates MoonBit's pragmatic approach: reuse existing optimized infrastructure when possible (String for 16-bit elements), but switch to appropriate alternatives when additional requirements (like sentinel values) demand it. The compiler intelligently balances code reuse, memory efficiency, and type safety requirements.
