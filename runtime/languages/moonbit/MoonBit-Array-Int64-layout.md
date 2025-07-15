## Array[Int64] and Array[Int64?] Memory Representation Analysis

Based on detailed examination of both the MoonBit source code comments and WAT implementation, I can now provide a comprehensive analysis of how MoonBit represents `Array[Int64]` and `Array[Int64?]` in WebAssembly linear memory.

### Array[Int64] (Non-Optional) - Type 241

**Memory Layout:**
Array[Int64] uses the standard MoonBit array structure with **direct 64-bit storage**:

```
Array[Int64] Object (16+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Element[0]      │ Element[1]      │
│ (4 bytes)       │ (4 bytes)       │ (8 bytes)       │ (8 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 241        Int64 Value      Int64 Value...
```

**Key Implementation Details:**

1. **Array Header**: Contains type 241 (FixedArray[Int]) with element count
2. **Element Storage**: Each Int64 stored directly as 8-byte values at 8-byte aligned offsets
3. **WAT Functions**:
   - `moonbit.int64_array_make`: Allocates `size * 8` bytes, uses `i32.const 3` (maps to type 241)
   - `moonbit.int64_array_item`: Direct 64-bit load with `i64.load offset=8` and 8-byte stride
4. **Memory Efficiency**: Optimal - no overhead per element beyond the 8 bytes needed

**Example Memory Layout** (Array [1, 2, 3]):
```
Offset: 0x0000BFA0 (Array wrapper)
[1 0 0 0] [0 2 0 0] [C0 BE 0 0] [3 0 0 0]
 RefCount   Type 0    Inner Ptr   Length

Offset: 0x0000BEC0 (FixedArray data)
[1 0 0 0] [241 3 0 0] [1 0 0 0 0 0 0 0] [2 0 0 0 0 0 0 0] [3 0 0 0 0 0 0 0]
 RefCount  Type+Len    Element[0]        Element[1]        Element[2]
```

### Array[Int64?] (Optional) - Type 242

**Memory Layout:**
Array[Int64?] uses a **fundamentally different strategy** with reference-based storage:

```
Array[Int64?] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Pointer[0]      │ Pointer[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 242         → None/Some      → None/Some...
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
    RefCount  Type+Len   Int64 Content
   ```

**Key Implementation Details:**

1. **Array Header**: Uses type 242 (FixedArray[String]) - reuses string infrastructure for pointer storage
2. **Element Storage**: 4-byte pointers to Option objects, not direct values
3. **None Optimization**: Shared singleton object with RefCount -1 (immortal)
4. **Some Objects**: Each requires separate 16-byte allocation
5. **No Sentinel Encoding**: Unlike Array[Int?] which uses in-band encoding with 2^32 sentinel

**Example Memory Layout** (Array [None, Some(1), Some(2)]):
```
Offset: 0x0000BFC0 (Array wrapper)
[1 0 0 0] [0 2 0 0] [C0 BE 0 0] [3 0 0 0]
 RefCount   Type 0    Inner Ptr   Length

Offset: 0x0000BEC0 (FixedArray data)
[1 0 0 0] [242 3 0 0] [98 29 0 0] [A0 BF 0 0] [C0 BF 0 0]
 RefCount  Type+Len    → None      → Some(1)   → Some(2)

Offset: 0x00002998 (Shared None)
[255 255 255 255] [0 0 0 0]

Offset: 0x0000BFA0 (Some(1))
[1 0 0 0] [1 2 0 0] [1 0 0 0 0 0 0 0]

Offset: 0x0000BFC0 (Some(2))
[1 0 0 0] [1 2 0 0] [2 0 0 0 0 0 0 0]
```

### Critical Design Differences

**Array[Int64] vs Array[Int64?]:**

1. **Storage Strategy**: Direct vs Reference-based
2. **Type System**: 241 (FixedArray[Int]) vs 242 (FixedArray[String])
3. **Memory Overhead**: 8 bytes/element vs 20+ bytes/element (4-byte pointer + 16-byte Option object)
4. **Access Pattern**: Single load vs pointer dereference + load
5. **None Handling**: Not applicable vs shared singleton optimization

**Comparison with Array[Int?]:**
- Array[Int?] uses clever **in-band encoding** with 2^32 as None sentinel in 64-bit storage
- Array[Int64?] **cannot use in-band encoding** because there's no unused bit pattern in Int64's full 64-bit range
- This necessitates the more expensive reference-based approach

This analysis demonstrates MoonBit's sophisticated approach to memory optimization, where the compiler chooses different strategies based on the specific constraints and characteristics of each type, prioritizing efficiency while maintaining type safety.
