package main

import (
	lib "__lib"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(lib.CheckUserHeader())

	e.GET("/greeting", func(c echo.Context) error {
		userId := c.Request().Header.Get("X-User-Id")
		return c.JSON(200, map[string]string{
			"message": fmt.Sprintf("Hello, %s!", userId),
		})
	})

	e.Logger.Fatal(e.Start(":5000"))
}
