CREATE SCHEMA IF NOT EXISTS delegations;

CREATE TABLE IF NOT EXISTS delegations.delegations (
    id           TEXT PRIMARY KEY,
    org_id       TEXT NOT NULL,
    delegator_id TEXT NOT NULL,
    delegatee_id TEXT NOT NULL,
    scope        TEXT NOT NULL DEFAULT '',
    status       TEXT NOT NULL DEFAULT 'approved',
    expires_at   TIMESTAMPTZ NOT NULL,
    revoked_at   TIMESTAMPTZ,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_delegations_org ON delegations.delegations (org_id);
