package in_memory

import (
	"GraphQLTestCase/internal/domain"
	"GraphQLTestCase/internal/usecases"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

var _ usecases.PostRepository = &PostInMemoryRepository{}

type PostInMemoryRepository struct {
	AbstractInMemoryRepository[*domain.Post]
}

func NewPostInMemoryRepository(logger *slog.Logger) *PostInMemoryRepository {
	return &PostInMemoryRepository{
		AbstractInMemoryRepository: NewAbstractInMemoryRepository[*domain.Post](logger),
	}
}

func (r *PostInMemoryRepository) GetByAuthorID(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*domain.Post, error) {
	r.m.Lock()
	defer r.m.Unlock()

	var posts []*domain.Post

	var keys []uuid.UUID
	for key := range r.entities {
		keys = append(keys, key)
	}

	i := offset
	for i < len(keys) && len(posts) < limit {
		if r.entities[keys[i]].AuthorID == userID {
			posts = append(posts, r.entities[keys[i]])
		}
		i++
	}

	return posts, nil
}
