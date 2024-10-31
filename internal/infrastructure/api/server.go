package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *App) setupServer() {
	a.Server = gin.New()
	a.Server.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	rootGroup := a.Server.Group("/")
	setupHealthCheckRoute(rootGroup)

	baseGroup := a.Server.Group("/api")
	a.setupRoutes(baseGroup)
}

func (a *App) setupRoutes(g *gin.RouterGroup) {
	handler := a.deps.UserHandler

	recordGroup := g.Group("/user")

	recordGroup.POST("", handler.CreateUser)
	recordGroup.GET("/:id", handler.GetUserByID)
	recordGroup.GET("/email", handler.GetUserByEmail)
}

func (a *App) startServer() {
	if err := a.Server.Run(fmt.Sprintf(":%s", getEnvFallback("PORT", "3000"))); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (a *App) stopApp() {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.deps.OtelShutdown(context.Background()); err != nil {
		log.Fatalf("Failed to shutdown observability: %v", err)
	}
}

func (a *App) StartApp() {
	a.startServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	a.stopApp()
}

func setupHealthCheckRoute(g *gin.RouterGroup) {
	g.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
}

func getEnvFallback(env string, fallback string) string {
	port := os.Getenv(env)
	if port == "" {
		return fallback
	}
	return port
}
