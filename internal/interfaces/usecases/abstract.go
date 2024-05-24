package usecases

import (
	"GraphQLTestCase/internal/domain"
	"context"
	"github.com/google/uuid"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=AbstractRepositoryInterface
type AbstractRepositoryInterface[T domain.Model] interface {
	Create(ctx context.Context, entity T) error
	Update(ctx context.Context, entity T) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (T, error)
	GetByIds(ctx context.Context, ids []uuid.UUID) ([]T, error)
	GetAll(ctx context.Context, limit int, offset int) ([]T, error)
}

type AbstractUseCaseInterface[T domain.Model] interface {
	Create(ctx context.Context, entity T) error
	Update(ctx context.Context, entity T) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (T, error)
	GetByIds(ctx context.Context, ids []uuid.UUID) ([]T, error)
	GetAll(ctx context.Context, limit int, offset int) ([]T, error)
}

var _ AbstractUseCaseInterface[domain.Model] = &AbstractUseCase[domain.Model]{}

type AbstractUseCase[T domain.Model] struct {
	repository AbstractRepositoryInterface[T]
}

func NewAbstractUseCase[T domain.Model](repository AbstractRepositoryInterface[T]) AbstractUseCase[T] {
	return AbstractUseCase[T]{repository: repository}
}

func (uc *AbstractUseCase[T]) Create(ctx context.Context, entity T) error {
	entity.SetID(uuid.New())
	return uc.repository.Create(ctx, entity)
}

func (uc *AbstractUseCase[T]) Update(ctx context.Context, entity T) error {
	return uc.repository.Update(ctx, entity)
}

func (uc *AbstractUseCase[T]) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repository.Delete(ctx, id)
}

func (uc *AbstractUseCase[T]) GetByID(ctx context.Context, id uuid.UUID) (T, error) {
	return uc.repository.GetByID(ctx, id)
}

func (uc *AbstractUseCase[T]) GetByIds(ctx context.Context, ids []uuid.UUID) ([]T, error) {
	return uc.repository.GetByIds(ctx, ids)
}

func (uc *AbstractUseCase[T]) GetAll(ctx context.Context, limit int, offset int) ([]T, error) {
	return uc.repository.GetAll(ctx, limit, offset)
}
