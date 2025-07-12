## Detailed Analysis: MoonBit Array[Bool] Memory Representation in WASM

After thorough analysis of the WAT file, MoonBit source code, and all Array[Bool] operations, I can provide a comprehensive explanation of how MoonBit represents `Array[Bool]` in WASM linear memory.

### Key Discovery: Double-Indirection Structure

MoonBit `Array[Bool]` uses a **two-level memory structure**:

1. **Outer wrapper**: An Array container (tuple-like structure)
2. **Inner array**: The actual FixedArray containing the boolean data

### Memory Layout

#### Outer Array Wrapper Structure
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (constant 1573120 = 0x180100)
0x08    4     Pointer to inner FixedArray
0x0C    4     Array length (number of elements)
```

#### Inner FixedArray Structure
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info + length (type 241 + embedded length)
0x08    N*4   Boolean data (N elements * 4 bytes each)
```

### Detailed Field Analysis

#### Outer Wrapper Fields
- **Reference count**: Standard garbage collection counter (typically 1)
- **Type info**: Constant `1573120` (0x180100) - identifies this as an Array wrapper
- **Data pointer**: 32-bit pointer to the inner FixedArray
- **Length**: Number of elements in the array

#### Inner FixedArray Fields
- **Reference count**: Independent reference counter (typically 1)
- **Type info**: Type ID 241 (FixedArray[Int]) with embedded length information
- **Boolean data**: Each boolean stored as a 4-byte integer (0 = false, 1 = true)

### Memory Examples from Analysis

#### Empty Array `[]`
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, ptr_to_inner, 0]
Inner array:   [1, 0, 0, 0, 241, 0, 0, 0]
```

#### Single element `[true]`
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, ptr_to_inner, 1]
Inner array:   [1, 0, 0, 0, 241, 1, 0, 0, 1, 0, 0, 0]
```

#### Two elements `[false, true]`
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, ptr_to_inner, 2]
Inner array:   [1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0]
```

### Element Access Pattern

Array elements are accessed using:
1. **Load outer wrapper**: Get pointer to inner array from offset 12
2. **Access inner array**: Use `moonbit.array_item` function
3. **Element calculation**: `inner_base + 8 + (index * 4)`
4. **Load value**: 32-bit integer (0 or 1)

The `moonbit.array_item` function performs:
```wasm
local.get $arr          ; inner array pointer
local.get $index        ; element index
i32.const 4            ; 4 bytes per element
i32.mul                ; index * 4
i32.add                ; arr + (index * 4)
i32.load offset=8      ; load from offset 8 (after header)
```

### Special Cases and Optimizations

#### Empty Array Optimization
- All empty arrays share the same pre-allocated inner FixedArray at address 19632
- The outer wrapper is still allocated per instance but points to the shared empty array
- This saves memory for frequently used empty arrays

#### Option Types `Array[Bool?]`
- For `Array[Bool?]`, the inner array uses `-1` to represent `None`
- `Some(false)` = 0, `Some(true)` = 1, `None` = -1
- Array creation uses `moonbit.i32_array_make` with initial value -1

### Type Information Encoding

#### Outer Wrapper Type (1573120 = 0x180100)
- This constant identifies the Array wrapper type
- Used consistently across all Array types (not just Bool)

#### Inner FixedArray Type (241)
- Type ID 241 represents `FixedArray[Int]`
- Length is encoded separately in the type header
- Uses the same encoding as other array types (but different from String encoding)

### Key Implementation Details

1. **Double allocation**: Each Array[Bool] requires two memory allocations
2. **4-byte elements**: Booleans are stored as full 32-bit integers for efficiency
3. **Shared empty arrays**: Memory optimization for empty arrays
4. **Reference counting**: Both wrapper and inner array have independent reference counts
5. **Type safety**: Distinct type IDs prevent confusion between different structures

### Verification Across All Operations

This structure is consistent across:
- ✅ **Array creation**: Both literal arrays and dynamic allocation
- ✅ **Element access**: Standard `moonbit.array_item` function
- ✅ **Array comparison**: `assert_eq` functions work on both levels
- ✅ **Option handling**: `-1` encoding for `None` values
- ✅ **Empty arrays**: Shared optimization for zero-length arrays
- ✅ **Memory management**: Reference counting for both wrapper and data

This double-indirection design allows MoonBit to maintain consistent Array semantics while providing efficient memory management and type safety in the WASM environment.
