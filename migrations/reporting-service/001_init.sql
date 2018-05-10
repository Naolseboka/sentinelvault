CREATE SCHEMA IF NOT EXISTS reporting;

CREATE TABLE IF NOT EXISTS reporting.reports (
    id           TEXT PRIMARY KEY,
    org_id       TEXT NOT NULL,
    report_type  TEXT NOT NULL,
    summary      TEXT NOT NULL DEFAULT '',
    stats        JSONB NOT NULL DEFAULT '{}',
    metrics      JSONB NOT NULL DEFAULT '{}',
    generated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_reporting_reports_org ON reporting.reports (org_id);
