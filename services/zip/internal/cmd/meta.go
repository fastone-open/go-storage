package main

import (
	def "github.com/fastone-open/go-storage/definitions"
	"github.com/fastone-open/go-storage/types"
)

var Metadata = def.Metadata{
	Name:  "zip",
	Pairs: []def.Pair{},
	Infos: []def.Info{},
	Factory: []def.Pair{
		def.PairWorkDir,
	},
	Service: def.Service{},
	Storage: def.Storage{
		Features: types.StorageFeatures{},
	},
}
