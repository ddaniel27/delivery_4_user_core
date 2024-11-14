package users

import (
	"context"
	"ddaniel27/usercore/internal/core/models"
)

func (us *UserService) GetUserByID(ctx context.Context, id int) (models.User, error) {
	u, err := us.UserRepository.GetUserByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	u.CredentialsID = -1 // don't return the credentials id

	return u, nil
}

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	u, err := us.UserRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}

	u.CredentialsID = -1 // don't return the credentials id

	return u, nil
}

func (us *UserService) GetUsers(ctx context.Context) ([]models.User, error) {
	users, err := us.UserRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	for i := range users {
		users[i].CredentialsID = -1 // don't return the credentials id
	}

	return users, nil
}
