package handlers

import (
	"log"

	"github.com/labstack/echo/v4"

	lib "__lib"
)

type LoginRequestBody struct {
	Email string `json:"email" validate:"required,email"`
}

func (h handler) login(c echo.Context) error {
	body := new(LoginRequestBody)
	if err := c.Bind(body); err != nil {
		return c.JSON(400, map[string]string{
			"message": "Bad Request",
		})
	}

	if err := c.Validate(body); err != nil {
		return c.JSON(400, map[string]string{
			"message": "Invalid Request",
		})
	}

	user, err := h.repository.Login(body.Email)
	if err != nil {
		return c.JSON(400, map[string]string{
			"message": "User not found",
		})
	}

	token, err := lib.GenerateJwtToken(user.ID, "some_secret")
	if err != nil {
		log.Println(err)
		return c.JSON(500, map[string]string{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(200, map[string]string{
		"token": token,
	})
}
