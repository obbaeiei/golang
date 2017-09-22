package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func add(c echo.Context) error {
	var a, b int
	var err error
	if a, err = strconv.Atoi(c.Param("a")); err != nil {
		return c.String(http.StatusBadRequest, "invalid value a")
	}
	if b, err = strconv.Atoi(c.Param("b")); err != nil {
		return c.String(http.StatusBadRequest, "invalid value b")
	}

	return c.String(http.StatusOK, strconv.Itoa(a+b))
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", hello)
	e.GET("/add/:a/:b", add)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
