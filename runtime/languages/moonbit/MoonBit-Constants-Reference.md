# MoonBit Handler Constants Reference

Quick reference for all constants used in MoonBit array handling.

## ClassID Constants

| Constant | Value | Used For | Array Function |
|----------|-------|----------|----------------|
| `TupleBlockType` | 0 | General tuples, some shared arrays | - |
| `StringBlockType` | 80 | Int16, UInt16 arrays | `moonbit_int16_array_make` |
| `BoolByteCharClassID` | 96 | Bool?, Byte?, Char?, Int16?, UInt arrays | `moonbit_i32_array_make` |
| `Int64DoubleClassID` | 112 | Int64, UInt64, Double arrays | Various | 
| `RefArrayClassID` | 160 | Double?, Float?, Int64?, UInt64? arrays | `moonbit_ref_array_make` |
| `FixedArrayPrimitiveBlockType` | 241 | General primitive arrays | Manual allocation |
| `PtrArrayBlockType` | 242 | Pointer arrays | Manual allocation |
| `FixedArrayByteBlockType` | 246 | Byte arrays | `moonbit_bytes_make` |

## None Values for Optional Types

| Constant | Value | Hex | Used For | Storage Type |
|----------|-------|-----|----------|-------------|
| `NoneSentinelUInt32` | 4294967295 | 0xFFFFFFFF | Bool?, Byte?, Char? | Direct 32-bit |
| `NoneValueInt16` | 32768 | 0x8000 | Int16? | Direct 32-bit |
| `NoneValueInt` | 4294967296 | 0x100000000 | Int? | Direct 64-bit |
| `NoneSingletonPointer` | 10248 | 0x2808 | Double?, Float?, Int64?, UInt64? | Pointer to singleton |

## Memory Layout Constants

| Constant | Value | Purpose |
|----------|-------|---------|
| `MemoryBlockHeaderSize` | 8 | Standard memory block header size |
| `MemoryBlockHeaderSizeLg` | 16 | Extended header size for some arrays |
| `MinValidMemoryOffset` | 1000 | Minimum valid memory address threshold |
| `EmptyArrayMarker1` | 4294967295 | Special marker for empty arrays |
| `EmptyArrayMarker2` | 1610612736 | Special marker for empty arrays |

## Type Size Constants

| Constant | Value | Purpose |
|----------|-------|---------|
| `MoonBitBoolSize` | 4 | MoonBit Bool size (vs Go's 1 byte) |
| `MoonBitCharSize` | 4 | MoonBit Char size (vs Go's 2 bytes) |
| `StandardPtrSize` | 4 | Standard pointer size in MoonBit |
| `Int64Size` | 8 | Size of Int64/UInt64 types |
| `StringCharSize` | 2 | UTF-16 character size in strings |

## Bit Manipulation Constants

| Constant | Value | Hex | Purpose |
|----------|-------|-----|----------|
| `ClassIDShift` | 24 | 0x18 | Bit shift for classID extraction |
| `StringLengthMask` | 268435455 | 0x0FFFFFFF | 28-bit mask for string length |
| `WordsCountMask` | 16777215 | 0x00ffffff | 24-bit mask for word count |
| `ClassIDMask` | 255 | 0xff | 8-bit mask for classID |
| `FourByteAlignmentMask` | 3 | 0x3 | Mask for 4-byte alignment |
| `ByteAlignmentIncrement` | 3 | 0x3 | Increment for alignment calculations |

## Array Type Categories

### Category A: Direct Storage (ClassID 96)
- **Types**: Bool?, Byte?, Char?, Int16?, UInt
- **Function**: `moonbit_i32_array_make`
- **Element Size**: 4 bytes
- **None Values**: Type-specific (see table above)

### Category B: Reference Storage (ClassID 160)
- **Types**: Double?, Float?, Int64?, UInt64?
- **Function**: `moonbit_ref_array_make`
- **Element Size**: 4 bytes (pointers)
- **None Value**: `NoneSingletonPointer` (10248)

### Category C: Non-Optional Primitives (Various ClassIDs)
- **Types**: Int16 (80), Byte (246), Int/Float (241), etc.
- **Functions**: Type-specific (`moonbit_int16_array_make`, `moonbit_bytes_make`, etc.)
- **Element Size**: Varies by type
- **None Value**: N/A (not nullable)

### Category D: Int64-Based Storage
- **Types**: Int?
- **Function**: `moonbit_int64_array_make`
- **Element Size**: 8 bytes
- **None Value**: `NoneValueInt` (4294967296)

## Usage Examples

```go
// Instead of magic numbers:
if value == 4294967296 {
    item = nil
}

// Use named constants:
if value == NoneValueInt {
    item = nil
}

// Instead of magic numbers:
if classID == 96 {
    // handle direct storage
}

// Use named constants:
if classID == BoolByteCharClassID {
    // handle direct storage
}
```

This approach makes the code self-documenting and easier to maintain.
