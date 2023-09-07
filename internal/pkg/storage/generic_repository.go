package storage

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// GenericRepository is a in memory repository using generics
type GenericRepository[T any] interface {
	Get() ([]T, error)
	GetByID(uuid.UUID) (T, error)
	Add(T) uuid.UUID
	Update(uuid.UUID, T) error
}

type genericRepository[T any] struct {
	Items map[uuid.UUID]T
	m     sync.RWMutex
}

func NewGenericRepository[T any]() GenericRepository[T] {
	return &genericRepository[T]{
		Items: make(map[uuid.UUID]T),
	}
}

func (r *genericRepository[T]) Get() ([]T, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	items := make([]T, 0, len(r.Items))
	for _, val := range r.Items {
		items = append(items, val)
	}

	return items, nil
}

func (r *genericRepository[T]) GetByID(id uuid.UUID) (T, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	_, found := r.Items[id]

	if !found {
		return *new(T), fmt.Errorf("item with ID %s not found", id)
	}

	return r.Items[id], nil
}

func (r *genericRepository[T]) Add(item T) uuid.UUID {
	r.m.Lock()
	defer r.m.Unlock()

	id := uuid.New()
	r.Items[id] = item

	return id
}

func (r *genericRepository[T]) Update(id uuid.UUID, item T) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, found := r.Items[id]

	if !found {
		return fmt.Errorf("item with ID %s not found", id)
	}

	r.Items[id] = item

	return nil
}
