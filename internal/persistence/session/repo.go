package session

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sentinelvault/sentinelvault/internal/common"
)

// SessionRecord is the persistence model for session-service.
type SessionRecord struct {
	ID        string          `json:"id"`
	OrgID     string          `json:"org_id"`
	Name      string          `json:"name"`
	Status    string          `json:"status"`
	Payload   json.RawMessage `json:"payload"`
	Version   int             `json:"version"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *time.Time      `json:"deleted_at,omitempty"`
}

// SessionRepository provides Postgres persistence for session-service.
type SessionRepository struct {
	pool   *pgxpool.Pool
	schema string
	table  string
}

// NewSessionRepository constructs a repository bound to schema session.
func NewSessionRepository(pool *pgxpool.Pool) *SessionRepository {
	return &SessionRepository{pool: pool, schema: "session", table: "session_records"}
}

func (r *SessionRepository) fq() string {
	return pgx.Identifier{r.schema, r.table}.Sanitize()
}

// EnsureSchema creates schema/table if missing (dev convenience).
func (r *SessionRepository) EnsureSchema(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS %s`, r.schema))
	if err != nil {
		return err
	}
	_, err = r.pool.Exec(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
  id TEXT PRIMARY KEY,
  org_id TEXT NOT NULL,
  name TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'active',
  payload JSONB NOT NULL DEFAULT '{}',
  version INT NOT NULL DEFAULT 1,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ
)`, r.fq()))
	if err != nil {
		return err
	}
	_, err = r.pool.Exec(ctx, fmt.Sprintf(`CREATE INDEX IF NOT EXISTS %s_org_idx ON %s(org_id) WHERE deleted_at IS NULL`, r.table, r.fq()))
	return err
}

// Insert stores a new record.
func (r *SessionRepository) Insert(ctx context.Context, rec *SessionRecord) error {
	if rec.ID == "" {
		rec.ID = common.NewID()
	}
	now := common.NowUTC()
	if rec.CreatedAt.IsZero() {
		rec.CreatedAt = now
	}
	rec.UpdatedAt = now
	if rec.Version == 0 {
		rec.Version = 1
	}
	if rec.Status == "" {
		rec.Status = "active"
	}
	if rec.Payload == nil {
		rec.Payload = json.RawMessage(`{}`)
	}
	_, err := r.pool.Exec(ctx, fmt.Sprintf(`
INSERT INTO %s (id, org_id, name, status, payload, version, created_at, updated_at)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`, r.fq()),
		rec.ID, rec.OrgID, rec.Name, rec.Status, rec.Payload, rec.Version, rec.CreatedAt, rec.UpdatedAt)
	return err
}

// Update mutates mutable fields and bumps version.
func (r *SessionRepository) Update(ctx context.Context, rec *SessionRecord) error {
	rec.UpdatedAt = common.NowUTC()
	rec.Version++
	tag, err := r.pool.Exec(ctx, fmt.Sprintf(`
UPDATE %s SET name=$2, status=$3, payload=$4, version=$5, updated_at=$6
WHERE id=$1 AND deleted_at IS NULL`, r.fq()),
		rec.ID, rec.Name, rec.Status, rec.Payload, rec.Version, rec.UpdatedAt)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return common.New(common.ErrCodeNotFound, "Session not found")
	}
	return nil
}

// SoftDelete marks deleted_at.
func (r *SessionRepository) SoftDelete(ctx context.Context, id string) error {
	now := common.NowUTC()
	tag, err := r.pool.Exec(ctx, fmt.Sprintf(`
UPDATE %s SET deleted_at=$2, updated_at=$2, status='deleted' WHERE id=$1 AND deleted_at IS NULL`, r.fq()), id, now)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return common.New(common.ErrCodeNotFound, "Session not found")
	}
	return nil
}

// Get fetches by ID.
func (r *SessionRepository) Get(ctx context.Context, id string) (*SessionRecord, error) {
	row := r.pool.QueryRow(ctx, fmt.Sprintf(`
SELECT id, org_id, name, status, payload, version, created_at, updated_at, deleted_at
FROM %s WHERE id=$1`, r.fq()), id)
	var rec SessionRecord
	if err := row.Scan(&rec.ID, &rec.OrgID, &rec.Name, &rec.Status, &rec.Payload, &rec.Version, &rec.CreatedAt, &rec.UpdatedAt, &rec.DeletedAt); err != nil {
		if err == pgx.ErrNoRows {
			return nil, common.New(common.ErrCodeNotFound, "Session not found")
		}
		return nil, err
	}
	return &rec, nil
}

// ListByOrg returns paginated active records.
func (r *SessionRepository) ListByOrg(ctx context.Context, orgID string, offset, limit int) ([]SessionRecord, int64, error) {
	if limit <= 0 {
		limit = 50
	}
	var total int64
	if err := r.pool.QueryRow(ctx, fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE org_id=$1 AND deleted_at IS NULL`, r.fq()), orgID).Scan(&total); err != nil {
		return nil, 0, err
	}
	rows, err := r.pool.Query(ctx, fmt.Sprintf(`
SELECT id, org_id, name, status, payload, version, created_at, updated_at, deleted_at
FROM %s WHERE org_id=$1 AND deleted_at IS NULL
ORDER BY created_at DESC OFFSET $2 LIMIT $3`, r.fq()), orgID, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	out := make([]SessionRecord, 0, limit)
	for rows.Next() {
		var rec SessionRecord
		if err := rows.Scan(&rec.ID, &rec.OrgID, &rec.Name, &rec.Status, &rec.Payload, &rec.Version, &rec.CreatedAt, &rec.UpdatedAt, &rec.DeletedAt); err != nil {
			return nil, 0, err
		}
		out = append(out, rec)
	}
	return out, total, rows.Err()
}

// SearchName performs ILIKE search within an org.
func (r *SessionRepository) SearchName(ctx context.Context, orgID, q string, limit int) ([]SessionRecord, error) {
	if limit <= 0 {
		limit = 25
	}
	rows, err := r.pool.Query(ctx, fmt.Sprintf(`
SELECT id, org_id, name, status, payload, version, created_at, updated_at, deleted_at
FROM %s WHERE org_id=$1 AND deleted_at IS NULL AND name ILIKE '%%' || $2 || '%%'
ORDER BY updated_at DESC LIMIT $3`, r.fq()), orgID, q, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := make([]SessionRecord, 0)
	for rows.Next() {
		var rec SessionRecord
		if err := rows.Scan(&rec.ID, &rec.OrgID, &rec.Name, &rec.Status, &rec.Payload, &rec.Version, &rec.CreatedAt, &rec.UpdatedAt, &rec.DeletedAt); err != nil {
			return nil, err
		}
		out = append(out, rec)
	}
	return out, rows.Err()
}
