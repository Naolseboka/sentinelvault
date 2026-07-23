package otp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Client talks to otp-service over HTTP.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

// New creates a otp-service client.
func New(baseURL, token string) *Client {
	return &Client{
		BaseURL:    stringsTrimSlash(baseURL),
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		Token:      token,
	}
}

func stringsTrimSlash(s string) string {
	for len(s) > 0 && s[len(s)-1] == '/' {
		s = s[:len(s)-1]
	}
	return s
}

// OTPRecord is the API resource.
type OTPRecord struct {
	ID        string            `json:"id"`
	OrgID     string            `json:"org_id"`
	Name      string            `json:"name"`
	Status    string            `json:"status"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// ListResult wraps paginated data.
type ListResult struct {
	Data   []OTPRecord `json:"data"`
	Total  int64       `json:"total"`
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
}

func (c *Client) do(ctx context.Context, method, path string, in any, out any) error {
	var body *bytes.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewReader(b)
	} else {
		body = bytes.NewReader(nil)
	}
	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, body)
	if err != nil {
		return err
	}
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		return fmt.Errorf("otp: unexpected status %d", res.StatusCode)
	}
	if out == nil {
		return nil
	}
	return json.NewDecoder(res.Body).Decode(out)
}

// Create creates a OTPRecord.
func (c *Client) Create(ctx context.Context, orgID, name string, metadata map[string]string) (*OTPRecord, error) {
	var out OTPRecord
	err := c.do(ctx, http.MethodPost, "/v1/otps", map[string]any{
		"org_id": orgID, "name": name, "metadata": metadata,
	}, &out)
	return &out, err
}

// Get fetches a OTPRecord by id.
func (c *Client) Get(ctx context.Context, id string) (*OTPRecord, error) {
	var out OTPRecord
	err := c.do(ctx, http.MethodGet, "/v1/otps/"+url.PathEscape(id), nil, &out)
	return &out, err
}

// List lists OTPRecord records for an org.
func (c *Client) List(ctx context.Context, orgID string, offset, limit int) (*ListResult, error) {
	q := url.Values{}
	q.Set("org_id", orgID)
	q.Set("offset", strconv.Itoa(offset))
	q.Set("limit", strconv.Itoa(limit))
	var out ListResult
	err := c.do(ctx, http.MethodGet, "/v1/otps?"+q.Encode(), nil, &out)
	return &out, err
}

// Delete deletes a OTPRecord.
func (c *Client) Delete(ctx context.Context, id string) error {
	return c.do(ctx, http.MethodDelete, "/v1/otps/"+url.PathEscape(id), nil, nil)
}

// maintenance note 83
