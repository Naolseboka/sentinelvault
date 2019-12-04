package app

import (
	"context"
	"net/http"
	"os"

	"github.com/sentinelvault/sentinelvault/apps/auth-service/internal/domain"
	httpapi "github.com/sentinelvault/sentinelvault/apps/auth-service/internal/http"
	"github.com/sentinelvault/sentinelvault/internal/platform"
	"github.com/sentinelvault/sentinelvault/internal/queue"
)

const serviceName = "auth-service"

// Run bootstraps and runs the service.
func Run(ctx context.Context, configPath string) error {
	svc, err := platform.Bootstrap(serviceName, configPath)
	if err != nil {
		return err
	}
	events := queue.NewMemoryProducer()
	service := domain.NewService(
		domain.NewMemoryCredentialRepository(),
		domain.NewMemorySessionRepository(),
		domain.ServiceConfig{
			JWTSecret: []byte(os.Getenv("JWT_SECRET")),
			Events:    events,
		},
	)
	handler := httpapi.NewHandler(svc, service)
	mux := http.NewServeMux()
	svc.Health.RegisterRoutes(mux)
	httpapi.RegisterRoutes(mux, handler)
	return svc.Run(ctx, mux)
}

// Service exposes the domain service for tests.
type Service = domain.Service

// NewService wires an in-memory auth service.
func NewService() *domain.Service {
	return domain.NewService(
		domain.NewMemoryCredentialRepository(),
		domain.NewMemorySessionRepository(),
		domain.ServiceConfig{JWTSecret: []byte("test-jwt-secret-for-unit-tests!!")},
	)
}
