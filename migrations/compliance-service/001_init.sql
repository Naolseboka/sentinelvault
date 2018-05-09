CREATE SCHEMA IF NOT EXISTS compliance;

CREATE TABLE IF NOT EXISTS compliance.evidence (
    id           TEXT PRIMARY KEY,
    org_id       TEXT NOT NULL,
    control_id   TEXT NOT NULL,
    source       TEXT NOT NULL DEFAULT '',
    summary      TEXT NOT NULL DEFAULT '',
    collected_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS compliance.reports (
    id             TEXT PRIMARY KEY,
    org_id         TEXT NOT NULL,
    framework      TEXT NOT NULL,
    generated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    total_controls INT NOT NULL DEFAULT 0,
    covered        INT NOT NULL DEFAULT 0,
    coverage_pct   DOUBLE PRECISION NOT NULL DEFAULT 0,
    gaps           TEXT[] NOT NULL DEFAULT '{}',
    evidence_ids   TEXT[] NOT NULL DEFAULT '{}'
);

CREATE INDEX IF NOT EXISTS idx_compliance_evidence_org ON compliance.evidence (org_id);
