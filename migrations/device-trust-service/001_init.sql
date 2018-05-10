CREATE SCHEMA IF NOT EXISTS device_trust;

CREATE TABLE IF NOT EXISTS device_trust.trust_records (
    id          TEXT PRIMARY KEY,
    org_id      TEXT NOT NULL,
    device_id   TEXT NOT NULL,
    user_id     TEXT NOT NULL DEFAULT '',
    trust_level TEXT NOT NULL DEFAULT 'pending',
    score       INT NOT NULL DEFAULT 0,
    risk_level  TEXT NOT NULL DEFAULT 'low',
    factors     JSONB NOT NULL DEFAULT '{}',
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_device_trust_device ON device_trust.trust_records (device_id);
