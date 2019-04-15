package crypto

import (
	"encoding/binary"
	"fmt"
)

const blobMagic = "SVEN"

// CiphertextBlob is the on-wire encrypted payload: nonce | wrapped DEK | ciphertext.
type CiphertextBlob struct {
	Version    byte
	KeyID      string
	Nonce      []byte
	WrapNonce  []byte
	WrappedDEK []byte
	Ciphertext []byte
}

// MarshalBlob serializes a ciphertext blob.
func MarshalBlob(b *CiphertextBlob) ([]byte, error) {
	if b == nil {
		return nil, fmt.Errorf("blob is nil")
	}
	keyID := []byte(b.KeyID)
	if len(keyID) > 65535 {
		return nil, fmt.Errorf("key id too long")
	}
	if len(b.WrappedDEK) > 65535 {
		return nil, fmt.Errorf("wrapped dek too long")
	}

	size := 4 + 1 + 2 + len(keyID) + len(b.Nonce) + len(b.WrapNonce) + 2 + len(b.WrappedDEK) + len(b.Ciphertext)
	out := make([]byte, 0, size)
	out = append(out, blobMagic...)
	out = append(out, b.Version)
	out = binary.BigEndian.AppendUint16(out, uint16(len(keyID)))
	out = append(out, keyID...)
	out = append(out, b.Nonce...)
	out = append(out, b.WrapNonce...)
	out = binary.BigEndian.AppendUint16(out, uint16(len(b.WrappedDEK)))
	out = append(out, b.WrappedDEK...)
	out = append(out, b.Ciphertext...)
	return out, nil
}

// UnmarshalBlob parses a serialized ciphertext blob.
func UnmarshalBlob(data []byte) (*CiphertextBlob, error) {
	if len(data) < 4+1+2+12+12+2 {
		return nil, fmt.Errorf("blob too short")
	}
	if string(data[:4]) != blobMagic {
		return nil, fmt.Errorf("invalid blob magic")
	}
	off := 4
	version := data[off]
	off++

	keyLen := int(binary.BigEndian.Uint16(data[off : off+2]))
	off += 2
	if off+keyLen > len(data) {
		return nil, fmt.Errorf("truncated key id")
	}
	keyID := string(data[off : off+keyLen])
	off += keyLen

	if off+12+12+2 > len(data) {
		return nil, fmt.Errorf("truncated header")
	}
	nonce := data[off : off+12]
	off += 12
	wrapNonce := data[off : off+12]
	off += 12

	wrappedLen := int(binary.BigEndian.Uint16(data[off : off+2]))
	off += 2
	if off+wrappedLen > len(data) {
		return nil, fmt.Errorf("truncated wrapped dek")
	}
	wrappedDEK := data[off : off+wrappedLen]
	off += wrappedLen

	return &CiphertextBlob{
		Version:    version,
		KeyID:      keyID,
		Nonce:      nonce,
		WrapNonce:  wrapNonce,
		WrappedDEK: wrappedDEK,
		Ciphertext: data[off:],
	}, nil
}
