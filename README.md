# SentinelVault

Enterprise Identity, Secrets & Privileged Access Management platform (Go monorepo).

## Phase Status

| Phase | Scope | Status |
|-------|-------|--------|
| 1 — Platform Foundation | Monorepo, shared libs, gateway, config/telemetry, CI/CD, Docker/K8s/Terraform, docs, 42 service shells | Complete |
| 2 — Core Security | Auth (OAuth2/JWT/MFA/WebAuthn), RBAC/ABAC, users, orgs, sessions, device trust | Complete |
| 3 — Vault | Encryption, KMS, password/secret vault, attachments, versioning (ciphertext at rest) | Complete |
| 4 — Enterprise | Audit hash-chain, search, notifications, reporting, compliance, rotation, PAM | Complete |
| 5 — Advanced Security | Risk/behavior/threat, PKI CA, SSH, secret engines (DB/AWS/Azure/GCP/K8s/Transit/GitHub/GitLab) | Complete |
| 6 — Operations | Backup/DR, billing/licensing, admin, Istio policies, GraphQL health, runbooks | Complete |

Authored Go: **~252k LOC** of domain logic, persistence, validators, engines, SDKs, and tests. Track with `make loc` / `./scripts/loc.sh`.

## Quick Start

```bash
cp .env.example .env
go mod tidy
make test
make build-gateway
./bin/api-gateway
```

Local infra:

```bash
make docker-up
```

MVP integration test (auth → KMS/encrypt → vault validation → signed audit):

```bash
go test ./tests/integration/...
```

Demo script (with auth-service + gateway running):

```bash
./scripts/demo-mvp.sh
```

Use the SDK for auth flows:

```go
client := sdk.New("http://localhost:8080")
auth := client.Auth()
_, err := auth.Register(ctx, sdk.RegisterRequest{Email: "you@example.com", Password: "secret", OrgName: "My Org"})
tokens, err := auth.Login(ctx, sdk.LoginRequest{Email: "you@example.com", Password: "secret"})
```

## Project Layout

```
apps/           # 42 microservices + api-gateway
internal/       # Shared platform libraries
pkg/sdk/        # Public HTTP client SDK
configs/        # Service YAML configuration
migrations/     # Per-service SQL migrations
proto/          # Protobuf service definitions
deploy/         # Docker, Helm, Terraform, Kubernetes
docs/           # Architecture, ADRs, runbooks
```

## Services

All services expose `/healthz`, `/readyz`, and `/v1/*` REST endpoints. The API gateway routes `/v1/*` to upstream services.

## Development

```bash
make bootstrap    # verify Go toolchain
make lint         # golangci-lint
make test-all     # all unit tests
make docker-up    # local infra stack
```

## License

Apache-2.0 — see [LICENSE](LICENSE).
