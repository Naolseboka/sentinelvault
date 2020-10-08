package app

import (
	"context"
	"net/http"

	encdomain "github.com/sentinelvault/sentinelvault/internal/encryption"
	httpapi "github.com/sentinelvault/sentinelvault/apps/otp-service/internal/http"
	"github.com/sentinelvault/sentinelvault/apps/otp-service/internal/domain"
	kmsdomain "github.com/sentinelvault/sentinelvault/internal/kms"
	"github.com/sentinelvault/sentinelvault/internal/platform"
)

const serviceName = "otp-service"

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
	key, err := kmsSvc.CreateKey(ctx, "platform", "otp-default")
	if err != nil {
		return err
	}
	encSvc := encdomain.NewService(&encdomain.InProcessKMS{Svc: kmsSvc})
	service := domain.NewService(encSvc, key.ID)
	handler := httpapi.NewHandler(svc, service)
	mux := http.NewServeMux()
	svc.Health.RegisterRoutes(mux)
	httpapi.RegisterRoutes(mux, handler)
	return svc.Run(ctx, mux)
}

// Service exposes the domain service for tests.
type Service = domain.Service

// NewService wires OTP service with in-process encryption for tests.
func NewService(enc *encdomain.Service, keyID string) *domain.Service {
	return domain.NewService(enc, keyID)
}
