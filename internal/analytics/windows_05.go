package analytics

import "time"

// Window05 is a predefined analytics window helper.
func Window05(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(6) * time.Hour)
	return start, end
}
