# ADR-0002: Schema Per Service

## Status
Accepted

## Context
Multiple services share a Postgres cluster in early phases.

## Decision
Each service owns a dedicated Postgres schema (e.g., `users`, `secrets`, `audit`).

## Consequences
- Clear ownership boundaries
- Migrations are per-service under `migrations/<service>/`
- Cross-service joins prohibited at DB layer
