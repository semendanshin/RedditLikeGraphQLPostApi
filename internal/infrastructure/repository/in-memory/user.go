package in_memory

import (
	"GraphQLTestCase/internal/domain"
	"GraphQLTestCase/internal/usecases"
	"log/slog"
)

var _ usecases.UserRepository = &UserInMemoryRepository{}

type UserInMemoryRepository struct {
	AbstractInMemoryRepository[*domain.User]
}

func NewUserInMemoryRepository(logger *slog.Logger) *UserInMemoryRepository {
	return &UserInMemoryRepository{
		AbstractInMemoryRepository: NewAbstractInMemoryRepository[*domain.User](logger),
	}
}
