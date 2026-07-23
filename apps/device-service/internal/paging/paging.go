package paging

import "github.com/sentinelvault/sentinelvault/internal/common"

type Page struct {
	Offset  int   `json:"offset"`
	Limit   int   `json:"limit"`
	Total   int64 `json:"total"`
	HasMore bool  `json:"has_more"`
}

func Normalize(offset, limit, defaultLimit, maxLimit int) (int, int, error) {
	if offset < 0 {
		return 0, 0, common.New(common.ErrCodeInvalidArgument, "offset must be >= 0")
	}
	if defaultLimit <= 0 {
		defaultLimit = 50
	}
	if maxLimit <= 0 {
		maxLimit = 500
	}
	if limit <= 0 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}
	return offset, limit, nil
}

func NewPage(offset, limit int, total int64) Page {
	return Page{Offset: offset, Limit: limit, Total: total, HasMore: int64(offset+limit) < total}
}

func Slice[T any](in []T, offset, limit int) []T {
	offset, limit, err := Normalize(offset, limit, 50, 500)
	if err != nil {
		return nil
	}
	if offset >= len(in) {
		return []T{}
	}
	end := offset + limit
	if end > len(in) {
		end = len(in)
	}
	return in[offset:end]
}

// maintenance note 87
