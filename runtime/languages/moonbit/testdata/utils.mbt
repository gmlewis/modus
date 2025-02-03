// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
fn fail(msg : String) -> Unit {
  @console.error(msg)
  @sys.exit(1)
}

// fn assert_equal[T any](expect, actual T) {
// 	if !reflect.DeepEqual(expect, actual) {
// 		fail(fmt.Sprintf("expected %v, got %v", expect, actual))
// 	}
// }

// fn assert_slices_equal[T comparable](expect, actual []T) {
// 	if (expect == nil && actual != nil) || (expect != nil && actual == nil) || !slices.Equal(expect, actual) {
// 		fail(fmt.Sprintf("expected %v, got %v", expect, actual))
// 	}
// }

// fn assert_maps_equal[T ~map[K]V, K, V comparable](expect, actual T) {
// 	if (expect == nil && actual != nil) || (expect != nil && actual == nil) || !maps.Equal(expect, actual) {
// 		fail(fmt.Sprintf("expected %v, got %v", expect, actual))
// 	}
// }

// fn assert_ptr_slices_equal[T comparable](expect, actual []*T) {
// 	if (expect == nil && actual != nil) || (expect != nil && actual == nil) ||
// 		!slices.EqualFunc(expect, actual, ptrValsEqual) {
// 		fail(fmt.Sprintf("expected %v, got %v", expect, actual))
// 	}
// }

// fn ptr_vals_equal[T comparable](a, b *T) bool {
// 	return a == nil && b == nil || a != nil && b != nil && *a == *b
// }

// fn assert_nil[T any](actual T) {
// 	if !hasNil(actual) {
// 		fail(fmt.Sprintf("expected nil, got %v", actual))
// 	}
// }

// // HasNil returns true if the given interface value is nil, or contains a pointer, interface, slice, or map that is nil.
// fn has_nil(x any) bool {
// 	if x == nil {
// 		return true
// 	}

// 	rv := reflect.ValueOf(x)
// 	if canBeNil(rv.Type()) {
// 		return rv.IsNil()
// 	}

// 	return false
// }

// fn can_be_nil(rt reflect.Type) bool {
// 	switch rt.Kind() {
// 	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
// 		return true
// 	default:
// 		return false
// 	}
// }
