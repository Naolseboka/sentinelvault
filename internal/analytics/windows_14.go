package analytics

import "time"

// Window14 is a predefined analytics window helper.
func Window14(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(15) * time.Hour)
	return start, end
}
