package domain

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"time"

	encdomain "github.com/sentinelvault/sentinelvault/internal/encryption"
	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/vault"
)

const defaultAttachmentDir = "/tmp/sv-attachments"

// AttachmentMeta is persisted attachment metadata.
type AttachmentMeta struct {
	ID          string    `json:"id"`
	OrgID       string    `json:"org_id"`
	VaultItemID string    `json:"vault_item_id"`
	Name        string    `json:"name"`
	ContentType string    `json:"content_type"`
	Size        int64     `json:"size"`
	StoragePath string    `json:"-"`
	KeyID       string    `json:"key_id"`
	CreatedAt   time.Time `json:"created_at"`
}

// Repository persists attachment metadata.
type Repository interface {
	Create(ctx context.Context, rec *AttachmentMeta) error
	Get(ctx context.Context, id string) (*AttachmentMeta, error)
	ListByItem(ctx context.Context, vaultItemID string) ([]AttachmentMeta, error)
}

// Service stores encrypted attachment blobs.
type Service struct {
	repo      Repository
	encryptor vault.Encryptor
	baseDir   string
}

// NewService creates an attachment service.
func NewService(repo Repository, enc vault.Encryptor, baseDir string) *Service {
	if baseDir == "" {
		baseDir = defaultAttachmentDir
	}
	return &Service{repo: repo, encryptor: enc, baseDir: baseDir}
}

// UploadRequest stores encrypted attachment bytes.
type UploadRequest struct {
	OrgID       string
	VaultItemID string
	KeyID       string
	Name        string
	ContentType string
	Data        []byte
}

// Upload encrypts and stores an attachment.
func (s *Service) Upload(ctx context.Context, req UploadRequest) (*AttachmentMeta, error) {
	if err := common.ValidateRequired("org_id", req.OrgID); err != nil {
		return nil, err
	}
	if err := common.ValidateRequired("vault_item_id", req.VaultItemID); err != nil {
		return nil, err
	}
	if err := common.ValidateRequired("key_id", req.KeyID); err != nil {
		return nil, err
	}
	if len(req.Data) == 0 {
		return nil, common.New(common.ErrCodeInvalidArgument, "data required")
	}
	if err := os.MkdirAll(s.baseDir, 0o700); err != nil {
		return nil, err
	}

	encrypted, err := s.encryptor.EnvelopeEncrypt(ctx, encdomain.EncryptRequest{
		Plaintext: req.Data,
		KeyID:     req.KeyID,
		AAD:       []byte(req.OrgID + ":" + req.VaultItemID),
	})
	if err != nil {
		return nil, err
	}

	id := common.NewID()
	path := filepath.Join(s.baseDir, id+".bin")
	if err := os.WriteFile(path, encrypted.Ciphertext, 0o600); err != nil {
		return nil, err
	}

	meta := &AttachmentMeta{
		ID:          id,
		OrgID:       req.OrgID,
		VaultItemID: req.VaultItemID,
		Name:        req.Name,
		ContentType: req.ContentType,
		Size:        int64(len(req.Data)),
		StoragePath: path,
		KeyID:       req.KeyID,
		CreatedAt:   common.NowUTC(),
	}
	if err := s.repo.Create(ctx, meta); err != nil {
		_ = os.Remove(path)
		return nil, err
	}
	return meta, nil
}

// Download decrypts attachment bytes for authorized org.
func (s *Service) Download(ctx context.Context, orgID, id string) ([]byte, *AttachmentMeta, error) {
	meta, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	if meta == nil || meta.OrgID != orgID {
		return nil, nil, common.New(common.ErrCodeNotFound, "attachment not found")
	}
	ciphertext, err := os.ReadFile(meta.StoragePath)
	if err != nil {
		return nil, nil, common.Wrap(common.ErrCodeInternal, "read attachment", err)
	}
	decrypted, err := s.encryptor.EnvelopeDecrypt(ctx, encdomain.DecryptRequest{
		Ciphertext: ciphertext,
		AAD:        []byte(meta.OrgID + ":" + meta.VaultItemID),
	})
	if err != nil {
		return nil, nil, err
	}
	return decrypted.Plaintext, meta, nil
}

// GetMetadata returns attachment metadata.
func (s *Service) GetMetadata(ctx context.Context, orgID, id string) (*AttachmentMeta, error) {
	meta, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if meta == nil || meta.OrgID != orgID {
		return nil, common.New(common.ErrCodeNotFound, "attachment not found")
	}
	return meta, nil
}

// ListByItem lists attachments for a vault item.
func (s *Service) ListByItem(ctx context.Context, orgID, vaultItemID string) ([]AttachmentMeta, error) {
	items, err := s.repo.ListByItem(ctx, vaultItemID)
	if err != nil {
		return nil, err
	}
	out := make([]AttachmentMeta, 0)
	for _, item := range items {
		if item.OrgID == orgID {
			out = append(out, item)
		}
	}
	return out, nil
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "attachments" }

// TableName returns the primary table name.
func TableName() string { return "attachment_records" }

// MemoryRepository stores attachment metadata in memory (files still on disk).
type MemoryRepository struct {
	mu   sync.RWMutex
	data map[string]*AttachmentMeta
}

// NewMemoryRepository creates an empty repository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{data: map[string]*AttachmentMeta{}}
}

func (r *MemoryRepository) Create(_ context.Context, rec *AttachmentMeta) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *rec
	r.data[rec.ID] = &cp
	return nil
}

func (r *MemoryRepository) Get(_ context.Context, id string) (*AttachmentMeta, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	rec, ok := r.data[id]
	if !ok {
		return nil, nil
	}
	cp := *rec
	return &cp, nil
}

func (r *MemoryRepository) ListByItem(_ context.Context, vaultItemID string) ([]AttachmentMeta, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]AttachmentMeta, 0)
	for _, rec := range r.data {
		if rec.VaultItemID == vaultItemID {
			out = append(out, *rec)
		}
	}
	return out, nil
}
