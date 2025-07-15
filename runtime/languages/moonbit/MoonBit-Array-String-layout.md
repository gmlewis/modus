## Array[String] and Array[String?] Memory Representation Analysis

Based on comprehensive examination of both MoonBit source code comments and WAT implementation, here's the detailed analysis of how MoonBit represents `Array[String]` and `Array[String?]` in WebAssembly linear memory.

### Array[String] (Non-Optional) - Type 242

**Memory Layout:**
Array[String] uses **reference-based storage** with pointers to String objects:

```
Array[String] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Pointer[0]      │ Pointer[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 242        → String obj     → String obj...
```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.ref_array_make` with parameters:
   - Size: number of elements
   - Reference value: `11936` (a shared default string reference)
2. **Array Header**: Type 242 (FixedArray[String]) - same type as used for Array[Int64?]
3. **Element Storage**: 4-byte pointers to String objects, allocated with `size * 4` bytes
4. **Reference Management**: Handles reference counting for String objects

**String Object Structure (Type 243):**
```
String Object (16+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ String Header   │ UTF-16 Data     │ Length/Flags    │
│ (4 bytes)       │ (4 bytes)       │ (variable)      │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
   -1 (immortal)    Type 243        Unicode chars     Encoded length
```

**Example Memory Layout** (Array ["abc", "def", "ghi"]):
```
Offset: 0x0000BF20 (Array wrapper)
[1 0 0 0] [0 2 0 0] [40 BE 0 0] [3 0 0 0]
 RefCount   Type 0    Inner Ptr   Length

Offset: 0x0000BE40 (FixedArray data)
[1 0 0 0] [242 3 0 0] [20 3B 0 0] [D0 3E 0 0] [68 4B 0 0]
 RefCount  Type+Len    → "abc"     → "def"     → "ghi"

Offset: 0x00003B20 ("abc" string)
[255 255 255 255] [243 2 0 0] [97 0 98 0 99 0 0 1]
 RefCount -1        Type+Len     UTF-16: a b c + flags

Offset: 0x00003ED0 ("def" string)
[255 255 255 255] [243 2 0 0] [100 0 101 0 102 0 0 1]
 RefCount -1        Type+Len     UTF-16: d e f + flags

Offset: 0x00004B68 ("ghi" string)
[255 255 255 255] [243 2 0 0] [103 0 104 0 105 0 0 1]
 RefCount -1        Type+Len     UTF-16: g h i + flags
```

### Array[String?] (Optional) - Type 242

**Memory Layout:**
Array[String?] uses the **same structure** but with **null pointer encoding** for None:

```
Array[String?] Object (12+ bytes):
┌─────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ RefCount        │ Array Header    │ Pointer[0]      │ Pointer[1]      │
│ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │ (4 bytes)       │
└─────────────────┴─────────────────┴─────────────────┴─────────────────┘
     1573120          Type 242      0=None/→String   0=None/→String...
```

**Key Implementation Details:**

1. **WAT Function**: Uses `moonbit.ref_array_make` with parameters:
   - Size: number of elements
   - Default value: `0` (null pointer for None)
2. **None Representation**: Direct null pointer (0), no separate Option objects
3. **Some Representation**: Direct pointer to String object
4. **Efficient Encoding**: Avoids the complex Option object overhead used by Array[Int64?]

**Example Memory Layout** (Array [Some("abc"), None, Some("ghi")]):
```
Offset: 0x0000BF20 (Array wrapper)
[1 0 0 0] [0 2 0 0] [40 BE 0 0] [3 0 0 0]
 RefCount   Type 0    Inner Ptr   Length

Offset: 0x0000BE40 (FixedArray data)
[1 0 0 0] [242 3 0 0] [20 3B 0 0] [0 0 0 0] [68 4B 0 0]
 RefCount  Type+Len    → "abc"     None      → "ghi"

String objects at 0x00003B20 and 0x00004B68 same as above
```

### Critical Design Insights

**Array[String] vs Array[String?]:**

1. **Unified Approach**: Both use identical infrastructure (Type 242, ref_array_make)
2. **Optimal None Encoding**: Array[String?] uses null pointers instead of separate Option objects
3. **Reference Management**: Both handle String reference counting automatically
4. **Memory Efficiency**: None values cost only 4 bytes (null pointer) vs 16+ bytes for Option objects

**Comparison with Other Array Types:**

1. **vs Array[Int64]**: Strings always require reference-based storage (variable size), while Int64 uses direct storage
2. **vs Array[Int64?]**: String? uses efficient null pointer encoding, while Int64? requires separate Option objects (due to no available sentinel value in Int64 range)
3. **String Reuse**: String objects use RefCount -1 (immortal), enabling safe sharing across multiple arrays

**Performance Characteristics:**

- **Array[String]**: 4-byte pointer per element + string object overhead
- **Array[String?]**: Same as Array[String], with 0 for None (no additional overhead)
- **Access Pattern**: Single pointer dereference to reach string data
- **Memory Footprint**: Variable based on string content, but None values are extremely efficient

This demonstrates MoonBit's sophisticated approach to optional types, where the compiler chooses the most efficient representation based on the underlying type's characteristics. For reference types like String, null pointer encoding provides optimal performance, while for value types without available sentinel values, separate Option objects are used.

The analysis of Array[String] and Array[String?] is now complete! This reveals another fascinating aspect of MoonBit's type system optimization:

**Key Takeaways:**
1. **String arrays use reference-based storage** due to variable-length string content
2. **Array[String?] achieves optimal None encoding** using null pointers (0) instead of separate Option objects
3. **String objects are immortal** (RefCount -1) and use UTF-16 encoding
4. **Both variants use the same Type 242** infrastructure, showing unified handling of reference types
5. **This is much more efficient than Array[Int64?]** which requires separate Option objects due to lack of sentinel values in the Int64 space

The contrast between Array[String?] (null pointer encoding) and Array[Int64?] (separate Option objects) perfectly illustrates how MoonBit's compiler optimizes based on the specific constraints of each type.
