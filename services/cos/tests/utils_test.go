package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"

	ps "github.com/fastone-open/go-storage/pairs"
	cos "github.com/fastone-open/go-storage/services/cos"
	"github.com/fastone-open/go-storage/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for cos")

	store, err := cos.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_COS_CREDENTIAL")),
		ps.WithName(os.Getenv("STORAGE_COS_NAME")),
		ps.WithLocation(os.Getenv("STORAGE_COS_LOCATION")),
		ps.WithWorkDir("/"+uuid.New().String()+"/"),
		ps.WithEnableVirtualDir(),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
