package users

import "ddaniel27/usercore/internal/core/ports/repositories"

type UserService struct {
	UserRepository       repositories.UsersRepository
	CredentialRepository repositories.CredentialsRepository
}

func NewUserService(ur repositories.UsersRepository, cr repositories.CredentialsRepository) *UserService {
	return &UserService{
		UserRepository:       ur,
		CredentialRepository: cr,
	}
}
