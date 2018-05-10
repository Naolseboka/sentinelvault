CREATE SCHEMA IF NOT EXISTS rotation;

CREATE TABLE IF NOT EXISTS rotation.policies (
    id                TEXT PRIMARY KEY,
    org_id            TEXT NOT NULL,
    name              TEXT NOT NULL,
    target_type       TEXT NOT NULL DEFAULT 'secret',
    target_id         TEXT NOT NULL,
    mode              TEXT NOT NULL DEFAULT 'automatic',
    interval_seconds  BIGINT NOT NULL DEFAULT 0,
    next_run_at       TIMESTAMPTZ,
    approval_required BOOLEAN NOT NULL DEFAULT FALSE,
    enabled           BOOLEAN NOT NULL DEFAULT TRUE,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS rotation.jobs (
    id          TEXT PRIMARY KEY,
    policy_id   TEXT NOT NULL REFERENCES rotation.policies(id),
    org_id      TEXT NOT NULL,
    status      TEXT NOT NULL DEFAULT 'pending',
    mode        TEXT NOT NULL DEFAULT '',
    error       TEXT,
    started_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    finished_at TIMESTAMPTZ,
    new_version TEXT
);

CREATE INDEX IF NOT EXISTS idx_rotation_policies_org ON rotation.policies (org_id);
