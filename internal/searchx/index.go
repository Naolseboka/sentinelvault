package searchx

import (
	"strings"
	"sync"
)

// Document is a searchable record.
type Document struct {
	ID       string            `json:"id"`
	OrgID    string            `json:"org_id"`
	Index    string            `json:"index"`
	Title    string            `json:"title"`
	Body     string            `json:"body"`
	Tags     []string          `json:"tags"`
	Fields   map[string]string `json:"fields,omitempty"`
}

// Hit is a search result with score.
type Hit struct {
	Document Document `json:"document"`
	Score    float64  `json:"score"`
}

// Index is an in-memory full-text/tag search (OpenSearch adapter later).
type Index struct {
	mu   sync.RWMutex
	docs map[string]Document
}

// NewIndex creates an empty index.
func NewIndex() *Index {
	return &Index{docs: map[string]Document{}}
}

// Upsert indexes a document.
func (idx *Index) Upsert(doc Document) {
	idx.mu.Lock()
	defer idx.mu.Unlock()
	idx.docs[doc.ID] = doc
}

// Delete removes a document.
func (idx *Index) Delete(id string) {
	idx.mu.Lock()
	defer idx.mu.Unlock()
	delete(idx.docs, id)
}

// Query performs fuzzy-ish substring search with tag filters.
func (idx *Index) Query(orgID, indexName, q string, tags []string, limit int) []Hit {
	idx.mu.RLock()
	defer idx.mu.RUnlock()
	if limit <= 0 {
		limit = 50
	}
	q = strings.ToLower(strings.TrimSpace(q))
	hits := make([]Hit, 0)
	for _, doc := range idx.docs {
		if orgID != "" && doc.OrgID != orgID {
			continue
		}
		if indexName != "" && doc.Index != indexName {
			continue
		}
		if len(tags) > 0 && !containsAllTags(doc.Tags, tags) {
			continue
		}
		score := scoreDoc(doc, q)
		if q != "" && score <= 0 {
			continue
		}
		hits = append(hits, Hit{Document: doc, Score: score})
		if len(hits) >= limit {
			break
		}
	}
	// simple sort by score desc
	for i := 0; i < len(hits); i++ {
		for j := i + 1; j < len(hits); j++ {
			if hits[j].Score > hits[i].Score {
				hits[i], hits[j] = hits[j], hits[i]
			}
		}
	}
	return hits
}

func containsAllTags(have, need []string) bool {
	set := map[string]struct{}{}
	for _, t := range have {
		set[strings.ToLower(t)] = struct{}{}
	}
	for _, t := range need {
		if _, ok := set[strings.ToLower(t)]; !ok {
			return false
		}
	}
	return true
}

func scoreDoc(doc Document, q string) float64 {
	if q == "" {
		return 1
	}
	score := 0.0
	title := strings.ToLower(doc.Title)
	body := strings.ToLower(doc.Body)
	if title == q {
		score += 5
	}
	if strings.Contains(title, q) {
		score += 3
	}
	if strings.Contains(body, q) {
		score += 1
	}
	for _, tag := range doc.Tags {
		if strings.ToLower(tag) == q {
			score += 2
		}
	}
	// fuzzy: token overlap
	for _, tok := range strings.Fields(q) {
		if strings.Contains(title, tok) {
			score += 0.5
		}
		if strings.Contains(body, tok) {
			score += 0.25
		}
	}
	return score
}
