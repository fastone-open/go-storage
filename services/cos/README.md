[![Services Test Cos](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-cos.yml/badge.svg)](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-cos.yml)

# cos

[COS(Cloud Object Storage)](https://cloud.tencent.com/product/cos) service support for [go-storage](https://git.fastonetech.com/fastone/go-storage).

## Install

```go
go get git.fastonetech.com/fastone/go-storage/services/cos/v3
```

## Usage

```go
import (
	"log"

	_ "github.com/fastone-open/go-storage/services/cos/v3"
	"github.com/fastone-open/go-storage/services"
)

func main() {
	store, err := services.NewStoragerFromString("cos://bucket_name/path/to/workdir?credential=hmac:<account_name>:<account_key>")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://git.fastonetech.com/fastone/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/cos) about go-service-cos.
