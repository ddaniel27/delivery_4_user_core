package api

import (
	"context"
	"ddaniel27/usercore/internal/core/services/users"
	userhandlers "ddaniel27/usercore/internal/infrastructure/api/handler/users"
	"ddaniel27/usercore/internal/infrastructure/observability"
	"log"
)

func (a *App) setupDependencies() {
	usersServices := users.NewUserService(a.infra.UserStorage, a.infra.CredentialStorage)

	uh := userhandlers.NewUserHandler(usersServices)

	otelShutdown, err := observability.SetupOtelSDK(context.Background())
	if err != nil {
		log.Fatalf("Failed to setup observability: %v", err)
	}

	a.deps = &dependencies{
		UserHandler:  uh,
		OtelShutdown: otelShutdown,
	}
}
