package handlers

import (
	"auth/repository"

	"github.com/labstack/echo/v4"
)

type handler struct {
	repository repository.Repository
}

func NewHandler() *handler {
	return &handler{
		repository: repository.NewMemoryRepository(),
	}
}

func RegisterRoutes(e *echo.Echo) {
	h := NewHandler()

	e.POST("/login", h.login)
	e.POST("/register", h.register)
	e.GET("/profile", h.profile)
	e.GET("/add_header", h.addHeader)
}
