package threat

import (
	"strings"
	"unicode"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Finding is a detected security issue.
type Finding struct {
	ID       string `json:"id"`
	OrgID    string `json:"org_id"`
	Type     string `json:"type"`
	Severity string `json:"severity"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Resource string `json:"resource"`
}

// Analyzer inspects credentials and access patterns.
type Analyzer struct{}

// NewAnalyzer creates an analyzer.
func NewAnalyzer() *Analyzer { return &Analyzer{} }

// AnalyzePassword checks strength and common weaknesses.
func (a *Analyzer) AnalyzePassword(orgID, resourceID, password string) []Finding {
	out := make([]Finding, 0)
	if len(password) < 12 {
		out = append(out, finding(orgID, resourceID, "weak_password", "high", "Password shorter than 12 characters"))
	}
	if isCommonPassword(password) {
		out = append(out, finding(orgID, resourceID, "common_password", "critical", "Password matches a known common password"))
	}
	if !hasComplexity(password) {
		out = append(out, finding(orgID, resourceID, "low_complexity", "medium", "Password lacks character class diversity"))
	}
	return out
}

// DetectReuse finds passwords reused across items.
func (a *Analyzer) DetectReuse(orgID string, itemPasswords map[string]string) []Finding {
	byHash := map[string][]string{}
	for id, pw := range itemPasswords {
		byHash[pw] = append(byHash[pw], id)
	}
	out := make([]Finding, 0)
	for _, ids := range byHash {
		if len(ids) > 1 {
			out = append(out, Finding{
				ID: common.NewID(), OrgID: orgID, Type: "password_reuse", Severity: "high",
				Title: "Password reused across vault items", Detail: strings.Join(ids, ","), Resource: ids[0],
			})
		}
	}
	return out
}

// DarkWebChecker is an interface for breach lookups.
type DarkWebChecker interface {
	IsExposed(passwordOrHash string) (bool, string, error)
}

// MemoryDarkWeb is a stub breach DB.
type MemoryDarkWeb struct {
	Exposed map[string]string
}

func (m MemoryDarkWeb) IsExposed(passwordOrHash string) (bool, string, error) {
	src, ok := m.Exposed[passwordOrHash]
	return ok, src, nil
}

// CheckExposure adds findings for exposed credentials.
func (a *Analyzer) CheckExposure(orgID, resourceID, secret string, db DarkWebChecker) ([]Finding, error) {
	exposed, src, err := db.IsExposed(secret)
	if err != nil {
		return nil, err
	}
	if !exposed {
		return nil, nil
	}
	return []Finding{{
		ID: common.NewID(), OrgID: orgID, Type: "credential_exposure", Severity: "critical",
		Title: "Credential found in breach corpus", Detail: src, Resource: resourceID,
	}}, nil
}

func finding(orgID, resource, typ, sev, title string) Finding {
	return Finding{
		ID: common.NewID(), OrgID: orgID, Type: typ, Severity: sev, Title: title, Resource: resource,
	}
}

func hasComplexity(pw string) bool {
	var lower, upper, digit, special bool
	for _, r := range pw {
		switch {
		case unicode.IsLower(r):
			lower = true
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsDigit(r):
			digit = true
		default:
			special = true
		}
	}
	n := 0
	for _, b := range []bool{lower, upper, digit, special} {
		if b {
			n++
		}
	}
	return n >= 3
}

func isCommonPassword(pw string) bool {
	commonPw := map[string]struct{}{
		"password": {}, "password123": {}, "123456": {}, "qwerty": {}, "letmein": {},
		"admin": {}, "welcome": {}, "iloveyou": {}, "monkey": {}, "dragon": {},
	}
	_, ok := commonPw[strings.ToLower(pw)]
	return ok
}
// maintenance note 17
