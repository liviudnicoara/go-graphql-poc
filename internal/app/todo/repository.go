// Package todo provides functionality related to todo management.

package todo

import (
	"github.com/liviudnicoara/go-graphql-poc/internal/pkg/storage"
	"github.com/liviudnicoara/go-graphql-poc/shared"
)

// TodoRepository defines the interface for todo-related database operations.
type TodoRepository interface {
	storage.GenericRepository[shared.Todo]
}

// NewTodoRepository creates a new TodoRepository instance using the shared.Todo type.
func NewTodoRepository() TodoRepository {
	return storage.NewGenericRepository[shared.Todo]()
}
