package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"

	ps "github.com/fastone-open/go-storage/pairs"
	"github.com/fastone-open/go-storage/services/qingstor"
	"github.com/fastone-open/go-storage/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for qingstor")

	store, err := qingstor.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_QINGSTOR_CREDENTIAL")),
		ps.WithEndpoint(os.Getenv("STORAGE_QINGSTOR_ENDPOINT")),
		ps.WithName(os.Getenv("STORAGE_QINGSTOR_NAME")),
		ps.WithWorkDir("/"+uuid.New().String()+"/"),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
