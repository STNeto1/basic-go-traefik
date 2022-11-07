package repository

import (
	errors "auth/__errors"
	"auth/models"

	"github.com/gofrs/uuid"
)

type MemoryRepository struct {
	users map[string]*models.User
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users: map[string]*models.User{},
	}
}

func (r *MemoryRepository) GetUser(id string) (*models.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, errors.ErrUserNotFound{}
	}

	return user, nil
}

func (r *MemoryRepository) Login(email string) (*models.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.ErrUserNotFound{}
}

func (r *MemoryRepository) Register(email string) (*models.User, error) {
	existingUser, err := r.Login(email)
	if err == nil {
		return existingUser, nil
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:    uuid.String(),
		Email: email,
	}

	r.users[user.ID] = user

	return user, nil
}
