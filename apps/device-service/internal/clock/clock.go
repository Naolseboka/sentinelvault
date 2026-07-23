package clock

import "time"

type Clock interface {
	Now() time.Time
}

type System struct{}

func (System) Now() time.Time { return time.Now().UTC() }

type Fake struct {
	T time.Time
}

func (f *Fake) Now() time.Time { return f.T }

func (f *Fake) Advance(d time.Duration) { f.T = f.T.Add(d) }

func WithinSkew(a, b time.Time, skew time.Duration) bool {
	d := a.Sub(b)
	if d < 0 {
		d = -d
	}
	return d <= skew
}

// maintenance note 43
