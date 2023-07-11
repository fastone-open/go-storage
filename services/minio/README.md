[![Services Test Minio](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-minio.yml/badge.svg)](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-minio.yml)

# minio

[MinIO](https://min.io/) is an open source cloud-native high-performance object storage service. 
This project will use minio's native SDK to implement [go-storage](https://git.fastonetech.com/fastone/go-storage/), 
enabling users to manipulate data on minio servers through a unified interface.

## Install

```go
go get git.fastonetech.com/fastone/go-storage/services/minio
```

## Usage

```go
import (
	"log"

	_ "github.com/fastone-open/go-storage/services/minio"
	"github.com/fastone-open/go-storage/services"
)

func main() {
	store, err := services.NewStoragerFromString("minio://<bucket_name>/<work_dir>?credential=hmac:<access_key>:<secret_key>&endpoint=https:<host>:<port>")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://git.fastonetech.com/fastone/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/minio) about go-service-minio.
