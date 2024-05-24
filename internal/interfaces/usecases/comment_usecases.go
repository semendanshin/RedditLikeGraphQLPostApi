package usecases

import (
	"GraphQLTestCase/internal/domain"
	"context"
	"github.com/google/uuid"
	"time"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=CommentUseCase
type CommentUseCase interface {
	AbstractUseCaseInterface[*domain.Comment]
	GetChildren(ctx context.Context, commentID uuid.UUID, limit int, offset int) ([]*domain.Comment, error)
	GetByPostID(ctx context.Context, postID uuid.UUID, limit int, offset int) ([]*domain.Comment, error)
	GetLastComment(ctx context.Context, postID uuid.UUID, lastSeen time.Time, limit int) ([]*domain.Comment, error)
}
