package app

import (
	"context"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/compliance-service/internal/domain"
	httpapi "github.com/sentinelvault/sentinelvault/apps/compliance-service/internal/http"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

const serviceName = "compliance-service"

// Run bootstraps and runs the service.
func Run(ctx context.Context, configPath string) error {
	svc, err := platform.Bootstrap(serviceName, configPath)
	if err != nil {
		return err
	}
	service := domain.NewService(domain.NewCollector())
	handler := httpapi.NewHandler(svc, service)
	mux := http.NewServeMux()
	svc.Health.RegisterRoutes(mux)
	httpapi.RegisterRoutes(mux, handler)
	return svc.Run(ctx, mux)
}

// Service exposes the domain service for tests.
type Service = domain.Service

// NewService wires an in-memory compliance service.
func NewService() *domain.Service {
	return domain.NewService(domain.NewCollector())
}
