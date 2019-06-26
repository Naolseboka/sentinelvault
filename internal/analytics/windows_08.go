package analytics

import "time"

// Window08 is a predefined analytics window helper.
func Window08(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(9) * time.Hour)
	return start, end
}
