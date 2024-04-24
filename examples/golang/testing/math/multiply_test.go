package math

import "testing"

func TestMultiply(t *testing.T) {
	result := Multiply(5, 3)
	var expected float32 = 15

	if result != expected {
		t.Errorf("Multiply(5, 3) = %f; want %f", result, expected)
	}
}