## FixedArray[Float] and FixedArray[Float?] Memory Representation Analysis

Based on comprehensive examination of the WAT implementation, here's the detailed analysis of how MoonBit represents `FixedArray[Float]` and `FixedArray[Float?]` in WebAssembly linear memory.

### FixedArray[Float] (Non-Optional) - Specialized Float Storage

**Memory Layout:**
FixedArray[Float] uses **specialized 4-byte IEEE 754 storage**:

```
FixedArray[Float] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
        1            Special Type    IEEE 754 Float  IEEE 754 Float...
```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.float32_array_make` (specialized for single-precision floating-point)
2. **Array Header**: `i32.const 1` and `i32.const 2` (different from doubles which use 3)
3. **Memory Allocation**: `size * 4` bytes (direct 4-byte allocation for floats)
4. **Element Storage**: Direct IEEE 754 single-precision storage using `f32.store`
5. **Element Access**: `moonbit.float32_array_item` with 4-byte stride and `f32.load`
6. **Value Range**: Full IEEE 754 single-precision range (±3.4e38, with special values for NaN, ±Infinity)

**Example Memory Layout** (FixedArray [1.0, 2.0, 3.0]):
```
Float Array Data (20 bytes):
[1 0 0 0] [Special Header] [1.0 IEEE 754] [2.0 IEEE 754] [3.0 IEEE 754]
 RefCount   Type+Length     4 bytes        4 bytes        4 bytes
```

### FixedArray[Float?] (Optional) - Type 242 with Compact Reference-Based Storage

**Memory Layout:**
FixedArray[Float?] uses **reference-based storage** with more compact Option objects than Double:

```
FixedArray[Float?] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Pointer[0]      │ Pointer[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
        1            Type 242        → None/Some      → None/Some...
```

**Option Object Representation:**

1. **None Values**: All point to shared singleton at offset 10248:
   ```
   None Object (8 bytes):
   [255 255 255 255] [0 0 0 0]
    RefCount -1       Type 0 (Tuple)
   ```

2. **Some Values**: Compact Option objects (12 bytes vs 16 for Double):
   ```
   Some Object (12 bytes):
   [1 0 0 0] [1572865] [IEEE 754 Float (4 bytes)]
    RefCount  Type+Len   Float Content
   ```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.ref_array_make` (shared with other reference-based optional arrays)
2. **Array Header**: Type 242 (FixedArray[String]) - reuses string infrastructure for pointer storage
3. **Element Storage**: 4-byte pointers to Option objects, not direct values
4. **None Optimization**: Shared singleton object with RefCount -1 (immortal)
5. **Compact Some Objects**: Only 12 bytes per Some value (vs 16 for Double Option objects)
6. **Option Header**: Uses `1572865` instead of `2097153` used by Double options

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

Offset: 0x0000BFA0 (Some(1.0) - 12 bytes)
[1 0 0 0] [1572865] [1.0 IEEE 754 (4 bytes)]

Offset: 0x0000BFC0 (Some(3.0) - 12 bytes)
[1 0 0 0] [1572865] [3.0 IEEE 754 (4 bytes)]
```

### Critical Design Insights

**FixedArray[Float] vs FixedArray[Float?]:**

1. **Fundamentally Different Approaches**: Direct 4-byte storage vs reference-based pointer storage
2. **Memory Efficiency Trade-off**:
   - FixedArray[Float]: 4 bytes per element (optimal for IEEE 754 single-precision)
   - FixedArray[Float?]: 4-byte pointer + 12-byte Option object (≥16 bytes per Some value)
3. **Compact Option Objects**: Float options use 12 bytes vs 16 bytes for Double options
4. **Type System Consistency**: Same Type 242 as other reference-based optional arrays

**Float vs Double Comparison:**

| Aspect | FixedArray[Float] | FixedArray[Double] |
|--------|-------------------|-------------------|
| **Element Size** | 4 bytes | 8 bytes |
| **Array Header** | `1, 2` | `1, 3` |
| **Option Object Size** | 12 bytes | 16 bytes |
| **Option Header** | `1572865` | `2097153` |
| **Memory Overhead** | 4x for optionality | 2.5x for optionality |

**Why No NaN Sentinel Encoding:**

Like Double, Float doesn't use NaN values as sentinels because:
1. **IEEE 754 Compliance**: All NaN representations should remain valid
2. **Semantic Clarity**: Some(NaN) vs None should be distinguishable
3. **Implementation Consistency**: Same reference-based approach as other floating-point optionals

**Performance Characteristics:**

- **FixedArray[Float]**: Optimal memory density and cache efficiency for 32-bit IEEE 754 values
- **FixedArray[Float?]**: 4x memory overhead but more efficient than Double? due to compact Option objects
- **Cache Friendliness**: Float arrays pack more elements per cache line than Double arrays
- **Access Pattern**: Optional variant requires pointer dereferencing for every access

**Design Trade-offs:**

1. **IEEE 754 Compliance vs Memory Efficiency**: Prioritizes standards compliance over sentinel-based optimization
2. **Precision vs Memory**: Single-precision uses half the memory of double-precision for both direct and optional storage
3. **Implementation Consistency**: Uses same reference-based pattern as Double? but with size optimizations

**Optimization Highlights:**

- **Compact Option Objects**: 12-byte Float options vs 16-byte Double options (25% memory savings)
- **Efficient Packing**: 4-byte elements enable better cache utilization than 8-byte doubles
- **Unified Infrastructure**: Shares Type 242 with other reference-based optionals

This analysis reveals MoonBit's attention to memory efficiency within its design constraints. While FixedArray[Float?] still uses the expensive reference-based approach (due to IEEE 754 compliance requirements), it optimizes Option object size compared to Double, demonstrating the compiler's awareness of the different memory requirements for different floating-point precisions.

The analysis of FixedArray[Float] and FixedArray[Float?] is now complete! This reveals the final piece of MoonBit's floating-point optimization strategy:

**Key Takeaways:**
1. **FixedArray[Float] uses specialized 4-byte IEEE 754 storage** with dedicated single-precision operations for optimal performance
2. **FixedArray[Float?] uses compact reference-based storage** with 12-byte Option objects (vs 16 bytes for Double)
3. **25% memory savings on Option objects** compared to Double? while maintaining the same reference-based approach
4. **IEEE 754 compliance preserved** - no NaN sentinel abuse, maintaining mathematical precision
5. **4x memory overhead for optionality** - still expensive but more efficient than Double? due to compact objects

This design showcases MoonBit's sophisticated memory optimization within design constraints:

- **Standards Compliance**: Full IEEE 754 support without compromising floating-point semantics
- **Size-Aware Optimization**: Compact Option objects for smaller floating-point types
- **Cache Efficiency**: 4-byte elements enable better cache line utilization than 8-byte doubles
- **Unified Infrastructure**: Consistent reference-based approach across all floating-point optionals

The Float vs Double comparison reveals MoonBit's attention to memory efficiency - while both use expensive reference-based optionals, Float achieves meaningful memory savings through compact Option objects, demonstrating that the compiler optimizes within its architectural constraints rather than applying one-size-fits-all solutions.
