package common

import "time"

// NowUTC returns the current time in UTC.
func NowUTC() time.Time {
	return time.Now().UTC()
}

// FormatRFC3339 formats t as RFC3339 in UTC.
func FormatRFC3339(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

// ParseRFC3339 parses an RFC3339 timestamp.
func ParseRFC3339(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// TruncateToDay returns midnight UTC for the given time.
func TruncateToDay(t time.Time) time.Time {
	utc := t.UTC()
	return time.Date(utc.Year(), utc.Month(), utc.Day(), 0, 0, 0, 0, time.UTC)
}
