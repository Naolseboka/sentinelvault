package domain

import (
	"context"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

const (
	JobStatusPending   = "pending"
	JobStatusRunning   = "running"
	JobStatusCompleted = "completed"
	JobStatusFailed    = "failed"

	ConflictDuplicateEmail = "duplicate_email"
	ConflictMissingField   = "missing_required"
	ConflictRoleMismatch   = "role_mismatch"
)

// LDAPConfig holds directory sync connection settings (no live LDAP required).
type LDAPConfig struct {
	OrgID    string `json:"org_id"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	BaseDN   string `json:"base_dn"`
	BindDN   string `json:"bind_dn"`
	Filter   string `json:"filter"`
	Enabled  bool   `json:"enabled"`
}

// SyncConflict describes a directory sync conflict.
type SyncConflict struct {
	ExternalID   string `json:"external_id"`
	Email        string `json:"email"`
	ConflictType string `json:"conflict_type"`
	Resolution   string `json:"resolution"`
}

// ConflictReport aggregates sync conflicts for a job.
type ConflictReport struct {
	JobID     string         `json:"job_id"`
	OrgID     string         `json:"org_id"`
	Conflicts []SyncConflict `json:"conflicts"`
	Total     int            `json:"total"`
}

// SyncJob represents a directory synchronization run.
type SyncJob struct {
	ID           string         `json:"id"`
	OrgID        string         `json:"org_id"`
	Connector    string         `json:"connector"` // ldap|azure_ad|okta|google
	Config       LDAPConfig     `json:"config"`
	Status       string         `json:"status"`
	UsersCreated int            `json:"users_created"`
	UsersUpdated int            `json:"users_updated"`
	UsersRemoved int            `json:"users_removed"`
	Conflicts    []SyncConflict `json:"conflicts,omitempty"`
	Error        string         `json:"error,omitempty"`
	StartedAt    time.Time      `json:"started_at"`
	FinishedAt   *time.Time     `json:"finished_at,omitempty"`
}

// ConfigureInput upserts LDAP sync configuration.
type ConfigureInput struct {
	OrgID   string `json:"org_id"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	BaseDN  string `json:"base_dn"`
	BindDN  string `json:"bind_dn"`
	Filter  string `json:"filter"`
	Enabled bool   `json:"enabled"`
}

// RunSyncInput starts a sync job using stored or inline config.
type RunSyncInput struct {
	OrgID     string `json:"org_id"`
	Connector string `json:"connector"`
}

// DirectoryEntry is a simulated LDAP user for in-memory sync.
type DirectoryEntry struct {
	ExternalID string
	Email      string
	DisplayName string
	Active     bool
}

// Service manages directory sync jobs and LDAP configuration.
type Service struct {
	mu      sync.RWMutex
	jobs    map[string]*SyncJob
	configs map[string]*LDAPConfig // orgID -> config
	entries map[string][]DirectoryEntry // orgID -> simulated directory
}

// NewService creates a directory sync domain service.
func NewService() *Service {
	return &Service{
		jobs:    map[string]*SyncJob{},
		configs: map[string]*LDAPConfig{},
		entries: map[string][]DirectoryEntry{},
	}
}

// Configure stores LDAP sync settings for an org.
func (s *Service) Configure(_ context.Context, in ConfigureInput) (*LDAPConfig, error) {
	if in.OrgID == "" || in.Host == "" || in.BaseDN == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id, host, and base_dn required")
	}
	port := in.Port
	if port == 0 {
		port = 389
	}
	filter := in.Filter
	if filter == "" {
		filter = "(objectClass=person)"
	}
	cfg := &LDAPConfig{
		OrgID: in.OrgID, Host: in.Host, Port: port, BaseDN: in.BaseDN,
		BindDN: in.BindDN, Filter: filter, Enabled: in.Enabled,
	}
	s.mu.Lock()
	s.configs[in.OrgID] = cfg
	s.mu.Unlock()
	cp := *cfg
	return &cp, nil
}

// SeedDirectory loads simulated LDAP entries for testing (no real LDAP).
func (s *Service) SeedDirectory(orgID string, entries []DirectoryEntry) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entries[orgID] = entries
}

// RunSync executes a directory sync job against simulated entries.
func (s *Service) RunSync(_ context.Context, in RunSyncInput) (*SyncJob, error) {
	if in.OrgID == "" || in.Connector == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and connector required")
	}
	s.mu.Lock()
	cfg, ok := s.configs[in.OrgID]
	if !ok {
		cfg = &LDAPConfig{OrgID: in.OrgID, Host: "ldap.internal", Port: 389, BaseDN: "dc=example,dc=com", Filter: "(objectClass=person)", Enabled: true}
	}
	entries := s.entries[in.OrgID]
	if len(entries) == 0 {
		entries = []DirectoryEntry{
			{ExternalID: "ldap-1", Email: "alice@example.com", DisplayName: "Alice", Active: true},
			{ExternalID: "ldap-2", Email: "bob@example.com", DisplayName: "Bob", Active: true},
			{ExternalID: "ldap-3", Email: "alice@example.com", DisplayName: "Alice Duplicate", Active: false},
		}
	}

	conflicts := detectConflicts(entries)
	created, updated := 0, 0
	seenEmails := map[string]bool{}
	for _, e := range entries {
		if e.Email == "" {
			conflicts = append(conflicts, SyncConflict{
				ExternalID: e.ExternalID, ConflictType: ConflictMissingField, Resolution: "skip",
			})
			continue
		}
		if seenEmails[e.Email] {
			continue
		}
		seenEmails[e.Email] = true
		if e.Active {
			created++
		} else {
			updated++
		}
	}

	now := common.NowUTC()
	job := &SyncJob{
		ID: common.NewID(), OrgID: in.OrgID, Connector: in.Connector,
		Config: *cfg, Status: JobStatusRunning, StartedAt: now,
		UsersCreated: created, UsersUpdated: updated, Conflicts: conflicts,
	}
	fin := common.NowUTC()
	job.FinishedAt = &fin
	job.Status = JobStatusCompleted
	s.jobs[job.ID] = job
	s.mu.Unlock()
	cp := *job
	return &cp, nil
}

func detectConflicts(entries []DirectoryEntry) []SyncConflict {
	conflicts := make([]SyncConflict, 0)
	seen := map[string]string{}
	for _, e := range entries {
		if e.Email == "" {
			continue
		}
		if prev, ok := seen[e.Email]; ok {
			conflicts = append(conflicts, SyncConflict{
				ExternalID: e.ExternalID, Email: e.Email,
				ConflictType: ConflictDuplicateEmail,
				Resolution:   "keep_" + prev,
			})
			continue
		}
		seen[e.Email] = e.ExternalID
		if !e.Active {
			conflicts = append(conflicts, SyncConflict{
				ExternalID: e.ExternalID, Email: e.Email,
				ConflictType: ConflictRoleMismatch, Resolution: "deactivate",
			})
		}
	}
	return conflicts
}

// GetConflictReport returns conflicts for a completed sync job.
func (s *Service) GetConflictReport(_ context.Context, jobID string) (*ConflictReport, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	job, ok := s.jobs[jobID]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "sync job not found")
	}
	return &ConflictReport{
		JobID: jobID, OrgID: job.OrgID,
		Conflicts: job.Conflicts, Total: len(job.Conflicts),
	}, nil
}

// Get returns a sync job by ID.
func (s *Service) Get(_ context.Context, id string) (*SyncJob, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	job, ok := s.jobs[id]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "sync job not found")
	}
	cp := *job
	return &cp, nil
}

// List returns sync jobs for an org.
func (s *Service) List(_ context.Context, orgID string) []SyncJob {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]SyncJob, 0)
	for _, j := range s.jobs {
		if orgID == "" || j.OrgID == orgID {
			out = append(out, *j)
		}
	}
	return out
}

// Start is an alias for RunSync.
func (s *Service) Start(ctx context.Context, in RunSyncInput) (*SyncJob, error) {
	return s.RunSync(ctx, in)
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "directory_sync" }

// TableName returns the primary table name.
func TableName() string { return "sync_jobs" }
