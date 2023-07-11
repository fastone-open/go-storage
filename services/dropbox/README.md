[![Services Test Dropbox](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-dropbox.yml/badge.svg)](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-dropbox.yml)

# dropbox

[Dropbox](https://www.dropbox.com) service support for [go-storage](https://git.fastonetech.com/fastone/go-storage).

## Install

```go
go get git.fastonetech.com/fastone/go-storage/services/dropbox/v3
```

## Usage

```go
import (
	"log"

	_ "github.com/fastone-open/go-storage/services/dropbox/v3"
	"github.com/fastone-open/go-storage/services"
)

func main() {
	store, err := services.NewStoragerFromString("dropbox:///path/to/workdir?credential=apikey:<apikey>")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://git.fastonetech.com/fastone/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/dropbox) about go-service-dropbox.
