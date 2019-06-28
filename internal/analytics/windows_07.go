package analytics

import "time"

// Window07 is a predefined analytics window helper.
func Window07(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(8) * time.Hour)
	return start, end
}
