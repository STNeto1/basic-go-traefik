package main

import (
	"auth/handlers"
	"auth/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/exp/slices"
)

func main() {
	e := echo.New()

	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("some_secret"),
		TokenLookup: "header:Authorization",
		ContextKey:  "user",
		Claims:      jwt.MapClaims{},
		Skipper: func(c echo.Context) bool {
			publicEndpoints := []string{"/login", "/register"}
			return slices.Contains(publicEndpoints, c.Path())
		},
		ErrorHandler: func(err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
				"message":     "Unauthorized",
				"status_code": http.StatusUnauthorized,
			})
		},
	}))

	handlers.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":4000"))
}
