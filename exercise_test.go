package golang

import (
	"errors"
	"strconv"
	"testing"
)

type number int

func (n number) String() string {
	return strconv.Itoa(int(n))
}

func TestMyPrintfErrorAndStringer(t *testing.T) {
	err := errors.New("something when wrong!!")
	n := number(12)

	s := myPrintf(err, n)
	if s != "something when wrong!! 12" {
		t.Error("it shhould work with fmt.Stringer and error as well")
	}
}

func TestMyPrintfAny(t *testing.T) {
	err := errors.New("something when wrong!!")
	n := number(12)

	s := myPrintf("abc", err, n, 123)
	if s != "abc something when wrong!! 12 123" {
		t.Error("it shhould work with fmt.Stringer and error as well")
	}
}
