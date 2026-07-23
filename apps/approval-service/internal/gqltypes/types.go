package gqltypes

import "time"

type Approval struct {
	ID        string            `json:"id"`
	OrgID     string            `json:"orgId"`
	Name      string            `json:"name"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
	Tags      []string          `json:"tags"`
	Metadata  map[string]string `json:"metadata"`
}

type ApprovalConnection struct {
	Nodes      []Approval `json:"nodes"`
	TotalCount int        `json:"totalCount"`
	PageInfo   PageInfo   `json:"pageInfo"`
}

type PageInfo struct {
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

type ApprovalEdge struct {
	Cursor string   `json:"cursor"`
	Node   Approval `json:"node"`
}

func Connect(items []Approval, offset, limit, total int) ApprovalConnection {
	if limit <= 0 {
		limit = 50
	}
	end := offset + limit
	if end > len(items) {
		end = len(items)
	}
	if offset > len(items) {
		offset = len(items)
	}
	nodes := items[offset:end]
	start, endc := "", ""
	if len(nodes) > 0 {
		start, endc = nodes[0].ID, nodes[len(nodes)-1].ID
	}
	return ApprovalConnection{
		Nodes: nodes, TotalCount: total,
		PageInfo: PageInfo{HasNextPage: end < total, HasPreviousPage: offset > 0, StartCursor: start, EndCursor: endc},
	}
}
