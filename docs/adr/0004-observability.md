# ADR-0004: Observability Stack

## Status
Accepted

## Context
42 services require unified logging, metrics, and tracing.

## Decision
- **Logging**: JSON slog to stdout
- **Metrics**: Prometheus counters (HTTP requests)
- **Tracing**: OpenTelemetry OTLP gRPC (noop when endpoint empty)

Bootstrap via `internal/telemetry.Start()`.

## Consequences
- Local stack includes Prometheus, Grafana, Loki, Tempo, Jaeger
- Services expose `/metrics` in Phase 2
