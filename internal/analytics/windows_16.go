package analytics

import "time"

// Window16 is a predefined analytics window helper.
func Window16(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(17) * time.Hour)
	return start, end
}
