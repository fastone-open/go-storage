package memory

import (
	"context"
	"path"

	"github.com/fastone-open/go-storage/services"
	"github.com/fastone-open/go-storage/types"
)

func (s *Service) create(ctx context.Context, name string, opt pairServiceCreate) (store types.Storager, err error) {
	// ServicePairCreate requires location, so we don't need to add location into pairs
	// pairs := append(opt.pairs, ps.WithName(name))
	st, err := s.newStorage(name)
	if err != nil {
		return
	}

	s.Lock()
	defer s.Unlock()
	s.buckets[name] = st
	return st, nil
}

func (s *Service) delete(ctx context.Context, name string, opt pairServiceDelete) (err error) {
	s.Lock()
	defer s.Unlock()
	delete(s.buckets, name)
	return nil
}

func (s *Service) formatError(op string, err error, name string) error {
	if err == nil {
		return nil
	}

	return services.ServiceError{
		Op:       op,
		Err:      formatError(err),
		Servicer: s,
		Name:     name,
	}
}

func (s *Service) get(ctx context.Context, name string, opt pairServiceGet) (store types.Storager, err error) {
	store, err = s.newStorage(name)
	if err != nil {
		return
	}
	return
}

func (s *Service) list(ctx context.Context, opt pairServiceList) (it *types.StoragerIterator, err error) {
	return types.NewStoragerIterator(ctx, s.nextStoragePage, nil), nil
}

func (s *Service) newStorage(name string) (store *Storage, err error) {
	s.Lock()
	defer s.Unlock()
	store, ok := s.buckets[name]
	if ok {
		return store, nil
	}

	store = &Storage{
		f:       s.f,
		workDir: path.Join(s.f.WorkDir, name),
	}
	return store, nil
}

func (s *Service) nextStoragePage(ctx context.Context, page *types.StoragerPage) error {
	for name := range s.buckets {
		store, err := s.newStorage(name)
		if err != nil {
			return err
		}
		page.Data = append(page.Data, store)
	}

	return types.IterateDone
}
