package scimx

import (
	"strings"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// User is a SCIM 2.0 User resource subset.
type User struct {
	ID          string   `json:"id"`
	Schemas     []string `json:"schemas"`
	UserName    string   `json:"userName"`
	DisplayName string   `json:"displayName"`
	Active      bool     `json:"active"`
	Emails      []Email  `json:"emails"`
	OrgID       string   `json:"org_id"`
	ExternalID  string   `json:"externalId,omitempty"`
	Meta        Meta     `json:"meta"`
}

// Email is a SCIM email.
type Email struct {
	Value   string `json:"value"`
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
}

// Meta is SCIM meta.
type Meta struct {
	ResourceType string    `json:"resourceType"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
}

// ListResponse is a SCIM ListResponse.
type ListResponse struct {
	Schemas      []string `json:"schemas"`
	TotalResults int      `json:"totalResults"`
	StartIndex   int      `json:"startIndex"`
	ItemsPerPage int      `json:"itemsPerPage"`
	Resources    []User   `json:"Resources"`
}

// Directory stores SCIM users.
type Directory struct {
	mu    sync.RWMutex
	users map[string]*User
}

// NewDirectory creates an empty SCIM directory.
func NewDirectory() *Directory {
	return &Directory{users: map[string]*User{}}
}

// CreateUser creates a SCIM user.
func (d *Directory) CreateUser(orgID, userName, displayName, email string) (*User, error) {
	if orgID == "" || userName == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and userName required")
	}
	now := common.NowUTC()
	u := &User{
		ID:       common.NewID(),
		Schemas:  []string{"urn:ietf:params:scim:schemas:core:2.0:User"},
		UserName: userName, DisplayName: displayName, Active: true, OrgID: orgID,
		Emails: []Email{{Value: email, Type: "work", Primary: true}},
		Meta:   Meta{ResourceType: "User", Created: now, LastModified: now},
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, existing := range d.users {
		if existing.OrgID == orgID && strings.EqualFold(existing.UserName, userName) {
			return nil, common.New(common.ErrCodeConflict, "userName already exists")
		}
	}
	d.users[u.ID] = u
	cp := *u
	return &cp, nil
}

// GetUser returns a user by ID.
func (d *Directory) GetUser(id string) (*User, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	u, ok := d.users[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "user not found")
	}
	cp := *u
	return &cp, nil
}

// PatchActive sets active flag.
func (d *Directory) PatchActive(id string, active bool) (*User, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	u, ok := d.users[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "user not found")
	}
	u.Active = active
	u.Meta.LastModified = common.NowUTC()
	cp := *u
	return &cp, nil
}

// ListUsers returns paginated users for an org, optional filter on userName.
func (d *Directory) ListUsers(orgID, filter string, start, count int) ListResponse {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if start < 1 {
		start = 1
	}
	if count <= 0 {
		count = 100
	}
	all := make([]User, 0)
	for _, u := range d.users {
		if orgID != "" && u.OrgID != orgID {
			continue
		}
		if filter != "" && !strings.Contains(strings.ToLower(u.UserName), strings.ToLower(filter)) {
			continue
		}
		all = append(all, *u)
	}
	total := len(all)
	from := start - 1
	if from > total {
		from = total
	}
	to := from + count
	if to > total {
		to = total
	}
	return ListResponse{
		Schemas:      []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"},
		TotalResults: total, StartIndex: start, ItemsPerPage: to - from, Resources: all[from:to],
	}
}
