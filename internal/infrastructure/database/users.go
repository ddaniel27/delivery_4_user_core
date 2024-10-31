package database

import (
	"context"
	"ddaniel27/usercore/internal/core/models"

	"github.com/uptrace/bun"
)

type UsersStorage struct {
	db *bun.DB
}

func NewUsersStorage(db *bun.DB) *UsersStorage {
	return &UsersStorage{
		db: db,
	}
}

func (s *UsersStorage) GetUserByID(ctx context.Context, id int) (models.User, error) {
	var user models.User
	err := s.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)

	return user, err
}

func (s *UsersStorage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := s.db.NewSelect().Model(&user).Where("email = ?", email).Scan(ctx)

	return user, err
}

func (s *UsersStorage) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	_, err := s.db.NewInsert().Model(&user).Exec(ctx)

	return user, err
}

func (s *UsersStorage) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	_, err := s.db.NewUpdate().Model(&user).WherePK().Exec(ctx)

	return user, err
}

func (s *UsersStorage) DeleteUserByID(ctx context.Context, id int) error {
	_, err := s.db.NewDelete().Model(&models.User{}).Where("id = ?", id).Exec(ctx)

	return err
}

func (s *UsersStorage) DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := s.db.NewDelete().Model(&models.User{}).Where("email = ?", email).Exec(ctx)

	return err
}
