# Architecture Overview

SentinelVault is a schema-per-service monorepo implementing enterprise secrets management.

## Layers

1. **API Gateway** — routing, rate limiting, JWT validation, OpenAPI
2. **Identity Services** — auth, users, orgs, teams, RBAC, sessions
3. **Vault Services** — secrets, encryption, KMS, certificates, attachments
4. **Security Services** — policy, threat, risk, behavior, device trust
5. **Operations Services** — audit, compliance, search, backup, billing

## Shared Platform (`internal/`)

| Package | Responsibility |
|---------|----------------|
| `common` | Errors, IDs, pagination, validation |
| `config` | YAML + env configuration |
| `crypto` | AES-GCM, Argon2, HKDF, Ed25519, envelope encryption |
| `storage` | Postgres pool, migrations, outbox, transactions |
| `cache` | Redis JSON cache |
| `queue` | Kafka producer/consumer, domain events |
| `telemetry` | slog, Prometheus, OpenTelemetry |
| `middleware` | Request ID, recovery, auth, rate limit, idempotency |
| `health` | Liveness/readiness probes |
| `httpx` | HTTP server, router, JSON responses |
| `platform` | Service bootstrap lifecycle |

## Data Flow

```
Client → API Gateway → Service HTTP → Domain → Repository → Postgres (schema)
                              ↓
                         Outbox → Kafka → Consumers
```

## Deployment

- **Local**: `deploy/docker/docker-compose.yml`
- **Kubernetes**: Helm charts under `deploy/helm/charts/`
- **Terraform**: Dev-oriented modules for Postgres, Redis, MinIO
