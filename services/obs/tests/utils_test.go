package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"

	ps "github.com/fastone-open/go-storage/pairs"
	obs "github.com/fastone-open/go-storage/services/obs"
	"github.com/fastone-open/go-storage/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for obs")

	store, err := obs.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_OBS_CREDENTIAL")),
		ps.WithName(os.Getenv("STORAGE_OBS_NAME")),
		ps.WithEndpoint(os.Getenv("STORAGE_OBS_ENDPOINT")),
		ps.WithWorkDir("/"+uuid.New().String()+"/"),
		ps.WithEnableVirtualDir(),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}

	return store
}
