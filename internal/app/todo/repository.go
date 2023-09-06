package todo

import (
	"github.com/liviudnicoara/go-graphql-poc/internal/pkg/storage"
	"github.com/liviudnicoara/go-graphql-poc/shared"
)

type TodoRepository interface {
	storage.GenericRepository[shared.Todo]
}

func NewToDoRepository() TodoRepository {
	return storage.NewGenericRepository[shared.Todo]()
}
