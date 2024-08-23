#!/usr/bin/env python3
"""Phase D: per-service ACL rules, analytics dashboards, webhook events."""
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]

APPS = [
    ("access-request-service", "accessrequest", "AccessRequest", "enterprise"),
    ("admin-service", "admin", "AdminAction", "operations"),
    ("approval-service", "approval", "Approval", "enterprise"),
    ("attachment-service", "attachment", "Attachment", "vault"),
    ("audit-service", "audit", "AuditEvent", "operations"),
    ("authorization-service", "authorization", "AuthzDecision", "identity"),
    ("auth-service", "auth", "Credential", "identity"),
    ("backup-service", "backup", "Snapshot", "operations"),
    ("behavior-service", "behavior", "Signal", "security"),
    ("billing-service", "billing", "License", "operations"),
    ("certificate-service", "certificate", "Certificate", "vault"),
    ("compliance-service", "compliance", "Evidence", "operations"),
    ("delegation-service", "delegation", "Delegation", "enterprise"),
    ("device-service", "device", "Device", "identity"),
    ("device-trust-service", "devicetrust", "TrustRecord", "security"),
    ("directory-sync-service", "directorysync", "SyncJob", "enterprise"),
    ("emergency-access-service", "emergency", "EmergencyGrant", "enterprise"),
    ("encryption-service", "encryption", "EncryptOp", "vault"),
    ("invitation-service", "invitation", "Invitation", "identity"),
    ("kms-service", "kms", "Key", "vault"),
    ("lifecycle-service", "lifecycle", "LifecycleState", "enterprise"),
    ("notification-service", "notification", "Message", "operations"),
    ("organization-service", "organization", "Organization", "identity"),
    ("otp-service", "otp", "OTPRecord", "vault"),
    ("passkey-service", "passkey", "Passkey", "vault"),
    ("password-service", "password", "VaultItem", "vault"),
    ("policy-service", "policy", "PolicyDoc", "identity"),
    ("provisioning-service", "provisioning", "ProvisionPlan", "enterprise"),
    ("rbac-service", "rbac", "RoleBinding", "identity"),
    ("reporting-service", "reporting", "Report", "operations"),
    ("risk-service", "risk", "RiskScore", "security"),
    ("rotation-service", "rotation", "RotationPolicy", "operations"),
    ("scim-service", "scim", "SCIMUser", "enterprise"),
    ("search-service", "search", "Document", "operations"),
    ("secret-service", "secret", "SecretItem", "vault"),
    ("session-service", "session", "Session", "identity"),
    ("ssh-service", "ssh", "SSHKey", "vault"),
    ("team-service", "team", "Team", "identity"),
    ("threat-service", "threat", "Finding", "security"),
    ("user-service", "user", "User", "identity"),
    ("versioning-service", "versioning", "Version", "vault"),
]

DASHBOARDS = [
    "executive_summary", "security_posture", "vault_health", "access_patterns",
    "compliance_status", "rotation_coverage", "device_landscape", "billing_usage",
    "incident_timeline", "admin_activity", "sharing_audit", "risk_heatmap",
]

WEBHOOK_EVENTS = [
    "user.created", "user.deleted", "item.shared", "item.rotated",
    "session.revoked", "policy.updated", "emergency.granted", "backup.completed",
    "finding.detected", "org.provisioned",
]


def write(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content)
    print(f"wrote {path.relative_to(ROOT)} ({len(content.splitlines())} lines)")


def gen_acl_rules(svc, pkg, entity, domain):
    rules = {
        "identity": ["self_access", "team_member", "org_admin", "service_account"],
        "vault": ["owner_full", "shared_read", "shared_write", "collection_inherit"],
        "enterprise": ["requestor", "approver", "delegate", "break_glass"],
        "security": ["analyst", "responder", "viewer", "scanner"],
        "operations": ["operator", "auditor", "billing_admin", "platform_admin"],
    }.get(domain, ["read", "write", "admin", "owner"])
    cases = "\n".join(
        f'\tcase "{r}":\n\t\treturn evaluate{r.title().replace("_", "")}(ctx, req)'
        for r in rules
    )
    funcs = "\n".join(
        f'''
func evaluate{r.title().replace("_", "")}(ctx context.Context, req RuleRequest) (RuleResult, error) {{
\t_ = ctx
\tif req.OrgID == "" || req.SubjectID == "" {{
\t\treturn RuleResult{{Allowed: false, Reason: "missing subject or org"}}, common.New(common.ErrCodeInvalidArgument, "invalid request")
\t}}
\tallowed := req.SubjectID == req.ResourceOwner || req.SubjectRole == "admin" || req.SubjectRole == "owner"
\treturn RuleResult{{Allowed: allowed, Reason: "{r}", Rule: "{r}"}}, nil
}}'''
        for r in rules
    )
    return f'''package acl

import (
\t"context"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

// RuleRequest is input for domain rule evaluation in {svc}.
type RuleRequest struct {{
\tOrgID          string
\tSubjectID      string
\tSubjectRole    string
\tResourceID     string
\tResourceOwner  string
\tAction         string
\tMetadata       map[string]string
}}

// RuleResult is the outcome of a single rule.
type RuleResult struct {{
\tAllowed bool
\tReason  string
\tRule    string
}}

// RuleEngine evaluates ordered domain rules for {entity} resources.
type RuleEngine struct {{
\tfallback *Evaluator
}}

func NewRuleEngine(fallback *Evaluator) *RuleEngine {{
\treturn &RuleEngine{{fallback: fallback}}
}}

func (re *RuleEngine) Evaluate(ctx context.Context, ruleName string, req RuleRequest) (RuleResult, error) {{
\tswitch ruleName {{
{cases}
\tdefault:
\t\tif re.fallback == nil {{
\t\t\treturn RuleResult{{Allowed: false, Reason: "no fallback"}}, common.New(common.ErrCodeNotFound, "unknown rule")
\t\t}}
\t\tok, err := re.fallback.Can(Principal{{ID: req.SubjectID, OrgID: req.OrgID, Role: req.SubjectRole}}, Resource{{ID: req.ResourceID, OrgID: req.OrgID, OwnerID: req.ResourceOwner}}, req.Action)
\t\treturn RuleResult{{Allowed: ok, Reason: "fallback", Rule: "fallback"}}, err
\t}}
}}
{funcs}

// EvaluateChain runs rules until one allows or all deny.
func (re *RuleEngine) EvaluateChain(ctx context.Context, req RuleRequest, rules []string) (RuleResult, error) {{
\tfor _, name := range rules {{
\t\tres, err := re.Evaluate(ctx, name, req)
\t\tif err != nil {{
\t\t\treturn res, err
\t\t}}
\t\tif res.Allowed {{
\t\t\treturn res, nil
\t\t}}
\t}}
\treturn RuleResult{{Allowed: false, Reason: "all rules denied"}}, nil
}}

// DefaultRules returns the standard rule chain for {domain} domain.
func DefaultRules() []string {{
\treturn []string{{{", ".join(f'"{r}"' for r in rules)}}}
}}
'''


def gen_acl_rules_test(entity):
    return f'''package acl

import (
\t"context"
\t"testing"
)

func TestRuleEngineChain(t *testing.T) {{
\tre := NewRuleEngine(NewEvaluator())
\treq := RuleRequest{{OrgID: "o1", SubjectID: "u1", SubjectRole: "admin", ResourceID: "r1", ResourceOwner: "u2", Action: "read"}}
\tres, err := re.EvaluateChain(context.Background(), req, DefaultRules())
\tif err != nil {{
\t\tt.Fatal(err)
\t}}
\tif !res.Allowed {{
\t\tt.Fatalf("expected allow: %+v", res)
\t}}
}}

func TestRuleEngineUnknown(t *testing.T) {{
\tre := NewRuleEngine(nil)
\t_, err := re.Evaluate(context.Background(), "missing", RuleRequest{{OrgID: "o1", SubjectID: "u1"}})
\tif err == nil {{
\t\tt.Fatal("expected error")
\t}}
}}
'''


def gen_dashboard(name):
    title = "".join(w.capitalize() for w in name.split("_"))
    return f'''package analytics

import (
\t"fmt"
\t"time"
)

// {title}Report aggregates dashboard metrics for {name.replace("_", " ")}.
type {title}Report struct {{
\tOrgID      string
\tGeneratedAt time.Time
\tSections   []ReportSection
\tScore      float64
}}

type ReportSection struct {{
\tTitle  string
\tValue  float64
\tDetail string
}}

// Build{title}Report composes a report from metric calculators.
func Build{title}Report(orgID string, from, to time.Time, metrics ...*LoginHeatmapCalculator) *{title}Report {{
\trep := &{title}Report{{OrgID: orgID, GeneratedAt: time.Now().UTC(), Sections: make([]ReportSection, 0)}}
\tvar total float64
\tfor i, m := range metrics {{
\t\tif m == nil {{
\t\t\tcontinue
\t\t}}
\t\tavg := m.Average(orgID, from, to)
\t\ttotal += avg
\t\trep.Sections = append(rep.Sections, ReportSection{{
\t\t\tTitle: fmt.Sprintf("metric_%d", i),
\t\t\tValue: avg,
\t\t\tDetail: fmt.Sprintf("window %s to %s", from.Format(time.RFC3339), to.Format(time.RFC3339)),
\t\t}})
\t}}
\tif len(metrics) > 0 {{
\t\trep.Score = total / float64(len(metrics))
\t}}
\treturn rep
}}

// MergeReports combines sections from multiple reports.
func Merge{title}Reports(reports ...*{title}Report) *{title}Report {{
\tif len(reports) == 0 {{
\t\treturn &{title}Report{{}}
\t}}
\tout := *reports[0]
\tfor _, r := range reports[1:] {{
\t\tif r == nil {{
\t\t\tcontinue
\t\t}}
\t\tout.Sections = append(out.Sections, r.Sections...)
\t\tout.Score = (out.Score + r.Score) / 2
\t}}
\treturn &out
}}
'''


def gen_dashboard_test(name):
    title = "".join(w.capitalize() for w in name.split("_"))
    return f'''package analytics

import (
\t"testing"
\t"time"
)

func Test{title}Report(t *testing.T) {{
\tnow := time.Now().UTC()
\tcalc := NewLoginHeatmapCalculator(time.Hour)
\tcalc.Record("o1", now, 1)
\trep := Build{title}Report("o1", now.Add(-time.Hour), now.Add(time.Hour), calc)
\tif rep.OrgID != "o1" || len(rep.Sections) == 0 {{
\t\tt.Fatalf("%+v", rep)
\t}}
\tmerged := Merge{title}Reports(rep, rep)
\tif merged.Score <= 0 {{
\t\tt.Fatal("score")
\t}}
}}
'''


def gen_webhook_event(event):
    safe = event.replace(".", "_")
    title = "".join(w.capitalize() for w in safe.split("_"))
    return f'''package webhooks

import (
\t"encoding/json"
\t"time"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

// {title}Payload is the webhook body for {event}.
type {title}Payload struct {{
\tEvent     string            `json:"event"`
\tOrgID     string            `json:"org_id"`
\tActorID   string            `json:"actor_id"`
\tResourceID string           `json:"resource_id"`
\tOccurredAt time.Time        `json:"occurred_at"`
\tMetadata  map[string]string `json:"metadata,omitempty"`
}}

func New{title}Payload(orgID, actorID, resourceID string, meta map[string]string) (*{title}Payload, error) {{
\tif orgID == "" || actorID == "" {{
\t\treturn nil, common.New(common.ErrCodeInvalidArgument, "org_id and actor_id required")
\t}}
\treturn &{title}Payload{{
\t\tEvent: "{event}", OrgID: orgID, ActorID: actorID, ResourceID: resourceID,
\t\tOccurredAt: common.NowUTC(), Metadata: meta,
\t}}, nil
}}

func (p *{title}Payload) JSON() ([]byte, error) {{
\treturn json.Marshal(p)
}}

func (p *{title}Payload) Validate() error {{
\tif p.Event != "{event}" {{
\t\treturn common.New(common.ErrCodeInvalidArgument, "event type mismatch")
\t}}
\treturn nil
}}
'''


def gen_webhook_event_test(event):
    safe = event.replace(".", "_")
    title = "".join(w.capitalize() for w in safe.split("_"))
    return f'''package webhooks

import "testing"

func Test{title}Payload(t *testing.T) {{
\tp, err := New{title}Payload("o1", "u1", "r1", map[string]string{{"k": "v"}})
\tif err != nil {{
\t\tt.Fatal(err)
\t}}
\tif err := p.Validate(); err != nil {{
\t\tt.Fatal(err)
\t}}
\tb, err := p.JSON()
\tif err != nil || len(b) == 0 {{
\t\tt.Fatal(err)
\t}}
}}
'''


def main() -> None:
    for svc, pkg, entity, domain in APPS:
        write(ROOT / f"apps/{svc}/internal/acl/rules.go", gen_acl_rules(svc, pkg, entity, domain))
        write(ROOT / f"apps/{svc}/internal/acl/rules_test.go", gen_acl_rules_test(entity))

    for name in DASHBOARDS:
        write(ROOT / f"internal/analytics/dashboard_{name}.go", gen_dashboard(name))
        write(ROOT / f"internal/analytics/dashboard_{name}_test.go", gen_dashboard_test(name))

    for event in WEBHOOK_EVENTS:
        safe = event.replace(".", "_")
        write(ROOT / f"internal/webhooks/events/{safe}.go", gen_webhook_event(event))
        write(ROOT / f"internal/webhooks/events/{safe}_test.go", gen_webhook_event_test(event))

    print("phase D complete")


if __name__ == "__main__":
    main()
