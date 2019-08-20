package transit

import (
	"encoding/base64"
	"sync"

	"github.com/sentinelvault/sentinelvault/internal/common"
	svcrypto "github.com/sentinelvault/sentinelvault/internal/crypto"
)

// KeyType for transit keys.
const (
	KeyTypeAES256 = "aes256-gcm96"
	KeyTypeEd25519 = "ed25519"
)

// KeyVersion holds key material.
type KeyVersion struct {
	Version int    `json:"version"`
	Key     []byte `json:"-"`
}

// Key is a named transit keyring.
type Key struct {
	Name     string       `json:"name"`
	OrgID    string       `json:"org_id"`
	Type     string       `json:"type"`
	Versions []KeyVersion `json:"versions"`
	Latest   int          `json:"latest"`
}

// Engine provides encrypt/decrypt/sign/verify (Vault transit-like).
type Engine struct {
	policies map[string]*PolicyDocument
	mu   sync.Mutex
	keys map[string]*Key
}

// NewEngine creates a transit engine.
func NewEngine() *Engine {
	return &Engine{policies: map[string]*PolicyDocument{}, keys: map[string]*Key{}}
}

func keyID(orgID, name string) string { return orgID + "/" + name }

// CreateKey creates a new named key.
func (e *Engine) CreateKey(orgID, name, keyType string) (*Key, error) {
	if orgID == "" || name == "" {
		return nil, common.New(common.ErrCodeInvalidArgument, "org_id and name required")
	}
	if keyType == "" {
		keyType = KeyTypeAES256
	}
	id := keyID(orgID, name)
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, exists := e.keys[id]; exists {
		return nil, common.New(common.ErrCodeConflict, "key exists")
	}
	var material []byte
	var err error
	switch keyType {
	case KeyTypeAES256:
		material, err = svcrypto.RandomBytes(svcrypto.AESKeySize)
	case KeyTypeEd25519:
		_, priv, e2 := svcrypto.GenerateEd25519KeyPair()
		err = e2
		material = priv
	default:
		return nil, common.New(common.ErrCodeInvalidArgument, "unsupported key type")
	}
	if err != nil {
		return nil, err
	}
	k := &Key{
		Name: name, OrgID: orgID, Type: keyType, Latest: 1,
		Versions: []KeyVersion{{Version: 1, Key: material}},
	}
	e.keys[id] = k
	return publicKey(k), nil
}

// Rotate appends a new key version.
func (e *Engine) Rotate(orgID, name string) (*Key, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	k, ok := e.keys[keyID(orgID, name)]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "key not found")
	}
	var material []byte
	var err error
	switch k.Type {
	case KeyTypeAES256:
		material, err = svcrypto.RandomBytes(svcrypto.AESKeySize)
	case KeyTypeEd25519:
		_, priv, e2 := svcrypto.GenerateEd25519KeyPair()
		err = e2
		material = priv
	}
	if err != nil {
		return nil, err
	}
	k.Latest++
	k.Versions = append(k.Versions, KeyVersion{Version: k.Latest, Key: material})
	return publicKey(k), nil
}

// Encrypt encrypts plaintext with latest AES key.
func (e *Engine) Encrypt(orgID, name string, plaintext []byte) (string, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	k, ok := e.keys[keyID(orgID, name)]
	if !ok {
		return "", common.New(common.ErrCodeNotFound, "key not found")
	}
	if k.Type != KeyTypeAES256 {
		return "", common.New(common.ErrCodeInvalidArgument, "key is not AES")
	}
	key := k.Versions[len(k.Versions)-1].Key
	aead, err := svcrypto.NewAEAD(key)
	if err != nil {
		return "", err
	}
	nonce, err := svcrypto.RandomBytes(aead.NonceSize())
	if err != nil {
		return "", err
	}
	ct, err := aead.Encrypt(nonce, plaintext, []byte(orgID+"/"+name))
	if err != nil {
		return "", err
	}
	payload := append(nonce, ct...)
	return "sv:v" + itoa(k.Latest) + ":" + base64.RawStdEncoding.EncodeToString(payload), nil
}

// Decrypt decrypts a transit ciphertext.
func (e *Engine) Decrypt(orgID, name, ciphertext string) ([]byte, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	k, ok := e.keys[keyID(orgID, name)]
	if !ok {
		return nil, common.New(common.ErrCodeNotFound, "key not found")
	}
	ver, payload, err := parseCipher(ciphertext)
	if err != nil {
		return nil, err
	}
	var key []byte
	for _, v := range k.Versions {
		if v.Version == ver {
			key = v.Key
			break
		}
	}
	if key == nil {
		return nil, common.New(common.ErrCodeNotFound, "key version not found")
	}
	aead, err := svcrypto.NewAEAD(key)
	if err != nil {
		return nil, err
	}
	ns := aead.NonceSize()
	if len(payload) < ns {
		return nil, common.New(common.ErrCodeInvalidArgument, "ciphertext too short")
	}
	return aead.Decrypt(payload[:ns], payload[ns:], []byte(orgID+"/"+name))
}

func publicKey(k *Key) *Key {
	cp := *k
	cp.Versions = make([]KeyVersion, len(k.Versions))
	for i, v := range k.Versions {
		cp.Versions[i] = KeyVersion{Version: v.Version}
	}
	return &cp
}

func parseCipher(s string) (int, []byte, error) {
	// format sv:vN:base64
	if len(s) < 6 || s[:4] != "sv:v" {
		return 0, nil, common.New(common.ErrCodeInvalidArgument, "invalid ciphertext")
	}
	rest := s[4:]
	i := 0
	for i < len(rest) && rest[i] != ':' {
		i++
	}
	if i == 0 || i >= len(rest)-1 {
		return 0, nil, common.New(common.ErrCodeInvalidArgument, "invalid ciphertext")
	}
	ver := 0
	for _, c := range rest[:i] {
		if c < '0' || c > '9' {
			return 0, nil, common.New(common.ErrCodeInvalidArgument, "invalid version")
		}
		ver = ver*10 + int(c-'0')
	}
	raw, err := base64.RawStdEncoding.DecodeString(rest[i+1:])
	if err != nil {
		return 0, nil, err
	}
	return ver, raw, nil
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b [16]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	return string(b[i:])
}
