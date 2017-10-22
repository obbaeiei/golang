package golang

import (
	"testing"
)

func Test(t *testing.T) {
	if s != expected {
		t.Errorf("%s is expected result but got %s", expected, s)
	}
}
