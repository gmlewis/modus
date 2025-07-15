## Array[UInt64] and Array[UInt64?] Memory Representation Analysis

Based on comprehensive examination of both MoonBit source code comments and WAT implementation, here's the detailed analysis of how MoonBit represents `Array[UInt64]` and `Array[UInt64?]` in WebAssembly linear memory.

### Array[UInt64] (Non-Optional) - Type 241

**Memory Layout:**
Array[UInt64] uses **direct 64-bit storage** identical to Array[Int64]:

```
Array[UInt64] Object (16+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (8 bytes)       │ (8 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 241        UInt64 Value    UInt64 Value...
```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.int64_array_make` (shared with Array[Int64])
2. **Array Header**: Type 241 (FixedArray[UInt]) - same type as Array[UInt]
3. **Element Storage**: Direct 8-byte unsigned integers, allocated with `size * 8` bytes
4. **Element Access**: `moonbit.int64_array_item` with 8-byte stride and `i64.load`
5. **Value Range**: 0 to 18446744073709551615 (0x0000000000000000 to 0xFFFFFFFFFFFFFFFF)

**Example Memory Layout** (Array [1, 2, 3]):
```
Offset: 0x0000BFC0 (Array wrapper)
[1 0 0 0] [0 2 0 0] [A0 BF 0 0] [3 0 0 0]
 RefCount   Type 0    Inner Ptr   Length

Offset: 0x0000BFA0 (FixedArray data)
[1 0 0 0] [241 3 0 0] [1 0 0 0 0 0 0 0] [2 0 0 0 0 0 0 0] [3 0 0 0 0 0 0 0]
 RefCount  Type+Len    Element[0]        Element[1]        Element[2]
```

### Array[UInt64?] (Optional) - Type 242 with Reference-Based Storage

**Memory Layout:**
Array[UInt64?] uses **identical structure to Array[Int64?]** with reference-based storage:

```
Array[UInt64?] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Pointer[0]      │ Pointer[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 242        → None/Some      → None/Some...
```

**Option Object Representation:**

1. **None Values**: All point to shared singleton at offset 10648:
   ```
   None Object (8 bytes):
   [255 255 255 255] [0 0 0 0]
    RefCount -1       Type 0 (Tuple)
   ```

2. **Some Values**: Individual Option objects (type 1):
   ```
   Some Object (16 bytes):
   [1 0 0 0] [1 2 0 0] [Value (8 bytes)]
    RefCount  Type+Len   UInt64 Content
   ```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.ref_array_make` (shared with Array[Int64?])
2. **Array Header**: Type 242 (FixedArray[String]) - reuses string infrastructure for pointer storage
3. **Element Storage**: 4-byte pointers to Option objects, not direct values
4. **None Optimization**: Shared singleton object with RefCount -1 (immortal)
5. **Some Objects**: Each requires separate 16-byte allocation with 2097153 header
6. **No Sentinel Encoding**: Like Array[Int64?], no available bit patterns for in-band encoding

**Example Memory Layout** (Array [Some(1), None, Some(2)]):
```
Offset: 0x0000BFC0 (Array wrapper)
[1 0 0 0] [0 2 0 0] [C0 BE 0 0] [3 0 0 0]
 RefCount   Type 0    Inner Ptr   Length

Offset: 0x0000BEC0 (FixedArray data)
[1 0 0 0] [242 3 0 0] [A0 BF 0 0] [98 29 0 0] [C0 BF 0 0]
 RefCount  Type+Len    → Some(1)   → None      → Some(2)

Offset: 0x00002998 (Shared None)
[255 255 255 255] [0 0 0 0]

Offset: 0x0000BFA0 (Some(1))
[1 0 0 0] [1 2 0 0] [1 0 0 0 0 0 0 0]

Offset: 0x0000BFC0 (Some(2))
[1 0 0 0] [1 2 0 0] [2 0 0 0 0 0 0 0]
```

### Critical Design Insights

**Array[UInt64] vs Array[UInt64?]:**

1. **Identical to Int64 Strategy**: Both UInt64 and Int64 arrays use the exact same implementation
2. **Shared Infrastructure**: Same WAT functions (`int64_array_make`, `ref_array_make`) and type IDs
3. **No Signed/Unsigned Distinction**: At the memory level, both are just 64-bit values
4. **Reference-Based Optionals**: Same expensive approach as Array[Int64?] due to lack of sentinel values

**Comparison with Other Array Types:**

1. **vs Array[UInt]/Array[UInt?]**: UInt has efficient in-band encoding (2^32 sentinel), UInt64 requires separate objects
2. **vs Array[UInt16]/Array[UInt16?]**: UInt16 has efficient sentinel encoding (-1), UInt64 has no available sentinels
3. **vs Array[String?]**: String? uses null pointer encoding, UInt64? uses separate Option objects

**Performance Characteristics:**

- **Array[UInt64]**: Optimal - direct 8-byte storage with no overhead
- **Array[UInt64?]**: Expensive - 4-byte pointer + 16-byte Option object per Some value (≥20 bytes per element)
- **Memory Overhead**: Significant - approximately 2.5x overhead for Some values vs direct storage
- **Access Pattern**: Double indirection (pointer dereference + Option object access)

**Why No In-Band Encoding:**

Unlike smaller integer types, UInt64 uses the full 64-bit space (0 to 2^64-1), leaving no unused bit patterns for sentinel values. This forces MoonBit to use the more expensive reference-based approach with separate Option objects.

This analysis demonstrates that while MoonBit provides consistent APIs across all integer types, the underlying implementations vary dramatically based on the mathematical properties and available optimization opportunities of each type. Array[UInt64?] represents the "worst case" scenario where no efficient sentinel encoding is possible, requiring the full complexity of separate object allocation and reference management.

The analysis of Array[UInt64] and Array[UInt64?] is now complete! This reveals the final piece of MoonBit's integer array optimization puzzle:

**Key Takeaways:**
1. **Array[UInt64] shares identical implementation with Array[Int64]** - same WAT functions, same Type 241, same direct 8-byte storage
2. **Array[UInt64?] also shares implementation with Array[Int64?]** - same reference-based approach with Type 242 and separate Option objects
3. **No signed/unsigned distinction at memory level** - both are just 64-bit values in WebAssembly
4. **Full 64-bit space prevents sentinel encoding** - unlike smaller types, no unused bit patterns exist for efficient None representation
5. **Most expensive optional representation** - requires ~2.5x memory overhead due to separate object allocation

This completes our understanding of MoonBit's integer array hierarchy:

- **Efficient in-band encoding**: Array[Int?], Array[UInt?] (use 2^32 sentinel)
- **Efficient sentinel encoding**: Array[UInt16?] (uses -1 sentinel)
- **Efficient null pointer encoding**: Array[String?] (uses 0 pointer)
- **Expensive reference-based**: Array[Int64?], Array[UInt64?] (no available sentinels)

The pattern shows MoonBit's compiler systematically choosing the most efficient representation possible based on the mathematical constraints of each type, falling back to more expensive approaches only when necessary.
