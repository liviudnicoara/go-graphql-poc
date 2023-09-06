package user

import (
	"github.com/liviudnicoara/go-graphql-poc/internal/pkg/storage"
	"github.com/liviudnicoara/go-graphql-poc/shared"
)

type UserRepository interface {
	storage.GenericRepository[shared.User]
}

func NewUserRepository() UserRepository {
	return storage.NewGenericRepository[shared.User]()
}
