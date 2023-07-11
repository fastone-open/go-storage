package tests

import (
	"os"
	"testing"

	ps "github.com/fastone-open/go-storage/pairs"
	"github.com/fastone-open/go-storage/services"
	_ "github.com/fastone-open/go-storage/services/ftp"
	"github.com/fastone-open/go-storage/types"
)

func initTest(t *testing.T) (store types.Storager) {
	t.Log("Setup test for ftp")

	store, err := services.NewStorager("ftp",
		ps.WithCredential(os.Getenv("STORAGE_FTP_CREDENTIAL")),
		ps.WithEndpoint(os.Getenv("STORAGE_FTP_ENDPOINT")),
	)
	if err != nil {
		t.Errorf("create storager: %v", err)
	}

	return
}
