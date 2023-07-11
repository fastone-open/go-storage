package main

import (
	"context"
	"github.com/fastone-open/go-storage/types"
)

type Service struct {
	defaultPairs types.DefaultServicePairs
	features     types.ServiceFeatures

	Pairs []types.Pair

	types.UnimplementedServicer
}

func (s *Service) delete(ctx context.Context, name string, opt pairServiceDelete) (err error) {
	panic("not implemented")
}
