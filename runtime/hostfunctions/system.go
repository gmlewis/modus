/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package hostfunctions

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gmlewis/modus/runtime/logger"
	"github.com/gmlewis/modus/runtime/timezones"
	"github.com/gmlewis/modus/runtime/utils"
)

func init() {
	const module_name = "modus_system"

	registerHostFunction(module_name, "logMessage", LogMessage)
	registerHostFunction(module_name, "getTimeInZone", GetTimeInZone)
	registerHostFunction(module_name, "getTimeZoneData", GetTimeZoneData)
}

func LogMessage(ctx context.Context, level, message string) {
	log.SetFlags(0)
	log.Printf("GML: hostfunctions/system.go: LogMessage: level: '%v', message: '%v'", level, message)

	// store messages in the context, so we can return them to the caller
	messages := ctx.Value(utils.FunctionMessagesContextKey).(*[]utils.LogMessage)
	*messages = append(*messages, utils.LogMessage{
		Level:   level,
		Message: message,
	})

	// If debugging, write debug messages to stderr instead of the logger
	if level == "debug" && utils.DebugModeEnabled() {
		fmt.Fprintln(os.Stderr, message)
		return
	}

	// write to the logger
	logger.Get(ctx).
		WithLevel(logger.ParseLevel(level)).
		Str("text", message).
		Bool("user_visible", true).
		Msg("Message logged from function.")
}

func GetTimeInZone(ctx context.Context, tz *string) *string {
	log.SetFlags(0)
	if tz != nil {
		log.Printf("GML: hostfunctions/system.go: GetTimeInZone: tz: '%v'", *tz)
	}
	now := time.Now()
	log.Printf("GML: hostfunctions/system.go: GetTimeInZone: now: '%v'", now)

	var loc *time.Location
	if tz != nil && *tz != "" {
		loc = timezones.GetLocation(*tz)
	} else if tz, ok := ctx.Value(utils.TimeZoneContextKey).(string); ok {
		log.Printf("GML: hostfunctions/system.go: GetTimeInZone: tz=nil: local tz: '%v'", tz)
		loc = timezones.GetLocation(tz)
	}

	if loc == nil {
		log.Printf("GML: hostfunctions/system.go: GetTimeInZone: loc: nil - returning nil")
		return nil
	}

	s := now.In(loc).Format(time.RFC3339Nano)
	log.Printf("GML: hostfunctions/system.go: GetTimeInZone: s: '%v'", s)
	return &s
}

func GetTimeZoneData(ctx context.Context, tz, format *string) []byte {
	log.SetFlags(0)
	if tz != nil {
		log.Printf("GML: hostfunctions/system.go: GetTimeZoneData: tz: '%v'", *tz)
	} else {
		if tz2, ok := ctx.Value(utils.TimeZoneContextKey).(string); ok {
			log.Printf("GML: hostfunctions/system.go: GetTimeZoneData: tz=nil: local tz: '%v'", tz2)
			tz = &tz2
		}
	}
	if format != nil {
		log.Printf("GML: hostfunctions/system.go: GetTimeZoneData: format: '%v'", *format)
	} else {
		log.Printf("GML: hostfunctions/system.go: GetTimeZoneData: format: nil - setting to 'tzif'")
		fmt := "tzif"
		format = &fmt
	}
	if tz == nil {
		return nil
	}

	result := timezones.GetTimeZoneData(*tz, *format)
	log.Printf("GML: hostfunctions/system.go: GetTimeZoneData: result: %+v", result)
	return result
}
