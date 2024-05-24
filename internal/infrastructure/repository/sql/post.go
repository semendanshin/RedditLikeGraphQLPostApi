package sql

import (
	"GraphQLTestCase/internal/domain"
	"GraphQLTestCase/internal/infrastructure/entities"
	"GraphQLTestCase/internal/usecases"
	"GraphQLTestCase/internal/utils/mappers"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log/slog"
)

var _ usecases.PostRepository = &PostSQLRepository{}

type PostSQLRepository struct {
	AbstractSQLRepository[*domain.Post, entities.Post]
}

func NewPostSQLRepository(db *gorm.DB, logger *slog.Logger) *PostSQLRepository {
	return &PostSQLRepository{
		AbstractSQLRepository: NewAbstractSQLRepository[*domain.Post, entities.Post](
			db, logger, mappers.DomainToEntityPost, mappers.EntityToDomainPost,
		),
	}
}

func (r *PostSQLRepository) GetByAuthorID(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*domain.Post, error) {
	const op = "PostSQLRepository.GetByAuthorID"
	var posts []*domain.Post
	var postEntities []*entities.Post
	if err := r.db.WithContext(ctx).Where("author_id = ?", userID).Limit(limit).Offset(offset).Find(&postEntities).Error; err != nil {
		r.logger.Error(op, slog.Any("error", err.Error()))
		return nil, err
	}
	for _, entity := range postEntities {
		posts = append(posts, r.entityToModel(entity))
	}
	return posts, nil

}
