# Coding Standards

## Go

- Go 1.22+ with standard `net/http` routing patterns
- Module path: `github.com/sentinelvault/sentinelvault`
- Typed errors via `internal/common.AppError`
- ULIDs for identifiers (`internal/common.NewID()`)
- Context passed as first parameter

## Services

Each service follows:

```
apps/<name>/
  cmd/main.go
  internal/app/app.go
  internal/domain/domain.go
  internal/http/handlers.go
```

## Testing

- Unit tests required for shared libraries
- In-memory repositories for service domain tests
- Table-driven tests preferred

## Security

- Envelope encryption for secrets at rest (see ADR-0003)
- Argon2id for password hashing
- No secrets in config files or logs
