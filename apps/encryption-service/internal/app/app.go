package app

import (
	"context"
	"net/http"

	httpapi "github.com/sentinelvault/sentinelvault/apps/encryption-service/internal/http"
	encdomain "github.com/sentinelvault/sentinelvault/internal/encryption"
	kmsdomain "github.com/sentinelvault/sentinelvault/internal/kms"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

const serviceName = "encryption-service"

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
	handler := httpapi.NewHandler(svc, encSvc)
	mux := http.NewServeMux()
	svc.Health.RegisterRoutes(mux)
	httpapi.RegisterRoutes(mux, handler)
	return svc.Run(ctx, mux)
}

// NewService wires in-process KMS + encryption for tests.
func NewService(kms *kmsdomain.Service) *encdomain.Service {
	return encdomain.NewService(&encdomain.InProcessKMS{Svc: kms})
}
