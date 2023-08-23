package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fastone-open/go-storage/services"
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

func TestService(t *testing.T) {
	servicer, err := memory.NewServicer()
	require.NoError(t, err)
	store, err := servicer.Create("demo")
	require.NoError(t, err)
	t.Log(store)
}

func TestService2(t *testing.T) {
	servicer, err := services.NewServicerFromString("memory:///mock/?credential=hmac::&endpoint=")
	require.NoError(t, err)
	store, err := servicer.Create("demo")
	require.NoError(t, err)
	t.Log(store)

}
