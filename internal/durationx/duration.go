package durationx

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

func Parse(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, common.New(common.ErrCodeInvalidArgument, "empty duration")
	}
	if d, err := time.ParseDuration(s); err == nil {
		return d, nil
	}
	// support days: 7d
	if strings.HasSuffix(s, "d") {
		n, err := strconv.Atoi(strings.TrimSuffix(s, "d"))
		if err != nil {
			return 0, common.New(common.ErrCodeInvalidArgument, "invalid days")
		}
		return time.Duration(n) * 24 * time.Hour, nil
	}
	if strings.HasSuffix(s, "w") {
		n, err := strconv.Atoi(strings.TrimSuffix(s, "w"))
		if err != nil {
			return 0, common.New(common.ErrCodeInvalidArgument, "invalid weeks")
		}
		return time.Duration(n) * 7 * 24 * time.Hour, nil
	}
	return 0, common.New(common.ErrCodeInvalidArgument, fmt.Sprintf("invalid duration %q", s))
}

func Humanize(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("%dh", int(d.Hours()))
	}
	return fmt.Sprintf("%dd", int(d.Hours()/24))
}
