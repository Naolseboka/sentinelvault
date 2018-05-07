CREATE SCHEMA IF NOT EXISTS backup;

CREATE TABLE IF NOT EXISTS backup.snapshots (
    id          TEXT PRIMARY KEY,
    org_id      TEXT NOT NULL,
    label       TEXT NOT NULL DEFAULT '',
    size_bytes  INT NOT NULL DEFAULT 0,
    checksum    TEXT NOT NULL,
    encrypted   BOOLEAN NOT NULL DEFAULT TRUE,
    blob        BYTEA NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_backup_snapshots_org ON backup.snapshots (org_id);
