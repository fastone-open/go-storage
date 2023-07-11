package cos

import "github.com/fastone-open/go-storage/services"

var (
	// ErrServerSideEncryptionCustomerKeyInvalid will be returned while server-side encryption customer key is invalid.
	ErrServerSideEncryptionCustomerKeyInvalid = services.NewErrorCode("invalid server-side encryption customer key")
)
