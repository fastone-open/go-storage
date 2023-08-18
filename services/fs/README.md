[![Services Test Fs](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-fs.yml/badge.svg)](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-fs.yml)

# fs

Local file system service support for [go-storage](https://git.fastonetech.com/fastone/go-storage).

## Install

```go
go get go.beyondstorage.io/services/fs/v4
```

## Usage

```go
import (
	"log"

	_ "go.beyondstorage.io/services/fs/v4"
	"github.com/fastone-open/go-storage/services"
)

func main() {
	store, err := services.NewStoragerFromString("fs:///path/to/workdir")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://git.fastonetech.com/fastone/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/fs) about go-service-fs.
