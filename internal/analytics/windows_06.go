package analytics

import "time"

// Window06 is a predefined analytics window helper.
func Window06(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(7) * time.Hour)
	return start, end
}
