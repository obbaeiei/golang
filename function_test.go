package golang

import "testing"

func TestSubtractFuncTenAndTwo(t *testing.T) {
	result := subtract(10, 2)

	if result != 8 {
		t.Error("Please make the right subtract function")
	}
}

func TestSubtractFuncTenAndFive(t *testing.T) {
	result := subtract(10, 5)

	if result != 5 {
		t.Error("Please make the right subtract function")
	}
}
