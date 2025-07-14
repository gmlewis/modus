## FixedArray[Double] and FixedArray[Double?] Memory Representation Analysis

Based on comprehensive examination of the WAT implementation, here's the detailed analysis of how MoonBit represents `FixedArray[Double]` and `FixedArray[Double?]` in WebAssembly linear memory.

### FixedArray[Double] (Non-Optional) - Specialized Double Storage

**Memory Layout:**
FixedArray[Double] uses **specialized 8-byte IEEE 754 storage**:

```
FixedArray[Double] Object (16+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (8 bytes)       │ (8 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
        1            Special Type    IEEE 754 Double IEEE 754 Double...
```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.float_array_make` (specialized for floating-point arrays)
2. **Array Header**: `i32.const 1` and `i32.const 3` (different type encoding than integer arrays)
3. **Memory Allocation**: `size * 8` bytes (direct 8-byte allocation for doubles)
4. **Element Storage**: Direct IEEE 754 double-precision storage using `f64.store`
5. **Element Access**: `moonbit.float_array_item` with 8-byte stride and `f64.load`
6. **Value Range**: Full IEEE 754 double-precision range (±1.8e308, with special values for NaN, ±Infinity)

**Example Memory Layout** (FixedArray [1.0, 2.0, 3.0]):
```
Double Array Data (32 bytes):
[1 0 0 0] [Special Header] [1.0 IEEE 754] [2.0 IEEE 754] [3.0 IEEE 754]
 RefCount   Type+Length     8 bytes        8 bytes        8 bytes
```

### FixedArray[Double?] (Optional) - Type 160 with Reference-Based Storage

**Memory Layout:**
FixedArray[Double?] uses **reference-based storage** identical to Array[Int64?] and Array[UInt64?]:

```
FixedArray[Double?] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Pointer[0]      │ Pointer[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
        1            Type 160        → None/Some      → None/Some...
```

**Option Object Representation:**

1. **None Values**: All point to shared singleton at offset 10248:
   ```
   None Object (8 bytes):
   [255 255 255 255] [0 0 0 0]
    RefCount -1       Type 0 (Tuple)
   ```

2. **Some Values**: Individual Option objects:
   ```
   Some Object (16 bytes):
   [1 0 0 0] [2097153] [IEEE 754 Double (8 bytes)]
    RefCount  Type+Len   Double Content
   ```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.ref_array_make` (shared with other reference-based optional arrays)
2. **Array Header**: Type 160 (64-bit reference types) - specialized classID for Double?/Float?/Int64?/UInt64?
3. **Element Storage**: 4-byte pointers to Option objects, not direct values
4. **None Optimization**: Shared singleton object with RefCount -1 (immortal)
5. **Some Objects**: Each requires separate 16-byte allocation with 2097153 header
6. **No Sentinel Encoding**: IEEE 754 has NaN values, but MoonBit doesn't use them as sentinels

**Example Memory Layout** (FixedArray [Some(1.0), None, Some(3.0)]):
```
Offset: 0x0000BFC0 (Array wrapper)
[1 0 0 0] [0 2 0 0] [C0 BE 0 0] [3 0 0 0]
 RefCount   Type 0    Inner Ptr   Length

Offset: 0x0000BEC0 (FixedArray data)
[1 0 0 0] [242 3 0 0] [A0 BF 0 0] [98 29 0 0] [C0 BF 0 0]
 RefCount  Type+Len    → Some(1.0) → None      → Some(3.0)

Offset: 0x00002998 (Shared None)
[255 255 255 255] [0 0 0 0]

Offset: 0x0000BFA0 (Some(1.0))
[1 0 0 0] [2097153] [1.0 IEEE 754 (8 bytes)]

Offset: 0x0000BFC0 (Some(3.0))
[1 0 0 0] [2097153] [3.0 IEEE 754 (8 bytes)]
```

### Critical Design Insights

**FixedArray[Double] vs FixedArray[Double?]:**

1. **Fundamentally Different Approaches**: Direct 8-byte storage vs reference-based pointer storage
2. **Memory Efficiency Trade-off**:
   - FixedArray[Double]: 8 bytes per element (optimal for IEEE 754)
   - FixedArray[Double?]: 4-byte pointer + 16-byte Option object (≥20 bytes per Some value)
3. **Type System Switch**: Changes from specialized double type to Type 160 (64-bit reference array)
4. **Access Pattern Change**: Direct f64.load vs pointer dereference + f64.load

**Why No NaN Sentinel Encoding:**

IEEE 754 doubles have many NaN representations that could theoretically serve as sentinels, but MoonBit doesn't use this approach because:
1. **Semantic Clarity**: Some(NaN) vs None should be distinguishable
2. **IEEE 754 Compliance**: All NaN values should be valid double values
3. **Implementation Consistency**: Uses same reference-based approach as other 64-bit optional types

**Comparison with Other Types:**

- **vs FixedArray[Float]**: Both floating-point types likely use reference-based optionals
- **vs Array[Double]**: Array[Double] shares the same implementation pattern
- **vs FixedArray[Int64]**: Int64 uses direct storage; Double also uses direct storage for non-optional
- **vs FixedArray[Char]**: Char uses sentinel encoding; Double uses reference-based encoding

**Performance Characteristics:**

- **FixedArray[Double]**: Optimal memory density and cache efficiency for IEEE 754 values
- **FixedArray[Double?]**: Significant memory overhead (≥2.5x) but maintains IEEE 754 compliance
- **Computational Cost**: Optional variant requires pointer dereferencing for every access
- **Memory Fragmentation**: Option objects scattered throughout heap vs contiguous array storage

**Design Trade-offs:**

1. **IEEE 754 Compliance vs Memory Efficiency**: Chooses standards compliance over sentinel-based optimization
2. **Semantic Clarity vs Performance**: Some(NaN) and None remain distinct concepts
3. **Implementation Consistency**: Uses same reference-based pattern as other 64-bit optional types

This analysis reveals MoonBit's commitment to IEEE 754 standards compliance and semantic clarity. Rather than using NaN values as sentinels (which would be memory-efficient but semantically confusing), the compiler uses the established reference-based approach for optional 64-bit types. This ensures that all valid IEEE 754 values, including various NaN representations, remain accessible through the Some() wrapper while maintaining clear None semantics.

The analysis of FixedArray[Double] and FixedArray[Double?] is now complete! This reveals another fascinating aspect of MoonBit's design philosophy:

**Key Takeaways:**
1. **FixedArray[Double] uses specialized 8-byte IEEE 754 storage** with dedicated floating-point operations for optimal performance
2. **FixedArray[Double?] uses reference-based storage** (Type 160) identical to other 64-bit optional types, avoiding NaN sentinels
3. **Standards compliance over optimization** - maintains full IEEE 754 semantics rather than using NaN values as None sentinels
4. **Significant memory overhead for optionality** - ≥2.5x memory cost due to separate Option object allocation
5. **Semantic clarity preserved** - Some(NaN) and None remain distinct, maintaining mathematical precision

This design demonstrates MoonBit's prioritization of:
- **Standards Compliance**: Full IEEE 754 support without sentinel abuse
- **Semantic Clarity**: Clear distinction between mathematical concepts (NaN) and language concepts (None)
- **Implementation Consistency**: Same reference-based approach for all 64-bit optional types
- **Mathematical Correctness**: All valid floating-point values remain accessible

The contrast between efficient direct storage for FixedArray[Double] and expensive reference-based storage for FixedArray[Double?] highlights the fundamental challenge of adding optionality to types that use their full bit space for valid values.

## Correction Notes (Updated After Successful Implementation)

**ClassID Correction**: The actual classID used by FixedArray[Double?] is **160**, not 242 as originally analyzed. This was discovered through:
1. Debug output showing `classID=160` in runtime execution
2. WAT analysis confirming the memory layout
3. Successful implementation requiring classID 160 handling

**Implementation Requirements**:
1. **Reading**: Add classID 160 to decoding conditions alongside 241/242/112
2. **Writing**: Use `moonbit_ref_array_make` function instead of manual memory allocation
3. **None Handling**: Special case for None singleton pointer at offset 10248
4. **Applies to**: Double?, Float?, Int64?, UInt64? (all 64-bit reference types)

**Working Go Implementation**: See commit "Fix FixedArrayOutput double option test" for complete implementation details.

This correction demonstrates the importance of runtime debugging over static analysis when dealing with compiler-generated memory layouts.
