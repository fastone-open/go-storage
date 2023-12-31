package main

import (
	def "github.com/fastone-open/go-storage/definitions"
	"github.com/fastone-open/go-storage/types"
)

var Metadata = def.Metadata{
	Name:  "memory",
	Pairs: []def.Pair{},
	Infos: []def.Info{},
	Factory: []def.Pair{
		def.PairWorkDir,
		def.PairName,
	},
	Service: def.Service{
		Features: types.ServiceFeatures{
			Create: true,
			Delete: true,
			Get:    true,
			List:   true,
		},
	},
	Storage: def.Storage{
		Features: types.StorageFeatures{
			WriteEmptyObject: true,

			Create:       true,
			CreateAppend: true,
			CreateDir:    true,
			CommitAppend: true,
			Copy:         true,
			Delete:       true,
			List:         true,
			Metadata:     true,
			Move:         true,
			Read:         true,
			Stat:         true,
			Write:        true,
			WriteAppend:  true,
		},

		Create: []def.Pair{
			def.PairObjectMode,
		},
		Delete: []def.Pair{
			def.PairObjectMode,
		},
		List: []def.Pair{
			def.PairListMode,
		},
		Read: []def.Pair{
			def.PairOffset,
			def.PairIoCallback,
			def.PairSize,
		},
		Write: []def.Pair{
			def.PairContentMD5,
			def.PairContentType,
			def.PairIoCallback,
		},
		Stat: []def.Pair{
			def.PairObjectMode,
		},
	},
}
