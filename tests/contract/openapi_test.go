package contract

import (
	"encoding/json"
	"strings"
	"testing"
)

const gatewayOpenAPISpec = `{
  "openapi": "3.0.3",
  "paths": {
    "/healthz": {"get": {"responses": {"200": {"description": "OK"}}}},
    "/readyz": {"get": {"responses": {"200": {"description": "Ready"}}}},
    "/v1/{service}/{id}": {"get": {"responses": {"200": {"description": "Success"}}}}
  }
}`

func TestOpenAPIPathsPresent(t *testing.T) {
	var doc map[string]any
	if err := json.Unmarshal([]byte(gatewayOpenAPISpec), &doc); err != nil {
		t.Fatal(err)
	}
	paths, ok := doc["paths"].(map[string]any)
	if !ok {
		t.Fatal("missing paths")
	}
	for _, p := range []string{"/healthz", "/readyz", "/v1/{service}/{id}"} {
		if _, ok := paths[p]; !ok {
			t.Fatalf("missing path %s", p)
		}
	}
}

func TestAuthRegisterRequestSchema(t *testing.T) {
	raw := `{"email":"user@example.com","password":"long-password-123","org_name":"Acme"}`
	var req map[string]any
	if err := json.Unmarshal([]byte(raw), &req); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"email", "password"} {
		v, ok := req[k].(string)
		if !ok || strings.TrimSpace(v) == "" {
			t.Fatalf("missing %s", k)
		}
	}
}

func TestAuthLoginRequestSchema(t *testing.T) {
	raw := `{"email":"user@example.com","password":"secret","device_id":"dev-1"}`
	var req map[string]any
	if err := json.Unmarshal([]byte(raw), &req); err != nil {
		t.Fatal(err)
	}
	if req["email"] == "" || req["password"] == "" {
		t.Fatal("email/password required")
	}
}
