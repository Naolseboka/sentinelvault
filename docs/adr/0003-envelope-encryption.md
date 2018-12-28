# ADR-0003: Envelope Encryption

## Status
Accepted

## Context
Vault items contain highly sensitive plaintext that must be encrypted at rest.

## Decision
Use envelope encryption: random DEK per item, AES-256-GCM for data, KEK wraps DEK.

Implementation: `internal/crypto/envelope.go`

## Consequences
- KEK rotation re-wraps DEKs without re-encrypting all data
- Associated data binds ciphertext to item metadata
<!-- touch 4513 -->
