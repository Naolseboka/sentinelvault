package app

import (
	"context"
	"net/http"

	"github.com/sentinelvault/sentinelvault/apps/attachment-service/internal/domain"
	httpapi "github.com/sentinelvault/sentinelvault/apps/attachment-service/internal/http"
	encdomain "github.com/sentinelvault/sentinelvault/internal/encryption"
	kmsdomain "github.com/sentinelvault/sentinelvault/internal/kms"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

const serviceName = "attachment-service"

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
	service := domain.NewService(domain.NewMemoryRepository(), &encdomain.InProcessEncryptor{Svc: encSvc}, "")
	handler := httpapi.NewHandler(svc, service)
	mux := http.NewServeMux()
	svc.Health.RegisterRoutes(mux)
	httpapi.RegisterRoutes(mux, handler)
	return svc.Run(ctx, mux)
}
