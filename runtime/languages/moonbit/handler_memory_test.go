/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package moonbit

import (
	"context"
	"errors"
	"math"
	"testing"

	"github.com/gmlewis/modus/runtime/utils"
	wasm "github.com/tetratelabs/wazero/api"
)

type myWasmMock struct {
	offset uint32
	m      *myWasmMockMemory
}
type myWasmMockMemory struct {
	offset uint32
	bytes  []byte
	wasm.Memory
}

func TestMyWasmMock(t *testing.T) {
	ctx := context.Background()
	m := &myWasmMock{offset: 46864}
	block1, _, _ := m.allocateAndPinMemory(ctx, 16, 0)
	m.Memory().WriteUint32Le(block1, 41)
	m.Memory().WriteUint32Le(block1+4, 42)
	m.Memory().WriteUint32Le(block1+8, 43)
	m.Memory().WriteUint32Le(block1+12, 44)
	block2, _, _ := m.allocateAndPinMemory(ctx, 160, 242)
	m.Memory().WriteUint32Le(block2, 161)
	m.Memory().WriteUint32Le(block2+4, 162)
	m.Memory().WriteUint32Le(block2+8, 163)
	m.Memory().WriteUint32Le(block2+12, 164)
	block3, _, _ := m.allocateAndPinMemory(ctx, 8, 243)
	m.Memory().WriteUint32Le(block3, 81)
	m.Memory().WriteUint32Le(block3+4, 82)

	if v, ok := m.Memory().ReadUint32Le(block1); !ok || v != 41 {
		t.Errorf("expected 41, got %d", v)
	}
	if v, ok := m.Memory().ReadUint32Le(block1 + 4); !ok || v != 42 {
		t.Errorf("expected 42, got %d", v)
	}
	if v, ok := m.Memory().ReadUint32Le(block1 + 8); !ok || v != 43 {
		t.Errorf("expected 43, got %d", v)
	}
	if v, ok := m.Memory().ReadUint32Le(block1 + 12); !ok || v != 44 {
		t.Errorf("expected 44, got %d", v)
	}
	if v, ok := m.Memory().ReadUint32Le(block2); !ok || v != 161 {
		t.Errorf("expected 161, got %d", v)
	}
	if v, ok := m.Memory().ReadUint32Le(block2 + 4); !ok || v != 162 {
		t.Errorf("expected 162, got %d", v)
	}
	if v, ok := m.Memory().ReadUint32Le(block2 + 8); !ok || v != 163 {
		t.Errorf("expected 163, got %d", v)
	}
	if v, ok := m.Memory().ReadUint32Le(block2 + 12); !ok || v != 164 {
		t.Errorf("expected 164, got %d", v)
	}
	if v, ok := m.Memory().ReadUint32Le(block3); !ok || v != 81 {
		t.Errorf("expected 81, got %d", v)
	}
	if v, ok := m.Memory().ReadUint32Le(block3 + 4); !ok || v != 82 {
		t.Errorf("expected 82, got %d", v)
	}
}

func (m *myWasmMock) allocateAndPinMemory(ctx context.Context, size, classID uint32) (uint32, utils.Cleaner, error) {
	if size == 0 {
		return 0, nil, errors.New("size must be greater than 0")
	}
	size = 4 * ((size + 3) / 4) // round up to nearest multiple of 4

	offset := m.offset
	if m.m == nil {
		m.m = &myWasmMockMemory{
			offset: m.offset,
			bytes:  make([]byte, size+8),
		}
	} else {
		offset += uint32(len(m.m.bytes))
		m.m.bytes = append(m.m.bytes, make([]byte, size+8)...)
	}
	refCount := uint32(1)
	memType := ((size / 4) << 8) | classID
	m.m.WriteUint32Le(m.offset, refCount)
	m.m.WriteUint32Le(m.offset+4, memType)
	// allocateAndPinMemory always returns a pointer _after_ the memory block header
	return offset + 8, nil, nil
}

func (m *myWasmMock) Memory() wasm.Memory                     { return m.m }
func (m *myWasmMockMemory) Definition() wasm.MemoryDefinition { return nil }
func (m *myWasmMockMemory) Grow(uint32) (uint32, bool)        { return 0, false }
func (m *myWasmMockMemory) Read(offset, size uint32) ([]byte, bool) {
	offset -= m.offset
	return m.bytes[offset : offset+size], true
}
func (m *myWasmMockMemory) ReadByte(offset uint32) (byte, bool) {
	return m.bytes[offset-m.offset], true
}
func (m *myWasmMockMemory) ReadFloat32Le(offset uint32) (float32, bool) {
	val, ok := m.ReadUint32Le(offset)
	return math.Float32frombits(val), ok
}
func (m *myWasmMockMemory) ReadFloat64Le(offset uint32) (float64, bool) {
	val, ok := m.ReadUint64Le(offset)
	return math.Float64frombits(val), ok
}
func (m *myWasmMockMemory) ReadUint16Le(offset uint32) (uint16, bool) {
	offset -= m.offset
	return uint16(m.bytes[offset]) | (uint16(m.bytes[offset+1]) << 8), true
}
func (m *myWasmMockMemory) ReadUint32Le(offset uint32) (uint32, bool) {
	offset -= m.offset
	return uint32(m.bytes[offset]) |
		(uint32(m.bytes[offset+1]) << 8) |
		(uint32(m.bytes[offset+2]) << 16) |
		(uint32(m.bytes[offset+3]) << 24), true
}
func (m *myWasmMockMemory) ReadUint64Le(offset uint32) (uint64, bool) {
	offset -= m.offset
	return uint64(m.bytes[offset]) |
		(uint64(m.bytes[offset+1]) << 8) |
		(uint64(m.bytes[offset+2]) << 16) |
		(uint64(m.bytes[offset+3]) << 24) |
		(uint64(m.bytes[offset+4]) << 32) |
		(uint64(m.bytes[offset+5]) << 40) |
		(uint64(m.bytes[offset+6]) << 48) |
		(uint64(m.bytes[offset+7]) << 56), true
}
func (m *myWasmMockMemory) Size() uint32 { return uint32(len(m.bytes)) }
func (m *myWasmMockMemory) Write(offset uint32, bytes []byte) bool {
	copy(m.bytes[offset-m.offset:], bytes)
	return true
}
func (m *myWasmMockMemory) WriteByte(offset uint32, b byte) bool {
	m.bytes[offset-m.offset] = b
	return true
}
func (m *myWasmMockMemory) WriteFloat32Le(offset uint32, val float32) bool {
	v := math.Float32bits(val)
	return m.WriteUint32Le(offset, v)
}
func (m *myWasmMockMemory) WriteFloat64Le(offset uint32, val float64) bool {
	v := math.Float64bits(val)
	return m.WriteUint64Le(offset, v)
}
func (m *myWasmMockMemory) WriteUint16Le(offset uint32, val uint16) bool {
	offset -= m.offset
	m.bytes[offset] = byte(val)
	m.bytes[offset+1] = byte(val >> 8)
	return true
}
func (m *myWasmMockMemory) WriteUint32Le(offset uint32, val uint32) bool {
	offset -= m.offset
	m.bytes[offset] = byte(val)
	m.bytes[offset+1] = byte(val >> 8)
	m.bytes[offset+2] = byte(val >> 16)
	m.bytes[offset+3] = byte(val >> 24)
	return true
}
func (m *myWasmMockMemory) WriteUint64Le(offset uint32, val uint64) bool {
	offset -= m.offset
	m.bytes[offset] = byte(val)
	m.bytes[offset+1] = byte(val >> 8)
	m.bytes[offset+2] = byte(val >> 16)
	m.bytes[offset+3] = byte(val >> 24)
	m.bytes[offset+4] = byte(val >> 32)
	m.bytes[offset+5] = byte(val >> 40)
	m.bytes[offset+6] = byte(val >> 48)
	m.bytes[offset+7] = byte(val >> 56)
	return true
}
func (m *myWasmMockMemory) WriteString(offset uint32, val string) bool { return false }
