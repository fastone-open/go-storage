package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"

	ps "github.com/fastone-open/go-storage/pairs"
	gcs "github.com/fastone-open/go-storage/services/gcs"
	"github.com/fastone-open/go-storage/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for gcs")

	store, err := gcs.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_GCS_CREDENTIAL")),
		ps.WithName(os.Getenv("STORAGE_GCS_NAME")),
		ps.WithWorkDir("/"+uuid.New().String()+"/"),
		gcs.WithProjectID(os.Getenv("STORAGE_GCS_PROJECT_ID")),
		ps.WithEnableVirtualDir(),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
