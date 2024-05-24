package usecases

import (
	"GraphQLTestCase/internal/domain"
	"context"
	"github.com/google/uuid"
)

//go:generate go run github.com/vektra/mockery/v2@v2.40.2 --name=PostUseCase
type PostUseCase interface {
	AbstractUseCaseInterface[*domain.Post]
	GetByAuthorID(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*domain.Post, error)
}
