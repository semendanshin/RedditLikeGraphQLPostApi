package usecases

import (
	"GraphQLTestCase/internal/domain"
	usecaseInterfaces "GraphQLTestCase/internal/interfaces/usecases"
	"context"
	"github.com/google/uuid"
	"time"
)

//go:generate go run github.com/vektra/mockery/v2@v2.40.2 --name=CommentRepository
type CommentRepository interface {
	usecaseInterfaces.AbstractRepositoryInterface[*domain.Comment]
	GetChildren(ctx context.Context, commentID uuid.UUID, limit int, offset int) ([]*domain.Comment, error)
	GetByPostID(ctx context.Context, postID uuid.UUID, limit int, offset int) ([]*domain.Comment, error)
	GetLastComment(ctx context.Context, postID uuid.UUID, lastSeen time.Time, limit int) ([]*domain.Comment, error)
}

var _ usecaseInterfaces.CommentUseCase = &CommentUseCase{}

type CommentUseCase struct {
	Repository CommentRepository
	usecaseInterfaces.AbstractUseCase[*domain.Comment]
}

func NewCommentUseCase(repository CommentRepository) *CommentUseCase {
	return &CommentUseCase{
		Repository:      repository,
		AbstractUseCase: usecaseInterfaces.NewAbstractUseCase[*domain.Comment](repository),
	}
}

func (uc *CommentUseCase) GetChildren(ctx context.Context, commentID uuid.UUID, limit int, offset int) ([]*domain.Comment, error) {
	return uc.Repository.GetChildren(ctx, commentID, limit, offset)
}

func (uc *CommentUseCase) GetByPostID(ctx context.Context, postID uuid.UUID, limit int, offset int) ([]*domain.Comment, error) {
	return uc.Repository.GetByPostID(ctx, postID, limit, offset)
}

func (uc *CommentUseCase) Create(ctx context.Context, entity *domain.Comment) error {
	if len(entity.Content) > 2000 {
		return domain.ErrCommentIsTooLong
	}
	entity.SetID(uuid.New())
	return uc.AbstractUseCase.Create(ctx, entity)
}

func (uc *CommentUseCase) GetLastComment(ctx context.Context, postID uuid.UUID, lastSeen time.Time, limit int) ([]*domain.Comment, error) {
	return uc.Repository.GetLastComment(ctx, postID, lastSeen, limit)
}
