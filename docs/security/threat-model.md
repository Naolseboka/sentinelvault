# Threat Model (Phase 1)

## Assets
- User credentials and vault secrets
- Encryption keys (DEK, KEK)
- Audit logs
- Session tokens

## Trust Boundaries
- Client ↔ API Gateway (TLS, JWT)
- API Gateway ↔ Services (mTLS in production)
- Services ↔ Postgres/Redis/Kafka (network isolation)

## Threats & Mitigations

| Threat | Mitigation |
|--------|------------|
| Secret exposure at rest | Envelope encryption (AES-256-GCM) |
| Password database leak | Argon2id hashing |
| Token forgery | Ed25519/JWT validation at gateway |
| Replay attacks | Idempotency keys, nonce in AEAD |
| Insider abuse | Audit service, RBAC, approval workflows |

## Out of Scope (Phase 1)
- HSM integration
- Zero-knowledge client-side encryption
- Formal penetration testing
