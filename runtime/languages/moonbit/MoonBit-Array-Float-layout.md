## Detailed Analysis: MoonBit Array[Float] Memory Representation in WASM

After thorough analysis of the WAT file, MoonBit source code, and all Array[Float] operations, I can provide a comprehensive explanation of how MoonBit represents `Array[Float]` in WASM linear memory.

### Key Discovery: Efficient Single-Precision Storage with Option Optimization

MoonBit `Array[Float]` employs the **same dual storage strategy** as `Array[Double]`, but optimized for **32-bit IEEE 754 single-precision** floating-point values:

1. **Array[Float]**: Direct 32-bit IEEE 754 storage for maximum efficiency
2. **Array[Float?]**: Reference-based storage with compact Option objects

### Memory Layout

#### Outer Array Wrapper Structure (Same as other Arrays)
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (constant 1573120 = 0x180100)
0x08    4     Array length (number of elements)
0x0C    4     Pointer to inner FixedArray
```

#### Inner Array Structures - TWO DISTINCT VARIANTS

**For Array[Float] (non-optional):**
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (kind=1, elem_size_shift=2, length)
0x08    N*4   Float data (N elements * 4 bytes each, IEEE 754 single)
```

**For Array[Float?] (optional):**
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type info (kind=2, elem_size_shift=2, length)
0x08    N*4   Pointer data (N pointers * 4 bytes each)
```

### Option Object Structure (for Array[Float?])

**None Object:**
- Shared object at address 10248 (same as Array[Double?])

**Some(Float) Object:**
```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (1)
0x04    4     Type info (constant 1572865 = 0x00180001)
0x08    4     Float value (IEEE 754 32-bit)
```

### Detailed Implementation Analysis

#### Array[Float] Implementation

**Creation Function:**
- Uses `moonbit.float32_array_make(size, 0.0f)`
- **Memory allocation**: `size * 4` bytes (4 bytes per float)
- **Header calculation**: `(1 << 30) | (2 << 28) | length`
- **elem_size_shift=2**: Indicates 4 bytes per element (2^2 = 4)

**Element Storage:**
```wasm
local.get $array_ptr
local.get $index
i32.const 4
i32.mul
i32.add
local.get $value
f32.store offset=8 align=1    ; Store at offset 8+(index*4)
```

**Element Access:**
- Uses `moonbit.float32_array_item`
- Direct `f32.load offset=8+(index*4)`
- Returns raw 32-bit float value

#### Array[Float?] Implementation

**Creation Function:**
- Uses `moonbit.ref_array_make(size, None_address)` (same as Double?)
- **Memory allocation**: `size * 4` bytes (4 bytes per pointer)
- **Header calculation**: `(2 << 30) | (2 << 28) | length`
- **Same structure as Array[Double?]**

**Element Storage:**
```wasm
; For Some(value):
; 1. Allocate 12-byte Some object (vs 16-byte for Double)
i32.const 12
call moonbit.gc.malloc
local.tee $some_ptr
i32.const 1572865
i32.store offset=4 align=1    ; Some(Float) object header
local.get $some_ptr
local.get $float_value
f32.store offset=8 align=1    ; Store float in Some object

; 2. Store pointer in array (same as Double?)
local.get $array_ptr
local.get $index
i32.const 4
i32.mul
i32.add
local.get $some_ptr
i32.store offset=8 align=1    ; Store pointer at offset 8+(index*4)
```

### Memory Examples

#### Array[Float] Examples

**Empty Array `[]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 0, ptr_to_empty]
Empty inner:   [stored at address 37856]
```

**Two floats `[1.0f, 2.0f]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 2, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0x60000002, 1.0_as_f32, 2.0_as_f32]
Binary layout: [1, 0, 0, 0, 0x60000002, 0x3F800000, 0x40000000]
```

#### Array[Float?] Examples

**Single None `[None]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0xA0000001, 10248]
None object:   [shared object at address 10248]
```

**Single Some `[Some(1.0f)]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 1, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0xA0000001, ptr_to_some]
Some object:   [1, 0, 0, 0, 0x00180001, 1.0_as_f32]
```

**Mixed options `[Some(1.0f), None, Some(3.0f)]`**
```
Outer wrapper: [1, 0, 0, 0, 0x00180100, 3, ptr_to_inner]
Inner array:   [1, 0, 0, 0, 0xA0000003, ptr1, 10248, ptr3]
Some object 1: [1, 0, 0, 0, 0x00180001, 1.0_as_f32]
Some object 3: [1, 0, 0, 0, 0x00180001, 3.0_as_f32]
```

### Performance and Memory Characteristics

#### Array[Float] Advantages
1. **Compact storage**: 4 bytes per element (half the size of doubles)
2. **Direct access**: No pointer indirection for element access
3. **Cache efficiency**: Better memory locality due to smaller size
4. **WASM optimized**: Direct f32.load/f32.store operations

#### Array[Float?] Optimization
1. **Compact Some objects**: 12 bytes vs 16 bytes for Some(Double)
2. **Shared None**: Same None object as other optional types
3. **Memory overhead**: 4 bytes per element + 12 bytes per Some object
4. **Better than Double?**: 25% less memory per Some object

### Type Header Analysis

| Array Type | Kind | Elem Size Shift | Element Size | Storage Type |
|------------|------|-----------------|--------------|--------------|
| Array[Float] | 1 | 2 | 4 bytes | Direct IEEE 754 single |
| Array[Float?] | 2 | 2 | 4 bytes | Pointers to Option |
| Array[Double] | 1 | 3 | 8 bytes | Direct IEEE 754 double |
| Array[Double?] | 2 | 2 | 4 bytes | Pointers to Option |

### Empty Array Specialization

Each floating-point array type has its own pre-allocated empty array:
- **Array[Float]**: Address 37856 (FixedArray[Float])
- **Array[Double]**: Address 40480 (FixedArray[Double])
- **Shared None**: Address 10248 (used by all optional types)

### Option Object Header Differentiation

```
Some(Double): 0x00200001 = 2097153
Some(Float):  0x00180001 = 1572865
Difference:   0x00080000 = 524288
```

The header values encode the contained type, allowing runtime type checking and proper garbage collection.

### Comparison with Array[Double]

| Aspect | Array[Float] | Array[Double] | Efficiency Gain |
|--------|--------------|---------------|-----------------|
| Element size | 4 bytes | 8 bytes | 50% less memory |
| Some object size | 12 bytes | 16 bytes | 25% less memory |
| Precision | ~7 decimal digits | ~15 decimal digits | Trade-off |
| WASM support | Native f32 | Native f64 | Both optimal |

### Verification Across All Operations

This dual storage strategy is consistent across:
- ✅ **Array creation**: Specialized functions for optimal storage
- ✅ **Element access**: Type-specific access functions
- ✅ **Memory allocation**: Optimized for both storage and Option overhead
- ✅ **Option handling**: Shared None object, compact Some objects
- ✅ **Type safety**: Distinct headers prevent type confusion
- ✅ **IEEE 754 compliance**: Standard single-precision representation

### Key Design Insights

1. **Precision vs Memory Trade-off**: Clear choice between 32-bit and 64-bit precision
2. **Consistent Architecture**: Same patterns as Array[Double] but optimized for size
3. **Option Efficiency**: Compact Some objects reduce memory overhead
4. **Type Safety**: Runtime type checking via distinct headers
5. **WASM Integration**: Leverages native f32.load/f32.store operations

### Use Case Optimization

**Array[Float] is optimal for:**
- Graphics and game development (coordinates, colors)
- Audio processing (sample values)
- Scientific computing with moderate precision requirements
- Memory-constrained environments

**Array[Double] is preferred for:**
- Financial calculations requiring high precision
- Scientific computing with strict accuracy requirements
- Mathematical algorithms sensitive to precision

This implementation demonstrates MoonBit's sophisticated approach to numerical computing, providing both memory-efficient single-precision arrays and high-precision double arrays while maintaining consistent programming interfaces and full Option type support.
