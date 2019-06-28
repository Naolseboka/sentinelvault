package analytics

import "time"

// Window04 is a predefined analytics window helper.
func Window04(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(5) * time.Hour)
	return start, end
}
