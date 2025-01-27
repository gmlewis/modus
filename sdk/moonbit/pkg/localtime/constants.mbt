// These constants are used by the Modus host to format time strings.
// They are based on the Go time package's constants located here:
// https://cs.opensource.google/go/go/+/refs/tags/go1.23.5:src/time/format.go;l=109
// which has copyright:
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

///| The reference time, in numerical order.
pub const Layout = "01/02 03:04:05PM '06 -0700"

///|
pub const ANSIC = "Mon Jan _2 15:04:05 2006"

///|
pub const UnixDate = "Mon Jan _2 15:04:05 MST 2006"

///|
pub const RubyDate = "Mon Jan 02 15:04:05 -0700 2006"

///|
pub const RFC822 = "02 Jan 06 15:04 MST"

///| RFC822 with numeric zone
pub const RFC822Z = "02 Jan 06 15:04 -0700"

///|
pub const RFC850 = "Monday, 02-Jan-06 15:04:05 MST"

///|
pub const RFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"

///| RFC1123 with numeric zone
pub const RFC1123Z = "Mon, 02 Jan 2006 15:04:05 -0700"

///|
pub const RFC3339 = "2006-01-02T15:04:05Z07:00"

///|
pub const RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"

///|
pub const Kitchen = "3:04PM"

// Handy time stamps.

///|
pub const Stamp = "Jan _2 15:04:05"

///|
pub const StampMilli = "Jan _2 15:04:05.000"

///|
pub const StampMicro = "Jan _2 15:04:05.000000"

///|
pub const StampNano = "Jan _2 15:04:05.000000000"

///|
pub const DateTime = "2006-01-02 15:04:05"

///|
pub const DateOnly = "2006-01-02"

///|
pub const TimeOnly = "15:04:05"
