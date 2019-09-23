package common

import (
	"net/http"
	"strconv"
)

const (
	DefaultPageSize = 20
	MaxPageSize     = 100
)

// PageParams holds pagination query parameters.
type PageParams struct {
	Page     int
	PageSize int
	Offset   int
}

// ParsePageParams extracts page and page_size from query string.
func ParsePageParams(r *http.Request) PageParams {
	page := parseIntDefault(r.URL.Query().Get("page"), 1)
	size := parseIntDefault(r.URL.Query().Get("page_size"), DefaultPageSize)
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = DefaultPageSize
	}
	if size > MaxPageSize {
		size = MaxPageSize
	}
	return PageParams{
		Page:     page,
		PageSize: size,
		Offset:   (page - 1) * size,
	}
}

// PageMeta describes pagination metadata in list responses.
type PageMeta struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalCount int64 `json:"total_count"`
	TotalPages int   `json:"total_pages"`
}

// NewPageMeta builds pagination metadata.
func NewPageMeta(params PageParams, total int64) PageMeta {
	pages := int(total) / params.PageSize
	if int(total)%params.PageSize > 0 {
		pages++
	}
	return PageMeta{
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalCount: total,
		TotalPages: pages,
	}
}

func parseIntDefault(s string, def int) int {
	if s == "" {
		return def
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return n
}
