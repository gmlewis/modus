/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package moonbit_test

import (
	"testing"
	"time"

	"github.com/gmlewis/modus/runtime/utils"
)

var testTimeStr = "2024-12-31T23:59:59.999999999Z"
var testTime, _ = time.Parse(time.RFC3339, testTimeStr)
var testDuration = time.Duration(5 * time.Second)

func TestTimeInput(t *testing.T) {
	fnName := "test_time_input"
	if _, err := fixture.CallFunction(t, fnName, testTime); err != nil {
		t.Error(err)
	}
}

func TestTimeStrInput(t *testing.T) {
	fnName := "test_time_input"
	if _, err := fixture.CallFunction(t, fnName, testTimeStr); err != nil {
		t.Error(err)
	}
}

func TestTimeOptionInput(t *testing.T) {
	fnName := "test_time_option_input"
	if _, err := fixture.CallFunction(t, fnName, testTime); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testTime); err != nil {
		t.Error(err)
	}
}

func TestTimeStrOptionInput(t *testing.T) {
	fnName := "test_time_option_input"
	if _, err := fixture.CallFunction(t, fnName, testTimeStr); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testTimeStr); err != nil {
		t.Error(err)
	}
}

func TestTimeOptionInput_none(t *testing.T) {
	fnName := "test_time_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestTimeOutput(t *testing.T) {
	fnName := "test_time_output"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(time.Time); !ok {
		t.Errorf("expected a time.Time, got %T", result)
	} else if r != testTime {
		t.Errorf("expected %v, got %v", true, r)
	}
}

func TestTimeOptionOutput(t *testing.T) {
	fnName := "test_time_option_output"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*time.Time); !ok {
		t.Errorf("expected a *time.Time, got %T", result)
	} else if *r != testTime {
		t.Errorf("expected %v, got %v", true, *r)
	}
}

func TestTimeOptionOutput_none(t *testing.T) {
	fnName := "test_time_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestDurationInput(t *testing.T) {
	fnName := "test_duration_input"
	if _, err := fixture.CallFunction(t, fnName, testDuration); err != nil {
		t.Error(err)
	}
}

func TestDurationOptionInput(t *testing.T) {
	fnName := "test_duration_option_input"
	if _, err := fixture.CallFunction(t, fnName, testDuration); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testDuration); err != nil {
		t.Error(err)
	}
}

func TestDurationOptionInput_none(t *testing.T) {
	fnName := "test_duration_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestDurationOutput(t *testing.T) {
	fnName := "test_duration_output"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(time.Duration); !ok {
		t.Errorf("expected a time.Duration, got %T", result)
	} else if r != testDuration {
		t.Errorf("expected %v, got %v", true, r)
	}
}

func TestDurationOptionOutput(t *testing.T) {
	fnName := "test_duration_option_output"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*time.Duration); !ok {
		t.Errorf("expected a *time.Duration, got %T", result)
	} else if *r != testDuration {
		t.Errorf("expected %v, got %v", true, *r)
	}
}

func TestDurationOptionOutput_none(t *testing.T) {
	fnName := "test_duration_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}
