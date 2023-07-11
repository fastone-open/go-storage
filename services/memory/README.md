[![Services Test Memory](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-memory.yml/badge.svg)](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-memory.yml)

# memory

memory service support for [go-storage](https://git.fastonetech.com/fastone/go-storage).

## Install

```go
go get git.fastonetech.com/fastone/go-storage/services/memory
```

## Usage

```go
import (
	"log"

	_ "github.com/fastone-open/go-storage/services/memory"
	"github.com/fastone-open/go-storage/services"
)

func main() {
	store, err := services.NewStoragerFromString("memory:///path/to/workdir")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://git.fastonetech.com/fastone/go-storage-example).
