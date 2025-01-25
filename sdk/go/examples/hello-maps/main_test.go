package main

import (
	"reflect"
	"testing"
)

func TestHelloMapsNItems(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want map[string]string
	}{
		{
			name: "0 items",
			n:    0,
			want: map[string]string{},
		},
		{
			name: "1 item",
			n:    1,
			want: map[string]string{"key1": "value1"},
		},
		{
			name: "2 items",
			n:    2,
			want: map[string]string{"key1": "value1", "key2": "value2"},
		},
		{
			name: "10 items",
			n:    10,
			want: map[string]string{
				"key1":  "value1",
				"key2":  "value2",
				"key3":  "value3",
				"key4":  "value4",
				"key5":  "value5",
				"key6":  "value6",
				"key7":  "value7",
				"key8":  "value8",
				"key9":  "value9",
				"key10": "value10",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HelloMapsNItems(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HelloMapsNItems(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
