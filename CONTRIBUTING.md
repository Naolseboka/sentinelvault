# Contributing to SentinelVault

Thank you for contributing. Please follow these guidelines.

## Development Setup

1. Install Go 1.22+
2. Clone the repository
3. Run `make bootstrap` to verify tooling
4. Copy `.env.example` to `.env`

## Code Standards

- Match existing package structure and naming
- Add unit tests for shared libraries (`internal/common`, `internal/crypto`, etc.)
- Use typed errors from `internal/common`
- Keep services thin — domain logic in `internal/domain`

## Pull Request Process

1. Create a feature branch from `main`
2. Run `make lint` and `make test`
3. Update docs/ADRs for architectural changes
4. Open a PR with a clear description and test plan

## Commit Messages

Use imperative mood: `add user-service pagination`, `fix envelope decrypt`.

## Security

Report vulnerabilities privately to security@sentinelvault.io. Do not open public issues for security findings.
