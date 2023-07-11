[![Services Test Bos](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-bos.yml/badge.svg)](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-bos.yml)

# bos

BOS(Baidu Object Storage) service support for [go-storage](https://git.fastonetech.com/fastone/go-storage).

## Install

```go
go get git.fastonetech.com/fastone/go-storage/services/bos/v2
```

## Usage

```go
import (
	"log"

	_ "github.com/fastone-open/go-storage/services/bos/v2"
	"github.com/fastone-open/go-storage/services"
)

func main() {
	store, err := services.NewStoragerFromString("bos://bucket_name/path/to/workdir")
	if err != nil {
		log.Fatal(err)
	}

	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://git.fastonetech.com/fastone/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/bos) about go-service-bos.
