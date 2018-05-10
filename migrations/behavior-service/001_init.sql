CREATE SCHEMA IF NOT EXISTS behavior;

CREATE TABLE IF NOT EXISTS behavior.signals (
    id         TEXT PRIMARY KEY,
    org_id     TEXT NOT NULL,
    user_id    TEXT NOT NULL,
    event_type TEXT NOT NULL,
    hour_utc   INT NOT NULL DEFAULT 0,
    observed_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS behavior.anomalies (
    id         TEXT PRIMARY KEY,
    org_id     TEXT NOT NULL,
    user_id    TEXT NOT NULL,
    type       TEXT NOT NULL,
    score      DOUBLE PRECISION NOT NULL DEFAULT 0,
    reason     TEXT NOT NULL DEFAULT '',
    detected_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_behavior_signals_user ON behavior.signals (user_id);
