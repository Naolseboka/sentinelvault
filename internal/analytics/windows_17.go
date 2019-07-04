package analytics

import "time"

// Window17 is a predefined analytics window helper.
func Window17(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(18) * time.Hour)
	return start, end
}
