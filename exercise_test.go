package golang

import "testing"

func TestMultipleReturnValues1(t *testing.T) {
	a, b := stepup(0, 1)
	if a != 1 || b != 1 {
		t.Error("it should return second argument to the first value and sum of arguments to the second")
	}
}

func TestMultipleReturnValues2(t *testing.T) {
	a, b := stepup(1, 1)
	if a != 1 || b != 2 {
		t.Error("it should return second argument to the first value and sum of arguments to the second")
	}
}

func TestMultipleReturnValues3(t *testing.T) {
	a, b := stepup(1, 2)
	if a != 2 || b != 3 {
		t.Error("it should return second argument to the first value and sum of arguments to the second")
	}
}

func TestMultipleReturnValues4(t *testing.T) {
	a, b := stepup(2, 3)
	if a != 3 || b != 5 {
		t.Error("it should return second argument to the first value and sum of arguments to the second")
	}
}

func TestMultipleReturnValues5(t *testing.T) {
	a, b := stepup(3, 5)
	if a != 5 || b != 8 {
		t.Error("it should return second argument to the first value and sum of arguments to the second")
	}
}
