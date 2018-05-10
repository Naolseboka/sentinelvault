CREATE SCHEMA IF NOT EXISTS threat;

CREATE TABLE IF NOT EXISTS threat.findings (
    id         TEXT PRIMARY KEY,
    org_id     TEXT NOT NULL,
    type       TEXT NOT NULL,
    severity   TEXT NOT NULL DEFAULT 'medium',
    title      TEXT NOT NULL DEFAULT '',
    detail     TEXT NOT NULL DEFAULT '',
    resource   TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_threat_findings_org ON threat.findings (org_id);
