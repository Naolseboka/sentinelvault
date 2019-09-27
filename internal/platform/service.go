package platform

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/config"
	"github.com/sentinelvault/sentinelvault/internal/health"
	"github.com/sentinelvault/sentinelvault/internal/httpx"
	"github.com/sentinelvault/sentinelvault/internal/middleware"
	"github.com/sentinelvault/sentinelvault/internal/telemetry"
)

// Service wires config, telemetry, health, and HTTP lifecycle.
type Service struct {
	Name   string
	Config *config.ServiceConfig
	Tel    *telemetry.Bundle
	Health *health.Registry
	HTTP   *httpx.Server
}

// Bootstrap loads config and initializes observability for a service.
func Bootstrap(name, configPath string) (*Service, error) {
	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}
	_ = os.Setenv("SERVICE_NAME", name)
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}
	cfg.Name = name
	ctx := context.Background()
	tel, err := telemetry.Start(ctx, name, cfg.LogLevel, cfg.OTelEndpoint)
	if err != nil {
		return nil, fmt.Errorf("telemetry: %w", err)
	}
	reg := health.NewRegistry()
	return &Service{
		Name:   name,
		Config: cfg,
		Tel:    tel,
		Health: reg,
	}, nil
}

// SetHTTP attaches the HTTP server with standard middleware.
func (s *Service) SetHTTP(handler http.Handler) {
	wrapped := middleware.RequestID(
		middleware.Recovery(s.Tel.Logger)(
			middleware.AccessLog(s.Tel.Logger)(handler),
		),
	)
	s.HTTP = httpx.NewServer(s.Config.HTTPAddr, wrapped)
}

// Run starts HTTP and blocks until ctx is cancelled or a signal arrives.
func (s *Service) Run(ctx context.Context, handler http.Handler) error {
	if s.HTTP == nil {
		s.SetHTTP(handler)
	}
	errCh := make(chan error, 1)
	go func() {
		s.Tel.Logger.Info("starting http server", "addr", s.Config.HTTPAddr)
		errCh <- s.HTTP.ListenAndServe()
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
	case sig := <-sigCh:
		s.Tel.Logger.Info("signal received", "signal", sig.String())
	case err := <-errCh:
		return err
	}
	shCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := s.HTTP.Shutdown(shCtx); err != nil {
		return err
	}
	if s.Tel.ShutdownTracer != nil {
		_ = s.Tel.ShutdownTracer(shCtx)
	}
	return nil
}
