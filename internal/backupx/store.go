package backupx

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	svcrypto "github.com/sentinelvault/sentinelvault/internal/crypto"
)

// Snapshot is an encrypted backup artifact metadata.
type Snapshot struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"org_id"`
	Label     string    `json:"label"`
	SizeBytes int       `json:"size_bytes"`
	Checksum  string    `json:"checksum"`
	CreatedAt time.Time `json:"created_at"`
	Encrypted bool      `json:"encrypted"`
}

// Store holds encrypted backup blobs in memory (MinIO adapter later).
type Store struct {
	mu    sync.Mutex
	kek   []byte
	meta  map[string]*Snapshot
	blobs map[string][]byte
}

// NewStore creates a store with a 32-byte KEK.
func NewStore(kek []byte) (*Store, error) {
	if len(kek) != svcrypto.AESKeySize {
		k, err := svcrypto.RandomBytes(svcrypto.AESKeySize)
		if err != nil {
			return nil, err
		}
		kek = k
	}
	return &Store{kek: kek, meta: map[string]*Snapshot{}, blobs: map[string][]byte{}}, nil
}

// Backup compresses and envelope-encrypts payload.
func (s *Store) Backup(orgID, label string, payload any) (*Snapshot, error) {
	if orgID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(raw); err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	env, err := svcrypto.EnvelopeEncrypt(buf.Bytes(), s.kek, []byte(orgID))
	if err != nil {
		return nil, err
	}
	blob, err := svcrypto.MarshalJSONEnvelope(env)
	if err != nil {
		return nil, err
	}
	sum, err := svcrypto.SHA256Hex(blob)
	if err != nil {
		// fallback hash via envelope nonce uniqueness
		sum = common.NewID()
	}
	snap := &Snapshot{
		ID: common.NewID(), OrgID: orgID, Label: label, SizeBytes: len(blob),
		Checksum: sum, CreatedAt: common.NowUTC(), Encrypted: true,
	}
	s.mu.Lock()
	s.meta[snap.ID] = snap
	s.blobs[snap.ID] = blob
	s.mu.Unlock()
	cp := *snap
	return &cp, nil
}

// Restore decrypts and unmarshals into dest.
func (s *Store) Restore(id string, dest any) (*Snapshot, error) {
	s.mu.Lock()
	snap, ok := s.meta[id]
	blob := s.blobs[id]
	s.mu.Unlock()
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "snapshot not found")
	}
	env, err := svcrypto.UnmarshalJSONEnvelope(blob)
	if err != nil {
		return nil, err
	}
	plain, err := svcrypto.EnvelopeDecrypt(env, s.kek, []byte(snap.OrgID))
	if err != nil {
		return nil, err
	}
	zr, err := gzip.NewReader(bytes.NewReader(plain))
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	dec := json.NewDecoder(zr)
	if err := dec.Decode(dest); err != nil {
		return nil, err
	}
	cp := *snap
	return &cp, nil
}

// List returns snapshots for an org.
func (s *Store) List(orgID string) []Snapshot {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Snapshot, 0)
	for _, snap := range s.meta {
		if orgID == "" || snap.OrgID == orgID {
			out = append(out, *snap)
		}
	}
	return out
}
