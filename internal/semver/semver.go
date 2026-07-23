package semver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type Version struct {
	Major int
	Minor int
	Patch int
	Pre   string
	Build string
}

func Parse(s string) (Version, error) {
	s = strings.TrimSpace(strings.TrimPrefix(s, "v"))
	if s == "" {
		return Version{}, common.New(common.ErrCodeInvalidArgument, "empty version")
	}
	build := ""
	if i := strings.IndexByte(s, '+'); i >= 0 {
		build = s[i+1:]
		s = s[:i]
	}
	pre := ""
	if i := strings.IndexByte(s, '-'); i >= 0 {
		pre = s[i+1:]
		s = s[:i]
	}
	parts := strings.Split(s, ".")
	if len(parts) < 1 || len(parts) > 3 {
		return Version{}, common.New(common.ErrCodeInvalidArgument, "invalid version")
	}
	nums := make([]int, 3)
	for i := 0; i < len(parts); i++ {
		n, err := strconv.Atoi(parts[i])
		if err != nil || n < 0 {
			return Version{}, common.New(common.ErrCodeInvalidArgument, "invalid version number")
		}
		nums[i] = n
	}
	return Version{Major: nums[0], Minor: nums[1], Patch: nums[2], Pre: pre, Build: build}, nil
}

func (v Version) String() string {
	s := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Pre != "" {
		s += "-" + v.Pre
	}
	if v.Build != "" {
		s += "+" + v.Build
	}
	return s
}

func Compare(a, b Version) int {
	if a.Major != b.Major {
		if a.Major < b.Major {
			return -1
		}
		return 1
	}
	if a.Minor != b.Minor {
		if a.Minor < b.Minor {
			return -1
		}
		return 1
	}
	if a.Patch != b.Patch {
		if a.Patch < b.Patch {
			return -1
		}
		return 1
	}
	if a.Pre == b.Pre {
		return 0
	}
	if a.Pre == "" {
		return 1
	}
	if b.Pre == "" {
		return -1
	}
	if a.Pre < b.Pre {
		return -1
	}
	if a.Pre > b.Pre {
		return 1
	}
	return 0
}

func (v Version) BumpMajor() Version { return Version{Major: v.Major + 1} }
func (v Version) BumpMinor() Version { return Version{Major: v.Major, Minor: v.Minor + 1} }
func (v Version) BumpPatch() Version {
	return Version{Major: v.Major, Minor: v.Minor, Patch: v.Patch + 1}
}
