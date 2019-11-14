package urlx

import (
	"net/url"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

func ParseAbsolute(raw string) (*url.URL, error) {
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "url", err)
	}
	if u.Scheme != "https" && u.Scheme != "http" {
		return nil, common.New(common.ErrCodeInvalidArgument, "url scheme must be http(s)")
	}
	if u.Host == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "url host required")
	}
	return u, nil
}

func HostAllowed(raw string, allow []string) bool {
	u, err := ParseAbsolute(raw)
	if err != nil {
		return false
	}
	host := strings.ToLower(u.Hostname())
	for _, a := range allow {
		a = strings.ToLower(strings.TrimSpace(a))
		if a == "" {
			continue
		}
		if host == a || strings.HasSuffix(host, "."+a) {
			return true
		}
	}
	return false
}

func JoinPath(base string, parts ...string) (string, error) {
	u, err := ParseAbsolute(base)
	if err != nil {
		return "", err
	}
	segs := make([]string, 0, len(parts)+1)
	if u.Path != "" {
		segs = append(segs, strings.Trim(u.Path, "/"))
	}
	for _, p := range parts {
		p = strings.Trim(p, "/")
		if p != "" {
			segs = append(segs, p)
		}
	}
	u.Path = "/" + strings.Join(segs, "/")
	return u.String(), nil
}
