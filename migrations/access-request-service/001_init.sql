CREATE SCHEMA IF NOT EXISTS access_requests;

CREATE TABLE IF NOT EXISTS access_requests.requests (
    id            TEXT PRIMARY KEY,
    org_id        TEXT NOT NULL,
    user_id       TEXT NOT NULL,
    resource_id   TEXT NOT NULL,
    permission    TEXT NOT NULL DEFAULT '',
    justification TEXT NOT NULL DEFAULT '',
    status        TEXT NOT NULL DEFAULT 'pending',
    duration_secs BIGINT NOT NULL DEFAULT 0,
    expires_at    TIMESTAMPTZ NOT NULL,
    approved_by   TEXT,
    decided_at    TIMESTAMPTZ,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_access_requests_org_status ON access_requests.requests (org_id, status);
