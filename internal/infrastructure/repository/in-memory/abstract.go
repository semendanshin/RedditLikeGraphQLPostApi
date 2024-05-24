package in_memory

import (
	"GraphQLTestCase/internal/domain"
	"GraphQLTestCase/internal/interfaces/usecases"
	"context"
	"github.com/google/uuid"
	"log/slog"
	"sync"
)

var _ usecases.AbstractRepositoryInterface[domain.Model] = &AbstractInMemoryRepository[domain.Model]{}

type AbstractInMemoryRepository[T domain.Model] struct {
	entities map[uuid.UUID]T
	m        sync.Mutex
	logger   *slog.Logger
}

func NewAbstractInMemoryRepository[T domain.Model](logger *slog.Logger) AbstractInMemoryRepository[T] {
	return AbstractInMemoryRepository[T]{
		entities: make(map[uuid.UUID]T),
		m:        sync.Mutex{},
		logger:   logger,
	}
}

func (r *AbstractInMemoryRepository[T]) Create(ctx context.Context, entity T) error {
	const op = "AbstractInMemoryRepository.Create"
	r.m.Lock()
	defer r.m.Unlock()
	if _, ok := r.entities[entity.GetID()]; ok {
		return domain.ErrAlreadyExists
	}

	r.entities[entity.GetID()] = entity
	return nil
}

func (r *AbstractInMemoryRepository[T]) Update(ctx context.Context, entity T) error {
	const op = "AbstractInMemoryRepository.Update"
	r.m.Lock()
	defer r.m.Unlock()
	if _, ok := r.entities[entity.GetID()]; !ok {
		return domain.ErrNotFound
	}
	r.entities[entity.GetID()] = entity
	return nil
}

func (r *AbstractInMemoryRepository[T]) Delete(ctx context.Context, id uuid.UUID) error {
	const op = "AbstractInMemoryRepository.Delete"
	r.m.Lock()
	defer r.m.Unlock()
	if _, ok := r.entities[id]; !ok {
		return nil
	}
	delete(r.entities, id)
	return nil
}

func (r *AbstractInMemoryRepository[T]) GetByID(ctx context.Context, id uuid.UUID) (T, error) {
	const op = "AbstractInMemoryRepository.GetByID"
	r.m.Lock()
	defer r.m.Unlock()
	if entity, ok := r.entities[id]; ok {
		return entity, nil
	}
	var entity T
	return entity, domain.ErrNotFound
}

func (r *AbstractInMemoryRepository[T]) GetByIds(ctx context.Context, ids []uuid.UUID) ([]T, error) {
	const op = "AbstractInMemoryRepository.GetByIds"
	r.m.Lock()
	defer r.m.Unlock()
	entities := make([]T, 0, len(ids))
	for _, id := range ids {
		if entity, ok := r.entities[id]; ok {
			entities = append(entities, entity)
		}
	}
	r.logger.Debug(op, slog.Any("ids", ids), slog.Any("entities", entities))
	return entities, nil
}

func (r *AbstractInMemoryRepository[T]) GetAll(ctx context.Context, limit int, offset int) ([]T, error) {
	const op = "AbstractInMemoryRepository.GetAll"
	r.m.Lock()
	defer r.m.Unlock()

	keys := make([]uuid.UUID, 0)
	for key := range r.entities {
		keys = append(keys, key)
	}

	i := offset
	entities := make([]T, 0)
	for i < len(keys) && len(entities) < limit {
		entities = append(entities, r.entities[keys[i]])
		i++
	}
	return entities, nil
}
