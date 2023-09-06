package graph

import (
	"github.com/liviudnicoara/go-graphql-poc/graph/model"
	"github.com/liviudnicoara/go-graphql-poc/repositories"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserRepository repositories.BaseRepository[model.User]
	TodoRepository repositories.BaseRepository[model.Todo]
}
