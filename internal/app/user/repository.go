package user

import (
	"github.com/liviudnicoara/go-graphql-poc/internal/pkg/storage"
	"github.com/liviudnicoara/go-graphql-poc/shared"
)

// UserRepository defines the interface for user-related database operations.
type UserRepository interface {
	storage.GenericRepository[shared.User]
}

// NewUserRepository creates a new UserRepository instance using the shared.User type.
func NewUserRepository() UserRepository {
	return storage.NewGenericRepository[shared.User]()
}
