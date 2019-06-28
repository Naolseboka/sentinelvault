package analytics

import "time"

// Window10 is a predefined analytics window helper.
func Window10(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(11) * time.Hour)
	return start, end
}
