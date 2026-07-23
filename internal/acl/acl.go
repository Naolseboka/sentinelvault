package acl

import (
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

const (
	PermOwner  = "owner"
	PermAdmin  = "admin"
	PermEditor = "editor"
	PermViewer = "viewer"
)

var permRank = map[string]int{PermViewer: 1, PermEditor: 2, PermAdmin: 3, PermOwner: 4}

// Entry is an ACL grant on a resource.
type Entry struct {
	ID            string
	OrgID         string
	ResourceType  string
	ResourceID    string
	PrincipalID   string
	Permission    string
	ExpiresAt     *time.Time
	InheritedFrom string
	CreatedAt     time.Time
}

// Store holds ACL entries.
type Store struct {
	entries map[string]*Entry
}

func NewStore() *Store { return &Store{entries: map[string]*Entry{}} }

func (s *Store) Grant(orgID, resourceType, resourceID, principalID, perm string, ttl time.Duration) (*Entry, error) {
	if orgID == "" || resourceID == "" || principalID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, resource_id, principal_id required")
	}
	if permRank[perm] == 0 {
		return nil, common.New(common.ErrCodeInvalidArgument, "invalid permission")
	}
	now := common.NowUTC()
	e := &Entry{ID: common.NewID(), OrgID: orgID, ResourceType: resourceType, ResourceID: resourceID,
		PrincipalID: principalID, Permission: perm, CreatedAt: now}
	if ttl > 0 {
		exp := now.Add(ttl)
		e.ExpiresAt = &exp
	}
	s.entries[e.ID] = e
	cp := *e
	return &cp, nil
}

func (s *Store) Revoke(id string) error {
	if _, ok := s.entries[id]; !ok {
		return common.New(common.ErrCodeNotFound, "acl entry not found")
	}
	delete(s.entries, id)
	return nil
}

func (s *Store) Effective(orgID, resourceType, resourceID, principalID string, parentIDs []string) (string, bool) {
	now := common.NowUTC()
	best := 0
	bestPerm := ""
	consider := func(e *Entry) {
		if e.OrgID != orgID || e.PrincipalID != principalID {
			return
		}
		if e.ExpiresAt != nil && now.After(*e.ExpiresAt) {
			return
		}
		if e.ResourceType == resourceType && e.ResourceID == resourceID {
			if permRank[e.Permission] > best {
				best = permRank[e.Permission]
				bestPerm = e.Permission
			}
		}
		for _, pid := range parentIDs {
			if e.ResourceID == pid && permRank[e.Permission] > best {
				best = permRank[e.Permission]
				bestPerm = e.Permission
				// inherited viewer downgrades editor? keep same for simplicity
			}
		}
	}
	for _, e := range s.entries {
		consider(e)
	}
	return bestPerm, best > 0
}

func (s *Store) Can(orgID, resourceType, resourceID, principalID, need string, parentIDs []string) bool {
	have, ok := s.Effective(orgID, resourceType, resourceID, principalID, parentIDs)
	if !ok {
		return false
	}
	return permRank[have] >= permRank[need]
}

func NormalizePerm(p string) string { return strings.ToLower(strings.TrimSpace(p)) }
