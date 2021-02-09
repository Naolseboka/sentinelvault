package app

import (
	"context"
	"net/http"

	encdomain "github.com/sentinelvault/sentinelvault/internal/encryption"
	httpapi "github.com/sentinelvault/sentinelvault/apps/password-service/internal/http"
	"github.com/sentinelvault/sentinelvault/apps/password-service/internal/domain"
	kmsdomain "github.com/sentinelvault/sentinelvault/internal/kms"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

const serviceName = "password-service"

// Run bootstraps and runs the service.
func Run(ctx context.Context, configPath string) error {
	svc, err := platform.Bootstrap(serviceName, configPath)
	if err != nil {
		return err
	}
	reg, err := kmsdomain.NewRegistry()
	if err != nil {
		return err
	}
	kmsSvc := kmsdomain.NewService(reg)
	encSvc := encdomain.NewService(&encdomain.InProcessKMS{Svc: kmsSvc})
	service := domain.NewService(domain.NewMemoryRepository(), &encdomain.InProcessEncryptor{Svc: encSvc})
	handler := httpapi.NewHandler(svc, service)
	mux := http.NewServeMux()
	svc.Health.RegisterRoutes(mux)
	httpapi.RegisterRoutes(mux, handler)
	return svc.Run(ctx, mux)
}

// NewService wires password vault with in-process encryption for tests.
func NewService(enc *encdomain.Service) *domain.Service {
	return domain.NewService(domain.NewMemoryRepository(), &encdomain.InProcessEncryptor{Svc: enc})
}
