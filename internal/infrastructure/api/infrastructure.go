package api

import (
	"database/sql"
	"ddaniel27/usercore/internal/infrastructure/database"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func (a *App) setupInfrastructure() {
	dsn := "postgres://root:123@localhost:5432/isbn?sslmode=disable"
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDB, pgdialect.New())

	userStorage := database.NewUsersStorage(db)
	credentialStorage := database.NewCredentialsStorage(db)

	a.infra = &infrastructures{
		UserStorage:       userStorage,
		CredentialStorage: credentialStorage,
	}
}
