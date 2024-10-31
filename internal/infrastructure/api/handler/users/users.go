package handler

import (
	"ddaniel27/usercore/internal/core/ports/services"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const name = "user"

type (
	UserHandler struct {
		UsersService services.UsersService
		observability
	}

	observability struct {
		tracer      trace.Tracer
		meter       metric.Meter
		logger      slog.Logger
		userCounter metric.Int64Counter
	}
)

func NewUserHandler(s services.UsersService) *UserHandler {
	var err error

	observability := observability{
		tracer: otel.Tracer(name),
		meter:  otel.Meter(name),
		logger: *otelslog.NewLogger(name),
	}

	observability.userCounter, err = observability.
		meter.
		Int64Counter(
			"users.counter",
			metric.WithDescription("Number of users created"),
			metric.WithUnit("{number}"),
		)
	if err != nil {
		observability.logger.Error(fmt.Sprintf("Failed to create counter: %s", err.Error()))
	}

	return &UserHandler{
		UsersService:  s,
		observability: observability,
	}
}
