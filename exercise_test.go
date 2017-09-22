package exercises

import "testing"

func TestInterface1(t *testing.T) {
	var i int

	s := what(i)

	if s != "int" {
		t.Error("it should be int type but got", s)
	}
}

func TestInterface2(t *testing.T) {
	var i string

	s := what(i)

	if s != "string" {
		t.Error("it should be string type but got", s)
	}
}

func TestInterface3(t *testing.T) {
	var i bool

	s := what(i)

	if s != "bool" {
		t.Error("it should be bool type but got", s)
	}
}
