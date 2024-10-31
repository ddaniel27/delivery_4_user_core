package users

import (
	"context"
	"ddaniel27/usercore/internal/core/models"
)

func (us *UserService) UpdateUser(_ context.Context, _ models.User) (models.User, error) {
	return models.User{}, nil
}
