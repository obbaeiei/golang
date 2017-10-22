package golang

import "testing"

func TestCase1Get1(t *testing.T) {
	result := text(1)
	expected := "1"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase2Get2(t *testing.T) {
	result := text(2)
	expected := "2"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase3GetFizz(t *testing.T) {
	result := text(3)
	expected := "fizz"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase4Get4(t *testing.T) {
	result := text(4)
	expected := "4"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase5GetBuzz(t *testing.T) {
	result := text(5)
	expected := "buzz"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestCase6GetFizz(t *testing.T) {
	result := text(6)
	expected := "fizz"
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}
