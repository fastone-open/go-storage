package azblob

import "github.com/fastone-open/go-storage/services"

var (
	// ErrEncryptionKeyInvalid will be returned while encryption key is invalid.
	// Encryption key must be a 32-byte AES-256 key.
	ErrEncryptionKeyInvalid = services.NewErrorCode("invalid encryption key")
)
