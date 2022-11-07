package lib

import "github.com/labstack/echo/v4"

func CheckUserHeader() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userId := c.Request().Header.Get("X-User-ID")
			if userId == "" {
				return c.JSON(400, map[string]string{
					"message": "Unauthorized",
				})
			}

			return next(c)
		}
	}
}
