package analytics

import "time"

// Window19 is a predefined analytics window helper.
func Window19(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(20) * time.Hour)
	return start, end
}
