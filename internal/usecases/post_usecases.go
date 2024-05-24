package usecases

import (
	"GraphQLTestCase/internal/domain"
	usecaseInterfaces "GraphQLTestCase/internal/interfaces/usecases"
	"context"
	"github.com/google/uuid"
)

//go:generate go run github.com/vektra/mockery/v2@v2.40.2 --name=PostRepository
type PostRepository interface {
	usecaseInterfaces.AbstractRepositoryInterface[*domain.Post]
	GetByAuthorID(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*domain.Post, error)
}

var _ usecaseInterfaces.PostUseCase = &PostUseCase{}

type PostUseCase struct {
	Repository PostRepository
	usecaseInterfaces.AbstractUseCase[*domain.Post]
}

func NewPostUseCase(repository PostRepository) *PostUseCase {
	return &PostUseCase{
		Repository:      repository,
		AbstractUseCase: usecaseInterfaces.NewAbstractUseCase[*domain.Post](repository),
	}
}

func (uc *PostUseCase) GetByAuthorID(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*domain.Post, error) {
	return uc.Repository.GetByAuthorID(ctx, userID, limit, offset)
}
