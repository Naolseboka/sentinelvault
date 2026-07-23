package audit

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	svcrypto "github.com/sentinelvault/sentinelvault/internal/crypto"
)

// Event is an immutable audit record.
type Event struct {
	ID           string            `json:"id"`
	OrgID        string            `json:"org_id"`
	ActorID      string            `json:"actor_id"`
	Action       string            `json:"action"`
	ResourceType string            `json:"resource_type"`
	ResourceID   string            `json:"resource_id"`
	Severity     string            `json:"severity"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	OccurredAt   time.Time         `json:"occurred_at"`
	PrevHash     string            `json:"prev_hash"`
	Hash         string            `json:"hash"`
	Signature    string            `json:"signature"`
}

// Chain maintains a hash-linked, signed audit log.
type Chain struct {
	priv     ed25519.PrivateKey
	pub      ed25519.PublicKey
	lastHash string
	events   []Event
}

// NewChain creates a chain with a fresh Ed25519 keypair.
func NewChain() (*Chain, error) {
	pub, priv, err := svcrypto.GenerateEd25519KeyPair()
	if err != nil {
		return nil, err
	}
	return &Chain{
		priv:     priv,
		pub:      pub,
		lastHash: genesisHash(),
		events:   make([]Event, 0, 64),
	}, nil
}

// NewChainWithKeys creates a chain from existing keys.
func NewChainWithKeys(pub ed25519.PublicKey, priv ed25519.PrivateKey) *Chain {
	return &Chain{
		priv:     priv,
		pub:      pub,
		lastHash: genesisHash(),
		events:   make([]Event, 0, 64),
	}
}

func genesisHash() string {
	sum := sha256.Sum256([]byte("sentinelvault-audit-genesis-v1"))
	return hex.EncodeToString(sum[:])
}

// Append adds a signed, hash-chained event.
func (c *Chain) Append(orgID, actorID, action, resourceType, resourceID, severity string, metadata map[string]string) (*Event, error) {
	if orgID == "" || action == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and action are required")
	}
	if severity == "" {
		severity = "info"
	}
	ev := Event{
		ID:           common.NewID(),
		OrgID:        orgID,
		ActorID:      actorID,
		Action:       action,
		ResourceType: resourceType,
		ResourceID:   resourceID,
		Severity:     severity,
		Metadata:     metadata,
		OccurredAt:   common.NowUTC(),
		PrevHash:     c.lastHash,
	}
	payload, err := canonicalPayload(ev)
	if err != nil {
		return nil, err
	}
	sum := sha256.Sum256(payload)
	ev.Hash = hex.EncodeToString(sum[:])
	sig, err := svcrypto.SignEd25519(c.priv, []byte(ev.Hash))
	if err != nil {
		return nil, err
	}
	ev.Signature = hex.EncodeToString(sig)
	c.events = append(c.events, ev)
	c.lastHash = ev.Hash
	return &ev, nil
}

func canonicalPayload(ev Event) ([]byte, error) {
	type wire struct {
		ID           string            `json:"id"`
		OrgID        string            `json:"org_id"`
		ActorID      string            `json:"actor_id"`
		Action       string            `json:"action"`
		ResourceType string            `json:"resource_type"`
		ResourceID   string            `json:"resource_id"`
		Severity     string            `json:"severity"`
		Metadata     map[string]string `json:"metadata"`
		OccurredAt   time.Time         `json:"occurred_at"`
		PrevHash     string            `json:"prev_hash"`
	}
	return json.Marshal(wire{
		ID: ev.ID, OrgID: ev.OrgID, ActorID: ev.ActorID, Action: ev.Action,
		ResourceType: ev.ResourceType, ResourceID: ev.ResourceID, Severity: ev.Severity,
		Metadata: ev.Metadata, OccurredAt: ev.OccurredAt, PrevHash: ev.PrevHash,
	})
}

// VerifyIntegrity checks hash chain and signatures.
func (c *Chain) VerifyIntegrity() error {
	prev := genesisHash()
	for i, ev := range c.events {
		if ev.PrevHash != prev {
			return fmt.Errorf("hash chain broken at index %d", i)
		}
		payload, err := canonicalPayload(ev)
		if err != nil {
			return err
		}
		sum := sha256.Sum256(payload)
		want := hex.EncodeToString(sum[:])
		if ev.Hash != want {
			return fmt.Errorf("hash mismatch at index %d", i)
		}
		sig, err := hex.DecodeString(ev.Signature)
		if err != nil {
			return err
		}
		if err := svcrypto.VerifyEd25519(c.pub, []byte(ev.Hash), sig); err != nil {
			return fmt.Errorf("invalid signature at index %d: %w", i, err)
		}
		prev = ev.Hash
	}
	return nil
}

// List returns a copy of events, optionally filtered by org.
func (c *Chain) List(orgID string, offset, limit int) ([]Event, int64) {
	filtered := make([]Event, 0, len(c.events))
	for _, ev := range c.events {
		if orgID == "" || ev.OrgID == orgID {
			filtered = append(filtered, ev)
		}
	}
	total := int64(len(filtered))
	if offset >= len(filtered) {
		return []Event{}, total
	}
	end := offset + limit
	if limit <= 0 || end > len(filtered) {
		end = len(filtered)
	}
	out := make([]Event, end-offset)
	copy(out, filtered[offset:end])
	return out, total
}

// PublicKey returns the verification key.
func (c *Chain) PublicKey() ed25519.PublicKey { return c.pub }

// Len returns event count.
func (c *Chain) Len() int { return len(c.events) }
