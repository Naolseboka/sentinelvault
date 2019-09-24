package encodingx

import (
	"encoding/base64"
	"encoding/hex"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

func HexEncode(b []byte) string { return hex.EncodeToString(b) }

func HexDecode(s string) ([]byte, error) {
	b, err := hex.DecodeString(strings.TrimSpace(s))
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "hex", err)
	}
	return b, nil
}

func Base64URLEncode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

func Base64URLDecode(s string) ([]byte, error) {
	b, err := base64.RawURLEncoding.DecodeString(strings.TrimSpace(s))
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "base64url", err)
	}
	return b, nil
}

func Base64StdEncode(b []byte) string { return base64.StdEncoding.EncodeToString(b) }

func Base64StdDecode(s string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(strings.TrimSpace(s))
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "base64", err)
	}
	return b, nil
}
