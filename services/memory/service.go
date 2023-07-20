package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/fastone-open/go-storage/types"
)

var _ types.Servicer = &BucketService{}

type BucketService struct {
	types.UnimplementedServicer
	sync.RWMutex
	buckets map[string]types.Storager
}

func NewBucketService() types.Servicer {
	return &BucketService{
		buckets: make(map[string]types.Storager),
	}
}

func (m *BucketService) String() string {
	return "BucketService"
}
func (m *BucketService) Create(name string, pairs ...types.Pair) (store types.Storager, err error) {
	m.Lock()
	defer m.Unlock()
	if _, found := m.buckets[name]; found {
		return nil, errors.New("bucket already exists")
	}
	sto, err := NewStorager(pairs...)
	m.buckets[name] = sto
	return sto, err
}
func (m *BucketService) CreateWithContext(ctx context.Context, name string, pairs ...types.Pair) (store types.Storager, err error) {
	return m.Create(name, pairs...)
}
func (m *BucketService) Delete(name string, pairs ...types.Pair) (err error) {
	m.Lock()
	defer m.Unlock()
	delete(m.buckets, name)
	return nil
}
func (m *BucketService) DeleteWithContext(ctx context.Context, name string, pairs ...types.Pair) (err error) {
	return m.Delete(name, pairs...)
}
func (m *BucketService) Get(name string, pairs ...types.Pair) (store types.Storager, err error) {
	m.RLock()
	defer m.RUnlock()
	if sto, found := m.buckets[name]; found {
		return sto, nil
	}
	return nil, errors.New("bucket not found")
}
func (m *BucketService) GetWithContext(ctx context.Context, name string, pairs ...types.Pair) (store types.Storager, err error) {
	return m.Get(name, pairs...)
}
