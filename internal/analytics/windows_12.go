package analytics

import "time"

// Window12 is a predefined analytics window helper.
func Window12(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(13) * time.Hour)
	return start, end
}
