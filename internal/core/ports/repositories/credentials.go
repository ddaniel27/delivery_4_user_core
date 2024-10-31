package repositories

import (
	"context"
	"ddaniel27/usercore/internal/core/models"
)

type CredentialsRepository interface {
	GetCredentialsByID(ctx context.Context, id int) (models.Credential, error)
	GetCredentialsByEmail(ctx context.Context, email string) (models.Credential, error)
	CreateCredentials(ctx context.Context, credential models.Credential) (models.Credential, error)
}
