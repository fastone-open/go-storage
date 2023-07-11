package tests

import (
	"testing"

	ps "github.com/fastone-open/go-storage/pairs"
	fs "github.com/fastone-open/go-storage/services/fs"
	"github.com/fastone-open/go-storage/types"
)

func setupTest(t *testing.T) types.Storager {
	tmpDir := t.TempDir()
	t.Logf("Setup test at %s", tmpDir)

	store, err := fs.NewStorager(ps.WithWorkDir(tmpDir))
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
