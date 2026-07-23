package analytics

import (
	"math"
	"sort"
	"time"
)

// Point is a time-bucketed metric point.
type Point struct {
	Bucket time.Time
	Value  float64
	Count  int64
}

// Series is a named metric series.
type Series struct {
	Name   string
	Points []Point
}

// Aggregator accumulates events into fixed buckets.
type Aggregator struct {
	bucket time.Duration
	data   map[string]map[int64]*Point // name -> unix bucket -> point
}

func NewAggregator(bucket time.Duration) *Aggregator {
	if bucket <= 0 {
		bucket = time.Hour
	}
	return &Aggregator{bucket: bucket, data: map[string]map[int64]*Point{}}
}

func (a *Aggregator) key(ts time.Time) int64 {
	return ts.UTC().Truncate(a.bucket).Unix()
}

func (a *Aggregator) Add(name string, ts time.Time, value float64) {
	if a.data[name] == nil {
		a.data[name] = map[int64]*Point{}
	}
	k := a.key(ts)
	p := a.data[name][k]
	if p == nil {
		p = &Point{Bucket: time.Unix(k, 0).UTC()}
		a.data[name][k] = p
	}
	p.Value += value
	p.Count++
}

func (a *Aggregator) Series(name string, from, to time.Time) Series {
	out := Series{Name: name, Points: []Point{}}
	m := a.data[name]
	if m == nil {
		return out
	}
	keys := make([]int64, 0, len(m))
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
	for _, k := range keys {
		out.Points = append(out.Points, *m[k])
	}
	return out
}

func (a *Aggregator) Rate(name string, from, to time.Time) float64 {
	s := a.Series(name, from, to)
	var sum float64
	var n int64
	for _, p := range s.Points {
		sum += p.Value
		n += p.Count
	}
	if n == 0 {
		return 0
	}
	return sum / float64(n)
}

// HygieneScorecard summarizes password posture.
type HygieneScorecard struct {
	TotalItems  int
	WeakCount   int
	ReusedCount int
	StaleCount  int
	Score       float64 // 0-100
}

func ComputeHygiene(total, weak, reused, stale int) HygieneScorecard {
	if total < 0 {
		total = 0
	}
	sc := HygieneScorecard{TotalItems: total, WeakCount: weak, ReusedCount: reused, StaleCount: stale}
	if total == 0 {
		sc.Score = 100
		return sc
	}
	penalty := float64(weak*3+reused*4+stale*2) / float64(total) * 10
	sc.Score = math.Max(0, 100-penalty)
	return sc
}

// SessionStats aggregates durations.
type SessionStats struct {
	Count  int
	AvgSec float64
	P95Sec float64
	MaxSec float64
}

func ComputeSessionStats(durations []time.Duration) SessionStats {
	if len(durations) == 0 {
		return SessionStats{}
	}
	secs := make([]float64, len(durations))
	var sum, max float64
	for i, d := range durations {
		s := d.Seconds()
		secs[i] = s
		sum += s
		if s > max {
			max = s
		}
	}
	sort.Float64s(secs)
	p95 := secs[int(math.Min(float64(len(secs)-1), math.Floor(float64(len(secs))*0.95)))]
	return SessionStats{Count: len(secs), AvgSec: sum / float64(len(secs)), P95Sec: p95, MaxSec: max}
}
