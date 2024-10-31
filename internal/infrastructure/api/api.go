package api

import (
	"context"
	"ddaniel27/usercore/internal/core/ports/repositories"
	userhandler "ddaniel27/usercore/internal/infrastructure/api/handler/users"

	"github.com/gin-gonic/gin"
)

type dependencies struct {
	UserHandler  *userhandler.UserHandler
	OtelShutdown func(context.Context) error
}

type infrastructures struct {
	UserStorage       repositories.UsersRepository
	CredentialStorage repositories.CredentialsRepository
}

type App struct {
	Server *gin.Engine
	deps   *dependencies
	infra  *infrastructures
}

func NewApp() *App {
	a := &App{}
	a.setupInfrastructure()
	a.setupDependencies()
	a.setupServer()

	return a
}
