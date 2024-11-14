package services

import (
	"context"
	"ddaniel27/usercore/internal/core/dto"
	"ddaniel27/usercore/internal/core/models"
)

type UsersService interface {
	GetUserByID(ctx context.Context, id int) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, user *dto.CreateUserDTO) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUserByID(ctx context.Context, id int) error
	DeleteUserByEmail(ctx context.Context, email string) error
}
