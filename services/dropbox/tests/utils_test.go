package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"

	ps "github.com/fastone-open/go-storage/pairs"
	dropbox "github.com/fastone-open/go-storage/services/dropbox"
	"github.com/fastone-open/go-storage/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for dropbox")

	store, err := dropbox.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_DROPBOX_CREDENTIAL")),
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
