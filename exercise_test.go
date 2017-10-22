package golang

import "testing"

func TestCase1Get1(t *testing.T) {
	result := text(1)
	expected := "1"

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}
