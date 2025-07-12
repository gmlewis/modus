Now I need to verify if my understanding is fully accurate. Based on my analysis, I found some important corrections to make:

## Updated and Verified Explanation: MoonBit String Memory Representation in WASM

After thorough analysis of both static string literals and dynamically generated strings (including those from `String::repeat`), I can provide a corrected and comprehensive explanation:

### Key Correction to My Previous Understanding

The **most important correction** is about the length encoding:

- **Offset 4-7 encoding**: The length is stored in the **lower 28 bits** of the 32-bit word at offset 4, not as a separate "word count"
- The `moonbit.array_length` function uses `i32.const 268435455` (0x0FFFFFFF) as a mask to extract the length

### Corrected Memory Block Structure

Each MoonBit `String` object has this layout:

```
Offset  Size  Description
------  ----  -----------
0x00    4     Reference count (32-bit little-endian)
0x04    4     Type ID + Length info (32-bit little-endian)
0x08    N*2   UTF-16 encoded string data (N characters * 2 bytes each)
0x08+N*2 Pad  Zero padding to align to 4-byte boundaries
```

### Detailed Field Analysis (Corrected)

#### 1. Reference Count (Bytes 0-3)
- **Unchanged**: 32-bit little-endian integer for garbage collection
- `0xFFFFFFFF` (-1): Immortal/static strings (string literals)
- `0x00000001` (1): Dynamic strings with reference count 1

#### 2. Type ID + Length (Bytes 4-7) - **CORRECTED**
- **Bits 0-27** (lower 28 bits): String length in **characters** (not words)
- **Bits 28-31**: Type information and flags
- For strings, the upper bits contain type ID information (including the value 243 that represents String type)

#### 3. String Data (Bytes 8+)
- **Encoding**: UTF-16 little-endian (confirmed by `i32.load16_u` usage)
- **Access pattern**: `base_address + (character_index << 1) + 8`
- **Padding**: Zero-padded to maintain 4-byte alignment

### Verification with All String Types

This corrected understanding works for **ALL** string types:

#### Static String Literals
- Pre-allocated in the data section starting at address 10000
- Reference count = -1 (immortal)
- Length stored in lower 28 bits of offset 4

#### Dynamic Strings (from `repeat`, `StringBuilder`, etc.)
- Allocated at runtime using `moonbit.gc.malloc`
- Reference count = 1 (or higher)
- Same memory layout as static strings
- Length calculated and stored in the same format

#### Empty Strings
- Length = 0 in the lower 28 bits
- Only 8-byte header (no character data)
- Still padded to 4-byte boundaries

### Key Functions That Confirm This Structure

1. **`moonbit.array_length`**: Extracts length using mask `0x0FFFFFFF` (28 bits)
2. **`moonbit.string_item_unsafe`**: Accesses characters at `offset=8 + (index << 1)`
3. **`String::repeat`**: Creates new strings with the same memory layout
4. **`moonbit.add_string`**: Concatenates strings maintaining the same structure

### Why This Works for All Strings

1. **Static strings**: Compiler pre-calculates length and embeds it correctly
2. **Dynamic strings**: Runtime functions (`StringBuilder`, `repeat`) calculate length and store it in the same format
3. **Unicode support**: UTF-16 encoding allows proper handling of all Unicode characters
4. **Memory efficiency**: 28-bit length allows strings up to 268M characters while reserving upper bits for type information

### Final Validation

The corrected understanding explains:
- Why `test_string_output_len_100000` (100,000 characters) works correctly
- How string concatenation and manipulation functions operate
- Why the same memory access patterns work for both static and dynamic strings
- How the garbage collection system manages both immortal and mutable strings

This memory representation is consistent across **all** MoonBit string types and operations, including dynamic string generation through functions like `repeat`.
