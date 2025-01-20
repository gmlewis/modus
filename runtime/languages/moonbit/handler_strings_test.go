package moonbit

import "testing"

func TestConvertMoonBitUTF16ToUTF8(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		data     []byte
		expected string
		wantErr  bool
	}{
		{
			name:     "Valid UTF-16 data",
			data:     []byte{0xff, 0xff, 0xff, 0xff, 0xf3, 0x07, 0x00, 0x00, 0x48, 0x00, 0x65, 0x00, 0x6c, 0x00, 0x6c, 0x00, 0x6f, 0x00, 0x2c, 0x00, 0x20, 0x00, 0x57, 0x00, 0x6f, 0x00, 0x72, 0x00, 0x6c, 0x00, 0x64, 0x00, 0x21, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
			expected: "Hello, World!",
			wantErr:  false,
		},
		{
			name:     "Data too short",
			data:     []byte{0xff, 0xff, 0xff, 0xff},
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Null terminator in the middle",
			data:     []byte{0xff, 0xff, 0xff, 0xff, 0xf3, 0x07, 0x00, 0x00, 0x48, 0x00, 0x65, 0x00, 0x6c, 0x00, 0x6c, 0x00, 0x00, 0x00, 0x6f, 0x00, 0x2c, 0x00, 0x20, 0x00, 0x57, 0x00, 0x6f, 0x00, 0x72, 0x00, 0x6c, 0x00, 0x64, 0x00, 0x21, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
			expected: "Hell",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := convertMoonBitUTF16ToUTF8(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertMoonBitUTF16ToUTF8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("convertMoonBitUTF16ToUTF8() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestConvertGoUTF8ToUTF16(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Simple ASCII",
			input: "Hello, World!",
		},
		{
			name:  "UTF-8 with emojis",
			input: "Hello, 🌍!",
		},
		{
			name:  "UTF-8 with non-Latin characters",
			input: "こんにちは世界",
		},
		{
			name:  "UTF-8 with mixed characters",
			input: "Hello, 世界! 🌍",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			encoded := convertGoUTF8ToUTF16(tt.input)
			decoded, err := convertMoonBitUTF16ToUTF8(encoded)
			if err != nil {
				t.Errorf("convertMoonBitUTF16ToUTF8() error = %v", err)
				return
			}
			if decoded != tt.input {
				t.Errorf("Round trip conversion failed: got = %v, want = %v", decoded, tt.input)
			}
		})
	}
}
