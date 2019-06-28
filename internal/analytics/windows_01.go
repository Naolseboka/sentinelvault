package analytics

import "time"

// Window01 is a predefined analytics window helper.
func Window01(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(2) * time.Hour)
	return start, end
}
