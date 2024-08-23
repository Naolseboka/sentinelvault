package gateway

const openAPISpec = `{
  "openapi": "3.0.3",
  "info": {
    "title": "SentinelVault API",
    "version": "0.1.0",
    "description": "Phase 1 API gateway surface"
  },
  "paths": {
    "/healthz": {
      "get": {
        "summary": "Liveness probe",
        "responses": {
          "200": { "description": "OK" }
        }
      }
    },
    "/readyz": {
      "get": {
        "summary": "Readiness probe",
        "responses": {
          "200": { "description": "Ready" }
        }
      }
    },
    "/v1/{service}/{id}": {
      "get": {
        "summary": "Proxy to downstream service",
        "parameters": [
          { "name": "service", "in": "path", "required": true, "schema": { "type": "string" } },
          { "name": "id", "in": "path", "required": true, "schema": { "type": "string" } }
        ],
        "responses": {
          "200": { "description": "Success" },
          "501": { "description": "Not implemented" }
        }
      }
    }
  },
  "components": {
    "securitySchemes": {
      "bearerAuth": {
        "type": "http",
        "scheme": "bearer",
        "bearerFormat": "JWT"
      }
    }
  },
  "security": [{ "bearerAuth": [] }]
}`

// OpenAPISpec returns the embedded OpenAPI document.
func OpenAPISpec() string { return openAPISpec }
