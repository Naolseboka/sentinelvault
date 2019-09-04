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