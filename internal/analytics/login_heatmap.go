package analytics

import (
	"sort"
	"time"
)

// LoginHeatmapCalculator tracks login heatmap.
type LoginHeatmapCalculator struct {
	bucket time.Duration
	data   map[string]map[int64]float64
	counts map[string]map[int64]int64
}

func NewLoginHeatmapCalculator(bucket time.Duration) *LoginHeatmapCalculator {
	if bucket <= 0 {
		bucket = time.Hour
	}
	return &LoginHeatmapCalculator{bucket: bucket, data: map[string]map[int64]float64{}, counts: map[string]map[int64]int64{}}
}

func (c *LoginHeatmapCalculator) key(ts time.Time) int64 {
	return ts.UTC().Truncate(c.bucket).Unix()
}

// Record adds an observation for orgID at timestamp.
func (c *LoginHeatmapCalculator) Record(orgID string, ts time.Time, value float64) {
	if orgID == "" {
		return
	}
	k := c.key(ts)
	if c.data[orgID] == nil {
		c.data[orgID] = map[int64]float64{}
		c.counts[orgID] = map[int64]int64{}
	}
	c.data[orgID][k] += value
	c.counts[orgID][k]++
}

// Window returns sorted points in [from,to].
func (c *LoginHeatmapCalculator) Window(orgID string, from, to time.Time) []Point {
	m := c.data[orgID]
	if m == nil {
		return nil
	}
	keys := make([]int64, 0)
	for k := range m {
		if !from.IsZero() && k < from.Unix() {
			continue
		}
		if !to.IsZero() && k > to.Unix() {
			continue
		}
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	out := make([]Point, 0, len(keys))
	for _, k := range keys {
		out = append(out, Point{Bucket: time.Unix(k, 0).UTC(), Value: m[k], Count: c.counts[orgID][k]})
	}
	return out
}

// Merge combines another calculator into this one (same bucket size assumed).
func (c *LoginHeatmapCalculator) Merge(other *LoginHeatmapCalculator) {
	if other == nil {
		return
	}
	for org, buckets := range other.data {
		if c.data[org] == nil {
			c.data[org] = map[int64]float64{}
			c.counts[org] = map[int64]int64{}
		}
		for k, v := range buckets {
			c.data[org][k] += v
			c.counts[org][k] += other.counts[org][k]
		}
	}
}

// Average computes mean over window.
func (c *LoginHeatmapCalculator) Average(orgID string, from, to time.Time) float64 {
	pts := c.Window(orgID, from, to)
	if len(pts) == 0 {
		return 0
	}
	var sum float64
	var n int64
	for _, p := range pts {
		sum += p.Value
		n += p.Count
	}
	if n == 0 {
		return 0
	}
	return sum / float64(n)
}
