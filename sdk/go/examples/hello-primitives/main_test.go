package main

import (
	"testing"
)

func TestHelloPrimitiveTimeDuration(t *testing.T) {
	got := HelloPrimitiveTimeDurationMin().String()
	if want := "0s"; got != want {
		t.Errorf("HelloPrimitiveTimeDurationMin = %q, want %q", got, want)
	}

	got = HelloPrimitiveTimeDurationMax().String()
	if want := "277777h46m39s"; got != want {
		t.Errorf("HelloPrimitiveTimeDurationMax = %q, want %q", got, want)
	}
}
