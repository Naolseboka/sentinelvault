CREATE SCHEMA IF NOT EXISTS audit;

CREATE TABLE IF NOT EXISTS audit.audit_events (
    id            TEXT PRIMARY KEY,
    org_id        TEXT NOT NULL,
    actor_id      TEXT NOT NULL DEFAULT '',
    action        TEXT NOT NULL,
    resource_type TEXT NOT NULL DEFAULT '',
    resource_id   TEXT NOT NULL DEFAULT '',
    severity      TEXT NOT NULL DEFAULT 'info',
    metadata      JSONB,
    occurred_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    prev_hash     TEXT NOT NULL,
    hash          TEXT NOT NULL,
    signature     TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_audit_events_org_id ON audit.audit_events (org_id);
CREATE INDEX IF NOT EXISTS idx_audit_events_occurred_at ON audit.audit_events (occurred_at DESC);
