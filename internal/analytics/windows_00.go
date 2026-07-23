package analytics

import "time"

// Window00 is a predefined analytics window helper.
func Window00(now time.Time) (time.Time, time.Time) {
	end := now.UTC().Truncate(time.Hour)
	start := end.Add(-time.Duration(1) * time.Hour)
	return start, end
}

// maintenance note 42
