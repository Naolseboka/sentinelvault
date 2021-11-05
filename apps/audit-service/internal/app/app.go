package app

import (
	"context"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/audit-service/internal/domain"
	httpapi "github.com/sentinelvault/sentinelvault/apps/audit-service/internal/http"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

const serviceName = "audit-service"

// Run bootstraps and runs the service.
func Run(ctx context.Context, configPath string) error {
	svc, err := platform.Bootstrap(serviceName, configPath)
	if err != nil {
		return err
	}
	chain, err := domain.NewChain()
	if err != nil {
		return err
	}
	service := domain.NewService(chain)
	handler := httpapi.NewHandler(svc, service)
	mux := http.NewServeMux()
	svc.Health.RegisterRoutes(mux)
	httpapi.RegisterRoutes(mux, handler)
	return svc.Run(ctx, mux)
}

// Service exposes the domain service for tests.
type Service = domain.Service

// NewService wires an in-memory audit service.
func NewService() (*domain.Service, error) {
	chain, err := domain.NewChain()
	if err != nil {
		return nil, err
	}
	return domain.NewService(chain), nil
}
