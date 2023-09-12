package gqltypes

import "time"

type Notification struct {
	ID string `json:"id"`
	OrgID string `json:"orgId"`
	Name string `json:"name"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tags []string `json:"tags"`
	Metadata map[string]string `json:"metadata"`
}

type NotificationConnection struct {
	Nodes []Notification `json:"nodes"`
	TotalCount int `json:"totalCount"`
	PageInfo PageInfo `json:"pageInfo"`
}

type PageInfo struct {
	HasNextPage bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	StartCursor string `json:"startCursor"`
	EndCursor string `json:"endCursor"`
}

type NotificationEdge struct {
	Cursor string `json:"cursor"`
	Node Notification `json:"node"`
}

func Connect(items []Notification, offset, limit, total int) NotificationConnection {
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
	return NotificationConnection{
		Nodes: nodes, TotalCount: total,
		PageInfo: PageInfo{HasNextPage: end < total, HasPreviousPage: offset > 0, StartCursor: start, EndCursor: endc},
	}
}
