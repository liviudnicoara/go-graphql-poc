package storage_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/liviudnicoara/go-graphql-poc/internal/pkg/storage"
)

func TestGenericRepositoryGet(t *testing.T) {
	repo := storage.NewGenericRepository[int]()
	repo.Add(1)
	repo.Add(2)
	repo.Add(3)

	items, err := repo.Get()
	if err != nil {
		t.Errorf("Expected no error; got %v", err)
	}

	if len(items) != 3 {
		t.Errorf("Expected 3 items; got %d", len(items))
	}
}

func TestGenericRepositoryGetByID(t *testing.T) {
	repo := storage.NewGenericRepository[int]()
	validID := repo.Add(1)
	repo.Add(2)
	repo.Add(3)

	id := uuid.New()
	_, err := repo.GetByID(id)
	if err == nil {
		t.Errorf("Expected an error for non-existing ID; got nil")
	}

	item, err := repo.GetByID(validID)
	if err != nil {
		t.Errorf("Expected no error; got %v", err)
	}

	if item != 1 {
		t.Errorf("Expected item 1; got %v", item)
	}
}

func TestGenericRepositoryAdd(t *testing.T) {
	repo := storage.NewGenericRepository[int]()
	repo.Add(1)

	items, _ := repo.Get()
	if len(items) != 1 {
		t.Errorf("Expected 1 item; got %d", len(items))
	}
}

func TestGenericRepositoryUpdate(t *testing.T) {
	repo := storage.NewGenericRepository[int]()
	id := repo.Add(1)

	err := repo.Update(uuid.New(), 2)
	if err == nil {
		t.Errorf("Expected an error for non-existing ID; got nil")
	}

	err = repo.Update(id, 3)
	if err != nil {
		t.Errorf("Expected no error; got %v", err)
	}

	items, _ := repo.Get()
	if len(items) != 1 || items[0] != 3 {
		t.Errorf("Expected item 3; got %v", items[0])
	}
}
