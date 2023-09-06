package repositories

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type BaseRepository[T any] interface {
	Get() ([]T, error)
	GetByID(uuid.UUID) (T, error)
	Add(T) error
	Update(uuid.UUID, T) error
}

type baseRepository[T any] struct {
	Items map[uuid.UUID]T
	m     sync.RWMutex
}

func NewBaseRepository[T any]() BaseRepository[T] {
	return &baseRepository[T]{
		Items: make(map[uuid.UUID]T),
	}
}

func (r *baseRepository[T]) Get() ([]T, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	items := make([]T, 0, len(r.Items))
	for _, val := range r.Items {
		items = append(items, val)
	}

	return items, nil
}

func (r *baseRepository[T]) GetByID(id uuid.UUID) (T, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	_, found := r.Items[id]

	if !found {
		return *new(T), fmt.Errorf("Item with ID %s not found", id)
	}

	return r.Items[id], nil
}

func (r *baseRepository[T]) Add(item T) error {
	r.m.Lock()
	defer r.m.Unlock()

	r.Items[uuid.New()] = item

	return nil
}

func (r *baseRepository[T]) Update(id uuid.UUID, item T) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, found := r.Items[id]

	if !found {
		return fmt.Errorf("Item with ID %s not found", id)
	}

	r.Items[id] = item

	return nil
}
