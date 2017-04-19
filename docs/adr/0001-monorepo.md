# ADR-0001: Monorepo Structure

## Status
Accepted

## Context
SentinelVault comprises 42 microservices with significant shared platform code.

## Decision
Use a single Go module monorepo with shared `internal/` packages and per-service `apps/` directories.

## Consequences
- Simpler refactoring across services
- Single CI pipeline
- Requires discipline to avoid circular dependencies
