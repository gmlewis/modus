package moonbit

import (
	"testing"

	"github.com/gmlewis/modus/runtime/langsupport"
	wasm "github.com/tetratelabs/wazero/api"

	"github.com/stretchr/testify/mock"
)

type mockWasmAdapter struct {
	mock.Mock
	langsupport.WasmAdapter
}

func (m *mockWasmAdapter) Memory() wasm.Memory {
	args := m.Called()
	return args.Get(0).(wasm.Memory)
}

type mockMemory struct {
	mock.Mock
	wasm.Memory
}

func (m *mockMemory) Read(offset, size uint32) ([]byte, bool) {
	args := m.Called(offset, size)
	return args.Get(0).([]byte), args.Bool(1)
}

func TestMemoryBlockAtOffset(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		offset       uint32
		memBlock     []byte
		expectedData uint32
		expectedSize uint32
		expectedErr  error
	}{
		{
			name:         "Valid memory block, UTF-16 String 'Hello, ...0!' with remainder 0",
			offset:       100,
			memBlock:     []byte{1, 0, 0, 0, 243, 7, 0, 0, 72, 0, 101, 0, 108, 0, 108, 0, 111, 0, 44, 0, 32, 0, 46, 0, 46, 0, 46, 0, 48, 0, 33, 0, 97, 0, 109, 3},
			expectedData: 108,
			expectedSize: 24,
			expectedErr:  nil,
		},
		{
			name:         "Valid memory block, UTF-16 String 'Hello, 2!' with remainder 2",
			offset:       100,
			memBlock:     []byte{1, 0, 0, 0, 243, 5, 0, 0, 72, 0, 101, 0, 108, 0, 108, 0, 111, 0, 44, 0, 32, 0, 50, 0, 33, 0, 32, 1},
			expectedData: 108,
			expectedSize: 18,
			expectedErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockMem := new(mockMemory)
			mockWA := new(mockWasmAdapter)
			mockWA.On("Memory").Return(mockMem)
			mockMem.On("Read", tt.offset, uint32(8)).Return(tt.memBlock[:8], true)
			mockMem.On("Read", tt.offset, uint32(len(tt.memBlock))).Return(tt.memBlock, true)

			data, size, err := memoryBlockAtOffset(mockWA, tt.offset)
			if data != tt.expectedData || size != tt.expectedSize || (err != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("memoryBlockAtOffset() = (data: %v, size: %v, err: %v), want (data: %v, size: %v, err: %v)",
					data, size, err, tt.expectedData, tt.expectedSize, tt.expectedErr)
			}
		})
	}
}
