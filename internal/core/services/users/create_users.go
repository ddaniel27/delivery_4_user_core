package users

import (
	"context"
	"ddaniel27/usercore/internal/core/dto"
	"ddaniel27/usercore/internal/core/models"
)

func (us *UserService) CreateUser(ctx context.Context, u *dto.CreateUserDTO) (models.User, error) {
	// Create credentials
	credential := models.Credential{
		Email:    u.Email,
		Password: u.Password,
	}
	c, err := us.CredentialRepository.CreateCredentials(ctx, credential)
	if err != nil {
		return models.User{}, err
	}

	// Create the user
	user := models.User{
		Email:         u.Email,
		Name:          u.Name,
		Institution:   u.Inst,
		City:          u.City,
		Birthdate:     u.Birthdate,
		CredentialsID: c.ID,
	}

	return us.UserRepository.CreateUser(ctx, user)
}
