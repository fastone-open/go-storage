package tests

import (
	"os"
	"testing"

	onedrive "github.com/fastone-open/go-storage/services/onedrive"
	"github.com/google/uuid"

	ps "github.com/fastone-open/go-storage/pairs"
	"github.com/fastone-open/go-storage/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for onedrive")

	store, err := onedrive.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_ONEDRIVE_CREDENTIAL")),
		ps.WithWorkDir("/"+uuid.New().String()),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}

	t.Cleanup(func() {
		err = store.Delete("")
		if err != nil {
			t.Errorf("cleanup: %v", err)
		}
	})
	return store
}
