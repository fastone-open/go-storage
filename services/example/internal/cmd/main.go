package main

import (
	def "github.com/fastone-open/go-storage/definitions"
)

func main() {
	def.GenerateService(Metadata, "generated.go")
}
