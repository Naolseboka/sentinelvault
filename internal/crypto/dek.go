package crypto

// GenerateDEK returns a fresh 256-bit data encryption key.
func GenerateDEK() ([]byte, error) {
	return RandomBytes(AESKeySize)
}
