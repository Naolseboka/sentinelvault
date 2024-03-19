package domain

import (
	"context"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/searchx"
)

// Document re-exports searchx.Document.
type Document = searchx.Document

// Hit re-exports searchx.Hit.
type Hit = searchx.Hit

// IndexInput holds document indexing fields.
type IndexInput struct {
	ID     string            `json:"id"`
	OrgID  string            `json:"org_id"`
	Index  string            `json:"index"`
	Title  string            `json:"title"`
	Body   string            `json:"body"`
	Tags   []string          `json:"tags"`
	Fields map[string]string `json:"fields,omitempty"`
}

// Service wraps the search index.
type Service struct {
	idx *searchx.Index
}

// NewService creates a search domain service.
func NewService(idx *searchx.Index) *Service {
	return &Service{idx: idx}
}

// NewIndex creates an empty search index.
func NewIndex() *searchx.Index {
	return searchx.NewIndex()
}

// Upsert indexes or updates a document.
func (s *Service) Upsert(_ context.Context, in IndexInput) (*Document, error) {
	if in.OrgID == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id required")
	}
	if in.ID == "" {
		in.ID = common.NewID()
	}
	doc := Document{
		ID: in.ID, OrgID: in.OrgID, Index: in.Index,
		Title: in.Title, Body: in.Body, Tags: in.Tags, Fields: in.Fields,
	}
	s.idx.Upsert(doc)
	return &doc, nil
}

// Query searches documents.
func (s *Service) Query(_ context.Context, orgID, indexName, q string, tags []string, limit int) []Hit {
	return s.idx.Query(orgID, indexName, q, tags, limit)
}

// SchemaName returns the Postgres schema for this service.
func SchemaName() string { return "search" }

// TableName returns the primary table name.
func TableName() string { return "documents" }
