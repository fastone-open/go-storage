[![Services Test Kodo](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-kodo.yml/badge.svg)](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-kodo.yml)

# kodo

[qiniu kodo](https://www.qiniu.com/products/kodo) service support for [go-storage](https://git.fastonetech.com/fastone/go-storage).

## Install

```go
go get git.fastonetech.com/fastone/go-storage/services/kodo/v3
```

## Usage

```go
import (
	"log"

	_ "github.com/fastone-open/go-storage/services/kodo/v3"
	"github.com/fastone-open/go-storage/services"
)

func main() {
	store, err := services.NewStoragerFromString("kodo://bucket_name/path/to/workdir?credential=hmac:<access_key>:<secret_key>&endpoint=http:<domain>")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://git.fastonetech.com/fastone/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/kodo) about go-service-kodo.
