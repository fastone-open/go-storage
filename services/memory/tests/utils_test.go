package tests

import (
	"testing"

	"github.com/fastone-open/go-storage/types"

	"github.com/fastone-open/go-storage/services/memory"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for memory")

	store, err := memory.NewStorager()
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
