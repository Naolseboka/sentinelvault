CREATE SCHEMA IF NOT EXISTS risk;

CREATE TABLE IF NOT EXISTS risk.evaluations (
    id              TEXT PRIMARY KEY,
    org_id          TEXT NOT NULL,
    user_id         TEXT NOT NULL DEFAULT '',
    score_value     INT NOT NULL DEFAULT 0,
    score_level     TEXT NOT NULL DEFAULT 'low',
    reasons         TEXT[] NOT NULL DEFAULT '{}',
    require_mfa     BOOLEAN NOT NULL DEFAULT FALSE,
    block           BOOLEAN NOT NULL DEFAULT FALSE,
    evaluated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_risk_evaluations_org ON risk.evaluations (org_id);
