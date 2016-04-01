package util

import (
	"reflect"
	"testing"
)

func TestFloat32frombytes(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		bytes []byte
		// Expected results.
		want float32
	}{
		// TODO: Add test cases.
		{"float32 转成 [4]byte", []byte{0x44, 0xc8, 0x00, 0x00}, 1600.00},
		{"float32 转成 [4]byte", []byte{0xbe, 0x80, 0x00, 0x00}, -0.25},
	}
	for _, tt := range tests {
		if got := Float32frombytes(tt.bytes); got != tt.want {
			t.Errorf("%q. Float32frombytes() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFloat32bytes(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		float float32
		// Expected results.
		want []byte
	}{
		// TODO: Add test cases.
		{" [4]byte 转 float32", 1600.00, []byte{0x44, 0xc8, 0x00, 0x00}},
		{" [4]byte 转 float32", -0.25, []byte{0xbe, 0x80, 0x00, 0x00}},
	}
	for _, tt := range tests {
		if got := Float32bytes(tt.float); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Float32bytes() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFloat64frombytes(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		bytes []byte
		// Expected results.
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Float64frombytes(tt.bytes); got != tt.want {
			t.Errorf("%q. Float64frombytes() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFloat64bytes(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		float float64
		// Expected results.
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Float64bytes(tt.float); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Float64bytes() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
