package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"

	ps "github.com/fastone-open/go-storage/pairs"
	minio "github.com/fastone-open/go-storage/services/minio"
	"github.com/fastone-open/go-storage/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for minio")

	store, err := minio.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_MINIO_CREDENTIAL")),
		ps.WithEndpoint(os.Getenv("STORAGE_MINIO_ENDPOINT")),
		ps.WithName(os.Getenv("STORAGE_MINIO_NAME")),
		ps.WithWorkDir("/"+uuid.New().String()),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
