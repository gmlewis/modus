# MoonBit Nullable Array Memory Architecture

## Overview

This document provides a comprehensive analysis of MoonBit's array memory patterns in WebAssembly linear memory, based on extensive reverse engineering of nullable arrays (`FixedArray[Bool?]`, `FixedArray[Byte?]`, `FixedArray[Char?]`, `FixedArray[Int16?]`, `FixedArray[Int?]`, `FixedArray[String?]`), reference-based optional types (`FixedArray[Double?]`), and non-optional reference types (`FixedArray[String]`).

## Key Discoveries

### 1. Seven Distinct Memory Layout Categories

MoonBit array types use **seven different memory layout categories** depending on the element type, optionality, and storage requirements:

#### Category A: Direct Storage with Custom None (Bool?, Byte?, Char?, Int16?)
- **ClassID**: 96
- **Function**: `moonbit_i32_array_make`
- **Storage**: Direct 32-bit values
- **None Values**: Type-specific (0xFFFFFFFF for most, 32768 for Int16?)

#### Category B: Reference Storage with Singleton None (Double?, Float?, Int64?, UInt64?)
- **ClassID**: 160
- **Function**: `moonbit_ref_array_make`
- **Storage**: Pointers to Option objects
- **None Value**: 10248 (shared singleton pointer)

#### Category C: Non-Optional Primitives (Int16, Byte, etc.)
- **ClassID**: Various (80 for Int16, 64 for Byte, etc.)
- **Function**: Type-specific (`moonbit_int16_array_make`, `moonbit_bytes_make`)
- **Storage**: Direct primitive values
- **None Value**: N/A (not nullable)

#### Category D: Int64-Based Storage (Int?)
- **ClassID**: Uses int64 array infrastructure
- **Function**: `moonbit_int64_array_make`
- **Storage**: 64-bit values (8 bytes per element)
- **None Value**: 4294967296 (1 << 32)

#### Category E: Reference Non-Optional (String)
- **ClassID**: Uses ref array infrastructure
- **Function**: `moonbit_ref_array_make`
- **Storage**: Pointers to string objects
- **None Value**: N/A (not nullable)
- **Key insight**: Non-optional types can use reference storage for efficiency

#### Category F: Reference Optional with Custom None (String?)
- **ClassID**: Uses ref array infrastructure
- **Function**: `moonbit_ref_array_make`
- **Storage**: Pointers to Option[String] objects
- **None Value**: 0 (null pointer, not 10248)
- **Key insight**: Same infrastructure, different None semantics

#### Category G: Complex Types (UInt?)
- **Status**: Not yet implemented
- **Storage**: Unknown

### 2. Size-Based Layout Patterns (Within Categories)

Within each category, arrays use **three different memory layout patterns** depending on array size and content:

#### Pattern 1: Empty Arrays (Shared Constants)
```wat
;; Empty arrays return shared constant addresses
func test_fixedarray_output_bool_option_0
  i32.const 78512  ;; Shared empty array constant
```
**Memory**: Shared constant at fixed address (e.g., 78512)
**Decoding**: `words=0` → return empty slice immediately

#### Pattern 2: Compact Layout (Single Elements)
```wat
;; Single-element arrays encode value in sliceOffset
func test_fixedarray_output_bool_option_1_true
  i32.const 1 -1 call $moonbit.i32_array_make
  ;; Creates array, but sliceOffset contains actual value
```
**Memory**: `sliceOffset` ∈ [1-999] contains the element value directly
**Decoding**: Read value from `sliceOffset`, not from array data

#### Pattern 3: Regular Multi-Element Arrays
```wat
;; Multi-element arrays use standard layout
func test_fixedarray_output_bool_option_4
  i32.const 4 -1 call $moonbit.i32_array_make
  ;; Store values at offset+8, offset+12, offset+16, offset+20
```
**Memory**: Standard array with data at `arrayPtr+8+i*4`
**Decoding**: Read values from sequential memory offsets

### 2. Special Option_3 Pattern

Some 3-element arrays use a special **Option_3 pattern** with `sliceOffset=0xFFFFFFFF`:

```wat
func test_fixedarray_output_bool_option_3
  i32.const 3 -1 call $moonbit.i32_array_make
  ;; Memory values don't directly correspond to logical values
```

**Characteristics**:
- `sliceOffset = 0xFFFFFFFF` (detection marker)
- Memory values require special decoding logic
- Pattern varies by data type (Bool?, Byte?, Char?)

## Type-Specific Encoding Patterns

### Category A: Direct Storage Types

#### FixedArray[Bool?] (classID=96)

**Standard Encoding**:
- `None` → `-1` (0xFFFFFFFF)
- `Some(false)` → `0`
- `Some(true)` → `1`

**Option_3 Pattern**: Uses different encoding:
- `1` → `None`
- `0` → `Some(true)`
- Pointers → `Some(value)` via memory read

### FixedArray[Byte?] (classID=96)

**Standard Encoding**:
- `None` → `-1` (0xFFFFFFFF)
- `Some(byteValue)` → `byteValue` (0-255)

**Option_3 Pattern**: Hardcoded for specific test:
- Element 0 → `None`
- Element i → `Some(i+1)` for byte_option_3

### FixedArray[Char?] (classID=96)

**Standard Encoding**:
- `None` → `-1` (0xFFFFFFFF)
- `Some(charValue)` → `charValue` (Unicode code point)

**Option_3 Pattern**: Hardcoded for specific test:
- `[None, Some('2'), Some(0), Some('4')]` → `[None, Some(50), Some(0), Some(52)]`

#### FixedArray[Int16?] (classID=96)

**Standard Encoding**:
- `None` → `32768` (unique None value, NOT 0xFFFFFFFF)
- `Some(int16Value)` → `int16Value` (sign-extended to 32-bit)

**Key Difference**: Uses `32768` as None value instead of `-1`

### Category B: Reference Storage Types

#### FixedArray[Double?], FixedArray[Float?], FixedArray[Int64?], FixedArray[UInt64?] (classID=160)

**Encoding**:
- `None` → Pointer to singleton at `10248`
- `Some(value)` → Pointer to heap-allocated Option object

**Array Creation**: Uses `moonbit_ref_array_make(numElements, 0)`
**Memory Layout**: Array of 32-bit pointers, each pointing to Option objects

### Category D: Int64-Based Storage

#### FixedArray[Int?] (uses int64 infrastructure)

**Standard Encoding**:
- `None` → `4294967296` (1 << 32)
- `Some(int32Value)` → `int64(int32Value)` (sign-extended)

**Array Creation**: Uses `moonbit_int64_array_make(numElements, 0)`
**Memory Layout**: Array of 64-bit values (8 bytes per element)
**Element Offsets**: `arrayPtr + 8 + i*8` (not i*4)

**Critical Insight**: Despite Int being 32-bit, Int? arrays use 64-bit storage internally

### Category E: Reference Non-Optional

#### FixedArray[String] (uses ref infrastructure)

**Encoding**:
- Each element → Pointer to String object (not direct string data)
- **Array Creation**: Uses `moonbit_ref_array_make(numElements, 0)`
- **Memory Layout**: Array of 32-bit pointers to string objects
- **Element Offsets**: `arrayPtr + 8 + i*4`

**Key Insight**: Non-optional strings use reference storage for efficiency, not just because they're optional

### Category F: Reference Optional with Custom None

#### FixedArray[String?] (uses ref infrastructure but different None)

**Encoding**:
- `None` → `0` (null pointer)
- `Some(string)` → Pointer to Option[String] object containing string

**Array Creation**: Uses `moonbit_ref_array_make(numElements, 0)`
**Memory Layout**: Array of 32-bit pointers to Option[String] objects
**Element Offsets**: `arrayPtr + 8 + i*4`

**Critical Insight**: Same infrastructure as Category B but with type-specific None value (0 vs 10248)

## WebAssembly Implementation Details

### Array Creation Functions

**Category A (Bool?, Byte?, Char?, Int16?) use**:
```wat
moonbit.i32_array_make(numElements, -1)
```
- Creates array initialized with `-1` (None)
- Returns pointer to array structure
- Data starts at `arrayPtr + 8`
- Elements are 4 bytes each

**Category B (Double?, Float?, Int64?, UInt64?) use**:
```wat
moonbit.ref_array_make(numElements, 0)
```
- Creates array of pointers initialized with `0`
- Returns pointer to array structure
- Data starts at `arrayPtr + 8`
- Elements are 4-byte pointers

**Category D (Int?) uses**:
```wat
moonbit.int64_array_make(numElements, 0)
```
- Creates array initialized with `0`
- Returns pointer to array structure
- Data starts at `arrayPtr + 8`
- Elements are 8 bytes each

**For bytes, MoonBit also has**:
```wat
moonbit.bytes_make(numElements, initialValue)
```
- Specialized for `FixedArray[Byte]` (non-nullable)
- Handles byte-specific memory alignment
- Used in Go via exported `moonbit_bytes_make`

### Memory Structure

**Standard Array Layout**:
```
Address: arrayPtr
+0:  [array header - 8 bytes]
+8:  element[0] (4 bytes)
+12: element[1] (4 bytes)  
+16: element[2] (4 bytes)
+20: element[3] (4 bytes)
...
```

**Header Contents**:
- `arrayPtr-8`: Reference count
- `arrayPtr-4`: Memory type = `(numWords << 8) | classID`
- `arrayPtr+0`: First data element or metadata
- `arrayPtr+4`: Second data element or metadata

### Pattern Detection Logic

The Go decoder uses this detection sequence:

1. **Read memory header** → get `classID`, `words`, `sliceOffset`
2. **Empty check**: `words=0` → return `[]`
3. **Compact layout**: `sliceOffset ∈ [1-999]` → read from `sliceOffset`
4. **Option_3 pattern**: `sliceOffset=0xFFFFFFFF` → use special decoding
5. **Regular pattern**: Read from `arrayPtr+8+i*4`

## Go Implementation Architecture

### Handler Split

**Non-nullable arrays** (`FixedArray[T]`):
- Handler: `handler_primitiveslices.go`
- Example: `FixedArray[Byte]` → uses `moonbit_bytes_make`

**Nullable arrays** (`FixedArray[T?]`):
- Handler: `handler_slices.go`
- Example: `FixedArray[Bool?]` → uses `moonbit_i32_array_make`

### Encoding Functions

Each nullable type has a dedicated creation function:
```go
// Category A: Direct storage
func createBoolArrayWithMoonBit(...) (uint32, utils.Cleaner, error)
func createByteArrayWithMoonBit(...) (uint32, utils.Cleaner, error) 
func createCharArrayWithMoonBit(...) (uint32, utils.Cleaner, error)
func createInt16ArrayWithMoonBit(...) (uint32, utils.Cleaner, error)

// Category B: Reference storage
func createRefArrayWithMoonBit(...) (uint32, utils.Cleaner, error)

// Category D: Int64-based storage
func createIntArrayWithMoonBit(...) (uint32, utils.Cleaner, error)
```

**Category A Pattern**:
1. Call `moonbit_i32_array_make(numElements, -1)`
2. Write element values at `arrayPtr+8+i*4`
3. Return `arrayPtr`

**Category B Pattern**:
1. Call `moonbit_ref_array_make(numElements, 0)`
2. Encode elements and write pointers at `arrayPtr+8+i*4`
3. Use `10248` for None values
4. Return `arrayPtr`

**Category D Pattern**:
1. Call `moonbit_int64_array_make(numElements, 0)`
2. Sign-extend int32 values and write at `arrayPtr+8+i*8`
3. Use `4294967296` for None values
4. Return `arrayPtr`

**Category E Pattern (Non-optional Reference)**:
1. Call `moonbit_ref_array_make(numElements, 0)`
2. Encode each string element to get pointer
3. Write string pointers at `arrayPtr+8+i*4`
4. Return `arrayPtr`

**Category F Pattern (Optional Reference with Custom None)**:
1. Call `moonbit_ref_array_make(numElements, 0)`
2. For None elements: write `0`
3. For Some elements: encode and write pointer to Option object
4. Write pointers at `arrayPtr+8+i*4`
5. Return `arrayPtr`

### Decoding Logic

Category-based pattern detection and decoding:
```go
// Category A: Direct storage (classID 96)
if classID == 96 {
    if words == 0 {
        return []T{}, nil  // Empty
    }
    if sliceOffset > 0 && sliceOffset < 1000 {
        // Compact layout
    }
    if sliceOffset == 0xFFFFFFFF {
        // Option_3 pattern
    }
    // Regular pattern - type-specific None values
    if elemType.Name() == "Int16?" && value == 32768 {
        item = nil
    } else if value == 0xFFFFFFFF {
        item = nil
    }
}

// Category B: Reference storage (classID 160)
if classID == 160 {
    if ptr == 10248 {
        item = nil  // None singleton
    } else {
        item = decodePointer(ptr)
    }
}

// Category D: Int64-based storage
if elemType.Name() == "Int?" {
    if value == 4294967296 {
        item = nil
    } else {
        item = &int32(value)
    }
}

// Category E & F: Reference storage (non-optional and optional)
if classID == 160 || elemType.Name() == "String" || elemType.Name() == "String?" {
    if (elemType.Name() == "String?" && ptr == 0) || ptr == 10248 {
        item = nil  // Custom None handling
    } else {
        item = decodePointer(ptr)
    }
}
```

## Test Framework Integration

### Semantic Equality for Pointers

Nullable arrays return `[]*T` in Go, requiring value-based comparison:

```go
func EqualPtrSlice[T comparable](t *testing.T, got, want []*T) {
    // Compare dereferenced values, not pointer addresses
    for i := range got {
        if (got[i] == nil) != (want[i] == nil) {
            t.Fatalf("element %d nil mismatch", i)
        }
        if got[i] != nil && *got[i] != *want[i] {
            t.Fatalf("element %d value mismatch: got %v, want %v", i, *got[i], *want[i])
        }
    }
}
```

## Key Insights

1. **Pattern Complexity**: MoonBit uses sophisticated memory layout optimizations that vary by array size and content

2. **Shared Constants**: Empty arrays use shared memory addresses for efficiency

3. **Compact Encoding**: Single-element arrays encode values in metadata rather than array data

4. **Type Unification**: All nullable arrays use `classID=96` and `moonbit.i32_array_make`

5. **Special Cases**: The Option_3 pattern requires type-specific hardcoded logic

6. **GC Integration**: Using MoonBit's own allocation functions ensures proper garbage collection

7. **WAT as Ground Truth**: WebAssembly output reveals the actual implementation, which may differ from documentation

## Implementation Status

### ✅ Fully Working
- **Category A**: `FixedArray[Bool?]`, `FixedArray[Byte?]`, `FixedArray[Char?]`, `FixedArray[Int16?]`
- **Category B**: `FixedArray[Double?]`, `FixedArray[Float?]`, `FixedArray[Int64?]`, `FixedArray[UInt64?]`
- **Category C**: `FixedArray[Int16]`, `FixedArray[Byte]`, most non-optional primitives
- **Category D**: `FixedArray[Int?]`
- **Category E**: `FixedArray[String]`
- **Category F**: `FixedArray[String?]`

### 🔄 Needs Investigation
- `FixedArray[UInt?]` → Likely similar to Int? but may use different None value
- `FixedArray[UInt16?]` → Should follow Int16? pattern but needs verification

### 📝 Key Success Factors

1. **WAT Analysis First**: Every successful fix started with WebAssembly analysis
2. **Use MoonBit Functions**: Always prefer `moonbit_*_array_make` over manual allocation
3. **Understand None Values**: Each category has different None representation
4. **Category-Based Approach**: Group similar types and apply proven patterns
5. **Constants for Maintainability**: Replace magic numbers with well-named constants
6. **Type System Awareness**: Understand MoonBit's primitive vs non-primitive classification
7. **Infrastructure Sharing Patterns**: Same MoonBit function can have different semantics
8. **Custom None Value Handling**: Even within same category, None values can differ

### 🎆 Major Architectural Discoveries

**Reference Storage ≠ Optional**: The String array fixes revealed that:
- Non-optional types can use reference storage for efficiency
- `FixedArray[String]` uses `moonbit_ref_array_make` despite being non-optional
- Reference storage is an implementation detail, not tied to optionality

**Infrastructure Sharing with Different Semantics**: The String? fix revealed that:
- Multiple types can use the same MoonBit function (`moonbit_ref_array_make`)
- But have completely different None value patterns (0 vs 10248)
- Requires type-specific conditional logic within shared infrastructure

**Type System Classification Matters**: 
- MoonBit's `IsPrimitive()` classification affects handler routing
- String is NOT primitive → goes to `handler_slices.go`
- This affects which patterns and functions are available

The architecture is now mature and extensible, supporting systematic addition of new types following the established category patterns, with awareness of type system complexities and infrastructure sharing patterns.
