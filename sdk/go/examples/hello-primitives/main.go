/*
 * This example is part of the Modus project, licensed under the Apache License 2.0.
 * You may modify and use this example in accordance with the license.
 * See the LICENSE file that accompanied this code for further details.
 */

package main

import (
	"time"
)

func HelloPrimitiveTimeDurationMin() time.Duration {
	var d time.Duration
	return d
}

func HelloPrimitiveTimeDurationMax() time.Duration {
	d, _ := time.ParseDuration("999999999s")
	return d
}
