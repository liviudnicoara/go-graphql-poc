package graphql

import (
	"github.com/liviudnicoara/go-graphql-poc/internal/app/todo"
	"github.com/liviudnicoara/go-graphql-poc/internal/app/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserRepo user.UserRepository
	TodoRepo todo.TodoRepository
}
