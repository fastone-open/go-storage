package fs

import (
	"context"
	"os"
	"path"

	"github.com/qingstor/qingstor-sdk-go/v4/service"

	ps "github.com/fastone-open/go-storage/pairs"
	"github.com/fastone-open/go-storage/services"
	typ "github.com/fastone-open/go-storage/types"
)

func (s *Service) create(ctx context.Context, name string, opt pairServiceCreate) (store typ.Storager, err error) {
	// ServicePairCreate requires location, so we don't need to add location into pairs
	pairs := append(opt.pairs, ps.WithName(name))

	st, err := s.newStorage(pairs...)
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

func (s *Service) get(ctx context.Context, name string, opt pairServiceGet) (store typ.Storager, err error) {
	pairs := append(opt.pairs, ps.WithName(name))

	store, err = s.newStorage(pairs...)
	if err != nil {
		return
	}
	return
}

func (s *Service) list(ctx context.Context, opt pairServiceList) (it *typ.StoragerIterator, err error) {
	input := &storagePageStatus{}

	if opt.HasLocation {
		input.location = opt.Location
	}

	return typ.NewStoragerIterator(ctx, s.nextStoragePage, input), nil
}

func (s *Service) newStorage(pairs ...typ.Pair) (store *Storage, err error) {
	opt, err := parsePairStorageNew(pairs)
	if err != nil {
		return
	}

	s.Lock()
	defer s.Unlock()
	store, ok := s.buckets[opt.Name]
	if ok {
		return store, nil
	}

	store = &Storage{
		workDir: "/",
	}
	workDir, err := evalSymlinks(path.Join(s.workDir, opt.Name))
	if err != nil {
		return nil, err
	}
	store.workDir = workDir
	s.buckets[opt.Name] = store

	if opt.HasDefaultStoragePairs {
		store.defaultPairs = opt.DefaultStoragePairs
	}
	if opt.HasStorageFeatures {
		store.features = opt.StorageFeatures
	}
	if opt.HasWorkDir {
		workDir, err := evalSymlinks(opt.WorkDir)
		if err != nil {
			return nil, err
		}
		store.workDir = workDir
	}

	// Check and create work dir
	err = os.MkdirAll(store.workDir, 0755)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Service) nextStoragePage(ctx context.Context, page *typ.StoragerPage) error {
	input := page.Status.(*storagePageStatus)

	serviceInput := &service.ListBucketsInput{
		Limit:  &input.offset,
		Offset: &input.limit,
	}
	if input.location != "" {
		serviceInput.Location = &input.location
	}
	for name := range s.buckets {
		store, err := s.newStorage(ps.WithName(name))
		if err != nil {
			return err
		}
		page.Data = append(page.Data, store)
	}

	return typ.IterateDone

}
