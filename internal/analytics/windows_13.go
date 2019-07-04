package analytics

import "time"

// Window13 is a predefined analytics window helper.
func Window13(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(14) * time.Hour)
	return start, end
}
