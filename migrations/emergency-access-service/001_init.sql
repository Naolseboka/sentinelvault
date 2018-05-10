CREATE SCHEMA IF NOT EXISTS emergency_access;

CREATE TABLE IF NOT EXISTS emergency_access.sessions (
    id           TEXT PRIMARY KEY,
    org_id       TEXT NOT NULL,
    requester_id TEXT NOT NULL,
    target_org_id TEXT NOT NULL DEFAULT '',
    reason       TEXT NOT NULL,
    status       TEXT NOT NULL DEFAULT 'pending',
    wait_period_secs BIGINT NOT NULL DEFAULT 3600,
    active_until TIMESTAMPTZ,
    approved_by  TEXT,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_emergency_access_org ON emergency_access.sessions (org_id);
