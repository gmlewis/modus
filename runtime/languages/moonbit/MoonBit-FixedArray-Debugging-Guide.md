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

**Category E: Other Types (Still Failing)**
- `Int?`, `UInt?`, `String?` - Various issues
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

For primitive types like `Int16` - use `moonbit_*_array_make` functions:

### Int16 Fix (Successfully Implemented)

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
- `moonbit_int16_array_make` ✅ (working)
- `moonbit_i32_array_make` ✅ (used by Bool?)
- `moonbit_float_array_make` ✅ (used by Double)
- `moonbit_float32_array_make` ✅ (used by Float)
- `moonbit_int64_array_make` ✅ (used by Int64)
- `moonbit_ref_array_make` ✅ (used by Double?/Float?/Int64?/UInt64?)
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

## Testing Strategy

1. **Test reading first:** Fix `TestFixedArrayOutput_*` before input tests
2. **Test in groups:** Fix all Category B types together (Double?, Float?, Int64?, UInt64?)
3. **Verify working tests still work:** Don't break Bool?, Byte?, Char?
4. **Use specific test runs:** `go test -run ^TestFixedArrayOutput_double_option_4`

This guide captures the most effective techniques learned from successfully fixing the Double? test and should significantly speed up debugging of remaining failures.