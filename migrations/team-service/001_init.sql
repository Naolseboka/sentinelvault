CREATE SCHEMA IF NOT EXISTS teams;

CREATE TABLE IF NOT EXISTS teams.teams (
    id          TEXT PRIMARY KEY,
    org_id      TEXT NOT NULL,
    name        TEXT NOT NULL,
    status      TEXT NOT NULL DEFAULT 'active',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS teams.team_members (
    id          TEXT PRIMARY KEY,
    team_id     TEXT NOT NULL REFERENCES teams.teams(id),
    user_id     TEXT NOT NULL,
    role        TEXT NOT NULL DEFAULT 'member',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (team_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_teams_org_id ON teams.teams (org_id);
CREATE INDEX IF NOT EXISTS idx_team_members_team_id ON teams.team_members (team_id);
