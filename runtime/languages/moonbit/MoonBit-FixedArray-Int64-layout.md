## FixedArray[Int64] and FixedArray[Int64?] Memory Representation Analysis

Based on comprehensive examination of the WAT implementation, here's the detailed analysis of how MoonBit represents `FixedArray[Int64]` and `FixedArray[Int64?]` in WebAssembly linear memory.

### FixedArray[Int64] (Non-Optional) - Type 241

**Memory Layout:**
FixedArray[Int64] uses **direct 64-bit storage** identical to Array[Int64]:

```
FixedArray[Int64] Object (16+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (8 bytes)       │ (8 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 241        Int64 Value     Int64 Value...
```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.int64_array_make` (shared with Array[Int64])
2. **Array Header**: Type 241 (FixedArray[Int]) - same type as other integer arrays
3. **Element Storage**: Direct 8-byte signed integers, allocated with `size * 8` bytes
4. **Element Access**: `moonbit.int64_array_item` with 8-byte stride and `i64.load`
5. **Value Range**: -9223372036854775808 to 9223372036854775807 (full 64-bit signed range)

**Example Memory Layout** (FixedArray [1, 2, 3]):
```
FixedArray Data (32 bytes):
[1 0 0 0] [241 3 0 0] [1 0 0 0 0 0 0 0] [2 0 0 0 0 0 0 0] [3 0 0 0 0 0 0 0]
 RefCount  Type+Len    Element[0]        Element[1]        Element[2]
```

### FixedArray[Int64?] (Optional) - Type 242 with Reference-Based Storage

**Memory Layout:**
FixedArray[Int64?] uses **identical structure to Array[Int64?]** with reference-based storage:

```
FixedArray[Int64?] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Pointer[0]      │ Pointer[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 242        → None/Some      → None/Some...
```

**Option Object Representation:**

1. **None Values**: All point to shared singleton at offset 10248:
   ```
   None Object (8 bytes):
   [255 255 255 255] [0 0 0 0]
    RefCount -1       Type 0 (Tuple)
   ```

2. **Some Values**: Individual Option objects (type 1):
   ```
   Some Object (16 bytes):
   [1 0 0 0] [2097153] [Value (8 bytes)]
    RefCount  Type+Len   Int64 Content
   ```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.ref_array_make` (shared with Array[Int64?])
2. **Array Header**: Type 242 (FixedArray[String]) - reuses string infrastructure for pointer storage
3. **Element Storage**: 4-byte pointers to Option objects, not direct values
4. **None Optimization**: Shared singleton object with RefCount -1 (immortal)
5. **Some Objects**: Each requires separate 16-byte allocation with 2097153 header
6. **No Sentinel Encoding**: Full 64-bit space used by valid Int64 values, no available sentinels

**Example Memory Layout** (FixedArray [Some(1), None, Some(2)]):
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
[1 0 0 0] [2097153] [1 0 0 0 0 0 0 0]

Offset: 0x0000BFC0 (Some(2))
[1 0 0 0] [2097153] [2 0 0 0 0 0 0 0]
```

### Critical Design Insights

**FixedArray[Int64] vs FixedArray[Int64?]:**

1. **Identical to Array Strategy**: Both FixedArray and Array use the same implementation for Int64 types
2. **Shared Infrastructure**: Same WAT functions (`int64_array_make`, `ref_array_make`) and type IDs
3. **No FixedArray-Specific Optimization**: Both variants behave identically to their Array counterparts
4. **Reference-Based Necessity**: No available sentinel values in 64-bit space forces expensive approach

**Why No In-Band Encoding:**

Unlike smaller integer types, Int64 uses the complete 64-bit space (-2^63 to 2^63-1), leaving no unused bit patterns for sentinel values. This forces MoonBit to use the more expensive reference-based approach with separate Option objects.

**Comparison Across Int64 Types:**

| Type | Implementation | Memory per Element |
|------|----------------|-------------------|
| **FixedArray[Int64]** | Direct 8-byte | 8 bytes |
| **Array[Int64]** | Direct 8-byte | 8 bytes |
| **FixedArray[Int64?]** | Reference-based | ≥20 bytes |
| **Array[Int64?]** | Reference-based | ≥20 bytes |

**Performance Characteristics:**

- **FixedArray[Int64]**: Optimal - direct 8-byte storage with no overhead
- **FixedArray[Int64?]**: Expensive - 4-byte pointer + 16-byte Option object per Some value
- **Memory Overhead**: Approximately 2.5x overhead for Some values vs direct storage
- **Access Pattern**: Double indirection (pointer dereference + Option object access)
- **Cache Impact**: Poor locality due to scattered Option objects vs contiguous array storage

**Design Trade-offs:**

1. **Mathematical Constraints**: Full 64-bit value space prevents sentinel-based optimization
2. **Implementation Consistency**: Same approach for both FixedArray and Array variants
3. **Type Safety**: Maintains clear distinction between valid Int64 values and None
4. **Memory vs Simplicity**: Chooses reference-based approach over complex bit manipulation

**Edge Case Handling:**

- **Boundary Values**: Int64.min_value and Int64.max_value both fully supported
- **All Bit Patterns**: Every possible 64-bit pattern represents a valid Int64 value
- **No Sentinel Space**: No mathematical gaps available for efficient None encoding

This analysis reveals that FixedArray[Int64] and FixedArray[Int64?] represent the "least optimized" case in MoonBit's type hierarchy - not due to poor design, but due to mathematical constraints. The full utilization of 64-bit space by valid Int64 values forces the use of expensive reference-based optionals, demonstrating how the mathematical properties of types fundamentally constrain optimization opportunities in systems programming.

The analysis of FixedArray[Int64] and FixedArray[Int64?] is now complete! This reveals the final piece of the puzzle in MoonBit's optimization hierarchy:

**Key Takeaways:**
1. **FixedArray[Int64] shares identical implementation with Array[Int64]** - same WAT functions, same Type 241, same direct 8-byte storage
2. **FixedArray[Int64?] also shares implementation with Array[Int64?]** - same reference-based approach with Type 242 and separate Option objects
3. **No FixedArray vs Array distinction** at the implementation level for 64-bit integers
4. **Mathematical constraints dominate** - full 64-bit space prevents any sentinel encoding optimizations
5. **Most expensive optional representation** - requires ~2.5x memory overhead due to separate object allocation

This completes our comprehensive understanding of MoonBit's type optimization hierarchy:

**Optimization Spectrum (Best to Worst Memory Efficiency for Optionals):**

1. **Zero Overhead**: FixedArray[Bool?] (fits in same 4-byte space with -1 sentinel)
2. **Minimal Overhead**: FixedArray[String?] (null pointer encoding)
3. **Modest Overhead**: FixedArray[UInt16?] (-1 sentinel), FixedArray[Int16?] (32768 sentinel)
4. **Significant Overhead**: FixedArray[Int?], FixedArray[UInt?] (2^32 sentinel, 100% increase)
5. **Maximum Overhead**: FixedArray[Int64?], FixedArray[Double?], FixedArray[Float?] (reference-based, 250%+ increase)

**The Fundamental Pattern:**
- **Available sentinel space** → Efficient in-band or direct encoding
- **No available sentinel space** → Expensive reference-based encoding

FixedArray[Int64] represents the mathematical limit case where type safety and full value range utilization prevent any memory optimization for optional variants. This demonstrates how mathematical properties of types fundamentally constrain optimization opportunities in systems programming languages.
