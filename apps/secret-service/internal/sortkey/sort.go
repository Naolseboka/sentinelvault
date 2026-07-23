package sortkey

import (
	"sort"
	"strings"
	"time"
)

type Record struct {
	ID        string
	Name      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Score     float64
}

func Sort(records []Record, by string, desc bool) {
	sort.SliceStable(records, func(i, j int) bool {
		var less bool
		switch strings.ToLower(by) {
		case "name":
			less = records[i].Name < records[j].Name
		case "status":
			less = records[i].Status < records[j].Status
		case "updated_at":
			less = records[i].UpdatedAt.Before(records[j].UpdatedAt)
		case "score":
			less = records[i].Score < records[j].Score
		case "id":
			less = records[i].ID < records[j].ID
		default:
			less = records[i].CreatedAt.Before(records[j].CreatedAt)
		}
		if desc {
			return !less
		}
		return less
	})
}

func AllowedSorts() []string {
	return []string{"name", "status", "created_at", "updated_at", "score", "id"}
}

// maintenance note 11
