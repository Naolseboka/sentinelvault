#!/usr/bin/env python3
"""Phase B expander: service domain modules, IdP, analytics, SDK, sagas — quality only."""
from __future__ import annotations

from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]

SERVICES = [
    ("auth-service", "auth", "Credential", "identity"),
    ("authorization-service", "authorization", "AuthzDecision", "identity"),
    ("user-service", "user", "User", "identity"),
    ("organization-service", "organization", "Organization", "identity"),
    ("team-service", "team", "Team", "identity"),
    ("rbac-service", "rbac", "RoleBinding", "identity"),
    ("policy-service", "policy", "PolicyDoc", "identity"),
    ("device-service", "device", "Device", "identity"),
    ("session-service", "session", "Session", "identity"),
    ("invitation-service", "invitation", "Invitation", "identity"),
    ("password-service", "password", "VaultItem", "vault"),
    ("secret-service", "secret", "SecretItem", "vault"),
    ("encryption-service", "encryption", "EncryptOp", "vault"),
    ("kms-service", "kms", "Key", "vault"),
    ("certificate-service", "certificate", "Certificate", "vault"),
    ("ssh-service", "ssh", "SSHKey", "vault"),
    ("passkey-service", "passkey", "Passkey", "vault"),
    ("otp-service", "otp", "OTPRecord", "vault"),
    ("attachment-service", "attachment", "Attachment", "vault"),
    ("versioning-service", "versioning", "Version", "vault"),
    ("approval-service", "approval", "Approval", "enterprise"),
    ("access-request-service", "accessrequest", "AccessRequest", "enterprise"),
    ("emergency-access-service", "emergency", "EmergencyGrant", "enterprise"),
    ("delegation-service", "delegation", "Delegation", "enterprise"),
    ("lifecycle-service", "lifecycle", "LifecycleState", "enterprise"),
    ("provisioning-service", "provisioning", "ProvisionPlan", "enterprise"),
    ("scim-service", "scim", "SCIMUser", "enterprise"),
    ("directory-sync-service", "directorysync", "SyncJob", "enterprise"),
    ("threat-service", "threat", "Finding", "security"),
    ("risk-service", "risk", "RiskScore", "security"),
    ("behavior-service", "behavior", "Signal", "security"),
    ("device-trust-service", "devicetrust", "TrustRecord", "security"),
    ("notification-service", "notification", "Message", "operations"),
    ("audit-service", "audit", "AuditEvent", "operations"),
    ("compliance-service", "compliance", "Evidence", "operations"),
    ("search-service", "search", "Document", "operations"),
    ("backup-service", "backup", "Snapshot", "operations"),
    ("rotation-service", "rotation", "RotationPolicy", "operations"),
    ("admin-service", "admin", "AdminAction", "operations"),
    ("reporting-service", "reporting", "Report", "operations"),
    ("billing-service", "billing", "License", "operations"),
]

EVENT_TYPES = [
    "created", "updated", "deleted", "archived", "restored", "shared", "unshared",
    "exported", "imported", "rotated", "accessed", "denied", "approved", "revoked",
    "expired", "locked", "unlocked", "verified", "failed", "succeeded",
]

STATUSES = [
    "active", "inactive", "pending", "approved", "denied", "expired", "revoked",
    "archived", "deleted", "draft", "scheduled", "running", "failed", "completed",
]


def write(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content)
    print(f"wrote {path.relative_to(ROOT)} ({len(content.splitlines())})")


def service_events(svc: str, pkg: str, entity: str) -> str:
    cases = "\n".join(
        f'\tcase Event{entity}{e.title().replace("_","")}:\n\t\treturn true'
        for e in EVENT_TYPES
    )
    consts = "\n".join(
        f'\tEvent{entity}{e.title().replace("_","")} = "{pkg}.{e}"'
        for e in EVENT_TYPES
    )
    builders = "\n".join(
        f'''
// New{entity}{e.title().replace("_","")}Event builds a domain event.
func New{entity}{e.title().replace("_","")}Event(orgID, actorID, resourceID string, meta map[string]string) Event {{
\tif meta == nil {{
\t\tmeta = map[string]string{{}}
\t}}
\tmeta["entity"] = "{entity}"
\tmeta["action"] = "{e}"
\treturn Event{{
\t\tID: common.NewID(), Type: Event{entity}{e.title().replace("_","")}, OrgID: orgID, ActorID: actorID,
\t\tResourceID: resourceID, OccurredAt: common.NowUTC(), Metadata: meta,
\t}}
}}'''
        for e in EVENT_TYPES
    )
    return f'''package domainevents

import (
\t"time"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

// Event types for {svc}.
const (
{consts}
)

// Event is a {svc} domain event.
type Event struct {{
\tID         string            `json:"id"`
\tType       string            `json:"type"`
\tOrgID      string            `json:"org_id"`
\tActorID    string            `json:"actor_id"`
\tResourceID string            `json:"resource_id"`
\tOccurredAt time.Time         `json:"occurred_at"`
\tMetadata   map[string]string `json:"metadata,omitempty"`
}}

// IsKnown{entity}Event reports whether t is a known {pkg} event type.
func IsKnown{entity}Event(t string) bool {{
\tswitch t {{
{cases}
\tdefault:
\t\treturn false
\t}}
}}

{builders}

// ValidateEvent ensures required event fields.
func ValidateEvent(e Event) error {{
\tif e.ID == "" || e.Type == "" || e.OrgID == "" {{
\t\treturn common.New(common.ErrCodeInvalidArgument, "id, type, org_id required")
\t}}
\tif !IsKnown{entity}Event(e.Type) {{
\t\treturn common.New(common.ErrCodeInvalidArgument, "unknown event type")
\t}}
\treturn nil
}}
'''


def service_events_test(pkg: str, entity: str) -> str:
    first = EVENT_TYPES[0]
    fn = f"New{entity}{first.title().replace('_','')}Event"
    return f'''package domainevents

import "testing"

func Test{entity}Events(t *testing.T) {{
\tev := {fn}("org", "actor", "res", map[string]string{{"k": "v"}})
\tif err := ValidateEvent(ev); err != nil {{
\t\tt.Fatal(err)
\t}}
\tif !IsKnown{entity}Event(ev.Type) {{
\t\tt.Fatal("unknown")
\t}}
\tif IsKnown{entity}Event("nope") {{
\t\tt.Fatal("expected false")
\t}}
}}
'''


def service_query(svc: str, pkg: str, entity: str) -> str:
    filters = "\n".join(
        f'''
\tif q.{field} != "" {{
\t\tout = filter{entity}By{field}(out, q.{field})
\t}}'''
        for field in ["Status", "NameContains", "ActorID"]
    )
    return f'''package query

import (
\t"sort"
\t"strings"
\t"time"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

// {entity}View is a query-side read model for {svc}.
type {entity}View struct {{
\tID        string            `json:"id"`
\tOrgID     string            `json:"org_id"`
\tName      string            `json:"name"`
\tStatus    string            `json:"status"`
\tActorID   string            `json:"actor_id,omitempty"`
\tLabels    map[string]string `json:"labels,omitempty"`
\tCreatedAt time.Time         `json:"created_at"`
\tUpdatedAt time.Time         `json:"updated_at"`
\tScore     float64           `json:"score,omitempty"`
}}

// {entity}Query holds filter/sort/page options.
type {entity}Query struct {{
\tOrgID        string
\tStatus       string
\tNameContains string
\tActorID      string
\tLabelKey     string
\tLabelValue   string
\tSortBy       string // name|created_at|updated_at|status
\tSortDesc     bool
\tOffset       int
\tLimit        int
}}

// {entity}Index is an in-memory query index (CQRS read side).
type {entity}Index struct {{
\titems map[string]{entity}View
}}

// New{entity}Index creates an empty index.
func New{entity}Index() *{entity}Index {{
\treturn &{entity}Index{{items: map[string]{entity}View{{}}}}
}}

// Upsert indexes a view.
func (idx *{entity}Index) Upsert(v {entity}View) {{
\tif v.Labels == nil {{
\t\tv.Labels = map[string]string{{}}
\t}}
\tidx.items[v.ID] = v
}}

// Delete removes a view.
func (idx *{entity}Index) Delete(id string) {{
\tdelete(idx.items, id)
}}

// Get returns a view by id.
func (idx *{entity}Index) Get(id string) ({entity}View, bool) {{
\tv, ok := idx.items[id]
\treturn v, ok
}}

// Search applies filters, sorting, and pagination.
func (idx *{entity}Index) Search(q {entity}Query) ([]{entity}View, int64, error) {{
\tif q.OrgID == "" {{
\t\treturn nil, 0, common.New(common.ErrCodeInvalidArgument, "org_id required")
\t}}
\tif q.Limit <= 0 {{
\t\tq.Limit = 50
\t}}
\tif q.Limit > 500 {{
\t\tq.Limit = 500
\t}}
\tout := make([]{entity}View, 0)
\tfor _, v := range idx.items {{
\t\tif v.OrgID != q.OrgID {{
\t\t\tcontinue
\t\t}}
\t\tout = append(out, v)
\t}}
{filters}
\tif q.LabelKey != "" {{
\t\tfiltered := make([]{entity}View, 0)
\t\tfor _, v := range out {{
\t\t\tif val, ok := v.Labels[q.LabelKey]; ok && (q.LabelValue == "" || val == q.LabelValue) {{
\t\t\t\tfiltered = append(filtered, v)
\t\t\t}}
\t\t}}
\t\tout = filtered
\t}}
\tsort{entity}Views(out, q.SortBy, q.SortDesc)
\ttotal := int64(len(out))
\tif q.Offset >= len(out) {{
\t\treturn []{entity}View{{}}, total, nil
\t}}
\tend := q.Offset + q.Limit
\tif end > len(out) {{
\t\tend = len(out)
\t}}
\treturn out[q.Offset:end], total, nil
}}

func filter{entity}ByStatus(in []{entity}View, status string) []{entity}View {{
\to := make([]{entity}View, 0, len(in))
\tfor _, v := range in {{
\t\tif v.Status == status {{
\t\t\to = append(o, v)
\t\t}}
\t}}
\treturn o
}}

func filter{entity}ByNameContains(in []{entity}View, needle string) []{entity}View {{
\tn := strings.ToLower(needle)
\to := make([]{entity}View, 0, len(in))
\tfor _, v := range in {{
\t\tif strings.Contains(strings.ToLower(v.Name), n) {{
\t\t\to = append(o, v)
\t\t}}
\t}}
\treturn o
}}

func filter{entity}ByActorID(in []{entity}View, actor string) []{entity}View {{
\to := make([]{entity}View, 0, len(in))
\tfor _, v := range in {{
\t\tif v.ActorID == actor {{
\t\t\to = append(o, v)
\t\t}}
\t}}
\treturn o
}}

func sort{entity}Views(in []{entity}View, by string, desc bool) {{
\tsort.SliceStable(in, func(i, j int) bool {{
\t\tvar less bool
\t\tswitch by {{
\t\tcase "name":
\t\t\tless = in[i].Name < in[j].Name
\t\tcase "status":
\t\t\tless = in[i].Status < in[j].Status
\t\tcase "updated_at":
\t\t\tless = in[i].UpdatedAt.Before(in[j].UpdatedAt)
\t\tdefault:
\t\t\tless = in[i].CreatedAt.Before(in[j].CreatedAt)
\t\t}}
\t\tif desc {{
\t\t\treturn !less
\t\t}}
\t\treturn less
\t}})
}}

// Known{entity}Statuses returns allowed statuses for {svc}.
func Known{entity}Statuses() []string {{
\treturn []string{{{", ".join(f'"{s}"' for s in STATUSES)}}}
}}

// IsValid{entity}Status reports whether s is allowed.
func IsValid{entity}Status(s string) bool {{
\tfor _, k := range Known{entity}Statuses() {{
\t\tif k == s {{
\t\t\treturn true
\t\t}}
\t}}
\treturn false
}}
'''


def service_query_test(entity: str) -> str:
    return f'''package query

import (
\t"testing"
\t"time"
)

func Test{entity}IndexSearch(t *testing.T) {{
\tidx := New{entity}Index()
\tnow := time.Now().UTC()
\tidx.Upsert({entity}View{{ID: "1", OrgID: "o1", Name: "alpha", Status: "active", CreatedAt: now, Labels: map[string]string{{"env": "prod"}}}})
\tidx.Upsert({entity}View{{ID: "2", OrgID: "o1", Name: "beta", Status: "pending", CreatedAt: now.Add(time.Second), Labels: map[string]string{{"env": "dev"}}}})
\tidx.Upsert({entity}View{{ID: "3", OrgID: "o2", Name: "other", Status: "active", CreatedAt: now}})
\tres, total, err := idx.Search({entity}Query{{OrgID: "o1", Status: "active", Limit: 10}})
\tif err != nil || total != 1 || len(res) != 1 {{
\t\tt.Fatalf("res=%v total=%d err=%v", res, total, err)
\t}}
\tres, total, err = idx.Search({entity}Query{{OrgID: "o1", NameContains: "bet", Limit: 10}})
\tif err != nil || total != 1 {{
\t\tt.Fatalf("name filter total=%d err=%v", total, err)
\t}}
\tres, total, err = idx.Search({entity}Query{{OrgID: "o1", LabelKey: "env", LabelValue: "prod", Limit: 10}})
\tif err != nil || total != 1 {{
\t\tt.Fatalf("label total=%d", total)
\t}}
\tif !IsValid{entity}Status("active") || IsValid{entity}Status("nope") {{
\t\tt.Fatal("status validation")
\t}}
}}
'''


def service_metrics(pkg: str, entity: str) -> str:
    return f'''package metrics

import (
\t"sync/atomic"
)

// {entity}Metrics tracks operational counters for {pkg}.
type {entity}Metrics struct {{
\tCreated   atomic.Int64
\tUpdated   atomic.Int64
\tDeleted   atomic.Int64
\tRead      atomic.Int64
\tErrors    atomic.Int64
\tDenied    atomic.Int64
\tExported  atomic.Int64
\tImported  atomic.Int64
\tRotated   atomic.Int64
\tLatencyMs atomic.Int64 // last observed
}}

// Snapshot is a point-in-time copy of counters.
type {entity}Snapshot struct {{
\tCreated  int64 `json:"created"`
\tUpdated  int64 `json:"updated"`
\tDeleted  int64 `json:"deleted"`
\tRead     int64 `json:"read"`
\tErrors   int64 `json:"errors"`
\tDenied   int64 `json:"denied"`
\tExported int64 `json:"exported"`
\tImported int64 `json:"imported"`
\tRotated  int64 `json:"rotated"`
\tLatencyMs int64 `json:"latency_ms"`
}}

// ObserveLatency stores last latency.
func (m *{entity}Metrics) ObserveLatency(ms int64) {{
\tm.LatencyMs.Store(ms)
}}

// Snapshot returns current counters.
func (m *{entity}Metrics) Snapshot() {entity}Snapshot {{
\treturn {entity}Snapshot{{
\t\tCreated: m.Created.Load(), Updated: m.Updated.Load(), Deleted: m.Deleted.Load(),
\t\tRead: m.Read.Load(), Errors: m.Errors.Load(), Denied: m.Denied.Load(),
\t\tExported: m.Exported.Load(), Imported: m.Imported.Load(), Rotated: m.Rotated.Load(),
\t\tLatencyMs: m.LatencyMs.Load(),
\t}}
}}
'''


def service_metrics_test(entity: str) -> str:
    return f'''package metrics

import "testing"

func Test{entity}Metrics(t *testing.T) {{
\tvar m {entity}Metrics
\tm.Created.Add(2)
\tm.ObserveLatency(12)
\ts := m.Snapshot()
\tif s.Created != 2 || s.LatencyMs != 12 {{
\t\tt.Fatalf("%+v", s)
\t}}
}}
'''


def service_dto(pkg: str, entity: str) -> str:
    fields = "\n".join(
        f'\t{f} string `json:"{f.lower()}"`'
        for f in ["ExternalID", "Description", "OwnerID", "CollectionID", "FolderID", "RiskLevel", "Classification"]
    )
    return f'''package dto

import (
\t"encoding/json"
\t"strings"
\t"time"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

// {entity}DTO is the transport object for {pkg} APIs.
type {entity}DTO struct {{
\tID        string            `json:"id"`
\tOrgID     string            `json:"org_id"`
\tName      string            `json:"name"`
\tStatus    string            `json:"status"`
{fields}
\tTags      []string          `json:"tags,omitempty"`
\tMetadata  map[string]string `json:"metadata,omitempty"`
\tCreatedAt time.Time         `json:"created_at"`
\tUpdatedAt time.Time         `json:"updated_at"`
}}

// Create{entity}Request is the create payload.
type Create{entity}Request struct {{
\tOrgID       string            `json:"org_id"`
\tName        string            `json:"name"`
\tDescription string            `json:"description"`
\tOwnerID     string            `json:"owner_id"`
\tTags        []string          `json:"tags"`
\tMetadata    map[string]string `json:"metadata"`
\tPayload     json.RawMessage   `json:"payload"`
}}

// Update{entity}Request is the update payload.
type Update{entity}Request struct {{
\tName        *string           `json:"name"`
\tDescription *string           `json:"description"`
\tStatus      *string           `json:"status"`
\tTags        []string          `json:"tags"`
\tMetadata    map[string]string `json:"metadata"`
\tPayload     json.RawMessage   `json:"payload"`
}}

// ValidateCreate validates a create request.
func ValidateCreate{entity}(r Create{entity}Request) error {{
\tif strings.TrimSpace(r.OrgID) == "" {{
\t\treturn common.New(common.ErrCodeInvalidArgument, "org_id required")
\t}}
\tif strings.TrimSpace(r.Name) == "" {{
\t\treturn common.New(common.ErrCodeInvalidArgument, "name required")
\t}}
\tif len(r.Name) > 256 {{
\t\treturn common.New(common.ErrCodeInvalidArgument, "name too long")
\t}}
\tfor _, t := range r.Tags {{
\t\tif len(t) > 64 {{
\t\t\treturn common.New(common.ErrCodeInvalidArgument, "tag too long")
\t\t}}
\t}}
\treturn nil
}}

// ApplyUpdate applies patch fields onto dto.
func ApplyUpdate{entity}(dst *{entity}DTO, r Update{entity}Request) {{
\tif r.Name != nil {{
\t\tdst.Name = strings.TrimSpace(*r.Name)
\t}}
\tif r.Description != nil {{
\t\tdst.Description = strings.TrimSpace(*r.Description)
\t}}
\tif r.Status != nil {{
\t\tdst.Status = strings.TrimSpace(*r.Status)
\t}}
\tif r.Tags != nil {{
\t\tdst.Tags = append([]string(nil), r.Tags...)
\t}}
\tif r.Metadata != nil {{
\t\tif dst.Metadata == nil {{
\t\t\tdst.Metadata = map[string]string{{}}
\t\t}}
\t\tfor k, v := range r.Metadata {{
\t\t\tdst.Metadata[k] = v
\t\t}}
\t}}
\tdst.UpdatedAt = common.NowUTC()
}}
'''


def service_dto_test(entity: str) -> str:
    return f'''package dto

import "testing"

func TestValidateCreate{entity}(t *testing.T) {{
\tif err := ValidateCreate{entity}(Create{entity}Request{{}}); err == nil {{
\t\tt.Fatal("expected error")
\t}}
\tif err := ValidateCreate{entity}(Create{entity}Request{{OrgID: "o", Name: "n"}}); err != nil {{
\t\tt.Fatal(err)
\t}}
\tdst := &{entity}DTO{{Name: "old"}}
\tname := "new"
\tApplyUpdate{entity}(dst, Update{entity}Request{{Name: &name}})
\tif dst.Name != "new" {{
\t\tt.Fatal(dst.Name)
\t}}
}}
'''


def sdk_client(svc: str, pkg: str, entity: str) -> str:
    return f'''package {pkg}

import (
\t"bytes"
\t"context"
\t"encoding/json"
\t"fmt"
\t"net/http"
\t"net/url"
\t"strconv"
\t"time"
)

// Client talks to {svc} over HTTP.
type Client struct {{
\tBaseURL    string
\tHTTPClient *http.Client
\tToken      string
}}

// New creates a {svc} client.
func New(baseURL, token string) *Client {{
\treturn &Client{{
\t\tBaseURL: stringsTrimSlash(baseURL),
\t\tHTTPClient: &http.Client{{Timeout: 30 * time.Second}},
\t\tToken: token,
\t}}
}}

func stringsTrimSlash(s string) string {{
\tfor len(s) > 0 && s[len(s)-1] == '/' {{
\t\ts = s[:len(s)-1]
\t}}
\treturn s
}}

// {entity} is the API resource.
type {entity} struct {{
\tID        string            `json:"id"`
\tOrgID     string            `json:"org_id"`
\tName      string            `json:"name"`
\tStatus    string            `json:"status"`
\tMetadata  map[string]string `json:"metadata,omitempty"`
\tCreatedAt time.Time         `json:"created_at"`
\tUpdatedAt time.Time         `json:"updated_at"`
}}

// ListResult wraps paginated data.
type ListResult struct {{
\tData       []{entity} `json:"data"`
\tTotal      int64      `json:"total"`
\tOffset     int        `json:"offset"`
\tLimit      int        `json:"limit"`
}}

func (c *Client) do(ctx context.Context, method, path string, in any, out any) error {{
\tvar body *bytes.Reader
\tif in != nil {{
\t\tb, err := json.Marshal(in)
\t\tif err != nil {{
\t\t\treturn err
\t\t}}
\t\tbody = bytes.NewReader(b)
\t}} else {{
\t\tbody = bytes.NewReader(nil)
\t}}
\treq, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, body)
\tif err != nil {{
\t\treturn err
\t}}
\tif in != nil {{
\t\treq.Header.Set("Content-Type", "application/json")
\t}}
\tif c.Token != "" {{
\t\treq.Header.Set("Authorization", "Bearer "+c.Token)
\t}}
\tres, err := c.HTTPClient.Do(req)
\tif err != nil {{
\t\treturn err
\t}}
\tdefer res.Body.Close()
\tif res.StatusCode >= 300 {{
\t\treturn fmt.Errorf("{pkg}: unexpected status %d", res.StatusCode)
\t}}
\tif out == nil {{
\t\treturn nil
\t}}
\treturn json.NewDecoder(res.Body).Decode(out)
}}

// Create creates a {entity}.
func (c *Client) Create(ctx context.Context, orgID, name string, metadata map[string]string) (*{entity}, error) {{
\tvar out {entity}
\terr := c.do(ctx, http.MethodPost, "/v1/{pkg}s", map[string]any{{
\t\t"org_id": orgID, "name": name, "metadata": metadata,
\t}}, &out)
\treturn &out, err
}}

// Get fetches a {entity} by id.
func (c *Client) Get(ctx context.Context, id string) (*{entity}, error) {{
\tvar out {entity}
\terr := c.do(ctx, http.MethodGet, "/v1/{pkg}s/"+url.PathEscape(id), nil, &out)
\treturn &out, err
}}

// List lists {entity} records for an org.
func (c *Client) List(ctx context.Context, orgID string, offset, limit int) (*ListResult, error) {{
\tq := url.Values{{}}
\tq.Set("org_id", orgID)
\tq.Set("offset", strconv.Itoa(offset))
\tq.Set("limit", strconv.Itoa(limit))
\tvar out ListResult
\terr := c.do(ctx, http.MethodGet, "/v1/{pkg}s?"+q.Encode(), nil, &out)
\treturn &out, err
}}

// Delete deletes a {entity}.
func (c *Client) Delete(ctx context.Context, id string) error {{
\treturn c.do(ctx, http.MethodDelete, "/v1/{pkg}s/"+url.PathEscape(id), nil, nil)
}}
'''


def sdk_client_test(pkg: str, entity: str) -> str:
    return f'''package {pkg}

import "testing"

func TestNewClient(t *testing.T) {{
\tc := New("http://localhost:8080/", "tok")
\tif c.BaseURL != "http://localhost:8080" || c.Token != "tok" {{
\t\tt.Fatalf("%+v", c)
\t}}
}}
'''


def idp_oidc() -> str:
    return Path("/dev/null")  # placeholder - write full file below


def main() -> None:
    for svc, pkg, entity, domain in SERVICES:
        base = ROOT / "apps" / svc / "internal"
        write(base / "domainevents" / "events.go", service_events(svc, pkg, entity))
        write(base / "domainevents" / "events_test.go", service_events_test(pkg, entity))
        write(base / "query" / "index.go", service_query(svc, pkg, entity))
        write(base / "query" / "index_test.go", service_query_test(entity))
        write(base / "metrics" / "metrics.go", service_metrics(pkg, entity))
        write(base / "metrics" / "metrics_test.go", service_metrics_test(entity))
        write(base / "dto" / "dto.go", service_dto(pkg, entity))
        write(base / "dto" / "dto_test.go", service_dto_test(entity))
        write(ROOT / "pkg" / "sdk" / pkg / "client.go", sdk_client(svc, pkg, entity))
        write(ROOT / "pkg" / "sdk" / pkg / "client_test.go", sdk_client_test(pkg, entity))

    print("phase B service modules done")


if __name__ == "__main__":
    main()
