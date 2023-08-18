package tests

import (
	"os"
	"testing"

	ps "github.com/fastone-open/go-storage/pairs"
	"github.com/fastone-open/go-storage/services"
	fs "github.com/fastone-open/go-storage/services/fs"
	"github.com/fastone-open/go-storage/tests"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	servicer, err := fs.NewServicer(ps.WithWorkDir("/tmp/hello"))
	require.NoError(t, err)
	store, err := servicer.Create("demo")
	require.NoError(t, err)
	t.Log(store)
}

func TestService2(t *testing.T) {
	servicer, err := services.NewServicerFromString("fs:///tmp/storage-store/gaia-bucket/dev?credential=hmac::&endpoint=")
	require.NoError(t, err)
	store, err := servicer.Create("demo")
	require.NoError(t, err)
	t.Log(store)

}

func TestStorage(t *testing.T) {
	if os.Getenv("STORAGE_FS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_FS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestStorager(t, setupTest(t))
}

func TestAppend(t *testing.T) {
	if os.Getenv("STORAGE_FS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_FS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestAppender(t, setupTest(t))
}

func TestDir(t *testing.T) {
	if os.Getenv("STORAGE_FS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_FS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestDirer(t, setupTest(t))
}

func TestCopy(t *testing.T) {
	if os.Getenv("STORAGE_FS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_FS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestCopier(t, setupTest(t))
	tests.TestCopierWithDir(t, setupTest(t))
}

func TestMove(t *testing.T) {
	if os.Getenv("STORAGE_FS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_FS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestMover(t, setupTest(t))
	tests.TestMoverWithDir(t, setupTest(t))
}

func TestLinker(t *testing.T) {
	if os.Getenv("STORAGE_FS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_FS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestLinker(t, setupTest(t))
}
