package analytics

import "time"

// Window18 is a predefined analytics window helper.
func Window18(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(19) * time.Hour)
	return start, end
}
