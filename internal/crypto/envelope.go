package crypto

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// Envelope holds encrypted data with wrapped DEK.
type Envelope struct {
	Version      int    `json:"version"`
	Algorithm    string `json:"algorithm"`
	Nonce        string `json:"nonce"`
	Ciphertext   string `json:"ciphertext"`
	WrappedDEK   string `json:"wrapped_dek"`
	WrapNonce    string `json:"wrap_nonce"`
	AssociatedID string `json:"associated_id,omitempty"`
}

// EnvelopeEncrypt encrypts plaintext using a random DEK, wrapping the DEK with kek.
func EnvelopeEncrypt(plaintext, kek, associatedData []byte) (*Envelope, error) {
	if len(kek) != AESKeySize {
		return nil, ErrInvalidKeySize
	}
	dek, err := RandomBytes(AESKeySize)
	if err != nil {
		return nil, err
	}
	dataAEAD, err := NewAEAD(dek)
	if err != nil {
		return nil, err
	}
	nonce, err := RandomBytes(dataAEAD.NonceSize())
	if err != nil {
		return nil, err
	}
	ciphertext, err := dataAEAD.Encrypt(nonce, plaintext, associatedData)
	if err != nil {
		return nil, err
	}

	kekAEAD, err := NewAEAD(kek)
	if err != nil {
		return nil, err
	}
	wrapNonce, err := RandomBytes(kekAEAD.NonceSize())
	if err != nil {
		return nil, err
	}
	wrappedDEK, err := kekAEAD.Encrypt(wrapNonce, dek, associatedData)
	if err != nil {
		return nil, err
	}

	return &Envelope{
		Version:    1,
		Algorithm:  "AES-256-GCM",
		Nonce:      base64.RawStdEncoding.EncodeToString(nonce),
		Ciphertext: base64.RawStdEncoding.EncodeToString(ciphertext),
		WrappedDEK: base64.RawStdEncoding.EncodeToString(wrappedDEK),
		WrapNonce:  base64.RawStdEncoding.EncodeToString(wrapNonce),
	}, nil
}

// EnvelopeDecrypt unwraps DEK and decrypts ciphertext.
func EnvelopeDecrypt(env *Envelope, kek, associatedData []byte) ([]byte, error) {
	if env == nil {
		return nil, fmt.Errorf("envelope is nil")
	}
	if env.Version != 1 || env.Algorithm != "AES-256-GCM" {
		return nil, fmt.Errorf("unsupported envelope version")
	}
	nonce, err := base64.RawStdEncoding.DecodeString(env.Nonce)
	if err != nil {
		return nil, err
	}
	ciphertext, err := base64.RawStdEncoding.DecodeString(env.Ciphertext)
	if err != nil {
		return nil, err
	}
	wrappedDEK, err := base64.RawStdEncoding.DecodeString(env.WrappedDEK)
	if err != nil {
		return nil, err
	}
	wrapNonce, err := base64.RawStdEncoding.DecodeString(env.WrapNonce)
	if err != nil {
		return nil, err
	}

	kekAEAD, err := NewAEAD(kek)
	if err != nil {
		return nil, err
	}
	dek, err := kekAEAD.Decrypt(wrapNonce, wrappedDEK, associatedData)
	if err != nil {
		return nil, err
	}
	dataAEAD, err := NewAEAD(dek)
	if err != nil {
		return nil, err
	}
	return dataAEAD.Decrypt(nonce, ciphertext, associatedData)
}

// MarshalJSONEnvelope serializes an envelope.
func MarshalJSONEnvelope(env *Envelope) ([]byte, error) {
	return json.Marshal(env)
}

// UnmarshalJSONEnvelope deserializes an envelope.
func UnmarshalJSONEnvelope(data []byte) (*Envelope, error) {
	var env Envelope
	if err := json.Unmarshal(data, &env); err != nil {
		return nil, err
	}
	return &env, nil
}
