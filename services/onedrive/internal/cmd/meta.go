package main

import (
	def "github.com/fastone-open/go-storage/definitions"
	"github.com/fastone-open/go-storage/types"
)

var Metadata = def.Metadata{
	Name: "onedrive",
	Pairs: []def.Pair{
		pairDescription,
	},
	Infos: []def.Info{},
	Factory: []def.Pair{
		def.PairCredential,
		def.PairWorkDir,
	},
	Service: def.Service{},
	Storage: def.Storage{
		Features: types.StorageFeatures{
			Create:   true,
			Delete:   true,
			List:     true,
			Metadata: true,
			Read:     true,
			Stat:     true,
			Write:    true,
		},

		Create: []def.Pair{
			def.PairObjectMode,
		},
		Delete: []def.Pair{
			def.PairObjectMode,
		},
		List: []def.Pair{
			def.PairListMode,
			def.PairContinuationToken,
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
			pairDescription,
		},
		Stat: []def.Pair{
			def.PairObjectMode,
		},
	},
}

var pairDescription = def.Pair{
	Name:        "description",
	Type:        def.Type{Name: "string"},
	Description: "description for target file/dir.",
}
