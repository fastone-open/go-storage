package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"

	ps "github.com/fastone-open/go-storage/pairs"
	storj "github.com/fastone-open/go-storage/services/storj"
	"github.com/fastone-open/go-storage/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for STORJ")

	store, err := storj.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_STORJ_CREDENTIAL")),
		ps.WithName(os.Getenv("STORAGE_STORJ_NAME")),
		ps.WithLocation(os.Getenv("STORAGE_STORJ_LOCATION")),
		ps.WithWorkDir("/"+uuid.New().String()+"/"),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
