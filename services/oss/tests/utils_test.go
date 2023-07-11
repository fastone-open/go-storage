package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"

	ps "github.com/fastone-open/go-storage/pairs"
	oss "github.com/fastone-open/go-storage/services/oss"
	"github.com/fastone-open/go-storage/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for oss")

	store, err := oss.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_OSS_CREDENTIAL")),
		ps.WithName(os.Getenv("STORAGE_OSS_NAME")),
		ps.WithEndpoint(os.Getenv("STORAGE_OSS_ENDPOINT")),
		ps.WithWorkDir("/"+uuid.New().String()+"/"),
		ps.WithEnableVirtualDir(),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
