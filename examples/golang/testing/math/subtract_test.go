package math

import "testing"

func TestSubtract(t *testing.T) {
	result := Subtract(5, 2)
	var expected float32 = 3

	if result != expected {
		t.Errorf("Subtract(5, 2) = %f; want %f", result, expected)
	}
}