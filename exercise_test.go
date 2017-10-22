package golang

import (
	"testing"
)

func TestCase1(t *testing.T) {
	result := text(1)
	expected := "1"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase2(t *testing.T) {
	result := text(2)
	expected := "2"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase3(t *testing.T) {
	result := text(3)
	expected := "fizz"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase4(t *testing.T) {
	result := text(4)
	expected := "4"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase5(t *testing.T) {
	result := text(5)
	expected := "buzz"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase6(t *testing.T) {
	result := text(6)
	expected := "fizz"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase7(t *testing.T) {
	result := text(7)
	expected := "7"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase8(t *testing.T) {
	result := text(8)
	expected := "8"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase9(t *testing.T) {
	result := text(9)
	expected := "fizz"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase10(t *testing.T) {
	result := text(10)
	expected := "buzz"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase15(t *testing.T) {
	result := text(15)
	expected := "fizzbuzz"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase30(t *testing.T) {
	result := text(30)
	expected := "fizzbuzz"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}
