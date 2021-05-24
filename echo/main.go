package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user/:id", getUser)
	e.GET("/ikan/:id", getUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id, git := c.Param("id"), "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"

	return c.JSON(http.StatusOK, id+git)
}
