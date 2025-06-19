/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// Tests FAIL with moonc v0.6.18+8382ed77e

package moonbit_test

import (
	"math"
	"reflect"
	"slices"
	"testing"
)

func TestArrayOutput_int64_0(t *testing.T) {
	fnName := "test_array_output_int64_0"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int64{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int64); !ok {
		t.Errorf("expected a []int64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_1(t *testing.T) {
	fnName := "test_array_output_int64_1"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int64{1}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int64); !ok {
		t.Errorf("expected a []int64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_1_min(t *testing.T) {
	fnName := "test_array_output_int64_1_min"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0 0 0 0 128]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int64{math.MinInt64}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int64); !ok {
		t.Errorf("expected a []int64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_1_max(t *testing.T) {
	fnName := "test_array_output_int64_1_max"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 255 255 255 255 127]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int64{math.MaxInt64}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int64); !ok {
		t.Errorf("expected a []int64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_2(t *testing.T) {
	fnName := "test_array_output_int64_2"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 0 0 0 0 0 0 0 128 255 255 255 255 255 255 255 127]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int64{math.MinInt64, math.MaxInt64}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int64); !ok {
		t.Errorf("expected a []int64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_3(t *testing.T) {
	fnName := "test_array_output_int64_3"

	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 32=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0 3 0 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int64{1, 2, 3}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int64); !ok {
		t.Errorf("expected a []int64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_4(t *testing.T) {
	fnName := "test_array_output_int64_4"

	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 4 0 0 0]
	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 40=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0 3 0 0 0 0 0 0 0 4 0 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int64{1, 2, 3, 4}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int64); !ok {
		t.Errorf("expected a []int64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_option_0(t *testing.T) {
	fnName := "test_array_output_int64_option_0"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 8=8+words*4), moonBitType=242(FixedArray[String]), words=0, memBlock=[1 0 0 0 242 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int64{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int64); !ok {
		t.Errorf("expected a []*int64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_option_1_none(t *testing.T) {
	fnName := "test_array_output_int64_option_1_none"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 152 41 0 0]
	// memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int64{nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int64); !ok {
		t.Errorf("expected a []*int64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_option_1_min(t *testing.T) {
	fnName := "test_array_output_int64_option_1_min"

	// memoryBlockAtOffset(offset: 49072=0x0000BFB0=[176 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 12=8+words*4), moonBitType=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 192 190 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 0 0 0 0 0 0 0 128]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int64{Ptr(int64(math.MinInt64))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int64); !ok {
		t.Errorf("expected a []*int64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_option_1_max(t *testing.T) {
	fnName := "test_array_output_int64_option_1_max"

	// memoryBlockAtOffset(offset: 49072=0x0000BFB0=[176 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 12=8+words*4), moonBitType=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 192 190 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 255 255 255 255 255 255 255 127]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int64{Ptr(int64(math.MaxInt64))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int64); !ok {
		t.Errorf("expected a []*int64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_option_2(t *testing.T) {
	fnName := "test_array_output_int64_option_2"

	// memoryBlockAtOffset(offset: 49120=0x0000BFE0=[224 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 191 0 0 2 0 0 0]
	// memoryBlockAtOffset(offset: 49088=0x0000BFC0=[192 191 0 0], size: 16=8+words*4), moonBitType=242(FixedArray[String]), words=2, memBlock=[1 0 0 0 242 2 0 0 192 190 0 0 160 191 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 1 0 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 2 0 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int64{Ptr(int64(1)), Ptr(int64(2))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int64); !ok {
		t.Errorf("expected a []*int64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_option_3(t *testing.T) {
	fnName := "test_array_output_int64_option_3"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 20=8+words*4), moonBitType=242(FixedArray[String]), words=3, memBlock=[1 0 0 0 242 3 0 0 152 41 0 0 152 41 0 0 152 41 0 0]
	// memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	// memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	// memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int64{nil, nil, nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int64); !ok {
		t.Errorf("expected a []*int64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int64_option_4(t *testing.T) {
	fnName := "test_array_output_int64_option_4"

	// memoryBlockAtOffset(offset: 49152=0x0000C000=[0 192 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 224 191 0 0 4 0 0 0]
	// memoryBlockAtOffset(offset: 49120=0x0000BFE0=[224 191 0 0], size: 24=8+words*4), moonBitType=242(FixedArray[String]), words=4, memBlock=[1 0 0 0 242 4 0 0 152 41 0 0 192 190 0 0 160 191 0 0 192 191 0 0]
	// memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 0 0 0 0 0 0 0 128]
	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 0 0 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 49088=0x0000BFC0=[192 191 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 255 255 255 255 255 255 255 127]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int64{nil, Ptr(int64(math.MinInt64)), Ptr(int64(0)), Ptr(int64(math.MaxInt64))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int64); !ok {
		t.Errorf("expected a []*int64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}
