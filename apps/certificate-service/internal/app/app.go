package app

import (
	"context"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/certificate-service/internal/domain"
	httpapi "github.com/sentinelvault/sentinelvault/apps/certificate-service/internal/http"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

const serviceName = "certificate-service"

// Run bootstraps and runs the service.
func Run(ctx context.Context, configPath string) error {
	svc, err := platform.Bootstrap(serviceName, configPath)
	if err != nil {
		return err
	}
	auth, err := domain.NewAuthority("SentinelVault CA")
	if err != nil {
		return err
	}
	service := domain.NewService(auth)
	handler := httpapi.NewHandler(svc, service)
	mux := http.NewServeMux()
	svc.Health.RegisterRoutes(mux)
	httpapi.RegisterRoutes(mux, handler)
	return svc.Run(ctx, mux)
}

// Service exposes the domain service for tests.
type Service = domain.Service

// NewService wires an in-memory certificate service.
func NewService() (*domain.Service, error) {
	auth, err := domain.NewAuthority("SentinelVault Test CA")
	if err != nil {
		return nil, err
	}
	return domain.NewService(auth), nil
}
