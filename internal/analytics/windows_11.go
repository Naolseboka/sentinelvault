package analytics

import "time"

// Window11 is a predefined analytics window helper.
func Window11(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(12) * time.Hour)
	return start, end
}

// maintenance note 93
