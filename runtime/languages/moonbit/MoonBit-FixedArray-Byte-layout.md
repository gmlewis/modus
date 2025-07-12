## FixedArray[Byte] and FixedArray[Byte?] Memory Representation Analysis

Based on comprehensive examination of the WAT implementation, here's the detailed analysis of how MoonBit represents `FixedArray[Byte]` and `FixedArray[Byte?]` in WebAssembly linear memory.

### FixedArray[Byte] (Non-Optional) - Specialized Byte Storage

**Memory Layout:**
FixedArray[Byte] uses **specialized 1-byte storage** with 4-byte padding alignment:

```
FixedArray[Byte] Object (12+ bytes with padding):
┌─────────────────┬─────────────────┬───┬───┬───┬───┬─────────────────┐
│ RefCount        │ Array Header    │B0 │B1 │B2 │B3 │ Padding         │
│ (4 bytes)       │ (4 bytes)       │(1)│(1)│(1)│(1)│ (to 4-byte)     │
└─────────────────┴─────────────────┴───┴───┴───┴───┴─────────────────┘
        1            Special Type    Byte Values      Alignment
```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.bytes_make` (specialized for byte arrays)
2. **Array Header**: `i32.const 1` and `i32.const 0` (different type encoding than other arrays)
3. **Memory Allocation**: `(size + 3) & -4` bytes (rounds up to 4-byte boundary)
4. **Element Storage**: Direct 1-byte storage using `i32.store8`
5. **Element Access**: `moonbit.bytes_item` with 1-byte stride and `i32.load8_u`
6. **Value Range**: 0 to 255 (standard byte range)

**Example Memory Layout** (FixedArray [1, 2, 3, 4]):
```
Byte Array Data (12 bytes):
[1 0 0 0] [Special Header] [1] [2] [3] [4]
 RefCount   Type+Length     B0  B1  B2  B3
```

### FixedArray[Byte?] (Optional) - Type 241 with Sentinel Encoding

**Memory Layout:**
FixedArray[Byte?] **switches to 4-byte integer storage** to accommodate sentinel values:

```
FixedArray[Byte?] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
        1            Type 241        32-bit Value    32-bit Value...
```

**Sentinel Encoding Strategy:**

1. **None Representation**: `0xFFFFFFFF` (-1 in signed 32-bit)
2. **Some(value) Representation**: `0x000000XX` (byte value zero-extended to 32 bits)
3. **Sentinel Rationale**: -1 is outside Byte's valid range (0-255), providing clean None encoding

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.i32_array_make` with `-1` as default value (same as other optional integer arrays)
2. **Array Header**: Type 241 (FixedArray[Int]) - shares type with other integer arrays
3. **Element Storage**: 4-byte elements with `size * 4` allocation
4. **Element Access**: Standard `moonbit.array_item` with 4-byte stride and `i32.load`
5. **Memory Overhead**: 4x memory usage vs non-optional variant

**Example Memory Layout** (FixedArray [Some(1), None, Some(3)]):
```
Byte Option Array Data (20 bytes):
[1 0 0 0] [241 3 0 0] [1 0 0 0] [255 255 255 255] [3 0 0 0]
 RefCount  Type+Len    Some(1)   None (-1)         Some(3)
```

### Critical Design Insights

**FixedArray[Byte] vs FixedArray[Byte?]:**

1. **Dual Infrastructure**: Non-optional uses specialized byte infrastructure, optional uses integer infrastructure
2. **Memory Efficiency Trade-off**:
   - FixedArray[Byte]: 1 byte per element (optimal density)
   - FixedArray[Byte?]: 4 bytes per element (400% overhead for optionality)
3. **Type System Switch**: Changes from specialized byte type to Type 241 (integer array)
4. **Access Pattern Change**: 1-byte vs 4-byte stride, different load instructions

**Comparison with Other Types:**

- **vs FixedArray[Bool]**: Both non-optional and optional Bool use 4-byte storage; Byte optimizes non-optional to 1-byte
- **vs Array[Byte]**: Array[Byte] also uses the dual encoding strategy (1-byte vs 4-byte)
- **vs FixedArray[UInt16]**: UInt16 reuses String infrastructure; Byte has dedicated byte infrastructure
- **vs FixedArray[String]**: String always uses reference-based storage; Byte switches approaches

**Performance Characteristics:**

- **FixedArray[Byte]**: Optimal memory density, excellent cache efficiency
- **FixedArray[Byte?]**: 4x memory overhead but maintains simple sentinel-based None detection
- **Infrastructure Complexity**: Requires two completely different code paths for optional vs non-optional

**Memory Efficiency Analysis:**

```
Array Size: N elements
FixedArray[Byte]:   8 + N bytes (rounded up to 4-byte boundary)
FixedArray[Byte?]:  8 + N*4 bytes (no rounding needed)
```

**Design Trade-offs:**

1. **Memory vs Complexity**: Chooses memory optimization for common case (non-optional) but falls back to simpler approach for optional
2. **Type System Consistency**: Optional variant integrates with existing integer array infrastructure
3. **Access Speed**: Both variants maintain fast direct indexing

This analysis reveals MoonBit's sophisticated dual-strategy approach for byte arrays. The compiler optimizes the common case (non-optional bytes) with specialized 1-byte storage and dedicated byte operations, but switches to the standard 4-byte integer infrastructure for optional variants to enable efficient sentinel-based None encoding. This represents a perfect balance between memory efficiency for dense byte arrays and implementation simplicity for optional variants.

The analysis of FixedArray[Byte] and FixedArray[Byte?] is now complete! This reveals another fascinating optimization strategy:

**Key Takeaways:**
1. **FixedArray[Byte] uses specialized 1-byte storage** with dedicated `bytes_make`/`bytes_item` functions for optimal memory density
2. **FixedArray[Byte?] switches to 4-byte integer arrays** (Type 241) with -1 sentinel encoding for simplicity
3. **400% memory overhead for optionality** - the highest we've seen, but enables simple and fast None detection
4. **Dual infrastructure approach** - specialized byte operations vs standard integer array operations
5. **Perfect sentinel choice** - -1 is cleanly outside Byte's valid range (0-255)

This design demonstrates MoonBit's pragmatic optimization philosophy: maximize memory efficiency for the common case (dense byte arrays are often used for binary data), but fall back to proven infrastructure for less common cases (optional bytes). The compiler chooses the best representation for each specific use case rather than forcing uniformity across all variants.

The contrast between 1-byte and 4-byte storage also highlights why optional types can have significant memory implications in systems programming - sometimes the overhead of supporting null values fundamentally changes the data structure's characteristics.
