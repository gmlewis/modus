## FixedArray[Bool] and FixedArray[Bool?] Memory Representation Analysis

Based on comprehensive examination of both MoonBit source code comments and WAT implementation, here's the detailed analysis of how MoonBit represents `FixedArray[Bool]` and `FixedArray[Bool?]` in WebAssembly linear memory.

### FixedArray[Bool] (Non-Optional) - Type 241

**Memory Layout:**
FixedArray[Bool] uses **4-byte integer storage** for Boolean values:

```
FixedArray[Bool] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
        1            Type 241        Bool Value      Bool Value...
```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.i32_array_make` with 0 as default value
2. **Array Header**: Type 241 (FixedArray[Int]) - shares type with other integer arrays
3. **Element Storage**: 4-byte integers with `size * 4` allocation
4. **Boolean Encoding**:
   - `false` → `0` (32-bit integer)
   - `true` → `1` (32-bit integer)
5. **Element Access**: Standard `moonbit.array_item` with 4-byte stride and `i32.load`

**Example Memory Layout** (FixedArray [false, true, false]):
```
FixedArray Data (20 bytes):
[1 0 0 0] [241 3 0 0] [0 0 0 0] [1 0 0 0] [0 0 0 0]
 RefCount  Type+Len    false     true      false
```

### FixedArray[Bool?] (Optional) - Type 241 with Sentinel Encoding

**Memory Layout:**
FixedArray[Bool?] uses **the same 4-byte integer storage** with `-1` as None sentinel:

```
FixedArray[Bool?] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
        1            Type 241        32-bit Value    32-bit Value...
```

**Sentinel Encoding Strategy:**

1. **None Representation**: `0xFFFFFFFF` (-1 in signed 32-bit)
2. **Some(false) Representation**: `0x00000000` (0)
3. **Some(true) Representation**: `0x00000001` (1)
4. **Sentinel Rationale**: -1 is outside Bool's valid range (0, 1), providing clean None encoding

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.i32_array_make` with `-1` as default value
2. **Array Header**: Same Type 241 as non-optional variant
3. **Element Storage**: Same 4-byte elements with `size * 4` allocation
4. **Element Access**: Same `moonbit.array_item` with 4-byte stride and `i32.load`
5. **No Additional Overhead**: Optional support adds zero memory overhead per element

**Example Memory Layout** (FixedArray [Some(true), None, Some(false)]):
```
FixedArray Data (20 bytes):
[1 0 0 0] [241 3 0 0] [1 0 0 0] [255 255 255 255] [0 0 0 0]
 RefCount  Type+Len    Some(true) None (-1)        Some(false)
```

### Critical Design Insights

**FixedArray[Bool] vs FixedArray[Bool?]:**

1. **Unified Infrastructure**: Both use identical Type 241 and WAT functions
2. **Zero Overhead Optionality**: Optional variant adds no memory cost per element
3. **Efficient Sentinel**: -1 provides clean None encoding without conflicting with valid Bool values (0, 1)
4. **Simple Decoding**: `value == -1 ? None : Some(value != 0)`

**Memory Efficiency Analysis:**

- **Per-element cost**: 4 bytes for both variants (no overhead for optionality)
- **Total overhead**: Only the 8-byte array header
- **Memory waste**: 3 bytes per Boolean (75% overhead) due to 32-bit storage of 1-bit values

**Comparison with Other Types:**

- **vs Array[Bool]**: Same implementation but different wrapper type
- **vs Array[UInt16?]**: Both use -1 sentinel but UInt16 stores in 2-byte elements in String infrastructure
- **vs Array[String?]**: Bool? uses in-value sentinel, String? uses null pointer
- **vs Array[Int64?]**: Bool? uses efficient sentinel, Int64? requires separate Option objects

**Performance Characteristics:**

- **Access Speed**: Optimal - direct array indexing with no indirection
- **Memory Density**: Poor - 4 bytes per bit of information
- **Cache Efficiency**: Good - contiguous memory layout with predictable access patterns
- **Branching**: Minimal - simple comparison for None detection

**Design Trade-offs:**

1. **Memory vs Simplicity**: Chooses implementation simplicity over memory density
2. **Unified Type System**: Reuses integer array infrastructure rather than bit-packing
3. **Consistent Interface**: Same access patterns as other FixedArray types

This analysis reveals MoonBit's prioritization of implementation consistency and simplicity over memory optimization for Boolean arrays. While the 4-byte-per-Boolean storage is memory-inefficient, it enables Bool arrays to seamlessly integrate with the existing integer array infrastructure, providing fast access and simple optional value handling through sentinel encoding.

The analysis of FixedArray[Bool] and FixedArray[Bool?] is now complete! This reveals an interesting design choice in MoonBit:

**Key Takeaways:**
1. **FixedArray[Bool] uses 4-byte integer storage** with false=0, true=1 - prioritizing simplicity over memory efficiency
2. **FixedArray[Bool?] uses -1 sentinel encoding** with zero memory overhead for optionality
3. **Both share Type 241 infrastructure** with other integer arrays for unified implementation
4. **Memory inefficient but performant** - 75% memory waste (4 bytes for 1 bit) but fast direct access
5. **Perfect sentinel choice** - -1 is cleanly outside Bool's valid range (0,1)

This design demonstrates MoonBit's philosophy of implementation consistency over micro-optimizations. Rather than implementing specialized bit-packed Boolean arrays, the compiler reuses the robust integer array infrastructure, gaining implementation simplicity, type system consistency, and excellent performance at the cost of memory density.

The optional variant showcases one of MoonBit's most efficient sentinel encodings - no additional memory cost and simple detection logic, making FixedArray[Bool?] nearly as efficient as the non-optional variant.
