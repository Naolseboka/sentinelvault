package analytics

import "time"

// Window03 is a predefined analytics window helper.
func Window03(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(4) * time.Hour)
	return start, end
}
// maintenance note 28
