package telemetry

import (
	"context"
	"log/slog"
)

// Bundle groups observability primitives for a service.
type Bundle struct {
	Logger         *slog.Logger
	ShutdownTracer func(context.Context) error
}

// Start initializes logging and tracing for a service.
func Start(ctx context.Context, serviceName, logLevel, otelEndpoint string) (*Bundle, error) {
	logger := NewLogger(logLevel).With("service", serviceName)
	shutdown, err := InitTracer(ctx, serviceName, otelEndpoint)
	if err != nil {
		return nil, err
	}
	return &Bundle{Logger: logger, ShutdownTracer: shutdown}, nil
}
