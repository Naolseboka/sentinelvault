CREATE SCHEMA IF NOT EXISTS lifecycle;

CREATE TABLE IF NOT EXISTS lifecycle.user_states (
    id         TEXT PRIMARY KEY,
    org_id     TEXT NOT NULL,
    user_id    TEXT NOT NULL,
    state      TEXT NOT NULL DEFAULT 'active',
    reason     TEXT,
    changed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (org_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_lifecycle_user_states_org ON lifecycle.user_states (org_id);
