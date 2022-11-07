package handlers

import (
	lib "__lib"
	"log"

	"github.com/labstack/echo/v4"
)

type RegisterRequestBody struct {
	Email string `json:"email"`
}

func (h handler) register(c echo.Context) error {
	body := new(RegisterRequestBody)
	if err := c.Bind(body); err != nil {
		return c.JSON(400, map[string]string{
			"message": "Bad Request",
		})
	}

	if err := c.Validate(body); err != nil {
		return c.JSON(400, map[string]string{
			"message": "Bad Request",
		})
	}

	user, err := h.repository.Register(body.Email)
	if err != nil {
		return c.JSON(400, map[string]string{
			"message": "Bad Request",
		})
	}

	token, err := lib.GenerateJwtToken(user.ID, "some_secret")
	if err != nil {
		log.Println(err)
		return c.JSON(500, map[string]string{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(201, map[string]string{
		"token": token,
	})
}
