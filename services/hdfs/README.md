[![Services Test Hdfs](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-hdfs.yml/badge.svg)](https://git.fastonetech.com/fastone/go-storage/actions/workflows/services-test-hdfs.yml)

# hdfs 

Hadoop Distributed File System (HDFS) support for [go-storage](https://git.fastonetech.com/fastone/go-storage).

## Install

```go
go get go.beyondstorage.io/services/hdfs
```

## Usage

```go
import (
	"log"
	
	_ "go.beyondstorage.io/services/hdfs"
	"github.com/fastone-open/go-storage/services"
)

func main() {
	store, err := services.NewStoragerFromString("hdfs:///path/to/workdir?endpoint=tcp:<host>:<port>")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://git.fastonetech.com/fastone/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/hdfs) about go-service-hdfs.