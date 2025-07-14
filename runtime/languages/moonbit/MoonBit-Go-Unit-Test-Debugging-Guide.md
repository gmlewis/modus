# MoonBit Go Unit Test Debugging Guide

## Overview

This guide provides proven methodologies for debugging and fixing failing Go unit tests in the MoonBit runtime integration, based on successful resolution of complex array type issues.

## Core Debugging Philosophy

### 1. WASM/WAT is Ground Truth

**Fundamental principle**: The WAT (WebAssembly Text) output is the authoritative source of truth, not Go code assumptions.

**Why this matters**:
- Go code makes assumptions about memory layout
- MoonBit compiler generates the actual memory structures
- WAT shows exactly what MoonBit produces
- Mismatches between Go assumptions and WAT reality cause test failures

**Always start here**:
```bash
cd /app/runtime/languages/moonbit/testdata
grep -A 30 "test_array_output_byte_" build/testdata.wat
```

### 2. Test Failure Pattern Analysis

**Group test failures by pattern**:
1. **Type-specific failures**: Only one type fails (e.g., Array[Byte])
2. **Size-specific failures**: Only certain sizes fail (e.g., 0-2 pass, 3-4 fail)
3. **Direction-specific failures**: Only input or output fails
4. **Infrastructure failures**: Whole categories fail

**Example analysis**:
```bash
# Check which Array[Byte] tests are failing
go test -run TestArrayOutput_byte_[0-4] -v
# Look for patterns in failure messages
```

### 3. Size-Mismatched Type Identification

**Key insight**: Types where Go and MoonBit representations differ in size need special handling.

**Size-mismatched types**:
- `Bool`: Go bool (1 byte) vs MoonBit Bool (4 bytes)
- `Byte`: Go uint8 (1 byte) vs MoonBit Byte (context-dependent)
- `Char`: Go rune (4 bytes) vs MoonBit Char (context-dependent)

**Size-matched types** (usually work with standard infrastructure):
- `Int16`, `Int64`, `Float`, `Double`: Direct correspondence

## Systematic Debugging Approach

### Phase 1: Problem Classification

#### Step 1: Identify the Failing Type

**Run targeted tests**:
```bash
# Test specific type
NO_COLOR=1 go test -timeout 30s -tags integration -run "TestArrayOutput_byte_[0-4]$" -v

# Compare with working type
NO_COLOR=1 go test -timeout 30s -tags integration -run "TestArrayOutput_int_[0-4]$" -v
```

**Classification questions**:
- Is this type "primitive" per MoonBit's definition?
- Does it have size mismatch between Go and MoonBit?
- Are there specialized MoonBit functions for this type?

#### Step 2: Analyze Failure Patterns

**Common failure patterns**:
1. **GC unreachable errors**: Usually manual memory allocation issues
2. **Wrong values returned**: Decode logic reading wrong memory layout
3. **Memory access errors**: Offset calculation problems
4. **Type assertion failures**: Wrong type conversion

**Example analysis**:
```
expected [1 2 3], got [1 13 0]  // Decode reading wrong offsets
vs
wasm error: unreachable         // GC issue with manual allocation
```

### Phase 2: WAT Analysis

#### Step 3: Examine WAT Output

**Find the specific test function**:
```bash
grep -A 50 "test_array_output_byte_2" build/testdata.wat
```

**Look for key patterns**:
1. **MoonBit function calls**: `moonbit.bytes_make`, `moonbit.i32_array_make`, etc.
2. **Memory allocation**: `moonbit.gc.malloc`, `moonbit.cabi_realloc`
3. **Type info constants**: `i32.const 1573120`, `i32.const 96`
4. **Data writing patterns**: `i32.store8`, `i32.store`, `i64.store`

**Extract critical information**:
- What MoonBit function creates the array?
- What values are used for initialization?
- How are elements stored (1-byte, 4-byte, 8-byte)?
- Are there conversion functions called?

#### Step 4: Identify MoonBit's Approach

**Check for specialized functions**:
```bash
grep -n "moonbit_.*_make" /app/runtime/languages/moonbit/adapter.go
grep -n "moonbit_.*_to_" /app/runtime/languages/moonbit/adapter.go
```

**Common MoonBit array functions**:
- `moonbit_bytes_make` + `moonbit_bytes_to_array` (for Array[Byte])
- `moonbit_i32_array_make` (for Bool?, Byte?, Char?)
- `moonbit_ref_array_make` (for reference types)
- `moonbit_int16_array_make` (for Int16)
- `moonbit_int64_array_make` (for Int64, Int?)

### Phase 3: Implementation Strategy

#### Step 5: Choose Implementation Approach

**Decision tree**:
1. **If MoonBit has specialized functions**: Use them (best approach)
2. **If type fits existing category**: Extend existing pattern
3. **If completely new pattern**: Create new category

**Example - Array[Byte] decision**:
```
WAT shows: moonbit.bytes_make + moonbit.bytes_to_array
→ Use specialized MoonBit functions
→ Don't force into fixed array infrastructure
```

#### Step 6: Implement Following MoonBit's Pattern

**For specialized function approach**:
```go
// Step 1: Create using MoonBit's function
fn := wasmAdapter.GetFunction("moonbit_bytes_make")
results, err := fn.Call(ctx, uint64(numElements), uint64(0))
bytesPtr := uint32(results[0])

// Step 2: Write data as MoonBit expects
for i, val := range slice {
    byteVal := byte(val)
    offset := bytesPtr + 8 + uint32(i)  // MoonBit's layout
    wa.Memory().Write(offset, []byte{byteVal})
}

// Step 3: Convert using MoonBit's converter
converterFn := wasmAdapter.GetFunction("moonbit_bytes_to_array")
results, err := converterFn.Call(ctx, uint64(bytesPtr))
```

### Phase 4: Decode Logic Implementation

#### Step 7: Understand the Memory Layout

**Add debug output to understand structure**:
```go
if elemType.Name() == "Byte" {
    // Read first 32 bytes to understand structure
    if debugBytes, ok := wa.Memory().Read(offset, 32); ok {
        fmt.Printf("DEBUG: Array[Byte] decode - first 32 bytes: %v\n", debugBytes)
        // Parse key fields
        typeInfo := binary.LittleEndian.Uint32(debugBytes[4:8])
        length := binary.LittleEndian.Uint32(debugBytes[8:12])
        dataPtr := binary.LittleEndian.Uint32(debugBytes[12:16])
        fmt.Printf("DEBUG: typeInfo=%d, length=%d, dataPtr=%d\n", typeInfo, length, dataPtr)
    }
}
```

#### Step 8: Implement Detection Logic

**Pattern recognition for special layouts**:
```go
if typeInfo == 1573120 { // Array wrapper
    dataPtr := readUint32(offset + 12)
    if elemType.Name() == "Byte" {
        bytesTypeInfo := readUint32(dataPtr + 4)
        if (bytesTypeInfo & 0xFF000000) == 0x40000000 {
            // This is a Bytes object, adjust to actual data
            offset = dataPtr + 8
            isBytesData = true
        }
    }
}
```

#### Step 9: Implement Specialized Reading

**Handle the detected layout**:
```go
if isBytesData && elemType.Name() == "Byte" {
    // Read raw byte data directly
    dataBytes, ok := wa.Memory().Read(offset, arrayLength)
    if !ok {
        return nil, fmt.Errorf("failed to read byte data")
    }
    
    // Convert to Go slice
    items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(arrayLength), int(arrayLength))
    for i := uint32(0); i < arrayLength; i++ {
        val := h.converter.Decode(uint64(dataBytes[i]))
        items.Index(int(i)).Set(reflect.ValueOf(val))
    }
    return items.Interface(), nil
}
```

### Phase 5: Testing and Verification

#### Step 10: Test Incrementally

**Test one size at a time**:
```bash
# Start with simplest case
go test -run TestArrayOutput_byte_0 -v

# Progress through sizes
go test -run TestArrayOutput_byte_1 -v
go test -run TestArrayOutput_byte_2 -v
```

**Look for patterns in success/failure**:
- Do all sizes work or only some?
- Are there edge cases (empty arrays)?
- Do input and output both work?

#### Step 11: Clean Up and Verify

**Remove debug output**:
```go
// Replace debug prints with comments
// fmt.Printf("DEBUG: ...") → // Successfully converted Bytes to Array[Byte]
```

**Test full suite**:
```bash
# Test all sizes
go test -run TestArrayOutput_byte_[0-4] -v

# Verify no regression
go test -run TestArrayOutput_int_[0-4] -v
```

## Common Pitfalls and Solutions

### Pitfall 1: Manual Memory Allocation

**❌ Don't do this**:
```go
// Manual allocation often causes GC issues
wrapperPtr, err := wa.allocateAndPinMemory(ctx, 16, 0)
wa.Memory().WriteUint32Le(wrapperPtr+4, 1573120)
```

**✅ Do this instead**:
```go
// Use MoonBit's own functions
fn := wasmAdapter.GetFunction("moonbit_bytes_make")
results, err := fn.Call(ctx, uint64(numElements), uint64(0))
```

### Pitfall 2: Ignoring WAT Analysis

**❌ Don't do this**:
```go
// Assuming all arrays work the same way
return h.createStandardArray(ctx, wa, slice, numElements)
```

**✅ Do this instead**:
```go
// Check WAT output first, then implement MoonBit's approach
switch elemType.Name() {
case "Byte":
    return h.createByteDataArray(ctx, wa, wasmAdapter, slice, numElements)
case "Bool":
    return h.createBoolDataArray(ctx, wa, wasmAdapter, slice, numElements)
}
```

### Pitfall 3: Wrong Size Assumptions

**❌ Don't do this**:
```go
// Assuming Go and MoonBit sizes match
for i, val := range slice {
    offset := arrayPtr + 8 + uint32(i) // Wrong for 4-byte MoonBit types
    wa.Memory().Write(offset, []byte{byte(val)})
}
```

**✅ Do this instead**:
```go
// Use MoonBit's actual size requirements
for i, val := range slice {
    if elemType.Name() == "Bool" {
        offset := arrayPtr + 8 + uint32(i)*4 // 4-byte MoonBit Bool
        wa.Memory().WriteUint32Le(offset, boolToUint32(val))
    } else {
        offset := arrayPtr + 8 + uint32(i) // 1-byte for actual bytes
        wa.Memory().Write(offset, []byte{byte(val)})
    }
}
```

## Success Metrics

### Array[Byte] Success Case

**Before**: All tests failing
```
--- FAIL: TestArrayOutput_byte_0 (0.01s)
--- FAIL: TestArrayOutput_byte_1 (0.00s)
--- FAIL: TestArrayOutput_byte_2 (0.00s)
--- FAIL: TestArrayOutput_byte_3 (0.00s)
--- FAIL: TestArrayOutput_byte_4 (0.00s)
```

**After**: All tests passing
```
--- PASS: TestArrayOutput_byte_0 (0.01s)
--- PASS: TestArrayOutput_byte_1 (0.00s)
--- PASS: TestArrayOutput_byte_2 (0.00s)
--- PASS: TestArrayOutput_byte_3 (0.00s)
--- PASS: TestArrayOutput_byte_4 (0.00s)
```

**Key metrics**:
- 100% success rate for target type
- No regression in other types
- Clean, maintainable code
- Proper integration with MoonBit runtime

## Advanced Debugging Techniques

### Memory Layout Visualization

**Add hex dump functionality**:
```go
func debugMemoryLayout(wa wasmMemoryReader, offset uint32, size uint32, label string) {
    if data, ok := wa.Memory().Read(offset, size); ok {
        fmt.Printf("DEBUG: %s at offset %d:\n", label, offset)
        for i := 0; i < len(data); i += 4 {
            if i+4 <= len(data) {
                val := binary.LittleEndian.Uint32(data[i:i+4])
                fmt.Printf("  [%d]: 0x%08X (%d)\n", i, val, val)
            }
        }
    }
}
```

### Type Info Pattern Recognition

**Build pattern database**:
```go
var typeInfoPatterns = map[uint32]string{
    1573120:   "Array wrapper",
    0x40000000: "Bytes object (length in low bits)",
    0x60000000: "ClassID 96 pattern",
    10248:      "None singleton pointer",
}

func identifyTypeInfo(typeInfo uint32) string {
    if pattern, exists := typeInfoPatterns[typeInfo]; exists {
        return pattern
    }
    for mask, pattern := range typeInfoPatterns {
        if (typeInfo & mask) == mask {
            return fmt.Sprintf("%s (variant: 0x%08X)", pattern, typeInfo)
        }
    }
    return fmt.Sprintf("Unknown type info: 0x%08X", typeInfo)
}
```

### Function Call Tracing

**Track MoonBit function usage**:
```go
func traceMoonBitCall(functionName string, args []uint64, results []uint64) {
    fmt.Printf("TRACE: %s(args: %v) → results: %v\n", functionName, args, results)
}

// Use in implementation
results, err := fn.Call(ctx, uint64(numElements), uint64(0))
traceMoonBitCall("moonbit_bytes_make", []uint64{uint64(numElements), 0}, results)
```

## Future Development Guidelines

### 1. Type Classification System

**Maintain type category database**:
```go
type TypeCategory struct {
    Name        string
    MoonBitFunc string
    Storage     string // "direct", "reference", "specialized"
    NoneValue   interface{}
    Status      string // "working", "failing", "not_implemented"
}

var typeCategories = map[string]TypeCategory{
    "Array[Byte]": {"Array[Byte]", "moonbit_bytes_make+fnBytes2Array", "specialized", nil, "working"},
    "Array[Bool]": {"Array[Bool]", "moonbit_i32_array_make", "direct", 0xFFFFFFFF, "failing"},
    // ... add more as discovered
}
```

### 2. WAT Analysis Automation

**Create analysis scripts**:
```bash
#!/bin/bash
# analyze_wat_function.sh
FUNCTION_NAME=$1
grep -A 50 "$FUNCTION_NAME" build/testdata.wat | grep -E "(moonbit\.|i32\.const|i32\.store|i64\.store)"
```

### 3. Test Coverage Tracking

**Maintain test status matrix**:
```
Type         | Size 0 | Size 1 | Size 2 | Size 3 | Size 4 | Status
-------------|--------|--------|--------|--------|--------|--------
Array[Byte]  |   ✅   |   ✅   |   ✅   |   ✅   |   ✅   | Fixed
Array[Bool]  |   ❌   |   ❌   |   ❌   |   ❌   |   ❌   | Failing
Array[Int16] |   ✅   |   ✅   |   ✅   |   ✅   |   ✅   | Working
```

## Conclusion

This guide provides a systematic approach to debugging MoonBit Go unit tests, with emphasis on:

1. **WAT analysis as ground truth**
2. **Understanding MoonBit's type system**
3. **Using MoonBit's own functions rather than manual implementation**
4. **Incremental testing and verification**
5. **Pattern recognition across similar types**

The Array[Byte] success case demonstrates that even complex failing tests can be resolved by following MoonBit's intended patterns rather than fighting against them.

For future work, always start with WAT analysis, identify the MoonBit approach, and implement using the language's own runtime functions for optimal integration and reliability.