package api

import (
	"database/sql"
	"ddaniel27/usercore/internal/infrastructure/database"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func (a *App) setupInfrastructure() {
	user := getEnvFallback("POSTGRES_USER", "root")
	password := getEnvFallback("POSTGRES_PASSWORD", "123")
	host := getEnvFallback("POSTGRES_HOST", "localhost")
	port := getEnvFallback("POSTGRES_PORT", "5432")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/isbn?sslmode=disable", user, password, host, port)
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDB, pgdialect.New())

	userStorage := database.NewUsersStorage(db)
	credentialStorage := database.NewCredentialsStorage(db)

	a.infra = &infrastructures{
		UserStorage:       userStorage,
		CredentialStorage: credentialStorage,
	}
}
