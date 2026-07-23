package domain

import (
	"context"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

// SCIMName is the SCIM name complex type subset.
type SCIMName struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	Formatted  string `json:"formatted,omitempty"`
}

// SCIMEmail is a SCIM email value.
type SCIMEmail struct {
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
	Type    string `json:"type,omitempty"`
}

// SCIMUser is a SCIM User resource subset.
type SCIMUser struct {
	ID         string      `json:"id"`
	OrgID      string      `json:"org_id"`
	ExternalID string      `json:"externalId,omitempty"`
	UserName   string      `json:"userName"`
	Active     bool        `json:"active"`
	Name       SCIMName    `json:"name"`
	Emails     []SCIMEmail `json:"emails"`
	Schemas    []string    `json:"schemas"`
	Meta       SCIMMeta    `json:"meta"`
}

// SCIMMeta holds SCIM resource metadata.
type SCIMMeta struct {
	ResourceType string    `json:"resourceType"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
}

// CreateUserInput holds SCIM user creation fields.
type CreateUserInput struct {
	OrgID      string      `json:"org_id"`
	ExternalID string      `json:"externalId"`
	UserName   string      `json:"userName"`
	Name       SCIMName    `json:"name"`
	Emails     []SCIMEmail `json:"emails"`
}

// Service manages SCIM User resources.
type Service struct {
	mu    sync.RWMutex
	users map[string]*SCIMUser
}

// NewService creates a SCIM domain service.
func NewService() *Service {
	return &Service{users: map[string]*SCIMUser{}}
}

// CreateUser provisions a SCIM user.
func (s *Service) CreateUser(_ context.Context, in CreateUserInput) (*SCIMUser, error) {
	if in.OrgID == "" || in.UserName == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and userName required")
	}
	now := common.NowUTC()
	user := &SCIMUser{
		ID: common.NewID(), OrgID: in.OrgID, ExternalID: in.ExternalID,
		UserName: in.UserName, Active: true, Name: in.Name, Emails: in.Emails,
		Schemas: []string{"urn:ietf:params:scim:schemas:core:2.0:User"},
		Meta:    SCIMMeta{ResourceType: "User", Created: now, LastModified: now},
	}
	s.mu.Lock()
	s.users[user.ID] = user
	s.mu.Unlock()
	cp := *user
	return &cp, nil
}

// GetUser returns a SCIM user by ID.
func (s *Service) GetUser(_ context.Context, id string) (*SCIMUser, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, ok := s.users[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "user not found")
	}
	cp := *user
	return &cp, nil
}

// DeactivateUser sets active=false.
func (s *Service) DeactivateUser(_ context.Context, id string) (*SCIMUser, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, ok := s.users[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "user not found")
	}
	user.Active = false
	user.Meta.LastModified = common.NowUTC()
	cp := *user
	return &cp, nil
}

// ListUsers returns users for an org.
func (s *Service) ListUsers(_ context.Context, orgID string) []SCIMUser {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]SCIMUser, 0)
	for _, u := range s.users {
		if orgID == "" || u.OrgID == orgID {
			out = append(out, *u)
		}
	}
	return out
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "scim" }

// TableName returns the primary table name.
func TableName() string { return "users" }
