# API Documentation

OpenAPI spec is served at `GET /openapi.json` on the API gateway.

## Authentication

All `/v1/*` routes expect:

```
Authorization: Bearer <token>
```

## Common Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /healthz | Liveness |
| GET | /readyz | Readiness |
| GET | /v1/{resource}s | List resources |
| POST | /v1/{resource}s | Create resource |
| GET | /v1/{resource}s/{id} | Get resource |

## SDK

See `pkg/sdk/client.go` for the Go HTTP client.

```go
client := sdk.New("http://localhost:8080", sdk.WithToken("..."))
var result map[string]any
err := client.Get(ctx, "/v1/users", &result)
```
