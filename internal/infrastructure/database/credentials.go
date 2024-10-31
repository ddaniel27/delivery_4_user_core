package database

import (
	"context"
	"ddaniel27/usercore/internal/core/models"

	"github.com/uptrace/bun"
)

type CredentialsStorage struct {
	db *bun.DB
}

func NewCredentialsStorage(db *bun.DB) *CredentialsStorage {
	return &CredentialsStorage{
		db: db,
	}
}

func (s *CredentialsStorage) GetCredentialsByID(ctx context.Context, id int) (models.Credential, error) {
	var credential models.Credential
	err := s.db.NewSelect().Model(&credential).Where("id = ?", id).Scan(ctx)

	return credential, err
}

func (s *CredentialsStorage) GetCredentialsByEmail(ctx context.Context, email string) (models.Credential, error) {
	var credential models.Credential
	err := s.db.NewSelect().Model(&credential).Where("email = ?", email).Scan(ctx)

	return credential, err
}

func (s *CredentialsStorage) CreateCredentials(ctx context.Context, credential models.Credential) (models.Credential, error) {
	_, err := s.db.NewInsert().Model(&credential).Exec(ctx)

	return credential, err
}
