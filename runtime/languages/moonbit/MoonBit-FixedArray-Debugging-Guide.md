# MoonBit FixedArray Debugging Guide

## Overview

This guide provides proven techniques for debugging and fixing MoonBit FixedArray issues based on successful resolution of the `TestFixedArrayOutput_double_option_4` test and related types.

## Most Effective Debugging Approach

### 1. WAT Analysis First (MOST IMPORTANT)

**Start with the WASM Text (WAT) output** - this is the single most effective technique:

```bash
# Generate WAT if needed
cd /app/runtime/languages/moonbit/testdata
./build.sh  # or similar build process

# Find the specific function
grep -A 30 "test_fixedarray_output_double_option_4" build/testdata.wat
```

**Why this works best:**
- WAT shows the exact memory layout and operations MoonBit generates
- Reveals the actual function calls used (`moonbit.ref_array_make` vs `moonbit.i32_array_make`)
- Shows real classID values and memory offsets
- Eliminates guesswork about memory structures

**Key things to look for in WAT:**
- Function calls: `moonbit.ref_array_make`, `moonbit.i32_array_make`, etc.
- ClassID constants: `i32.const 160`, `i32.const 242`, etc.
- Memory allocation patterns
- None singleton usage: `i32.const 10248`

### 2. Add Targeted Debug Output

**Focus debug on the decoding path first:**
```go
fmt.Printf("DEBUG: %s memoryBlockAtOffset returned classID=%d, words=%d\n", h.typeDef.Name, classID, words)
```

**Then trace element processing:**
```go
fmt.Printf("DEBUG: Element %d: ptr=%d (0x%X)\n", i, ptr, ptr)
```

### 3. Understand Memory Layout Categories

Based on successful fixes, FixedArray types fall into these categories:

**Category A: Simple Types (Working)**
- `Bool?`, `Byte?`, `Char?` 
- Use classID 96
- Use `moonbit_i32_array_make`
- Have custom encoding logic

**Category B: 64-bit Reference Types (Fixed)**
- `Double?`, `Float?`, `Int64?`, `UInt64?`
- Use classID 160 (NOT 242 as originally documented)
- Use `moonbit_ref_array_make` 
- Need None singleton handling (offset 10248)

**Category C: Primitive Non-Optional Types (Fixed)**
- `Int16` - Use `moonbit_int16_array_make` (classID 80)
- Other primitive types work with various `ptr2*_array` functions

**Category D: ClassID 96 Optional Types (Fixed)**
- `Int16?` - Use direct storage with 32768 as None value (classID 96)
- Other classID 96 types: `Bool?`, `Byte?`, `Char?` - all working

**Category E: Int64-based Optional Types (Fixed)**
- `Int?` - Uses `moonbit_int64_array_make` with 4294967296 (1 << 32) as None value
- **Key insight**: Int? elements stored as 64-bit values despite being 32-bit integers
- Sign-extended encoding: int32 values become int64 in memory

**Category F: Reference Non-optional Types (Fixed)**
- `String` - Uses `moonbit_ref_array_make` to store pointers to string objects
- **Key insight**: Non-optional strings use reference storage, not direct storage

**Category G: Reference Optional with Custom None (Fixed)**
- `String?` - Uses `moonbit_ref_array_make` with 0 as None value (not 10248)
- **Critical insight**: Not all Category B types use the same None value
- Demonstrates type-specific None patterns within same infrastructure

**Category H: Other Types (Status Unknown)**
- `UInt?` - Likely similar to Int? but needs verification
- `UInt16?` - Similar to Int16? but needs verification
- `UInt16` - Needs manual handling (no moonbit_uint16_array_make)

## Proven Fix Pattern for Category B Types (64-bit Reference Optionals)

For `Double?`, `Float?`, `Int64?`, `UInt64?` - use `moonbit_ref_array_make`

### Reading Side Fix

1. **Add classID to condition:**
```go
if classID == FixedArrayPrimitiveBlockType || classID == PtrArrayBlockType || classID == 112 || classID == 160 {
```

2. **Handle None singleton in element decoding:**
```go
if ptr == 10248 {
    // None singleton pointer - return nil
    item = nil
} else {
    item, err = h.elementHandler.Decode(ctx, wasmAdapter, []uint64{uint64(ptr)})
    // ...
}
```

3. **Add None singleton handling in memoryBlockAtOffset:**
```go
if offset == 10248 {
    // Return mock data that represents None singleton structure
    return []byte{255, 255, 255, 255, 0, 0, 0, 0}, 0, 0, nil
}
```

### Writing Side Fix

1. **Use moonbit_ref_array_make instead of manual memory allocation:**
```go
else if elemType.Name() == "Double?" || elemType.Name() == "Float?" || elemType.Name() == "Int64?" || elemType.Name() == "UInt64?" {
    return h.createRefArrayWithMoonBit(ctx, wasmAdapter, slice, numElements)
}
```

2. **Implement createRefArrayWithMoonBit function:**
```go
func (h *sliceHandler) createRefArrayWithMoonBit(ctx context.Context, wasmAdapter langsupport.WasmAdapter, slice []any, numElements uint32) (uint32, utils.Cleaner, error) {
    // Use moonbit_ref_array_make
    fn := wasmAdapter.GetFunction("moonbit_ref_array_make")
    results, err := fn.Call(ctx, uint64(numElements), uint64(0))
    
    // Write element pointers
    for i, val := range slice {
        var elementPtr uint32
        if utils.HasNil(val) {
            elementPtr = 10248  // None singleton
        } else {
            // Encode element and get pointer
            results, _, err := h.elementHandler.Encode(ctx, wasmAdapter, val)
            elementPtr = uint32(results[0])
        }
        // Write to array[8 + i*4]
    }
}
```

## Common Mistakes to Avoid

### ❌ Don't Do This

1. **Don't assume documented classID is correct** - always verify with WAT output
2. **Don't manually construct memory headers** - use MoonBit runtime functions when possible
3. **Don't ignore None singleton addresses** - they're critical for optional types
4. **Don't try to fix encoding before fixing decoding** - reading is usually simpler
5. **Don't assume all optional types work the same way** - there are distinct categories

### ✅ Do This Instead

1. **Always start with WAT analysis** - it shows exactly what MoonBit generates
2. **Use MoonBit runtime functions** - `moonbit_ref_array_make`, `moonbit_i32_array_make`, etc.
3. **Handle None singleton explicitly** - offset 10248 is special
4. **Fix reading first, then writing** - decoding reveals the memory structure
5. **Group similar types together** - they often share the same patterns

## Proven Fix Pattern for Category C Types (Primitive Non-Optional)

For primitive types like `Int16` - use `moonbit_*_array_make` functions:

### Int16 Fix (Successfully Implemented)

1. **Use moonbit_int16_array_make instead of ptr2*_array:**
```go
case "Int16":
    // Use moonbit_int16_array_make
    arrayPtr, err = concreteWa.fnMakeArrayInt16.Call(ctx, uint64(numElements), 0)
    // Write data to the created array
```

## Proven Fix Pattern for Category D Types (ClassID 96 Optional)

For types like `Int16?` - use `moonbit_i32_array_make` functions:

### Int16? Fix (Successfully Implemented)

**Pattern**: ClassID 96 + Direct Storage + Custom None Value

1. **Reading Side:**
```go
if elemType.Name() == "Int16?" && classID == 96 {
    switch value {
    case 32768:  // NoneValueInt16 constant
        item = nil
    default:
        i := int16(value)
        item = &i
    }
}
```

2. **Writing Side:**
```go
func (h *sliceHandler) createInt16ArrayWithMoonBit(...) {
    // Use moonbit_i32_array_make(numElements, -1)
    fn := wasmAdapter.GetFunction("moonbit_i32_array_make")
    results, err := fn.Call(ctx, uint64(numElements), uint64(0xFFFFFFFF))
    
    // Write elements as 32-bit values with custom None
    for i, val := range slice {
        var encodedValue uint32
        if utils.HasNil(val) {
            encodedValue = 32768  // None value for Int16?
        } else {
            encodedValue = uint32(*val.(*int16))
        }
        // Write at arrayPtr + 8 + i*4
    }
}
```

## Proven Fix Pattern for Category E Types (Int64-based Optional)

For `Int?` arrays - use `moonbit_int64_array_make` with special None handling:

### Int? Fix (Successfully Implemented)

**Pattern**: Int64 Storage + 4294967296 None Value

1. **Reading Side:**
```go
if elemType.Name() == "Int?" {
    if value == 4294967296 {  // NoneValueInt constant (1 << 32)
        // None value, leave as nil
    } else {
        intVal := int32(value)  // Convert from 64-bit to 32-bit
        items.Index(int(i)).Set(reflect.ValueOf(&intVal))
    }
}
```

2. **Writing Side uses moonbit_int64_array_make with 64-bit storage**

**Key insights for Int?:**
- Uses 64-bit storage even though Int is 32-bit
- None value is `1 << 32` (4294967296), not 0xFFFFFFFF
- Elements stored 8 bytes apart (not 4)
- Sign extension required for negative values

## Proven Fix Pattern for Category F Types (Reference Non-optional)

For `String` arrays - use `moonbit_ref_array_make` with pointer storage:

### String Fix (Successfully Implemented)

**Pattern**: Reference Storage for Non-optional Strings

**Writing Side:**
```go
func (h *sliceHandler) createStringArrayWithMoonBit(...) {
    // Use moonbit_ref_array_make(numElements, 0)
    fn := wasmAdapter.GetFunction("moonbit_ref_array_make")
    results, err := fn.Call(ctx, uint64(numElements), uint64(0))
    
    // Encode strings and write pointers
    for i, val := range slice {
        results, _, err := h.elementHandler.Encode(ctx, wasmAdapter, val)
        stringPtr := uint32(results[0])
        // Write at arrayPtr + 8 + i*4
        offset := arrayPtr + 8 + uint32(i)*4
        wa.Memory().WriteUint32Le(offset, stringPtr)
    }
}
```

**Key insights for String:**
- Non-optional strings still use reference storage (unexpected!)
- Uses same `moonbit_ref_array_make` as optional reference types
- Demonstrates that reference storage isn't exclusive to optional types
- String is NOT considered "primitive" by MoonBit's type system

## Proven Fix Pattern for Category G Types (Reference Optional with Custom None)

For `String?` arrays - use `moonbit_ref_array_make` with type-specific None handling:

### String? Fix (Successfully Implemented)

**Pattern**: Reference Storage + Custom None Value

1. **Remove from primitive handling**: String? was incorrectly treated as primitive
2. **Add to reference array creation**: Use same infrastructure as Double?/Float?
3. **Custom None value handling**:

```go
// Writing Side
if utils.HasNil(val) {
    if h.typeInfo.ListElementType().Name() == "String?" {
        elementPtr = 0  // String? uses 0 as None
    } else {
        elementPtr = NoneSingletonPointer  // Others use 10248
    }
}

// Reading Side  
if (elemType.Name() == "String?" && ptr == 0) || ptr == NoneSingletonPointer {
    item = nil  // Handle both None patterns
}
```

**Key insights for String?:**
- Uses `moonbit_ref_array_make` like other Category B types
- **Critical discovery**: Different None value (0 vs 10248) within same category
- Proves that even within infrastructure categories, type-specific handling is needed
- WAT analysis showed exact None value pattern: `i32.const 0` for None elements

### Int16 Fix (Category C - Successfully Implemented)

1. **Use moonbit_int16_array_make instead of ptr2*_array:**
```go
case "Int16":
    // Use moonbit_int16_array_make
    arrayPtr, err = concreteWa.fnMakeArrayInt16.Call(ctx, uint64(numElements), 0)
    if err != nil {
        return 0, cln, fmt.Errorf("failed to call moonbit_int16_array_make: %w", err)
    }
    // Write data to the created array
    if len(arrayPtr) > 0 && arrayPtr[0] != 0 {
        int16ArrayPtr := uint32(arrayPtr[0])
        // Write individual int16 values at offset+8+i*2
        for i := uint32(0); i < numElements; i++ {
            val := binary.LittleEndian.Uint16(dataBuffer[i*2:])
            int16Addr := int16ArrayPtr + 8 + i*2
            wa.Memory().WriteUint16Le(int16Addr, val)
        }
        offset = int16ArrayPtr
        return offset, cln, nil
    }
```

2. **Pattern works because:**
   - Uses MoonBit's own array creation function
   - Proper GC integration
   - Correct memory layout (classID 80 for Int16)
   - Direct data writing at proper offsets

### Available MoonBit Array Functions

From `adapter.go`, these functions are available:
- `moonbit_int16_array_make` ✅ (working for Int16)
- `moonbit_i32_array_make` ✅ (used by Bool?, Byte?, Char?, Int16?)
- `moonbit_int64_array_make` ✅ (used by Int64 and Int?)
- `moonbit_float_array_make` ✅ (used by Double)
- `moonbit_float32_array_make` ✅ (used by Float)
- `moonbit_ref_array_make` ✅ (used by Double?/Float?/Int64?/UInt64?, String, String?)
- `moonbit_bytes_make` ✅ (used by Byte)

**Missing:** `moonbit_uint16_array_make` - UInt16 still needs manual handling

## Debugging Remaining Failures

### For Int16? (offset 32768 issue)
1. Analyze WAT for `test_fixedarray_output_int16_option_4`
2. Look for different None singleton or memory layout
3. Check if it uses different MoonBit runtime functions

### For Int?/UInt? (length mismatch issues)
1. These might need different `elemTypeSize` calculation
2. Check if they use 8-byte vs 4-byte storage
3. Verify header format differences

### For String? (similar to fixed types but different issues)
1. String arrays have different memory layout entirely
2. Check if they use UTF-16 encoding or other string-specific handling

## Code Quality: Magic Numbers → Constants Refactoring

**Completed**: All magic numbers in moonbit handlers have been replaced with well-named constants.

### Constants Categories

**ClassID Constants:**
```go
BoolByteCharClassID = 96   // Bool?, Byte?, Char?, Int16?, UInt arrays
Int64DoubleClassID  = 112  // Int64, UInt64, Double arrays  
RefArrayClassID     = 160  // Double?, Float?, Int64?, UInt64? arrays
StringBlockType     = 80   // Int16/UInt16 arrays
```

**None Singleton Values:**
```go
NoneSentinelUInt32   = 0xFFFFFFFF // Bool?, Byte?, Char?
NoneSingletonPointer = 10248      // 64-bit reference types
NoneValueInt16       = 32768      // Int16?
NoneValueInt         = 4294967296 // Int? (1 << 32)
```

**Memory Layout Constants:**
```go
MemoryBlockHeaderSize = 8    // Standard header size
MinValidMemoryOffset  = 1000 // Valid address threshold
MoonBitBoolSize      = 4    // MoonBit Bool size vs Go's 1 byte
StandardPtrSize      = 4    // Standard pointer size
Int64Size           = 8    // Int64/UInt64 size
```

**Benefits:**
- **Maintainability**: Clear meaning instead of scattered magic numbers
- **Debugging**: Easy to identify None values and classIDs
- **Consistency**: Centralized definitions prevent errors
- **Documentation**: Constants serve as inline documentation

**Usage Pattern:**
```go
// Before (unclear)
if value == 4294967296 {
    item = nil
}

// After (self-documenting)
if value == NoneValueInt {
    item = nil
}
```

## Key Insights from Successful Fixes

### Double? Fix (Category B)
1. **Real classID was 160, not 242** - documentation was wrong
2. **None singleton at 10248 is shared** - all Category B types use it
3. **moonbit_ref_array_make is the key** - manual memory allocation didn't work
4. **Element pointers vs values** - Category B stores pointers to Option objects
5. **WAT analysis revealed everything** - without it, would still be guessing

### Int16 Fix (Category C)
1. **Use MoonBit runtime functions** - `moonbit_int16_array_make` vs manual allocation
2. **ClassID 80 is correct** - StringBlockType for Int16/UInt16
3. **Direct data writing works** - write at offset+8+i*elemSize
4. **No ptr2*_array needed** - MoonBit functions handle GC integration
5. **Pattern applies to other primitives** - each type has its own `moonbit_*_array_make`

### Int16? Fix (Category D)
1. **ClassID 96 needs direct memory reading** - must add to memory reading condition
2. **32768 is inline None value** - stored directly, not dereferenced like 10248
3. **Uses moonbit_i32_array_make** - same as Bool?/Char?, not moonbit_int16_array_make
4. **Direct value storage pattern** - None=32768, Some values as 32-bit integers
5. **Critical insight**: WAT analysis revealed exact memory layout and None value

### Int? Fix (Category E)
1. **Uses moonbit_int64_array_make** - NOT moonbit_i32_array_make despite being 32-bit int
2. **64-bit storage for 32-bit values** - each Int? element takes 8 bytes in memory
3. **None value is 4294967296** - (1 << 32), completely different from other None values
4. **Sign extension required** - int32 values must be sign-extended to int64 for storage
5. **Memory layout**: Elements at offsets 8, 16, 24, 32 (8-byte spacing)
6. **WAT analysis critical**: Without it, would never have discovered the int64 storage pattern
7. **Pattern unique**: Only Int? uses this storage approach so far

### String Fix (Category F)
1. **Non-optional but uses reference storage** - Unexpected! String isn't "primitive"
2. **Uses moonbit_ref_array_make** - Same as optional reference types
3. **Stores pointers to string objects** - Not direct string data
4. **Routed to handler_slices.go** - Because String is not considered primitive
5. **WAT analysis revealed pattern**: `call $moonbit.ref_array_make` with string pointers
6. **Demonstrates reference storage ≠ optional**: Reference storage used for efficiency, not just optionality

### String? Fix (Category G)
1. **Initially misclassified as primitive** - Was treated like Int?/UInt? incorrectly
2. **Actually uses reference storage** - Like other Category B types but with twist
3. **Custom None value (0)** - Different from NoneSingletonPointer (10248) used by numeric types
4. **Type-specific None handling required** - Even within same infrastructure category
5. **WAT analysis showed None pattern**: `i32.const 0` for None elements
6. **Proves category complexity**: Same moonbit function, different None semantics
7. **Required dual None value logic**: Handle both 0 and 10248 in decoding

## Advanced Insights from String Array Fixes

### Type Classification Surprises

**Discovered**: MoonBit's type system classifications don't always match intuition:

1. **String is NOT primitive** - Routed to `handler_slices.go` not `handler_primitiveslices.go`
2. **Reference storage ≠ Optional** - Non-optional String uses reference storage for efficiency
3. **Infrastructure sharing with different semantics** - String? uses `moonbit_ref_array_make` but with different None value

**Routing Logic** (from `planner.go`):
```go
if !elemType.IsNullable() && elemType.IsPrimitive() {
    return p.NewPrimitiveSliceHandler(ti)  // → handler_primitiveslices.go
} else {
    return p.NewSliceHandler(ctx, ti)      // → handler_slices.go
}
```

**Key insight**: String fails `IsPrimitive()` check, so goes to slice handler.

### None Value Complexity

**Discovered**: None values are not uniform even within same infrastructure:

| Type Category | MoonBit Function | None Value | Storage |
|---------------|------------------|------------|----------|
| Bool?, Byte?, Char? | `moonbit_i32_array_make` | `0xFFFFFFFF` | Direct 32-bit |
| Int16? | `moonbit_i32_array_make` | `32768` | Direct 32-bit |
| Double?, Float?, Int64?, UInt64? | `moonbit_ref_array_make` | `10248` | Pointer to singleton |
| String? | `moonbit_ref_array_make` | `0` | Pointer (null) |
| Int? | `moonbit_int64_array_make` | `4294967296` | Direct 64-bit |

**Critical insight**: Even types using the same MoonBit function can have different None semantics.

### WAT Analysis Patterns

**Proven technique**: Look for these WAT patterns to understand array types:

1. **Function call patterns**:
   ```wat
   call $moonbit.i32_array_make     ; Direct storage, 32-bit elements
   call $moonbit.ref_array_make     ; Reference storage, pointer elements  
   call $moonbit.int64_array_make   ; Direct storage, 64-bit elements
   ```

2. **Initialization patterns**:
   ```wat
   i32.const 4 i32.const -1    ; Initialize with -1 (0xFFFFFFFF)
   i32.const 4 i32.const 0     ; Initialize with 0
   i32.const 4 i32.const 12768 ; Initialize with specific value
   ```

3. **Storage patterns**:
   ```wat
   i32.store offset=8  align=1     ; Store at 8-byte intervals
   i32.store offset=12 align=1     ; 4-byte pointer storage
   i64.store offset=16 align=1     ; 8-byte direct storage
   ```

### Debugging Strategy Evolution

**Level 1: Basic WAT Analysis**
- Find the function name
- Identify storage pattern (direct vs reference)
- Note initialization values

**Level 2: None Value Detection**
- Analyze what values represent None
- Check if None is stored inline or as pointer
- Identify singleton addresses vs direct values

**Level 3: Type System Understanding**
- Understand MoonBit's primitive vs non-primitive classification
- Learn routing logic in planner
- Recognize infrastructure sharing patterns

**Level 4: Cross-Type Pattern Recognition**
- Identify when types share infrastructure but differ in semantics
- Understand type-specific handling within shared functions
- Design conditional logic for mixed patterns

## Testing Strategy

1. **Test reading first:** Fix `TestFixedArrayOutput_*` before input tests
2. **Test in groups:** Fix all Category types together, but expect type-specific variations
3. **Verify working tests still work:** Don't break Bool?, Byte?, Char?
4. **Use specific test runs:** `go test -run ^TestFixedArrayOutput_double_option_4`
5. **Test both String and String?:** Verify non-optional and optional variants separately
6. **Check None value handling:** Test arrays with None elements specifically

## Future Type Implementation Guide

### Step 1: WAT Analysis
1. Find the `test_fixedarray_output_*` function
2. Identify MoonBit function used (`moonbit.*_array_make`)
3. Note initialization value and storage offsets
4. Analyze None value patterns

### Step 2: Classification
1. Determine if type is "primitive" per MoonBit
2. Check routing destination (primitiveslices vs slices)
3. Identify category based on MoonBit function used

### Step 3: Implementation
1. Add to appropriate creation function condition
2. Handle type-specific None values if needed
3. Update both encoding and decoding logic
4. Add constants for any new magic numbers

### Step 4: Testing
1. Test specific failing case first
2. Run full type suite (0, 1, 2, 3, 4 variants)
3. Verify no regressions in existing types
4. Test both input and output directions

This guide now captures the complete debugging methodology from basic WAT analysis through complex type-specific None value handling, based on successful fixes across all major array categories.