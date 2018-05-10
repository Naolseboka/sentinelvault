CREATE SCHEMA IF NOT EXISTS directory_sync;

CREATE TABLE IF NOT EXISTS directory_sync.sync_jobs (
    id            TEXT PRIMARY KEY,
    org_id        TEXT NOT NULL,
    connector     TEXT NOT NULL,
    status        TEXT NOT NULL DEFAULT 'pending',
    users_created INT NOT NULL DEFAULT 0,
    users_updated INT NOT NULL DEFAULT 0,
    users_removed INT NOT NULL DEFAULT 0,
    error         TEXT,
    started_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    finished_at   TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_directory_sync_jobs_org ON directory_sync.sync_jobs (org_id);
