package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"

	"github.com/fastone-open/go-storage/pairs"
	hdfs "github.com/fastone-open/go-storage/services/hdfs"
	"github.com/fastone-open/go-storage/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for HDFS")

	store, err := hdfs.NewStorager(
		pairs.WithEndpoint(os.Getenv("STORAGE_HDFS_ENDPOINT")),
		pairs.WithWorkDir("/"+uuid.New().String()+"/"),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}

	return store
}
