#!/usr/bin/env python3
"""Generate SentinelVault microservice skeletons."""

from __future__ import annotations

import os
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]

SERVICES = [
    "api-gateway",
    "auth-service",
    "authorization-service",
    "user-service",
    "organization-service",
    "team-service",
    "rbac-service",
    "policy-service",
    "device-service",
    "session-service",
    "invitation-service",
    "password-service",
    "secret-service",
    "encryption-service",
    "kms-service",
    "certificate-service",
    "ssh-service",
    "passkey-service",
    "otp-service",
    "attachment-service",
    "versioning-service",
    "approval-service",
    "access-request-service",
    "emergency-access-service",
    "delegation-service",
    "lifecycle-service",
    "provisioning-service",
    "scim-service",
    "directory-sync-service",
    "threat-service",
    "risk-service",
    "behavior-service",
    "device-trust-service",
    "notification-service",
    "audit-service",
    "compliance-service",
    "search-service",
    "backup-service",
    "rotation-service",
    "admin-service",
    "reporting-service",
    "billing-service",
]

# domain metadata per service
SERVICE_META = {
    "api-gateway": {"domain": "gateway", "entity": "Route", "schema": "gateway", "table": "routes"},
    "auth-service": {"domain": "identity", "entity": "Credential", "schema": "auth", "table": "credentials"},
    "authorization-service": {"domain": "identity", "entity": "Decision", "schema": "authorization", "table": "decisions"},
    "user-service": {"domain": "identity", "entity": "User", "schema": "users", "table": "users"},
    "organization-service": {"domain": "identity", "entity": "Organization", "schema": "organizations", "table": "organizations"},
    "team-service": {"domain": "identity", "entity": "Team", "schema": "teams", "table": "teams"},
    "rbac-service": {"domain": "identity", "entity": "RoleBinding", "schema": "rbac", "table": "role_bindings"},
    "policy-service": {"domain": "security", "entity": "Policy", "schema": "policy", "table": "policies"},
    "device-service": {"domain": "identity", "entity": "Device", "schema": "devices", "table": "devices"},
    "session-service": {"domain": "identity", "entity": "Session", "schema": "sessions", "table": "sessions"},
    "invitation-service": {"domain": "identity", "entity": "Invitation", "schema": "invitations", "table": "invitations"},
    "password-service": {"domain": "vault", "entity": "PasswordEntry", "schema": "passwords", "table": "password_entries"},
    "secret-service": {"domain": "vault", "entity": "Secret", "schema": "secrets", "table": "secrets"},
    "encryption-service": {"domain": "vault", "entity": "EncryptionKey", "schema": "encryption", "table": "encryption_keys"},
    "kms-service": {"domain": "vault", "entity": "MasterKey", "schema": "kms", "table": "master_keys"},
    "certificate-service": {"domain": "vault", "entity": "Certificate", "schema": "certificates", "table": "certificates"},
    "ssh-service": {"domain": "vault", "entity": "SSHKey", "schema": "ssh", "table": "ssh_keys"},
    "passkey-service": {"domain": "identity", "entity": "Passkey", "schema": "passkeys", "table": "passkeys"},
    "otp-service": {"domain": "identity", "entity": "OTPCode", "schema": "otp", "table": "otp_codes"},
    "attachment-service": {"domain": "vault", "entity": "Attachment", "schema": "attachments", "table": "attachments"},
    "versioning-service": {"domain": "vault", "entity": "Version", "schema": "versioning", "table": "versions"},
    "approval-service": {"domain": "operations", "entity": "Approval", "schema": "approvals", "table": "approvals"},
    "access-request-service": {"domain": "operations", "entity": "AccessRequest", "schema": "access_requests", "table": "access_requests"},
    "emergency-access-service": {"domain": "operations", "entity": "EmergencyGrant", "schema": "emergency_access", "table": "emergency_grants"},
    "delegation-service": {"domain": "operations", "entity": "Delegation", "schema": "delegations", "table": "delegations"},
    "lifecycle-service": {"domain": "operations", "entity": "LifecyclePolicy", "schema": "lifecycle", "table": "lifecycle_policies"},
    "provisioning-service": {"domain": "operations", "entity": "ProvisionJob", "schema": "provisioning", "table": "provision_jobs"},
    "scim-service": {"domain": "operations", "entity": "SCIMResource", "schema": "scim", "table": "scim_resources"},
    "directory-sync-service": {"domain": "operations", "entity": "DirectorySync", "schema": "directory_sync", "table": "directory_syncs"},
    "threat-service": {"domain": "security", "entity": "ThreatSignal", "schema": "threat", "table": "threat_signals"},
    "risk-service": {"domain": "security", "entity": "RiskScore", "schema": "risk", "table": "risk_scores"},
    "behavior-service": {"domain": "security", "entity": "BehaviorEvent", "schema": "behavior", "table": "behavior_events"},
    "device-trust-service": {"domain": "security", "entity": "TrustAssessment", "schema": "device_trust", "table": "trust_assessments"},
    "notification-service": {"domain": "operations", "entity": "Notification", "schema": "notifications", "table": "notifications"},
    "audit-service": {"domain": "audit", "entity": "AuditEvent", "schema": "audit", "table": "audit_events"},
    "compliance-service": {"domain": "audit", "entity": "ComplianceReport", "schema": "compliance", "table": "compliance_reports"},
    "search-service": {"domain": "operations", "entity": "SearchDocument", "schema": "search", "table": "search_documents"},
    "backup-service": {"domain": "operations", "entity": "Backup", "schema": "backups", "table": "backups"},
    "rotation-service": {"domain": "vault", "entity": "RotationJob", "schema": "rotation", "table": "rotation_jobs"},
    "admin-service": {"domain": "operations", "entity": "AdminAction", "schema": "admin", "table": "admin_actions"},
    "reporting-service": {"domain": "operations", "entity": "Report", "schema": "reporting", "table": "reports"},
    "billing-service": {"domain": "operations", "entity": "Subscription", "schema": "billing", "table": "subscriptions"},
}


def pkg(name: str) -> str:
    return name.replace("-", "")


def write(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    if not path.exists() or path.read_text() != content:
        path.write_text(content)
        print(f"wrote {path.relative_to(ROOT)}")


def gen_main(name: str) -> str:
    p = pkg(name)
    if name == "api-gateway":
        return f'''package main

import (
\t"context"
\t"log"
\t"os"

\t"github.com/sentinelvault/sentinelvault/apps/{name}/internal/gateway"
)

func main() {{
\tctx := context.Background()
\tconfigPath := os.Getenv("CONFIG_PATH")
\tif configPath == "" {{
\t\tconfigPath = "configs/{name}.yaml"
\t}}
\tif err := gateway.Run(ctx, configPath); err != nil {{
\t\tlog.Fatal(err)
\t}}
}}
'''
    return f'''package main

import (
\t"context"
\t"log"
\t"os"

\t"github.com/sentinelvault/sentinelvault/apps/{name}/internal/app"
)

func main() {{
\tctx := context.Background()
\tconfigPath := os.Getenv("CONFIG_PATH")
\tif configPath == "" {{
\t\tconfigPath = "configs/{name}.yaml"
\t}}
\tif err := app.Run(ctx, configPath); err != nil {{
\t\tlog.Fatal(err)
\t}}
}}
'''


def gen_app(name: str, meta: dict) -> str:
    entity = meta["entity"]
    return f'''package app

import (
\t"context"
\t"net/http"

\t"github.com/sentinelvault/sentinelvault/apps/{name}/internal/domain"
\thttpapi "github.com/sentinelvault/sentinelvault/apps/{name}/internal/http"
\t"github.com/sentinelvault/sentinelvault/internal/platform"
)

const serviceName = "{name}"

// Run bootstraps and runs the service.
func Run(ctx context.Context, configPath string) error {{
\tsvc, err := platform.Bootstrap(serviceName, configPath)
\tif err != nil {{
\t\treturn err
\t}}
\trepo := domain.NewMemoryRepository()
\tservice := domain.NewService(repo)
\thandler := httpapi.NewHandler(svc, service)
\tmux := http.NewServeMux()
\tsvc.Health.RegisterRoutes(mux)
\thttpapi.RegisterRoutes(mux, handler)
\treturn svc.Run(ctx, mux)
}}

// Service exposes the domain service for tests.
type Service = domain.Service

// NewService wires an in-memory {entity} service.
func NewService() *domain.Service {{
\treturn domain.NewService(domain.NewMemoryRepository())
}}
'''


def gen_domain(name: str, meta: dict) -> str:
    entity = meta["entity"]
    schema = meta["schema"]
    table = meta["table"]
    return f'''package domain

import (
\t"context"
\t"sync"
\t"time"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

// {entity} is the core aggregate for {name}.
type {entity} struct {{
\tID        string    `json:"id"`
\tOrgID     string    `json:"org_id"`
\tName      string    `json:"name"`
\tStatus    string    `json:"status"`
\tCreatedAt time.Time `json:"created_at"`
\tUpdatedAt time.Time `json:"updated_at"`
}}

// Repository persists {entity} records.
type Repository interface {{
\tCreate(ctx context.Context, item *{entity}) error
\tGet(ctx context.Context, id string) (*{entity}, error)
\tList(ctx context.Context, orgID string, offset, limit int) ([]{entity}, int64, error)
}}

// Service implements {name} business logic.
type Service struct {{
\trepo Repository
}}

// NewService creates a domain service.
func NewService(repo Repository) *Service {{
\treturn &Service{{repo: repo}}
}}

// Create registers a new {entity}.
func (s *Service) Create(ctx context.Context, orgID, name string) (*{entity}, error) {{
\tif err := common.ValidateRequired("name", name); err != nil {{
\t\treturn nil, err
\t}}
\tnow := common.NowUTC()
\titem := &{entity}{{
\t\tID:        common.NewID(),
\t\tOrgID:     orgID,
\t\tName:      name,
\t\tStatus:    "active",
\t\tCreatedAt: now,
\t\tUpdatedAt: now,
\t}}
\tif err := s.repo.Create(ctx, item); err != nil {{
\t\treturn nil, err
\t}}
\treturn item, nil
}}

// Get returns a {entity} by ID.
func (s *Service) Get(ctx context.Context, id string) (*{entity}, error) {{
\titem, err := s.repo.Get(ctx, id)
\tif err != nil {{
\t\treturn nil, err
\t}}
\tif item == nil {{
\t\treturn nil, common.New(common.ErrCodeNotFound, "{entity} not found")
\t}}
\treturn item, nil
}}

// List returns paginated {entity} records for an org.
func (s *Service) List(ctx context.Context, orgID string, offset, limit int) ([]{entity}, int64, error) {{
\treturn s.repo.List(ctx, orgID, offset, limit)
}}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string {{ return "{schema}" }}

// TableName returns the primary table name.
func TableName() string {{ return "{table}" }}

// MemoryRepository is an in-memory Repository for development.
type MemoryRepository struct {{
\tmu    sync.RWMutex
\tdata map[string]*{entity}
}}

// NewMemoryRepository creates an empty in-memory store.
func NewMemoryRepository() *MemoryRepository {{
\treturn &MemoryRepository{{data: map[string]*{entity}{{}}}}
}}

func (r *MemoryRepository) Create(_ context.Context, item *{entity}) error {{
\tr.mu.Lock()
\tdefer r.mu.Unlock()
\tcp := *item
\tr.data[item.ID] = &cp
\treturn nil
}}

func (r *MemoryRepository) Get(_ context.Context, id string) (*{entity}, error) {{
\tr.mu.RLock()
\tdefer r.mu.RUnlock()
\titem, ok := r.data[id]
\tif !ok {{
\t\treturn nil, nil
\t}}
\tcp := *item
\treturn &cp, nil
}}

func (r *MemoryRepository) List(_ context.Context, orgID string, offset, limit int) ([]{entity}, int64, error) {{
\tr.mu.RLock()
\tdefer r.mu.RUnlock()
\tall := make([]{entity}, 0)
\tfor _, item := range r.data {{
\t\tif orgID == "" || item.OrgID == orgID {{
\t\t\tall = append(all, *item)
\t\t}}
\t}}
\ttotal := int64(len(all))
\tif offset >= len(all) {{
\t\treturn []{entity}{{}}, total, nil
\t}}
\tend := offset + limit
\tif end > len(all) {{
\t\tend = len(all)
\t}}
\treturn all[offset:end], total, nil
}}
'''


def gen_handlers(name: str, meta: dict) -> str:
    entity = meta["entity"]
    route = name.replace("-service", "").replace("api-gateway", "gateway")
    return f'''package http

import (
\t"encoding/json"
\t"net/http"

\t"github.com/sentinelvault/sentinelvault/apps/{name}/internal/domain"
\t"github.com/sentinelvault/sentinelvault/internal/common"
\t"github.com/sentinelvault/sentinelvault/internal/httpx"
\t"github.com/sentinelvault/sentinelvault/internal/platform"
)

// Handler serves HTTP endpoints for {name}.
type Handler struct {{
\tplatform *platform.Service
\tservice  *domain.Service
}}

// NewHandler creates an HTTP handler.
func NewHandler(p *platform.Service, svc *domain.Service) *Handler {{
\treturn &Handler{{platform: p, service: svc}}
}}

// RegisterRoutes mounts service routes on mux.
func RegisterRoutes(mux *http.ServeMux, h *Handler) {{
\tmux.HandleFunc("GET /v1/{route}s", h.list)
\tmux.HandleFunc("POST /v1/{route}s", h.create)
\tmux.HandleFunc("GET /v1/{route}s/{{id}}", h.get)
\tmux.HandleFunc("GET /v1/meta", h.meta)
}}

func (h *Handler) meta(w http.ResponseWriter, r *http.Request) {{
\thttpx.WriteJSON(w, http.StatusOK, map[string]string{{
\t\t"service": "{name}",
\t\t"schema":  domain.SchemaName(),
\t\t"table":   domain.TableName(),
\t}})
}}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {{
\tvar req struct {{
\t\tOrgID string `json:"org_id"`
\t\tName  string `json:"name"`
\t}}
\tif err := json.NewDecoder(r.Body).Decode(&req); err != nil {{
\t\thttpx.WriteError(w, http.StatusBadRequest, common.New(common.ErrCodeInvalidArgument, "invalid json"))
\t\treturn
\t}}
\titem, err := h.service.Create(r.Context(), req.OrgID, req.Name)
\tif err != nil {{
\t\thttpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
\t\treturn
\t}}
\thttpx.WriteJSON(w, http.StatusCreated, item)
}}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {{
\tid := r.PathValue("id")
\titem, err := h.service.Get(r.Context(), id)
\tif err != nil {{
\t\thttpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
\t\treturn
\t}}
\thttpx.WriteJSON(w, http.StatusOK, item)
}}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {{
\tparams := common.ParsePageParams(r)
\torgID := r.URL.Query().Get("org_id")
\titems, total, err := h.service.List(r.Context(), orgID, params.Offset, params.PageSize)
\tif err != nil {{
\t\thttpx.WriteError(w, httpx.ErrorStatus(common.AsAppError(err)), common.AsAppError(err))
\t\treturn
\t}}
\thttpx.WriteJSON(w, http.StatusOK, map[string]any{{
\t\t"data":       items,
\t\t"pagination": common.NewPageMeta(params, total),
\t}})
}}
'''


def gen_dockerfile(name: str) -> str:
    return f'''FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/{name} ./apps/{name}/cmd

FROM gcr.io/distroless/static-debian12
COPY --from=builder /bin/{name} /{name}
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/{name}"]
'''


def gen_config(name: str) -> str:
    port_base = 8080 + (hash(name) % 1000)
    return f'''name: {name}
http_addr: ":{port_base}"
grpc_addr: ":{port_base + 1000}"
database_url: postgres://sentinel:sentinel@localhost:5432/sentinelvault?sslmode=disable
redis_url: redis://localhost:6379/0
kafka_brokers:
  - localhost:9092
log_level: info
otel_endpoint: ""
feature_flags:
  phase1: true
'''


def gen_migration(name: str, meta: dict) -> str:
    schema = meta["schema"]
    table = meta["table"]
    return f'''CREATE SCHEMA IF NOT EXISTS {schema};

CREATE TABLE IF NOT EXISTS {schema}.{table} (
    id          TEXT PRIMARY KEY,
    org_id      TEXT NOT NULL,
    name        TEXT NOT NULL,
    status      TEXT NOT NULL DEFAULT 'active',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_{table}_org_id ON {schema}.{table} (org_id);
'''


def gen_proto(name: str, meta: dict) -> str:
    entity = meta["entity"]
    domain = meta["domain"]
    svc = "".join(p.capitalize() for p in name.replace("-service", "").split("-"))
    return f'''syntax = "proto3";

package sentinelvault.{domain}.v1;

option go_package = "github.com/sentinelvault/sentinelvault/gen/{domain}/v1;{domain}v1";

message {entity} {{
  string id = 1;
  string org_id = 2;
  string name = 3;
  string status = 4;
}}

message List{entity}sRequest {{
  string org_id = 1;
  int32 page = 2;
  int32 page_size = 3;
}}

message List{entity}sResponse {{
  repeated {entity} items = 1;
}}

service {svc}Service {{
  rpc List{entity}s(List{entity}sRequest) returns (List{entity}sResponse);
}}
'''


def gen_helm_chart(name: str) -> None:
    chart_dir = ROOT / "deploy" / "helm" / "charts" / name
    write(chart_dir / "Chart.yaml", f'''apiVersion: v2
name: {name}
description: Helm chart for {name}
type: application
version: 0.1.0
appVersion: "0.1.0"
''')
    write(chart_dir / "values.yaml", f'''replicaCount: 1
image:
  repository: sentinelvault/{name}
  tag: latest
  pullPolicy: IfNotPresent
service:
  type: ClusterIP
  port: 8080
resources: {{}}
''')
    write(chart_dir / "templates" / "deployment.yaml", f'''apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{{{ .Chart.Name }}}}
spec:
  replicas: {{{{ .Values.replicaCount }}}}
  selector:
    matchLabels:
      app: {{%{{ .Chart.Name }}}}
  template:
    metadata:
      labels:
        app: {{{{ .Chart.Name }}}}
    spec:
      containers:
        - name: {name}
          image: "{{{{ .Values.image.repository }}}}:{{{{ .Values.image.tag }}}}"
          ports:
            - containerPort: {{{{ .Values.service.port }}}}
          env:
            - name: CONFIG_PATH
              value: /etc/sentinel/{name}.yaml
''')
    write(chart_dir / "templates" / "service.yaml", f'''apiVersion: v1
kind: Service
metadata:
  name: {{%{{ .Chart.Name }}}}
spec:
  selector:
    app: {{{{ .Chart.Name }}}}
  ports:
    - port: {{{{ .Values.service.port }}}}
      targetPort: {{{{ .Values.service.port }}}}
''')


def gen_k8s(name: str) -> None:
    write(ROOT / "deploy" / "kubernetes" / f"{name}-deployment.yaml", f'''apiVersion: apps/v1
kind: Deployment
metadata:
  name: {name}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {name}
  template:
    metadata:
      labels:
        app: {name}
    spec:
      containers:
        - name: {name}
          image: sentinelvault/{name}:latest
          ports:
            - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: {name}
spec:
  selector:
    app: {name}
  ports:
    - port: 8080
''')


def gen_service(name: str) -> None:
    meta = SERVICE_META.get(name, {"domain": "operations", "entity": "Record", "schema": name.replace("-", "_"), "table": "records"})
    if name != "api-gateway":
        write(ROOT / f"apps/{name}/cmd/main.go", gen_main(name))
        write(ROOT / f"apps/{name}/internal/app/app.go", gen_app(name, meta))
        write(ROOT / f"apps/{name}/internal/domain/domain.go", gen_domain(name, meta))
        write(ROOT / f"apps/{name}/internal/http/handlers.go", gen_handlers(name, meta))
    write(ROOT / f"apps/{name}/Dockerfile", gen_dockerfile(name))
    write(ROOT / f"deploy/docker/Dockerfile.{name}", gen_dockerfile(name))
    write(ROOT / f"configs/{name}.yaml", gen_config(name))
    write(ROOT / f"migrations/{name}/001_init.sql", gen_migration(name, meta))
    write(ROOT / f"proto/{meta['domain']}/{name}.proto", gen_proto(name, meta))
    gen_helm_chart(name)
    gen_k8s(name)


def main() -> None:
    for svc in SERVICES:
        gen_service(svc)
    print(f"Generated {len(SERVICES)} services")


if __name__ == "__main__":
    main()
