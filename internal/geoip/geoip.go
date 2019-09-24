package geoip

import (
	"encoding/binary"
	"net"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type Range struct {
	Start uint32
	End   uint32
	CC    string
}

type DB struct {
	ranges []Range
}

func NewDB(ranges []Range) *DB {
	return &DB{ranges: append([]Range(nil), ranges...)}
}

func DefaultDB() *DB {
	return NewDB(defaultRanges())
}

func ParseIPv4(ip string) (uint32, error) {
	parsed := net.ParseIP(strings.TrimSpace(ip))
	if parsed == nil {
		return 0, common.New(common.ErrCodeInvalidArgument, "invalid ip")
	}
	v4 := parsed.To4()
	if v4 == nil {
		return 0, common.New(common.ErrCodeInvalidArgument, "ipv4 required")
	}
	return binary.BigEndian.Uint32(v4), nil
}

func (db *DB) Lookup(ip string) (string, error) {
	n, err := ParseIPv4(ip)
	if err != nil {
		return "", err
	}
	for _, r := range db.ranges {
		if n >= r.Start && n <= r.End {
			return r.CC, nil
		}
	}
	return "ZZ", nil
}

func (db *DB) SameCountry(ip1, ip2 string) (bool, error) {
	a, err := db.Lookup(ip1)
	if err != nil {
		return false, err
	}
	b, err := db.Lookup(ip2)
	if err != nil {
		return false, err
	}
	return a == b && a != "ZZ", nil
}
