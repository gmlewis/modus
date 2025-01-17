/*
 * This example is part of the Modus project, licensed under the Apache License 2.0.
 * You may modify and use this example in accordance with the license.
 * See the LICENSE file that accompanied this code for further details.
 */

package main

// HelloOptionString demonstrates how the Modus Runtime for Go handles
// optional strings as return values.
func HelloOptionString(nilString, emptyString bool) *string {
	switch {
	case nilString:
		return nil
	case emptyString:
		s := ""
		return &s
	default:
		s := "Hello, World!"
		return &s
	}
}
