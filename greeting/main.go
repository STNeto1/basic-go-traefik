package main

import (
	lib "__lib"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(lib.CheckUserHeader())

	e.GET("/", func(c echo.Context) error {
		userId := c.Request().Header.Get("X-User-Id")
		return c.JSON(200, map[string]string{
			"message": fmt.Sprintf("Hello, %s!", userId),
		})
	})

	e.Logger.Fatal(e.Start(":5000"))
}
