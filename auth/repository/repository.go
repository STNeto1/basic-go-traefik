package repository

import "auth/models"

type Repository interface {
	GetUser(id string) (*models.User, error)
	Login(email string) (*models.User, error)
	Register(email string) (*models.User, error)
}
