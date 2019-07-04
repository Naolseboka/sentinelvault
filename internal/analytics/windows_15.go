package analytics

import "time"

// Window15 is a predefined analytics window helper.
func Window15(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(16) * time.Hour)
	return start, end
}
