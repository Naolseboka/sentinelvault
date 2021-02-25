package app

import (
	"context"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/invitation-service/internal/domain"
	httpapi "github.com/sentinelvault/sentinelvault/apps/invitation-service/internal/http"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

const serviceName = "invitation-service"

// Run bootstraps and runs the service.
func Run(ctx context.Context, configPath string) error {
	svc, err := platform.Bootstrap(serviceName, configPath)
	if err != nil {
		return err
	}
	repo := domain.NewMemoryRepository()
	service := domain.NewService(repo)
	handler := httpapi.NewHandler(svc, service)
	mux := http.NewServeMux()
	svc.Health.RegisterRoutes(mux)
	httpapi.RegisterRoutes(mux, handler)
	return svc.Run(ctx, mux)
}

// Service exposes the domain service for tests.
type Service = domain.Service

// NewService wires an in-memory Invitation service.
func NewService() *domain.Service {
	return domain.NewService(domain.NewMemoryRepository())
}

// maintenance note 24
