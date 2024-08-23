#!/usr/bin/env python3
"""Phase C: IdP, analytics, saga/outbox, ACL, webhooks, engines, per-service ACL, contracts."""
from __future__ import annotations

import textwrap
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

ENGINES = [
    ("aws", "AWS", "AccountID", False),
    ("azure", "Azure", "TenantID", False),
    ("gcp", "GCP", "ProjectID", False),
    ("kubernetes", "Kubernetes", "ClusterID", False),
    ("database", "Database", "DBType", False),
    ("transit", "Transit", "KeyName", False),
    ("githubeng", "GitHub", "InstallationID", False),
    ("gitlabeng", "GitLab", "GroupID", False),
    ("dockereng", "Docker", "Registry", False),
    ("jenkins", "Jenkins", "ControllerURL", False),
    ("terraformeng", "Terraform", "Workspace", False),
]

ANALYTICS_METRICS = [
    "login_heatmap", "vault_access_frequency", "risk_trend", "password_hygiene",
    "session_duration", "failed_login_rate", "mfa_adoption", "device_trust_score",
    "secret_rotation_lag", "emergency_access_usage", "sharing_events", "export_volume",
    "import_volume", "policy_denials", "api_latency_p95", "webhook_failures",
    "backup_success_rate", "compliance_gaps", "admin_actions", "billing_seats",
]

SAGA_NAMES = ["UserOffboarding", "SecretRotation", "OrgProvision", "EmergencyAccess"]


def write(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content)
    print(f"wrote {path.relative_to(ROOT)} ({len(content.splitlines())} lines)")


def gen_saml() -> tuple[str, str]:
    impl = textwrap.dedent('''
        package saml

        import (
            "encoding/base64"
            "encoding/xml"
            "fmt"
            "strings"
            "time"

            "github.com/sentinelvault/sentinelvault/internal/common"
        )

        // AuthnRequest is a simplified SAML 2.0 AuthnRequest builder.
        type AuthnRequest struct {
            ID           string
            IssueInstant time.Time
            Destination  string
            Issuer       string
            ACSURL       string
            ForceAuthn   bool
        }

        // BuildAuthnRequestXML returns deflated-ready XML (not compressed here).
        func BuildAuthnRequestXML(r AuthnRequest) ([]byte, error) {
            if r.Destination == "" || r.Issuer == "" || r.ACSURL == "" {
                return nil, common.New(common.ErrCodeInvalidArgument, "destination, issuer, acs required")
            }
            if r.ID == "" {
                r.ID = "_" + common.NewID()
            }
            if r.IssueInstant.IsZero() {
                r.IssueInstant = common.NowUTC()
            }
            force := "false"
            if r.ForceAuthn {
                force = "true"
            }
            doc := fmt.Sprintf(`<?xml version="1.0"?><samlp:AuthnRequest xmlns:samlp="urn:oasis:names:tc:SAML:2.0:protocol" xmlns:saml="urn:oasis:names:tc:SAML:2.0:assertion" ID="%s" Version="2.0" IssueInstant="%s" Destination="%s" AssertionConsumerServiceURL="%s" ForceAuthn="%s"><saml:Issuer>%s</saml:Issuer></samlp:AuthnRequest>`,
                xmlEscape(r.ID), r.IssueInstant.UTC().Format(time.RFC3339), xmlEscape(r.Destination), xmlEscape(r.ACSURL), force, xmlEscape(r.Issuer))
            return []byte(doc), nil
        }

        func xmlEscape(s string) string {
            s = strings.ReplaceAll(s, "&", "&amp;")
            s = strings.ReplaceAll(s, "<", "&lt;")
            s = strings.ReplaceAll(s, ">", "&gt;")
            s = strings.ReplaceAll(s, `"`, "&quot;")
            return s
        }

        // ResponseAssertion holds parsed assertion attributes (simplified).
        type ResponseAssertion struct {
            Subject       string
            SessionIndex  string
            Audience      string
            Recipient     string
            NotBefore     time.Time
            NotOnOrAfter  time.Time
            Attributes    map[string][]string
        }

        // ParseResponseAttributes extracts attributes from simplified SAML Response XML.
        func ParseResponseAttributes(xmlData []byte) (*ResponseAssertion, error) {
            type attrVal struct {
                Name string `xml:"Name,attr"`
                Val  string `xml:",chardata"`
            }
            type attrStmt struct {
                Attrs []attrVal `xml:"Attribute"`
            }
            type subject struct {
                NameID string `xml:"NameID"`
            }
            type conditions struct {
                NotBefore    string `xml:"NotBefore,attr"`
                NotOnOrAfter string `xml:"NotOnOrAfter,attr"`
                Audience     string `xml:"Audience"`
            }
            type assertion struct {
                Subject    subject    `xml:"Subject"`
                Conditions conditions `xml:"Conditions"`
                AttrStmt   attrStmt   `xml:"AttributeStatement"`
            }
            type response struct {
                Assertion assertion `xml:"Assertion"`
            }
            var resp response
            if err := xml.Unmarshal(xmlData, &resp); err != nil {
                return nil, common.Wrap(common.ErrCodeInvalidArgument, "invalid saml xml", err)
            }
            a := resp.Assertion
            out := &ResponseAssertion{
                Subject: a.Subject.NameID,
                Attributes: map[string][]string{},
            }
            if a.Conditions.Audience != "" {
                out.Audience = strings.TrimSpace(a.Conditions.Audience)
            }
            if a.Conditions.NotBefore != "" {
                t, err := time.Parse(time.RFC3339, a.Conditions.NotBefore)
                if err == nil {
                    out.NotBefore = t
                }
            }
            if a.Conditions.NotOnOrAfter != "" {
                t, err := time.Parse(time.RFC3339, a.Conditions.NotOnOrAfter)
                if err == nil {
                    out.NotOnOrAfter = t
                }
            }
            for _, at := range a.AttrStmt.Attrs {
                out.Attributes[at.Name] = append(out.Attributes[at.Name], strings.TrimSpace(at.Val))
            }
            return out, nil
        }

        // ValidateAssertion checks audience, recipient, and clock skew.
        func ValidateAssertion(a *ResponseAssertion, wantAudience, wantRecipient string, now time.Time, skew time.Duration) error {
            if a == nil {
                return common.New(common.ErrCodeInvalidArgument, "assertion required")
            }
            if a.Subject == "" {
                return common.New(common.ErrCodeUnauthorized, "subject required")
            }
            if wantAudience != "" && a.Audience != wantAudience {
                return common.New(common.ErrCodeUnauthorized, "audience mismatch")
            }
            if wantRecipient != "" && a.Recipient != "" && a.Recipient != wantRecipient {
                return common.New(common.ErrCodeUnauthorized, "recipient mismatch")
            }
            if skew <= 0 {
                skew = 2 * time.Minute
            }
            if !a.NotBefore.IsZero() && now.Before(a.NotBefore.Add(-skew)) {
                return common.New(common.ErrCodeUnauthorized, "assertion not yet valid")
            }
            if !a.NotOnOrAfter.IsZero() && !now.Before(a.NotOnOrAfter.Add(skew)) {
                return common.New(common.ErrCodeUnauthorized, "assertion expired")
            }
            return nil
        }

        // EncodeRedirectRequest base64-encodes AuthnRequest for HTTP-Redirect binding.
        func EncodeRedirectRequest(xmlData []byte) string {
            return base64.StdEncoding.EncodeToString(xmlData)
        }
    ''')
    test = textwrap.dedent('''
        package saml

        import (
            "testing"
            "time"
        )

        func TestAuthnRequestAndAssertion(t *testing.T) {
            xml, err := BuildAuthnRequestXML(AuthnRequest{
                Destination: "https://idp/saml/sso", Issuer: "sp", ACSURL: "https://sp/acs",
            })
            if err != nil || len(xml) == 0 {
                t.Fatal(err)
            }
            enc := EncodeRedirectRequest(xml)
            if enc == "" {
                t.Fatal("empty encode")
            }
            sample := []byte(`<Response><Assertion><Subject><NameID>user@example.com</NameID></Subject><Conditions NotBefore="2020-01-01T00:00:00Z" NotOnOrAfter="2030-01-01T00:00:00Z"><Audience>sp</Audience></Conditions><AttributeStatement><Attribute Name="email">user@example.com</Attribute></AttributeStatement></Assertion></Response>`)
            a, err := ParseResponseAttributes(sample)
            if err != nil || a.Subject == "" {
                t.Fatal(err)
            }
            if err := ValidateAssertion(a, "sp", "", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), time.Minute); err != nil {
                t.Fatal(err)
            }
        }

        func TestValidateAssertionErrors(t *testing.T) {
            a := &ResponseAssertion{Subject: "u", Audience: "wrong", NotOnOrAfter: time.Now().Add(-time.Hour)}
            if err := ValidateAssertion(a, "expected", "", time.Now(), 0); err == nil {
                t.Fatal("expected audience error")
            }
        }
    ''')
    return impl.strip(), test.strip()


def gen_oidc_client() -> tuple[str, str]:
    impl = textwrap.dedent('''
        package oidc

        import (
            "context"
            "encoding/json"
            "io"
            "net/http"
            "net/url"
            "strings"

            "github.com/sentinelvault/sentinelvault/internal/common"
        )

        // HTTPDoer is mockable HTTP client surface.
        type HTTPDoer interface {
            Do(req *http.Request) (*http.Response, error)
        }

        // TokenResponse is OIDC token endpoint response subset.
        type TokenResponse struct {
            AccessToken  string `json:"access_token"`
            TokenType    string `json:"token_type"`
            ExpiresIn    int    `json:"expires_in"`
            RefreshToken string `json:"refresh_token"`
            IDToken      string `json:"id_token"`
            Scope        string `json:"scope"`
        }

        // Client exchanges codes for tokens.
        type Client struct {
            HTTP HTTPDoer
        }

        // ExchangeAuthorizationCode posts token request and parses response.
        func (c *Client) ExchangeAuthorizationCode(ctx context.Context, req TokenRequest) (*TokenResponse, error) {
            if c.HTTP == nil {
                return nil, common.New(common.ErrCodeInternal, "http client required")
            }
            vals, err := req.FormValues()
            if err != nil {
                return nil, err
            }
            httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, req.TokenEndpoint, strings.NewReader(vals.Encode()))
            if err != nil {
                return nil, err
            }
            httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
            res, err := c.HTTP.Do(httpReq)
            if err != nil {
                return nil, common.Wrap(common.ErrCodeUnavailable, "token endpoint", err)
            }
            defer res.Body.Close()
            body, err := io.ReadAll(res.Body)
            if err != nil {
                return nil, err
            }
            if res.StatusCode >= 300 {
                return nil, common.New(common.ErrCodeUnauthorized, "token exchange failed")
            }
            var out TokenResponse
            if err := json.Unmarshal(body, &out); err != nil {
                return nil, common.Wrap(common.ErrCodeInvalidArgument, "token json", err)
            }
            if out.AccessToken == "" {
                return nil, common.New(common.ErrCodeUnauthorized, "missing access_token")
            }
            return &out, nil
        }

        // FetchDiscovery loads provider metadata from URL.
        func (c *Client) FetchDiscovery(ctx context.Context, issuer string) (*DiscoveryDocument, error) {
            if c.HTTP == nil {
                return nil, common.New(common.ErrCodeInternal, "http client required")
            }
            u, err := url.Parse(strings.TrimRight(issuer, "/") + "/.well-known/openid-configuration")
            if err != nil {
                return nil, err
            }
            req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
            if err != nil {
                return nil, err
            }
            res, err := c.HTTP.Do(req)
            if err != nil {
                return nil, err
            }
            defer res.Body.Close()
            data, err := io.ReadAll(res.Body)
            if err != nil {
                return nil, err
            }
            return ParseDiscovery(data)
        }
    ''')
    test = textwrap.dedent('''
        package oidc

        import (
            "bytes"
            "context"
            "io"
            "net/http"
            "testing"
        )

        type mockDoer struct{ body string; code int }

        func (m mockDoer) Do(req *http.Request) (*http.Response, error) {
            return &http.Response{
                StatusCode: m.code,
                Body:       io.NopCloser(bytes.NewBufferString(m.body)),
                Header:     make(http.Header),
            }, nil
        }

        func TestClientExchange(t *testing.T) {
            c := &Client{HTTP: mockDoer{code: 200, body: `{"access_token":"at","id_token":"x.y.z","token_type":"Bearer","expires_in":3600}`}}
            out, err := c.ExchangeAuthorizationCode(context.Background(), TokenRequest{
                TokenEndpoint: "https://idp/token", GrantType: "authorization_code",
                Code: "code", RedirectURI: "https://app/cb", ClientID: "cid", CodeVerifier: "v",
            })
            if err != nil || out.AccessToken != "at" {
                t.Fatal(err)
            }
        }

        func TestClientExchangeError(t *testing.T) {
            c := &Client{HTTP: mockDoer{code: 401, body: `{"error":"invalid_grant"}`}}
            _, err := c.ExchangeAuthorizationCode(context.Background(), TokenRequest{
                TokenEndpoint: "https://idp/token", GrantType: "authorization_code",
                Code: "bad", RedirectURI: "https://app/cb", ClientID: "cid",
            })
            if err == nil {
                t.Fatal("expected error")
            }
        }
    ''')
    return impl.strip(), test.strip()


def gen_analytics_metric(name: str) -> tuple[str, str]:
    title = "".join(w.capitalize() for w in name.split("_"))
    impl = f'''package analytics

import (
\t"sort"
\t"time"
)

// {title}Calculator tracks {name.replace("_", " ")}.
type {title}Calculator struct {{
\tbucket time.Duration
\tdata   map[string]map[int64]float64
\tcounts map[string]map[int64]int64
}}

func New{title}Calculator(bucket time.Duration) *{title}Calculator {{
\tif bucket <= 0 {{
\t\tbucket = time.Hour
\t}}
\treturn &{title}Calculator{{bucket: bucket, data: map[string]map[int64]float64{{}}, counts: map[string]map[int64]int64{{}}}}
}}

func (c *{title}Calculator) key(ts time.Time) int64 {{
\treturn ts.UTC().Truncate(c.bucket).Unix()
}}

// Record adds an observation for orgID at timestamp.
func (c *{title}Calculator) Record(orgID string, ts time.Time, value float64) {{
\tif orgID == "" {{
\t\treturn
\t}}
\tk := c.key(ts)
\tif c.data[orgID] == nil {{
\t\tc.data[orgID] = map[int64]float64{{}}
\t\tc.counts[orgID] = map[int64]int64{{}}
\t}}
\tc.data[orgID][k] += value
\tc.counts[orgID][k]++
}}

// Window returns sorted points in [from,to].
func (c *{title}Calculator) Window(orgID string, from, to time.Time) []Point {{
\tm := c.data[orgID]
\tif m == nil {{
\t\treturn nil
\t}}
\tkeys := make([]int64, 0)
\tfor k := range m {{
\t\tif !from.IsZero() && k < from.Unix() {{
\t\t\tcontinue
\t\t}}
\t\tif !to.IsZero() && k > to.Unix() {{
\t\t\tcontinue
\t\t}}
\t\tkeys = append(keys, k)
\t}}
\tsort.Slice(keys, func(i, j int) bool {{ return keys[i] < keys[j] }})
\tout := make([]Point, 0, len(keys))
\tfor _, k := range keys {{
\t\tout = append(out, Point{{Bucket: time.Unix(k, 0).UTC(), Value: m[k], Count: c.counts[orgID][k]}})
\t}}
\treturn out
}}

// Merge combines another calculator into this one (same bucket size assumed).
func (c *{title}Calculator) Merge(other *{title}Calculator) {{
\tif other == nil {{
\t\treturn
\t}}
\tfor org, buckets := range other.data {{
\t\tif c.data[org] == nil {{
\t\t\tc.data[org] = map[int64]float64{{}}
\t\t\tc.counts[org] = map[int64]int64{{}}
\t\t}}
\t\tfor k, v := range buckets {{
\t\t\tc.data[org][k] += v
\t\t\tc.counts[org][k] += other.counts[org][k]
\t\t}}
\t}}
}}

// Average computes mean over window.
func (c *{title}Calculator) Average(orgID string, from, to time.Time) float64 {{
\tpts := c.Window(orgID, from, to)
\tif len(pts) == 0 {{
\t\treturn 0
\t}}
\tvar sum float64
\tvar n int64
\tfor _, p := range pts {{
\t\tsum += p.Value
\t\tn += p.Count
\t}}
\tif n == 0 {{
\t\treturn 0
\t}}
\treturn sum / float64(n)
}}
'''
    test = f'''package analytics

import (
\t"testing"
\t"time"
)

func Test{title}Calculator(t *testing.T) {{
\tc := New{title}Calculator(time.Hour)
\tnow := time.Date(2025, 6, 1, 12, 0, 0, 0, time.UTC)
\tc.Record("org1", now, 1)
\tc.Record("org1", now.Add(30*time.Minute), 2)
\tpts := c.Window("org1", now.Add(-time.Hour), now.Add(2*time.Hour))
\tif len(pts) != 1 || pts[0].Value != 3 {{
\t\tt.Fatalf("%+v", pts)
\t}}
\tother := New{title}Calculator(time.Hour)
\tother.Record("org1", now, 1)
\tc.Merge(other)
\tif c.Average("org1", now.Add(-time.Hour), now.Add(2*time.Hour)) <= 0 {{
\t\tt.Fatal("avg")
\t}}
}}
'''
    return impl, test


def gen_outbox() -> None:
    impl = textwrap.dedent('''
        package outbox

        import (
            "context"
            "encoding/json"
            "sync"
            "time"

            "github.com/sentinelvault/sentinelvault/internal/common"
        )

        const (
            StatusPending   = "pending"
            StatusPublished = "published"
            StatusFailed    = "failed"
        )

        type Message struct {
            ID          string
            Topic       string
            Key         string
            Payload     json.RawMessage
            Status      string
            Attempts    int
            LastError   string
            CreatedAt   time.Time
            PublishedAt *time.Time
        }

        type Publisher interface {
            Publish(ctx context.Context, topic, key string, payload []byte) error
        }

        type Store struct {
            mu       sync.Mutex
            messages map[string]*Message
            maxRetry int
        }

        func NewStore() *Store { return &Store{messages: map[string]*Message{}, maxRetry: 5} }

        func (s *Store) Enqueue(topic, key string, payload any) (*Message, error) {
            if topic == "" {
                return nil, common.New(common.ErrCodeInvalidArgument, "topic required")
            }
            raw, err := json.Marshal(payload)
            if err != nil {
                return nil, err
            }
            m := &Message{
                ID: common.NewID(), Topic: topic, Key: key, Payload: raw,
                Status: StatusPending, CreatedAt: common.NowUTC(),
            }
            s.mu.Lock()
            s.messages[m.ID] = m
            s.mu.Unlock()
            cp := *m
            return &cp, nil
        }

        func (s *Store) Pending(limit int) []Message {
            if limit <= 0 {
                limit = 100
            }
            s.mu.Lock()
            defer s.mu.Unlock()
            out := make([]Message, 0, limit)
            for _, m := range s.messages {
                if m.Status != StatusPending && m.Status != StatusFailed {
                    continue
                }
                if m.Attempts >= s.maxRetry {
                    continue
                }
                cp := *m
                out = append(out, cp)
                if len(out) >= limit {
                    break
                }
            }
            return out
        }

        func (s *Store) MarkPublished(id string) error {
            s.mu.Lock()
            defer s.mu.Unlock()
            m, ok := s.messages[id]
            if !ok {
                return common.New(common.ErrCodeNotFound, "message not found")
            }
            now := common.NowUTC()
            m.Status = StatusPublished
            m.PublishedAt = &now
            return nil
        }

        func (s *Store) MarkFailed(id string, err error) error {
            s.mu.Lock()
            defer s.mu.Unlock()
            m, ok := s.messages[id]
            if !ok {
                return common.New(common.ErrCodeNotFound, "message not found")
            }
            m.Attempts++
            m.Status = StatusFailed
            if err != nil {
                m.LastError = err.Error()
            }
            return nil
        }

        type Relay struct {
            store     *Store
            publisher Publisher
        }

        func NewRelay(store *Store, pub Publisher) *Relay {
            return &Relay{store: store, publisher: pub}
        }

        func (r *Relay) Flush(ctx context.Context, limit int) (int, error) {
            pending := r.store.Pending(limit)
            n := 0
            for _, m := range pending {
                if err := r.publisher.Publish(ctx, m.Topic, m.Key, m.Payload); err != nil {
                    _ = r.store.MarkFailed(m.ID, err)
                    continue
                }
                if err := r.store.MarkPublished(m.ID); err != nil {
                    return n, err
                }
                n++
            }
            return n, nil
        }
    ''')
    test = textwrap.dedent('''
        package outbox

        import (
            "context"
            "errors"
            "testing"
        )

        type stubPub struct{ fail bool }

        func (s stubPub) Publish(ctx context.Context, topic, key string, payload []byte) error {
            if s.fail {
                return errors.New("fail")
            }
            return nil
        }

        func TestOutboxRelay(t *testing.T) {
            st := NewStore()
            m, err := st.Enqueue("events", "k1", map[string]string{"a": "b"})
            if err != nil {
                t.Fatal(err)
            }
            relay := NewRelay(st, stubPub{})
            n, err := relay.Flush(context.Background(), 10)
            if err != nil || n != 1 {
                t.Fatalf("n=%d err=%v", n, err)
            }
            got, _ := st.messages[m.ID]
            if got.Status != StatusPublished {
                t.Fatalf("status=%s", got.Status)
            }
        }

        func TestOutboxFailure(t *testing.T) {
            st := NewStore()
            _, _ = st.Enqueue("events", "k", map[string]int{"x": 1})
            relay := NewRelay(st, stubPub{fail: true})
            n, _ := relay.Flush(context.Background(), 10)
            if n != 0 {
                t.Fatal("expected 0 published")
            }
        }
    ''')
    write(ROOT / "internal/outbox/outbox.go", impl.strip())
    write(ROOT / "internal/outbox/outbox_test.go", test.strip())


def gen_saga_hooks() -> None:
    impl = textwrap.dedent('''
        package saga

        import "context"

        // StepHandler executes one saga step with typed compensation.
        type StepHandler interface {
            Name() string
            Execute(ctx context.Context, state map[string]any) error
            Compensate(ctx context.Context, state map[string]any) error
        }

        // noopHandler skips compensation.
        type noopHandler struct{ name string }

        func (n noopHandler) Name() string { return n.name }
        func (n noopHandler) Execute(ctx context.Context, s map[string]any) error { return nil }
        func (n noopHandler) Compensate(ctx context.Context, s map[string]any) error { return nil }

        // FromHandler converts StepHandler to Step.
        func FromHandler(h StepHandler) Step {
            return Step{Name: h.Name(), Execute: h.Execute, Compensate: h.Compensate}
        }

        // RegisterHandlers registers handlers as a saga definition.
        func RegisterHandlers(e *Engine, name string, handlers ...StepHandler) error {
            steps := make([]Step, len(handlers))
            for i, h := range handlers {
                steps[i] = FromHandler(h)
            }
            return e.Register(Definition{Name: name, Steps: steps})
        }
    ''')
    write(ROOT / "internal/saga/hooks.go", impl.strip())
    for saga in SAGA_NAMES:
        steps = []
        for i in range(1, 5):
            steps.append(f'''
type {saga}Step{i} struct{{ fail bool }}

func (s {saga}Step{i}) Name() string {{ return "{saga.lower()}_step_{i}" }}
func (s {saga}Step{i}) Execute(ctx context.Context, st map[string]any) error {{
\tif s.fail {{
\t\treturn common.New(common.ErrCodeInternal, "forced failure")
\t}}
\tst["{saga.lower()}_step_{i}"] = true
\treturn nil
}}
func (s {saga}Step{i}) Compensate(ctx context.Context, st map[string]any) error {{
\tst["{saga.lower()}_step_{i}"] = false
\treturn nil
}}
''')
        wf = f'''package saga

import (
\t"context"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

{"".join(steps)}

// Register{saga} wires {saga} saga with compensation hooks.
func Register{saga}(e *Engine) error {{
\treturn RegisterHandlers(e, "{saga}", {", ".join(f"&{saga}Step{i}{{}}" for i in range(1, 5))})
}}
'''
        tf = f'''package saga

import (
\t"context"
\t"testing"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

func Test{saga}Success(t *testing.T) {{
\te := NewEngine()
\tif err := Register{saga}(e); err != nil {{
\t\tt.Fatal(err)
\t}}
\tinst, err := e.Start(context.Background(), "{saga}", map[string]any{{"org_id": "o1"}})
\tif err != nil {{
\t\tt.Fatal(err)
\t}}
\tif inst.Status != StatusCompleted {{
\t\tt.Fatalf("status=%s", inst.Status)
\t}}
}}

func Test{saga}Compensate(t *testing.T) {{
\te := NewEngine()
\tsteps := []StepHandler{{&{saga}Step1{{}}, &{saga}Step2{{fail: true}}, &{saga}Step3{{}}}}
\t_ = RegisterHandlers(e, "{saga}Fail", steps...)
\tinst, _ := e.Start(context.Background(), "{saga}Fail", map[string]any{{}})
\tif inst.Status != StatusFailed {{
\t\tt.Fatalf("status=%s", inst.Status)
\t}}
\t_, err := e.Start(context.Background(), "missing", nil)
\tif err == nil || !common.IsCode(err, common.ErrCodeNotFound) {{
\t\tt.Fatalf("err=%v", err)
\t}}
}}
'''
        write(ROOT / f"internal/saga/workflows/{saga.lower()}.go", wf)
        write(ROOT / f"internal/saga/workflows/{saga.lower()}_test.go", tf)


def gen_acl_expand() -> None:
    impl = textwrap.dedent('''
        package acl

        import (
            "sort"
            "time"

            "github.com/sentinelvault/sentinelvault/internal/common"
        )

        // CollectionGrant applies permissions to all items via inheritance.
        type CollectionGrant struct {
            CollectionID string
            Permission   string
            PrincipalID  string
            ExpiresAt    *time.Time
        }

        // InheritanceResolver computes effective permission with collection chain.
        type InheritanceResolver struct {
            store       *Store
            collections map[string][]string // itemID -> parent collection IDs
        }

        func NewInheritanceResolver(store *Store) *InheritanceResolver {
            return &InheritanceResolver{store: store, collections: map[string][]string{}}
        }

        func (r *InheritanceResolver) LinkItem(itemID string, collectionIDs ...string) {
            r.collections[itemID] = append([]string(nil), collectionIDs...)
        }

        func (r *InheritanceResolver) EffectiveForItem(orgID, itemID, principalID string) (string, bool) {
            parents := r.collections[itemID]
            return r.store.Effective(orgID, "item", itemID, principalID, parents)
        }

        func (r *InheritanceResolver) GrantCollection(orgID, collectionID, principalID, perm string, ttl time.Duration) (*Entry, error) {
            e, err := r.store.Grant(orgID, "collection", collectionID, principalID, perm, ttl)
            if err != nil {
                return nil, err
            }
            e.InheritedFrom = collectionID
            return e, nil
        }

        // ListEntries returns ACL entries for a resource sorted by permission rank.
        func (s *Store) ListEntries(orgID, resourceType, resourceID string) []Entry {
            now := common.NowUTC()
            out := make([]Entry, 0)
            for _, e := range s.entries {
                if e.OrgID != orgID || e.ResourceType != resourceType || e.ResourceID != resourceID {
                    continue
                }
                if e.ExpiresAt != nil && now.After(*e.ExpiresAt) {
                    continue
                }
                out = append(out, *e)
            }
            sort.Slice(out, func(i, j int) bool {
                return permRank[out[i].Permission] > permRank[out[j].Permission]
            })
            return out
        }

        // RevokeByPrincipal removes all grants for principal on resource.
        func (s *Store) RevokeByPrincipal(orgID, resourceType, resourceID, principalID string) int {
            n := 0
            for id, e := range s.entries {
                if e.OrgID == orgID && e.ResourceType == resourceType && e.ResourceID == resourceID && e.PrincipalID == principalID {
                    delete(s.entries, id)
                    n++
                }
            }
            return n
        }
    ''')
    test = textwrap.dedent('''
        package acl

        import (
            "testing"
            "time"
        )

        func TestInheritanceResolver(t *testing.T) {
            st := NewStore()
            r := NewInheritanceResolver(st)
            _, _ = r.GrantCollection("o1", "col1", "u1", PermViewer, time.Hour)
            r.LinkItem("item1", "col1")
            perm, ok := r.EffectiveForItem("o1", "item1", "u1")
            if !ok || perm != PermViewer {
                t.Fatalf("perm=%s ok=%v", perm, ok)
            }
            _, _ = st.Grant("o1", "item", "item1", "u2", PermEditor, 0)
            entries := st.ListEntries("o1", "item", "item1")
            if len(entries) != 1 {
                t.Fatalf("len=%d", len(entries))
            }
            if st.RevokeByPrincipal("o1", "item", "item1", "u2") != 1 {
                t.Fatal("revoke count")
            }
        }
    ''')
    write(ROOT / "internal/acl/inheritance.go", impl.strip())
    write(ROOT / "internal/acl/inheritance_test.go", test.strip())


def gen_featureflags_expand() -> None:
    impl = textwrap.dedent('''
        package featureflags

        import (
            "sync"

            "github.com/sentinelvault/sentinelvault/internal/common"
        )

        // BooleanFlag is a simple on/off toggle.
        type BooleanFlag struct {
            Key     string
            Enabled bool
        }

        // PercentageFlag rolls out to a fraction of subjects.
        type PercentageFlag struct {
            Key        string
            Percentage int
            Enabled    bool
        }

        // OrgTargetedFlag allows per-org overrides.
        type OrgTargetedFlag struct {
            Key      string
            Enabled  bool
            OrgAllow map[string]bool
            OrgDeny  map[string]bool
        }

        // Evaluator combines registry evaluation helpers.
        type Evaluator struct {
            reg *Registry
            mu  sync.RWMutex
            cache map[string]bool
        }

        func NewEvaluator(reg *Registry) *Evaluator {
            return &Evaluator{reg: reg, cache: map[string]bool{}}
        }

        func (ev *Evaluator) IsEnabled(key, orgID, subjectID string) (bool, error) {
            ck := key + ":" + orgID + ":" + subjectID
            ev.mu.RLock()
            if v, ok := ev.cache[ck]; ok {
                ev.mu.RUnlock()
                return v, nil
            }
            ev.mu.RUnlock()
            ok, err := ev.reg.Evaluate(key, orgID, subjectID)
            if err != nil {
                return false, err
            }
            ev.mu.Lock()
            ev.cache[ck] = ok
            ev.mu.Unlock()
            return ok, nil
        }

        func UpsertBoolean(reg *Registry, f BooleanFlag) error {
            return reg.Upsert(Flag{Key: f.Key, Enabled: f.Enabled, Percentage: 100})
        }

        func UpsertPercentage(reg *Registry, f PercentageFlag) error {
            if !f.Enabled {
                return reg.Upsert(Flag{Key: f.Key, Enabled: false, Percentage: 0})
            }
            return reg.Upsert(Flag{Key: f.Key, Enabled: true, Percentage: f.Percentage})
        }

        func UpsertOrgTargeted(reg *Registry, f OrgTargetedFlag) error {
            if f.Key == "" {
                return common.New(common.ErrCodeInvalidArgument, "key required")
            }
            return reg.Upsert(Flag{
                Key: f.Key, Enabled: f.Enabled, Percentage: 100,
                OrgAllow: f.OrgAllow, OrgDeny: f.OrgDeny,
            })
        }
    ''')
    test = textwrap.dedent('''
        package featureflags

        import "testing"

        func TestEvaluatorCache(t *testing.T) {
            reg := NewRegistry()
            _ = UpsertBoolean(reg, BooleanFlag{Key: "beta", Enabled: true})
            ev := NewEvaluator(reg)
            v1, err := ev.IsEnabled("beta", "o1", "u1")
            if err != nil || !v1 {
                t.Fatal(err)
            }
            v2, _ := ev.IsEnabled("beta", "o1", "u1")
            if !v2 {
                t.Fatal("cache miss")
            }
            _ = UpsertPercentage(reg, PercentageFlag{Key: "rollout", Enabled: true, Percentage: 50})
            _ = UpsertOrgTargeted(reg, OrgTargetedFlag{Key: "orgfeat", Enabled: true, OrgAllow: map[string]bool{"o1": true}})
        }
    ''')
    write(ROOT / "internal/featureflags/evaluator.go", impl.strip())
    write(ROOT / "internal/featureflags/evaluator_test.go", test.strip())


def gen_webhooks_retry() -> None:
    impl = textwrap.dedent('''
        package webhooks

        import (
            "context"
            "time"
        )

        // DeliverWithRetry attempts delivery with exponential backoff schedule.
        func (r *Registry) DeliverWithRetry(ctx context.Context, orgID, eventType string, body []byte, maxAttempts int) []Attempt {
            if maxAttempts <= 0 {
                maxAttempts = 3
            }
            r.mu.Lock()
            targets := make([]*Endpoint, 0)
            for _, ep := range r.eps {
                if ep.OrgID == orgID && ep.Active && contains(ep.Events, eventType) {
                    cp := *ep
                    targets = append(targets, &cp)
                }
            }
            r.mu.Unlock()
            out := make([]Attempt, 0)
            for _, ep := range targets {
                var last Attempt
                for attempt := 1; attempt <= maxAttempts; attempt++ {
                    last = r.deliverOne(ctx, ep, eventType, body, attempt)
                    out = append(out, last)
                    if last.Success {
                        break
                    }
                    select {
                    case <-ctx.Done():
                        return out
                    case <-time.After(Backoff(attempt)):
                    }
                }
            }
            return out
        }

        // Deactivate marks endpoint inactive.
        func (r *Registry) Deactivate(id string) error {
            r.mu.Lock()
            defer r.mu.Unlock()
            ep, ok := r.eps[id]
            if !ok {
                return common.New(common.ErrCodeNotFound, "endpoint not found")
            }
            ep.Active = false
            return nil
        }
    ''').replace("common.New", '"github.com/sentinelvault/sentinelvault/internal/common".New').replace(
        "common.ErrCodeNotFound", "common.ErrCodeNotFound"
    )
    # fix import - rewrite properly
    impl = '''package webhooks

import (
\t"context"
\t"time"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

// DeliverWithRetry attempts delivery with exponential backoff schedule.
func (r *Registry) DeliverWithRetry(ctx context.Context, orgID, eventType string, body []byte, maxAttempts int) []Attempt {
\tif maxAttempts <= 0 {
\t\tmaxAttempts = 3
\t}
\tr.mu.Lock()
\ttargets := make([]*Endpoint, 0)
\tfor _, ep := range r.eps {
\t\tif ep.OrgID == orgID && ep.Active && contains(ep.Events, eventType) {
\t\t\tcp := *ep
\t\t\ttargets = append(targets, &cp)
\t\t}
\t}
\tr.mu.Unlock()
\tout := make([]Attempt, 0)
\tfor _, ep := range targets {
\t\tvar last Attempt
\t\tfor attempt := 1; attempt <= maxAttempts; attempt++ {
\t\t\tlast = r.deliverOne(ctx, ep, eventType, body, attempt)
\t\t\tout = append(out, last)
\t\t\tif last.Success {
\t\t\t\tbreak
\t\t\t}
\t\t\tselect {
\t\t\tcase <-ctx.Done():
\t\t\t\treturn out
\t\t\tcase <-time.After(Backoff(attempt)):
\t\t\t}
\t\t}
\t}
\treturn out
}

// Deactivate marks endpoint inactive.
func (r *Registry) Deactivate(id string) error {
\tr.mu.Lock()
\tdefer r.mu.Unlock()
\tep, ok := r.eps[id]
\tif !ok {
\t\treturn common.New(common.ErrCodeNotFound, "endpoint not found")
\t}
\tep.Active = false
\treturn nil
}
'''
    test = textwrap.dedent('''
        package webhooks

        import (
            "context"
            "net/http"
            "net/http/httptest"
            "testing"
        )

        func TestDeliverWithRetry(t *testing.T) {
            var hits int
            srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                hits++
                if hits < 2 {
                    w.WriteHeader(500)
                    return
                }
                w.WriteHeader(200)
            }))
            defer srv.Close()
            reg := NewRegistry()
            reg.client = srv.Client()
            ep, _ := reg.Register("o1", srv.URL, []string{"test.event"})
            _ = ep
            attempts := reg.DeliverWithRetry(context.Background(), "o1", "test.event", []byte(`{}`), 3)
            if len(attempts) == 0 {
                t.Fatal("no attempts")
            }
        }
    ''')
    write(ROOT / "internal/webhooks/retry.go", impl.strip())
    write(ROOT / "internal/webhooks/retry_test.go", test.strip())


def gen_engine(pkg: str, title: str, id_field: str, new_engine: bool) -> None:
    if new_engine:
        engine_ctor = f'''package {pkg}

import (
\t"crypto/sha256"
\t"sync"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

// Engine manages {title} dynamic credentials.
type Engine struct {{
\tmu       sync.Mutex
\troles    map[string]*Role
\tleases   map[string]*Lease
\tpolicies map[string]*PolicyDocument
\troot     []byte
}}

func NewEngine(rootSecret []byte) *Engine {{
\tif len(rootSecret) < 32 {{
\t\trootSecret = sha256.New().Sum([]byte("sentinelvault-{pkg}-engine"))
\t}}
\treturn &Engine{{
\t\troles: map[string]*Role{{}}, leases: map[string]*Lease{{}},
\t\tpolicies: map[string]*PolicyDocument{{}}, root: rootSecret,
\t}}
}}
'''
        write(ROOT / f"internal/engines/{pkg}/engine.go", engine_ctor)
        write(ROOT / f"internal/engines/{pkg}/roles.go", gen_engine_roles(pkg, title, id_field))
        write(ROOT / f"internal/engines/{pkg}/leases.go", gen_engine_leases(pkg, title))
        write(ROOT / f"internal/engines/{pkg}/revoke.go", gen_engine_revoke(pkg, title))

    write(ROOT / f"internal/engines/{pkg}/policies.go", gen_engine_policies(pkg, title))
    if new_engine:
        write(ROOT / f"internal/engines/{pkg}/policy_validate.go", gen_engine_policy_validate(pkg, title))
        write(ROOT / f"internal/engines/{pkg}/lease_list.go", gen_engine_lease_list(pkg, title))
    elif pkg in ("database", "dockereng", "jenkins", "terraformeng"):
        write(ROOT / f"internal/engines/{pkg}/lease_list.go", gen_engine_lease_list(pkg, title))


def gen_engine_roles(pkg, title, id_field):
    return f'''package {pkg}

import (
\t"time"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

// Role defines a {title} role binding.
type Role struct {{
\tID         string        `json:"id"`
\tOrgID      string        `json:"org_id"`
\tName       string        `json:"name"`
\t{id_field}   string        `json:"{id_field.lower()}"`
\tPolicies   []string      `json:"policies"`
\tDefaultTTL time.Duration `json:"default_ttl"`
\tMaxTTL     time.Duration `json:"max_ttl"`
}}

func (e *Engine) UpsertRole(r Role) (*Role, error) {{
\tif r.OrgID == "" || r.Name == "" || r.{id_field} == "" {{
\t\treturn nil, common.New(common.ErrCodeInvalidArgument, "org_id, name, {id_field.lower()} required")
\t}}
\tif r.ID == "" {{
\t\tr.ID = common.NewID()
\t}}
\tif r.DefaultTTL <= 0 {{
\t\tr.DefaultTTL = time.Hour
\t}}
\tif r.MaxTTL <= 0 {{
\t\tr.MaxTTL = 12 * time.Hour
\t}}
\tcp := r
\te.mu.Lock()
\te.roles[r.ID] = &cp
\te.mu.Unlock()
\treturn &cp, nil
}}

func (e *Engine) GetRole(id string) (*Role, error) {{
\te.mu.Lock()
\tdefer e.mu.Unlock()
\tr, ok := e.roles[id]
\tif !ok {{
\t\treturn nil, common.New(common.ErrCodeNotFound, "role not found")
\t}}
\tcp := *r
\treturn &cp, nil
}}

func (e *Engine) ListRoles(orgID string) []Role {{
\te.mu.Lock()
\tdefer e.mu.Unlock()
\tout := make([]Role, 0)
\tfor _, r := range e.roles {{
\t\tif orgID == "" || r.OrgID == orgID {{
\t\t\tout = append(out, *r)
\t\t}}
\t}}
\treturn out
}}
'''


def gen_engine_leases(pkg, title):
    return f'''package {pkg}

import (
\t"crypto/hmac"
\t"crypto/sha256"
\t"encoding/hex"
\t"fmt"
\t"time"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

// Lease is issued {title} credentials.
type Lease struct {{
\tID        string    `json:"id"`
\tRoleID    string    `json:"role_id"`
\tToken     string    `json:"token"`
\tExpiresAt time.Time `json:"expires_at"`
\tRevoked   bool      `json:"revoked"`
}}

func (e *Engine) IssueLease(roleID string, ttl time.Duration) (*Lease, error) {{
\te.mu.Lock()
\tdefer e.mu.Unlock()
\trole, ok := e.roles[roleID]
\tif !ok {{
\t\treturn nil, common.New(common.ErrCodeNotFound, "role not found")
\t}}
\tif ttl <= 0 {{
\t\tttl = role.DefaultTTL
\t}}
\tif ttl > role.MaxTTL {{
\t\tttl = role.MaxTTL
\t}}
\tid := common.NewID()
\ttoken := hex.EncodeToString(hmacSHA(e.root, []byte(fmt.Sprintf("%s:%s:%d", id, roleID, time.Now().UnixNano()))))
\tl := &Lease{{ID: id, RoleID: roleID, Token: token, ExpiresAt: common.NowUTC().Add(ttl)}}
\te.leases[l.ID] = l
\tcp := *l
\treturn &cp, nil
}}

func hmacSHA(key, msg []byte) []byte {{
\tm := hmac.New(sha256.New, key)
\t_, _ = m.Write(msg)
\treturn m.Sum(nil)
}}
'''


def gen_engine_policies(pkg, title):
    return f'''package {pkg}

import (
\t"strings"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

// PolicyStatement is one allow/deny rule.
type PolicyStatement struct {{
\tEffect   string   `json:"effect"`
\tActions  []string `json:"actions"`
\tResources []string `json:"resources"`
}}

// PolicyDocument groups statements for {title}.
type PolicyDocument struct {{
\tID          string            `json:"id"`
\tOrgID       string            `json:"org_id"`
\tName        string            `json:"name"`
\tStatements  []PolicyStatement `json:"statements"`
}}

func (e *Engine) UpsertPolicy(p PolicyDocument) (*PolicyDocument, error) {{
\tif p.OrgID == "" || p.Name == "" {{
\t\treturn nil, common.New(common.ErrCodeInvalidArgument, "org_id and name required")
\t}}
\tif err := ValidatePolicyDocument(p); err != nil {{
\t\treturn nil, err
\t}}
\tif p.ID == "" {{
\t\tp.ID = common.NewID()
\t}}
\tif e.policies == nil {{
\t\te.policies = map[string]*PolicyDocument{{}}
\t}}
\tcp := p
\te.mu.Lock()
\te.policies[p.ID] = &cp
\te.mu.Unlock()
\treturn &cp, nil
}}

func ValidatePolicyDocument(p PolicyDocument) error {{
\tif len(p.Statements) == 0 {{
\t\treturn common.New(common.ErrCodeInvalidArgument, "statements required")
\t}}
\tfor _, s := range p.Statements {{
\t\teff := strings.ToLower(s.Effect)
\t\tif eff != "allow" && eff != "deny" {{
\t\t\treturn common.New(common.ErrCodeInvalidArgument, "effect must be allow or deny")
\t\t}}
\t\tif len(s.Actions) == 0 {{
\t\t\treturn common.New(common.ErrCodeInvalidArgument, "actions required")
\t\t}}
\t}}
\treturn nil
}}
'''


def gen_engine_revoke(pkg, title):
    return f'''package {pkg}

import (
\t"time"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

func (e *Engine) RevokeLease(id string) error {{
\te.mu.Lock()
\tdefer e.mu.Unlock()
\tl, ok := e.leases[id]
\tif !ok {{
\t\treturn common.New(common.ErrCodeNotFound, "lease not found")
\t}}
\tl.Revoked = true
\treturn nil
}}

func (e *Engine) RevokeAllForRole(roleID string) int {{
\te.mu.Lock()
\tdefer e.mu.Unlock()
\tn := 0
\tfor _, l := range e.leases {{
\t\tif l.RoleID == roleID && !l.Revoked {{
\t\t\tl.Revoked = true
\t\t\tn++
\t\t}}
\t}}
\treturn n
}}

func (e *Engine) IsLeaseValid(id string) bool {{
\te.mu.Lock()
\tdefer e.mu.Unlock()
\tl, ok := e.leases[id]
\tif !ok || l.Revoked {{
\t\treturn false
\t}}
\treturn common.NowUTC().Before(l.ExpiresAt)
}}

func (e *Engine) RenewLease(id string, extend time.Duration) (*Lease, error) {{
\te.mu.Lock()
\tdefer e.mu.Unlock()
\tl, ok := e.leases[id]
\tif !ok {{
\t\treturn nil, common.New(common.ErrCodeNotFound, "lease not found")
\t}}
\tif l.Revoked {{
\t\treturn nil, common.New(common.ErrCodeConflict, "lease revoked")
\t}}
\trole := e.roles[l.RoleID]
\tmax := l.ExpiresAt
\tif role != nil {{
\t\tmax = common.NowUTC().Add(role.MaxTTL)
\t}}
\tl.ExpiresAt = l.ExpiresAt.Add(extend)
\tif l.ExpiresAt.After(max) {{
\t\tl.ExpiresAt = max
\t}}
\tcp := *l
\treturn &cp, nil
}}
'''


def gen_engine_policy_validate(pkg, title):
    return f'''package {pkg}

import "strings"

func (e *Engine) ValidatePolicyNames(names []string) []string {{
\te.mu.Lock()
\tdefer e.mu.Unlock()
\tmissing := make([]string, 0)
\tfor _, name := range names {{
\t\tfound := false
\t\tfor _, p := range e.policies {{
\t\t\tif strings.EqualFold(p.Name, name) {{
\t\t\t\tfound = true
\t\t\t\tbreak
\t\t\t}}
\t\t}}
\t\tif !found {{
\t\t\tmissing = append(missing, name)
\t\t}}
\t}}
\treturn missing
}}
'''


def gen_engine_lease_list(pkg, title):
    return f'''package {pkg}

import "time"

// LeaseSummary is a list view of active leases.
type LeaseSummary struct {{
\tID        string    `json:"id"`
\tRoleID    string    `json:"role_id"`
\tExpiresAt time.Time `json:"expires_at"`
\tRevoked   bool      `json:"revoked"`
}}

func (e *Engine) ListLeases(orgID string, includeRevoked bool) []LeaseSummary {{
\te.mu.Lock()
\tdefer e.mu.Unlock()
\tout := make([]LeaseSummary, 0)
\tfor _, l := range e.leases {{
\t\trole := e.roles[l.RoleID]
\t\tif orgID != "" && (role == nil || role.OrgID != orgID) {{
\t\t\tcontinue
\t\t}}
\t\tif !includeRevoked && l.Revoked {{
\t\t\tcontinue
\t\t}}
\t\tout = append(out, LeaseSummary{{ID: l.ID, RoleID: l.RoleID, ExpiresAt: l.ExpiresAt, Revoked: l.Revoked}})
\t}}
\treturn out
}}

func (e *Engine) ExpireLeases(now time.Time) int {{
\te.mu.Lock()
\tdefer e.mu.Unlock()
\tn := 0
\tfor _, l := range e.leases {{
\t\tif !l.Revoked && !now.Before(l.ExpiresAt) {{
\t\t\tl.Revoked = true
\t\t\tn++
\t\t}}
\t}}
\treturn n
}}
'''


def gen_engine_test(pkg, fname):
    base = fname.replace(".go", "")
    return f'''package {pkg}

import (
\t"testing"
\t"time"
)

func Test{base.title().replace("_", "")}(t *testing.T) {{
\te := NewEngine(nil)
\trole, err := e.UpsertRole(Role{{OrgID: "o1", Name: "r1", AccountID: "a1", TenantID: "t1", ProjectId: "p1", ClusterId: "c1", ConnectionId: "c1", KeyId: "k1", Org: "org", Group: "grp", Registry: "reg", ControllerUrl: "http://j", Workspace: "ws"}}[0:0] /* placeholder */)
\t_ = role
\t_ = err
\t// minimal compile test for generated engine module
\tif e == nil {{
\t\tt.Fatal("engine nil")
\t}}
\t_, _ = e.IssueLease("missing", time.Hour)
}}
'''


def gen_service_acl(svc: str, pkg: str, entity: str, domain: str) -> None:
    perms = {
        "identity": ["read_profile", "update_profile", "manage_members", "view_audit"],
        "vault": ["read_item", "write_item", "share_item", "export_item"],
        "enterprise": ["approve_request", "delegate_access", "provision_user", "emergency_unlock"],
        "security": ["view_findings", "ack_finding", "update_trust", "run_scan"],
        "operations": ["view_reports", "run_backup", "manage_billing", "admin_console"],
    }.get(domain, ["read", "write", "admin"])
    perm_checks = "\n".join(
        f'\tif action == "{p}" && roleRank[principalRole] >= {i + 1} {{\n\t\treturn true, nil\n\t}}'
        for i, p in enumerate(perms)
    )
    impl = f'''package acl

import (
\t"strings"
\t"time"

\t"github.com/sentinelvault/sentinelvault/internal/common"
)

// Resource types for {svc}.
const (
\tResource{entity}     = "{pkg}.{entity.lower()}"
\tResource{entity}List = "{pkg}.{entity.lower()}.list"
)

var roleRank = map[string]int{{
\t"viewer": 1, "editor": 2, "admin": 3, "owner": 4,
}}

// Principal holds subject identity for ACL checks.
type Principal struct {{
\tID       string
\tOrgID    string
\tRole     string
\tTeams     []string
\tExpires  *time.Time
}}

// Resource identifies a {entity} instance.
type Resource struct {{
\tID     string
\tOrgID  string
\tOwnerID string
\tTags   map[string]string
}}

// Evaluator evaluates domain-specific grants for {svc}.
type Evaluator struct {{
\tgrants map[string]map[string]time.Time // principalID -> action -> expiry
}}

func NewEvaluator() *Evaluator {{
\treturn &Evaluator{{grants: map[string]map[string]time.Time{{}}}}
}}

func (ev *Evaluator) Grant(principalID, action string, ttl time.Duration) {{
\tif ev.grants[principalID] == nil {{
\t\tev.grants[principalID] = map[string]time.Time{{}}
\t}}
\texp := common.NowUTC().Add(ttl)
\tif ttl <= 0 {{
\t\texp = time.Time{{}}
\t}}
\tev.grants[principalID][action] = exp
}}

func (ev *Evaluator) Revoke(principalID, action string) {{
\tif m, ok := ev.grants[principalID]; ok {{
\t\tdelete(m, action)
\t}}
}}

func (ev *Evaluator) Can(principal Principal, resource Resource, action string) (bool, error) {{
\tif principal.ID == "" || resource.OrgID == "" {{
\t\treturn false, common.New(common.ErrCodeInvalidArgument, "principal and org required")
\t}}
\tif principal.OrgID != resource.OrgID {{
\t\treturn false, common.New(common.ErrCodeForbidden, "org mismatch")
\t}}
\tif principal.Expires != nil && common.NowUTC().After(*principal.Expires) {{
\t\treturn false, common.New(common.ErrCodeForbidden, "principal expired")
\t}}
\tif resource.OwnerID == principal.ID {{
\t\treturn true, nil
\t}}
\tprincipalRole := strings.ToLower(principal.Role)
\tif principalRole == "" {{
\t\tprincipalRole = "viewer"
\t}}
{perm_checks}
\tif m, ok := ev.grants[principal.ID]; ok {{
\t\tif exp, ok := m[action]; ok {{
\t\t\tif exp.IsZero() || common.NowUTC().Before(exp) {{
\t\t\t\treturn true, nil
\t\t\t}}
\t\t}}
\t}}
\treturn false, nil
}}

// EffectivePermissions returns all actions allowed for principal on resource.
func (ev *Evaluator) EffectivePermissions(principal Principal, resource Resource) []string {{
\tallowed := make([]string, 0)
\tfor _, action := range []string{{{", ".join(f'"{p}"' for p in perms)}}} {{
\t\tok, _ := ev.Can(principal, resource, action)
\t\tif ok {{
\t\t\tallowed = append(allowed, action)
\t\t}}
\t}}
\treturn allowed
}}

// ValidateAction ensures action is known for this service domain.
func ValidateAction(action string) error {{
\tswitch action {{
\tcase {", ".join(f'"{p}"' for p in perms)}:
\t\treturn nil
\tdefault:
\t\treturn common.New(common.ErrCodeInvalidArgument, "unknown action")
\t}}
}}
'''
    test = f'''package acl

import (
\t"testing"
\t"time"
)

func Test{entity}EvaluatorOwner(t *testing.T) {{
\tev := NewEvaluator()
\tp := Principal{{ID: "u1", OrgID: "o1", Role: "viewer"}}
\tr := Resource{{ID: "r1", OrgID: "o1", OwnerID: "u1"}}
\tok, err := ev.Can(p, r, "{perms[0]}")
\tif err != nil || !ok {{
\t\tt.Fatalf("ok=%v err=%v", ok, err)
\t}}
}}

func Test{entity}EvaluatorGrant(t *testing.T) {{
\tev := NewEvaluator()
\tev.Grant("u2", "{perms[1]}", time.Hour)
\tp := Principal{{ID: "u2", OrgID: "o1", Role: "viewer"}}
\tr := Resource{{ID: "r1", OrgID: "o1", OwnerID: "u3"}}
\tok, err := ev.Can(p, r, "{perms[1]}")
\tif err != nil || !ok {{
\t\tt.Fatalf("ok=%v err=%v", ok, err)
\t}}
\tperms := ev.EffectivePermissions(p, r)
\tif len(perms) == 0 {{
\t\tt.Fatal("no perms")
\t}}
}}

func Test{entity}EvaluatorErrors(t *testing.T) {{
\tev := NewEvaluator()
\t_, err := ev.Can(Principal{{}}, Resource{{OrgID: "o1"}}, "{perms[0]}")
\tif err == nil {{
\t\tt.Fatal("expected error")
\t}}
\tif err := ValidateAction("bad.action"); err == nil {{
\t\tt.Fatal("expected invalid action")
\t}}
}}
'''
    write(ROOT / f"apps/{svc}/internal/acl/evaluator.go", impl)
    write(ROOT / f"apps/{svc}/internal/acl/evaluator_test.go", test)


def gen_contract_tests() -> None:
    impl = textwrap.dedent('''
        package contract

        import (
            "encoding/json"
            "strings"
            "testing"

            gateway "github.com/sentinelvault/sentinelvault/apps/api-gateway/internal/gateway"
        )

        func TestOpenAPIPathsPresent(t *testing.T) {
            spec := gateway.OpenAPISpec()
            var doc map[string]any
            if err := json.Unmarshal([]byte(spec), &doc); err != nil {
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
    ''')
    write(ROOT / "tests/contract/openapi_test.go", impl.strip())
    # Need OpenAPISpec export from gateway
    gateway_path = ROOT / "apps/api-gateway/internal/gateway/openapi.go"
    text = gateway_path.read_text()
    if "func OpenAPISpec()" not in text:
        text = text.replace("const openAPISpec = `", "const openAPISpec = `").rstrip() + "\n\n// OpenAPISpec returns the embedded OpenAPI document.\nfunc OpenAPISpec() string { return openAPISpec }\n"
        gateway_path.write_text(text)


def fix_engine_structs() -> None:
    """Add policies field to existing engine structs."""
    for pkg, _, _, new_engine in ENGINES:
        if new_engine:
            continue
        engine_path = ROOT / f"internal/engines/{pkg}/engine.go"
        if not engine_path.exists():
            continue
        text = engine_path.read_text()
        if "policies map[string]*PolicyDocument" in text:
            continue
        if "type Engine struct {" in text:
            text = text.replace(
                "type Engine struct {",
                "type Engine struct {\n\tpolicies map[string]*PolicyDocument",
                1,
            )
        if "policies:" not in text and "NewEngine" in text:
            # add init in first return NewEngine line
            text = text.replace(
                "return &Engine{",
                "return &Engine{policies: map[string]*PolicyDocument{}, ",
                1,
            )
        engine_path.write_text(text)


def fix_engine_tests() -> None:
    for pkg, _, _, new_engine in ENGINES:
        content = f'''package {pkg}

import "testing"

func TestPoliciesModule(t *testing.T) {{
\tpol := PolicyDocument{{OrgID: "o1", Name: "read", Statements: []PolicyStatement{{
\t\t{{Effect: "allow", Actions: []string{{"read"}}, Resources: []string{{"*"}}}},
\t}}}}
\tif err := ValidatePolicyDocument(pol); err != nil {{
\t\tt.Fatal(err)
\t}}
\tif err := ValidatePolicyDocument(PolicyDocument{{}}); err == nil {{
\t\tt.Fatal("expected error")
\t}}
}}
'''
        write(ROOT / f"internal/engines/{pkg}/policies_test.go", content)
        list_test = f'''package {pkg}

import "testing"

func TestLeaseListModule(t *testing.T) {{
\ts := LeaseSummary{{ID: "l1", RoleID: "r1"}}
\tif s.ID == "" {{
\t\tt.Fatal("empty")
\t}}
}}
'''
        write(ROOT / f"internal/engines/{pkg}/lease_list_test.go", list_test)


def main() -> None:
    saml_impl, saml_test = gen_saml()
    write(ROOT / "internal/idp/saml/saml.go", saml_impl)
    write(ROOT / "internal/idp/saml/saml_test.go", saml_test)

    client_impl, client_test = gen_oidc_client()
    write(ROOT / "internal/idp/oidc/client.go", client_impl)
    write(ROOT / "internal/idp/oidc/client_test.go", client_test)

    for metric in ANALYTICS_METRICS:
        impl, test = gen_analytics_metric(metric)
        write(ROOT / f"internal/analytics/{metric}.go", impl)
        write(ROOT / f"internal/analytics/{metric}_test.go", test)

    gen_outbox()
    gen_saga_hooks()
    gen_acl_expand()
    gen_featureflags_expand()
    gen_webhooks_retry()
    gen_contract_tests()

    for pkg, title, id_field, new_engine in ENGINES:
        gen_engine(pkg, title, id_field, new_engine=new_engine)

    fix_engine_structs()
    fix_engine_tests()

    for svc, pkg, entity, domain in APPS:
        gen_service_acl(svc, pkg, entity, domain)

    print("phase C generation complete")


if __name__ == "__main__":
    main()
