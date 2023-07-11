package tests

import (
	"os"
	"testing"

	ps "github.com/fastone-open/go-storage/pairs"
	azblob "github.com/fastone-open/go-storage/services/azblob"
	"github.com/fastone-open/go-storage/types"
	"github.com/google/uuid"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for azblob")

	store, err := azblob.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_AZBLOB_CREDENTIAL")),
		ps.WithName(os.Getenv("STORAGE_AZBLOB_NAME")),
		ps.WithEndpoint(os.Getenv("STORAGE_AZBLOB_ENDPOINT")),
		ps.WithWorkDir("/"+uuid.New().String()+"/"),
		ps.WithEnableVirtualDir(),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
