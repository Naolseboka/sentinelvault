package domain

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
	encdomain "github.com/sentinelvault/sentinelvault/internal/encryption"
	"github.com/sentinelvault/sentinelvault/internal/vault"
)

// CreateRequest creates a typed human vault item.
type CreateRequest struct {
	OrgID        string             `json:"org_id"`
	KeyID        string             `json:"key_id"`
	Type         vault.ItemType     `json:"type"`
	Name         string             `json:"name"`
	Metadata     vault.ItemMetadata `json:"metadata"`
	Payload      vault.ItemPayload  `json:"payload"`
	CollectionID string             `json:"collection_id,omitempty"`
	FolderID     string             `json:"folder_id,omitempty"`
	CreatedBy    string             `json:"created_by,omitempty"`
}

// UpdateRequest updates an item and bumps version.
type UpdateRequest struct {
	OrgID     string              `json:"org_id"`
	KeyID     string              `json:"key_id"`
	Name      string              `json:"name,omitempty"`
	Metadata  *vault.ItemMetadata `json:"metadata,omitempty"`
	Payload   vault.ItemPayload   `json:"payload,omitempty"`
	UpdatedBy string              `json:"updated_by,omitempty"`
}

// ItemView is returned to authorized callers with decrypted payload.
type ItemView struct {
	vault.VaultItem
	Payload vault.ItemPayload `json:"payload"`
}

// Repository persists encrypted password vault records.
type Repository interface {
	Create(ctx context.Context, rec *vault.StoredRecord) error
	Update(ctx context.Context, rec *vault.StoredRecord) error
	Get(ctx context.Context, id string) (*vault.StoredRecord, error)
	List(ctx context.Context, orgID string, offset, limit int) ([]vault.StoredRecord, int64, error)
	CreateCollection(ctx context.Context, c *vault.Collection) error
	GetCollection(ctx context.Context, id string) (*vault.Collection, error)
	CreateFolder(ctx context.Context, f *vault.Folder) error
	CreateShareGrant(ctx context.Context, g *vault.ShareGrant) error
	ListShareGrants(ctx context.Context, itemID string) ([]vault.ShareGrant, error)
	SaveVersion(ctx context.Context, snap *vault.VersionSnapshot) error
}

// Service implements password-service business logic.
type Service struct {
	repo      Repository
	encryptor vault.Encryptor
}

// NewService creates a password vault service.
func NewService(repo Repository, enc vault.Encryptor) *Service {
	return &Service{repo: repo, encryptor: enc}
}

// Create stores a new encrypted vault item.
func (s *Service) Create(ctx context.Context, req CreateRequest) (*ItemView, error) {
	if err := common.ValidateRequired("org_id", req.OrgID); err != nil {
		return nil, err
	}
	if err := common.ValidateRequired("name", req.Name); err != nil {
		return nil, err
	}
	if err := common.ValidateRequired("key_id", req.KeyID); err != nil {
		return nil, err
	}
	if !vault.IsHumanType(req.Type) {
		return nil, common.New(common.ErrCodeInvalidArgument, "invalid human item type")
	}
	if err := vault.ValidatePayload(req.Type, req.Payload); err != nil {
		return nil, err
	}
	if err := vault.ValidateMetadata(req.Type, req.Metadata); err != nil {
		return nil, err
	}

	plain, err := json.Marshal(req.Payload)
	if err != nil {
		return nil, err
	}
	aad := []byte(req.OrgID)
	encrypted, err := s.encryptor.EnvelopeEncrypt(ctx, encdomain.EncryptRequest{
		Plaintext: plain,
		KeyID:     req.KeyID,
		AAD:       aad,
	})
	if err != nil {
		return nil, err
	}

	now := common.NowUTC()
	rec := &vault.StoredRecord{
		ID:           common.NewID(),
		OrgID:        req.OrgID,
		Type:         req.Type,
		Name:         req.Name,
		Version:      1,
		Status:       vault.ItemStatusActive,
		Metadata:     req.Metadata,
		Ciphertext:   encrypted.Ciphertext,
		KeyID:        req.KeyID,
		CollectionID: req.CollectionID,
		FolderID:     req.FolderID,
		CreatedBy:    req.CreatedBy,
		UpdatedBy:    req.CreatedBy,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	if err := s.repo.Create(ctx, rec); err != nil {
		return nil, err
	}
	return s.toView(rec, req.Payload), nil
}

// Get decrypts an item for an authorized org caller.
func (s *Service) Get(ctx context.Context, orgID, id string) (*ItemView, error) {
	rec, err := s.loadActive(ctx, orgID, id)
	if err != nil {
		return nil, err
	}
	payload, err := s.decryptPayload(ctx, rec)
	if err != nil {
		return nil, err
	}
	return s.toView(rec, payload), nil
}

// List returns metadata-only views without decrypting payloads.
func (s *Service) List(ctx context.Context, orgID string, offset, limit int) ([]vault.VaultItem, int64, error) {
	recs, total, err := s.repo.List(ctx, orgID, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	out := make([]vault.VaultItem, 0, len(recs))
	for i := range recs {
		out = append(out, recordToItem(&recs[i]))
	}
	return out, total, nil
}

// Update replaces payload and metadata, saving a version snapshot.
func (s *Service) Update(ctx context.Context, id string, req UpdateRequest) (*ItemView, error) {
	rec, err := s.loadActive(ctx, req.OrgID, id)
	if err != nil {
		return nil, err
	}
	if err := s.repo.SaveVersion(ctx, snapshotFromRecord(rec)); err != nil {
		return nil, err
	}

	if req.Name != "" {
		rec.Name = req.Name
	}
	if req.Metadata != nil {
		if err := vault.ValidateMetadata(rec.Type, *req.Metadata); err != nil {
			return nil, err
		}
		rec.Metadata = *req.Metadata
	}
	var payload vault.ItemPayload
	if req.Payload != nil {
		if err := vault.ValidatePayload(rec.Type, req.Payload); err != nil {
			return nil, err
		}
		payload = req.Payload
	} else {
		payload, err = s.decryptPayload(ctx, rec)
		if err != nil {
			return nil, err
		}
	}
	keyID := req.KeyID
	if keyID == "" {
		keyID = rec.KeyID
	}
	plain, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	encrypted, err := s.encryptor.EnvelopeEncrypt(ctx, encdomain.EncryptRequest{
		Plaintext: plain,
		KeyID:     keyID,
		AAD:       []byte(rec.OrgID),
	})
	if err != nil {
		return nil, err
	}
	rec.Version++
	rec.Ciphertext = encrypted.Ciphertext
	rec.KeyID = keyID
	rec.UpdatedBy = req.UpdatedBy
	rec.UpdatedAt = common.NowUTC()
	if err := s.repo.Update(ctx, rec); err != nil {
		return nil, err
	}
	return s.toView(rec, payload), nil
}

// Delete soft-deletes an item.
func (s *Service) Delete(ctx context.Context, orgID, id string) error {
	rec, err := s.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if rec == nil || rec.OrgID != orgID {
		return common.New(common.ErrCodeNotFound, "item not found")
	}
	now := common.NowUTC()
	rec.Status = vault.ItemStatusDeleted
	rec.DeletedAt = &now
	rec.UpdatedAt = now
	return s.repo.Update(ctx, rec)
}

// Archive marks an item archived.
func (s *Service) Archive(ctx context.Context, orgID, id string) error {
	rec, err := s.loadActive(ctx, orgID, id)
	if err != nil {
		return err
	}
	rec.Status = vault.ItemStatusArchived
	rec.UpdatedAt = common.NowUTC()
	return s.repo.Update(ctx, rec)
}

// Restore reactivates archived or deleted items.
func (s *Service) Restore(ctx context.Context, orgID, id string) (*ItemView, error) {
	rec, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if rec == nil || rec.OrgID != orgID {
		return nil, common.New(common.ErrCodeNotFound, "item not found")
	}
	rec.Status = vault.ItemStatusActive
	rec.DeletedAt = nil
	rec.UpdatedAt = common.NowUTC()
	if err := s.repo.Update(ctx, rec); err != nil {
		return nil, err
	}
	payload, err := s.decryptPayload(ctx, rec)
	if err != nil {
		return nil, err
	}
	return s.toView(rec, payload), nil
}

// CreateCollection registers a collection.
func (s *Service) CreateCollection(ctx context.Context, orgID, name, description string) (*vault.Collection, error) {
	if err := common.ValidateRequired("name", name); err != nil {
		return nil, err
	}
	now := common.NowUTC()
	c := &vault.Collection{
		ID:          common.NewID(),
		OrgID:       orgID,
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := s.repo.CreateCollection(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

// CreateFolder registers a folder in a collection.
func (s *Service) CreateFolder(ctx context.Context, orgID, collectionID, name, parentID string) (*vault.Folder, error) {
	if err := common.ValidateRequired("name", name); err != nil {
		return nil, err
	}
	col, err := s.repo.GetCollection(ctx, collectionID)
	if err != nil {
		return nil, err
	}
	if col == nil || col.OrgID != orgID {
		return nil, common.New(common.ErrCodeNotFound, "collection not found")
	}
	now := common.NowUTC()
	f := &vault.Folder{
		ID:           common.NewID(),
		OrgID:        orgID,
		CollectionID: collectionID,
		ParentID:     parentID,
		Name:         name,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	if err := s.repo.CreateFolder(ctx, f); err != nil {
		return nil, err
	}
	return f, nil
}

// CreateShareGrant grants time-limited access.
func (s *Service) CreateShareGrant(ctx context.Context, orgID, itemID, granteeID, permission string, expiresAt time.Time) (*vault.ShareGrant, error) {
	if _, err := s.loadActive(ctx, orgID, itemID); err != nil {
		return nil, err
	}
	now := common.NowUTC()
	g := &vault.ShareGrant{
		ID:         common.NewID(),
		OrgID:      orgID,
		ItemID:     itemID,
		GranteeID:  granteeID,
		Permission: permission,
		ExpiresAt:  expiresAt,
		CreatedAt:  now,
	}
	if err := s.repo.CreateShareGrant(ctx, g); err != nil {
		return nil, err
	}
	return g, nil
}

func (s *Service) loadActive(ctx context.Context, orgID, id string) (*vault.StoredRecord, error) {
	rec, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if rec == nil || rec.OrgID != orgID || rec.Status == vault.ItemStatusDeleted {
		return nil, common.New(common.ErrCodeNotFound, "item not found")
	}
	return rec, nil
}

func (s *Service) decryptPayload(ctx context.Context, rec *vault.StoredRecord) (vault.ItemPayload, error) {
	decrypted, err := s.encryptor.EnvelopeDecrypt(ctx, encdomain.DecryptRequest{
		Ciphertext: rec.Ciphertext,
		AAD:        []byte(rec.OrgID),
	})
	if err != nil {
		return nil, err
	}
	var payload vault.ItemPayload
	if err := json.Unmarshal(decrypted.Plaintext, &payload); err != nil {
		return nil, err
	}
	return payload, nil
}

func (s *Service) toView(rec *vault.StoredRecord, payload vault.ItemPayload) *ItemView {
	item := recordToItem(rec)
	return &ItemView{VaultItem: item, Payload: payload}
}

func recordToItem(rec *vault.StoredRecord) vault.VaultItem {
	return vault.VaultItem{
		ID:           rec.ID,
		OrgID:        rec.OrgID,
		Type:         rec.Type,
		Name:         rec.Name,
		Version:      rec.Version,
		Status:       rec.Status,
		Metadata:     rec.Metadata,
		CollectionID: rec.CollectionID,
		FolderID:     rec.FolderID,
		KeyID:        rec.KeyID,
		CreatedBy:    rec.CreatedBy,
		UpdatedBy:    rec.UpdatedBy,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
	}
}

func snapshotFromRecord(rec *vault.StoredRecord) *vault.VersionSnapshot {
	return &vault.VersionSnapshot{
		ID:         common.NewID(),
		ItemID:     rec.ID,
		OrgID:      rec.OrgID,
		Version:    rec.Version,
		Type:       rec.Type,
		Name:       rec.Name,
		Metadata:   rec.Metadata,
		Ciphertext: append([]byte(nil), rec.Ciphertext...),
		KeyID:      rec.KeyID,
		CreatedBy:  rec.UpdatedBy,
		CreatedAt:  rec.UpdatedAt,
	}
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "passwords" }

// TableName returns the primary table name.
func TableName() string { return "vault_items" }

// MemoryRepository is an in-memory store for development and tests.
type MemoryRepository struct {
	mu          sync.RWMutex
	items       map[string]*vault.StoredRecord
	collections map[string]*vault.Collection
	folders     map[string]*vault.Folder
	grants      map[string]*vault.ShareGrant
	versions    map[string]*vault.VersionSnapshot
}

// NewMemoryRepository creates an empty repository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		items:       map[string]*vault.StoredRecord{},
		collections: map[string]*vault.Collection{},
		folders:     map[string]*vault.Folder{},
		grants:      map[string]*vault.ShareGrant{},
		versions:    map[string]*vault.VersionSnapshot{},
	}
}

func (r *MemoryRepository) Create(_ context.Context, rec *vault.StoredRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *rec
	cp.Ciphertext = append([]byte(nil), rec.Ciphertext...)
	r.items[rec.ID] = &cp
	return nil
}

func (r *MemoryRepository) Update(_ context.Context, rec *vault.StoredRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *rec
	cp.Ciphertext = append([]byte(nil), rec.Ciphertext...)
	r.items[rec.ID] = &cp
	return nil
}

func (r *MemoryRepository) Get(_ context.Context, id string) (*vault.StoredRecord, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	rec, ok := r.items[id]
	if !ok {
		return nil, nil
	}
	cp := *rec
	cp.Ciphertext = append([]byte(nil), rec.Ciphertext...)
	return &cp, nil
}

func (r *MemoryRepository) List(_ context.Context, orgID string, offset, limit int) ([]vault.StoredRecord, int64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	all := make([]vault.StoredRecord, 0)
	for _, rec := range r.items {
		if rec.OrgID == orgID && rec.Status != vault.ItemStatusDeleted {
			cp := *rec
			all = append(all, cp)
		}
	}
	total := int64(len(all))
	if offset >= len(all) {
		return []vault.StoredRecord{}, total, nil
	}
	end := offset + limit
	if end > len(all) {
		end = len(all)
	}
	return all[offset:end], total, nil
}

func (r *MemoryRepository) CreateCollection(_ context.Context, c *vault.Collection) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *c
	r.collections[c.ID] = &cp
	return nil
}

func (r *MemoryRepository) GetCollection(_ context.Context, id string) (*vault.Collection, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	c, ok := r.collections[id]
	if !ok {
		return nil, nil
	}
	cp := *c
	return &cp, nil
}

func (r *MemoryRepository) CreateFolder(_ context.Context, f *vault.Folder) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *f
	r.folders[f.ID] = &cp
	return nil
}

func (r *MemoryRepository) CreateShareGrant(_ context.Context, g *vault.ShareGrant) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *g
	r.grants[g.ID] = &cp
	return nil
}

func (r *MemoryRepository) ListShareGrants(_ context.Context, itemID string) ([]vault.ShareGrant, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]vault.ShareGrant, 0)
	for _, g := range r.grants {
		if g.ItemID == itemID {
			out = append(out, *g)
		}
	}
	return out, nil
}

func (r *MemoryRepository) SaveVersion(_ context.Context, snap *vault.VersionSnapshot) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *snap
	cp.Ciphertext = append([]byte(nil), snap.Ciphertext...)
	r.versions[snap.ID] = &cp
	return nil
}
