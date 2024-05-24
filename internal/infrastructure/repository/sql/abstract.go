package sql

import (
	"GraphQLTestCase/internal/domain"
	"GraphQLTestCase/internal/interfaces/usecases"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log/slog"
)

type Entity interface{}

var _ usecases.AbstractRepositoryInterface[domain.Model] = &AbstractSQLRepository[domain.Model, Entity]{}

type AbstractSQLRepository[TModel domain.Model, TEntity Entity] struct {
	db            *gorm.DB
	logger        *slog.Logger
	modelToEntity func(model TModel) *TEntity
	entityToModel func(entity *TEntity) TModel
}

func NewAbstractSQLRepository[TModel domain.Model, TEntity Entity](
	db *gorm.DB,
	logger *slog.Logger,
	modelToEntity func(model TModel) *TEntity,
	entityToModel func(entity *TEntity) TModel,
) AbstractSQLRepository[TModel, TEntity] {
	return AbstractSQLRepository[TModel, TEntity]{
		db:            db,
		logger:        logger,
		modelToEntity: modelToEntity,
		entityToModel: entityToModel,
	}
}

func (r *AbstractSQLRepository[TModel, TEntity]) Create(ctx context.Context, model TModel) error {
	const op = "AbstractSQLRepository.Create"
	entity := r.modelToEntity(model)
	if err := r.db.WithContext(ctx).Create(&entity).Error; err != nil {
		r.logger.Error(op, slog.Any("error", err.Error()))
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.ErrAlreadyExists
		}
		return err
	}
	return nil
}

func (r *AbstractSQLRepository[TModel, TEntity]) Update(ctx context.Context, model TModel) error {
	const op = "AbstractSQLRepository.Update"
	entity := r.modelToEntity(model)
	if err := r.db.WithContext(ctx).Model(entity).Updates(entity).Error; err != nil {
		r.logger.Error(op, slog.Any("error", err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *AbstractSQLRepository[TModel, TEntity]) Delete(ctx context.Context, id uuid.UUID) error {
	const op = "AbstractSQLRepository.Delete"
	var entity TEntity
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity).Error; err != nil {
		r.logger.Error(op, slog.Any("error", err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *AbstractSQLRepository[TModel, TEntity]) GetByID(ctx context.Context, id uuid.UUID) (TModel, error) {
	const op = "AbstractSQLRepository.GetByID"
	var entity TEntity
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&entity).Error; err != nil {
		r.logger.Error(op, slog.Any("error", err.Error()))
		var model TModel
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model, domain.ErrNotFound
		}
		return model, err
	}
	return r.entityToModel(&entity), nil
}

func (r *AbstractSQLRepository[TModel, TEntity]) GetByIds(ctx context.Context, ids []uuid.UUID) ([]TModel, error) {
	const op = "AbstractSQLRepository.GetByIds"
	var entities []*TEntity
	if err := r.db.WithContext(ctx).Where("id IN (?)", ids).Find(&entities).Error; err != nil {
		r.logger.Error(op, slog.Any("error", err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	var models []TModel
	for _, entity := range entities {
		models = append(models, r.entityToModel(entity))
	}
	return models, nil
}

func (r *AbstractSQLRepository[TModel, TEntity]) GetAll(ctx context.Context, limit int, offset int) ([]TModel, error) {
	const op = "AbstractSQLRepository.GetAll"
	var entities []*TEntity
	if err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&entities).Error; err != nil {
		r.logger.Error(op, slog.Any("error", err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	var models []TModel
	for _, entity := range entities {
		models = append(models, r.entityToModel(entity))
	}
	return models, nil
}
