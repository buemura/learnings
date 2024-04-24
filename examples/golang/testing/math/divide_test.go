package math

import "testing"

func TestDivide(t *testing.T) {
	result := Divide(15, 3)
	var expected float32 = 5

	if result != expected {
		t.Errorf("Divide(5, 3) = %f; want %f", result, expected)
	}
}