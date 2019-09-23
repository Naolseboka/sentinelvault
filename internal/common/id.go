package common

import (
	"crypto/rand"
	"io"
	"time"

	"github.com/oklog/ulid/v2"
)

// NewID generates a new ULID string.
func NewID() string {
	return ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader).String()
}

// ParseID validates and parses a ULID string.
func ParseID(s string) (ulid.ULID, error) {
	return ulid.Parse(s)
}

// MustParseID parses a ULID or panics.
func MustParseID(s string) ulid.ULID {
	id, err := ulid.Parse(s)
	if err != nil {
		panic(err)
	}
	return id
}

// NewEntropy returns a cryptographically secure entropy source for ULIDs.
func NewEntropy() io.Reader {
	return rand.Reader
}
