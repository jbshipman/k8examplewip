package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// time
var t = time.Now()

// alltings api structure
type allthing struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// allthings payload
var allthings = []allthing{
	{
		Message:   "Automate all the things!",
		Timestamp: t.UTC(),
	},
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, allthings)
	})

	e.GET("/easteregg", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct { Status string }{Status: "This is the easter egg payload."})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
