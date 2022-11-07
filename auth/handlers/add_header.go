package handlers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (h handler) addHeader(c echo.Context) error {
	payload := c.Get("user").(*jwt.Token)
	claims := payload.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)

	user, err := h.repository.GetUser(id)
	if err != nil {
		return c.JSON(400, map[string]string{
			"message": "Bad Request",
		})
	}

	c.Response().Header().Set("X-User-ID", user.ID)

	return c.NoContent(200)
}
