package main

import (
	def "github.com/fastone-open/go-storage/definitions"
)

var Metadata = def.Metadata{
	Name:  "webdav",
	Pairs: []def.Pair{},
	Infos: []def.Info{},
	Factory: []def.Pair{
		def.PairWorkDir,
	},
	Service: def.Service{},
	Storage: def.Storage{},
}
