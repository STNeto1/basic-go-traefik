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

	group := e.Group("/auth")

	group.POST("/login", h.login)
	group.POST("/register", h.register)
	group.GET("/profile", h.profile)
	group.GET("/add_header", h.addHeader)
}
