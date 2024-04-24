package math

import "testing"

func TestSum(t *testing.T) {
	result := Sum(1, 5)
	var expected float32 = 6

	if result != expected {
		t.Errorf("Sum(1, 5) = %f; want %f", result, expected)
	}
}