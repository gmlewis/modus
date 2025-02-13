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
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"
)

func (p *planner) NewTimeHandler(ti langsupport.TypeInfo) (langsupport.TypeHandler, error) {
	handler := &timeHandler{*NewTypeHandler(ti)}
	p.AddHandler(handler)
	return handler, nil
}

type timeHandler struct {
	typeHandler
}

func (h *timeHandler) Read(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {
	gmlPrintf("GML: handler_time.go: timeHandler.Read(offset: %v)", offset)

	if offset == 0 {
		return nil, nil
	}

	return h.Decode(ctx, wa, []uint64{uint64(offset)})

	// wall, ok := wa.Memory().ReadUint64Le(offset)
	// if !ok {
	// 	return nil, errors.New("failed to read time.Time.wall from WASM memory")
	// }

	// x, ok := wa.Memory().ReadUint64Le(offset + 8)
	// if !ok {
	// 	return nil, errors.New("failed to read time.Time.ext from WASM memory")
	// }
	// ext := int64(x)

	// // skip loc - we only support UTC

	// return timeFromVals(wall, ext), nil
}

func (h *timeHandler) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	gmlPrintf("GML: handler_time.go: timeHandler.Write(offset: %v, obj: %v)", offset, obj)
	res, _, err := h.Encode(ctx, wa, obj)
	if err != nil || len(res) != 1 {
		return nil, err
	}

	if !wa.Memory().WriteUint64Le(offset, res[0]) {
		return nil, errors.New("failed to write @time.ZonedDateTime to WASM memory")
	}

	return nil, nil
}

func (h *timeHandler) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	gmlPrintf("GML: handler_time.go: timeHandler.Decode(vals: %+v)", vals)

	if len(vals) != 1 {
		return nil, fmt.Errorf("MoonBit: expected 1 value when decoding time but got %v: %+v", len(vals), vals)
	}

	// MoonBit is configured to return a @time.ZonedDateTime from the moonbitlang/x/time package:
	// A datetime with a time zone and offset in the ISO 8601 calendar system.
	// struct ZonedDateTime {
	//   datetime : PlainDateTime
	//   zone : Zone
	//   offset : ZoneOffset
	// }
	// So although a lot more information is returned, we only need the wall and ext values to create a time.Time.
	memBlock, _, err := memoryBlockAtOffset(wa, uint32(vals[0]), 0, true)
	if err != nil {
		return nil, err
	}
	if len(memBlock) != 20 {
		return nil, fmt.Errorf("MoonBit: expected 20 bytes when decoding time but got %v: %+v", len(memBlock), memBlock)
	}

	datetimePtr := binary.LittleEndian.Uint32(memBlock[8:])
	// For reverse-engineering purposes only:
	zonePtr := binary.LittleEndian.Uint32(memBlock[12:])
	_, _, _ = memoryBlockAtOffset(wa, zonePtr, 0, true)
	offsetPtr := binary.LittleEndian.Uint32(memBlock[16:])
	_, _, _ = memoryBlockAtOffset(wa, offsetPtr, 0, true)
	// end of reverse-engineering section.

	datetime, _, err := memoryBlockAtOffset(wa, datetimePtr, 0, true)
	if err != nil {
		return nil, err
	}

	plainDatePtr := binary.LittleEndian.Uint32(datetime[8:])
	plainTimePtr := binary.LittleEndian.Uint32(datetime[12:])
	plainDate, _, err := memoryBlockAtOffset(wa, plainDatePtr, 0, true)
	if err != nil {
		return nil, err
	}
	plainTime, _, err := memoryBlockAtOffset(wa, plainTimePtr, 0, true)
	if err != nil {
		return nil, err
	}
	year := binary.LittleEndian.Uint32(plainDate[8:])
	month := binary.LittleEndian.Uint32(plainDate[12:])
	day := binary.LittleEndian.Uint32(plainDate[16:])
	hour := binary.LittleEndian.Uint32(plainTime[8:])
	minute := binary.LittleEndian.Uint32(plainTime[12:])
	second := binary.LittleEndian.Uint32(plainTime[16:])
	nanosecond := binary.LittleEndian.Uint32(plainTime[20:])
	// gmlPrintf("GML: plainDate=%+v, plainTime=%+v, year=%v, month=%v, day=%v, hour=%v, minute=%v, second=%v, nanosecond=%v", plainDate, plainTime, year, month, day, hour, minute, second, nanosecond)

	return time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), int(nanosecond), time.UTC), nil

	// zone, _, err := memoryBlockAtOffset(wa, zonePtr)
	// if err != nil {
	// return nil, err
	// }

	// zoneIDPtr := binary.LittleEndian.Uint32(zone[8:])
	// zoneOffsetsPtr := binary.LittleEndian.Uint32(zone[12:])
	// zoneID, _, err := memoryBlockAtOffset(wa, zoneIDPtr)
	// if err != nil {
	// return nil, err
	// }
	// zoneOffsets, _, err := memoryBlockAtOffset(wa, zoneOffsetsPtr)
	// if err != nil {
	// return nil, err
	// }
	// gmlPrintf("GML: zoneID=%+v, zoneOffsets=%+v", zoneID, zoneOffsets)

	// offset, _, err := memoryBlockAtOffset(wa, offsetPtr)
	// if err != nil {
	// return nil, err
	// }

	// zoneOffsetIDPtr := binary.LittleEndian.Uint32(offset[8:])
	// zoneOffsetSecondsPtr := binary.LittleEndian.Uint32(offset[12:])
	// zoneOffsetDSTPtr := binary.LittleEndian.Uint32(offset[16:])
	// zoneOffsetID, _, err := memoryBlockAtOffset(wa, zoneOffsetIDPtr)
	// if err != nil {
	// return nil, err
	// }
	// zoneOffsetSeconds, _, err := memoryBlockAtOffset(wa, zoneOffsetSecondsPtr)
	// if err != nil {
	// return nil, err
	// }
	// zoneOffsetDST, _, err := memoryBlockAtOffset(wa, zoneOffsetDSTPtr)
	// if err != nil {
	// return nil, err
	// }
	// gmlPrintf("GML: zoneOffsetID=%+v, zoneOffsetSeconds=%+v, zoneOffsetDST=%+v", zoneOffsetID, zoneOffsetSeconds, zoneOffsetDST)

	// return memBlock, nil

	// wall, ext := vals[0], int64(vals[1])
	// // skip loc - we only support UTC

	// return timeFromVals(wall, ext), nil
}

// This converts a Go `time.Time` to a MoonBit `@time.ZonedDateTime`.
func (h *timeHandler) Encode(ctx context.Context, wa langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	if obj == nil || utils.HasNil(obj) {
		return []uint64{0}, nil, nil
	}

	tm, err := utils.ConvertToTimestamp(obj)
	if err != nil {
		return nil, nil, err
	}

	seconds := tm.Unix()
	nanos := tm.UnixNano() % 1_000_000_000
	gmlPrintf("GML: handler_time.go: timeHandler.Encode(obj: %v): seconds=%v, nanos=%v", tm, seconds, nanos)
	res, err := wa.(*wasmAdapter).fnZonedDateTimeFromUnixSecondsAndNanos.Call(ctx, uint64(seconds), uint64(nanos))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert time.Time to ZonedDateTime: %w", err)
	}
	gmlPrintf("GML: handler_time.go: timeHandler.Encode: res=%+v", res)
	if len(res) != 2 || res[0] == 0 {
		return nil, nil, fmt.Errorf("failed to convert time.Time to ZonedDateTime: %+v", res)
	}

	return res[1:], nil, nil
}

// func timeFromVals(wall uint64, ext int64) time.Time {
// 	type tm struct {
// 		wall uint64
// 		ext  int64
// 		loc  *time.Location
// 	}
//
// 	t := tm{wall, ext, nil}
// 	return *(*time.Time)(unsafe.Pointer(&t))
// }
//
// func getTimeVals(t time.Time) (uint64, int64) {
// 	type tm struct {
// 		wall uint64
// 		ext  int64
// 		loc  *time.Location
// 	}
//
// 	s := *(*tm)(unsafe.Pointer(&t))
// 	return s.wall, s.ext
// }
