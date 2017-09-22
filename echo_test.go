package main

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

type mockHandler struct {
	echo.Context
	code int
	v    string
}

func (c *mockHandler) Param(s string) string {
	if s == "a" {
		return "1"
	}
	return "2"
}

func (c *mockHandler) String(code int, s string) error {
	c.code = code
	c.v = s
	return nil
}

func TestHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/add/1/2", nil)
	rec := httptest.NewRecorder()
	c := mockHandler{e.NewContext(req, rec), 0, ""}

	err := add(&c)

	fmt.Println(err)
	fmt.Println(c.v)
}
