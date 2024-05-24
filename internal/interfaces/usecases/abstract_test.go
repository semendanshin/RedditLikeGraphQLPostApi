package usecases

import (
	"GraphQLTestCase/internal/interfaces/usecases/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type model struct {
	ID uuid.UUID
}

func (m model) GetID() uuid.UUID {
	return m.ID
}

func (m model) SetID(id uuid.UUID) {
	m.ID = id
}

func TestAbstractUseCase_Create(t *testing.T) {
	repo := &mocks.AbstractRepositoryInterface[model]{}
	uc := NewAbstractUseCase[model](repo)

	repo.On("Create", mock.Anything, mock.Anything).Return(nil)

	entity := model{}
	err := uc.Create(nil, entity)

	assert.NoError(t, err)
}

func TestAbstractUseCase_Update(t *testing.T) {
	repo := &mocks.AbstractRepositoryInterface[model]{}
	uc := NewAbstractUseCase[model](repo)

	repo.On("Update", mock.Anything, mock.Anything).Return(nil)

	entity := model{}
	err := uc.Update(nil, entity)

	assert.NoError(t, err)
}

func TestAbstractUseCase_Delete(t *testing.T) {
	repo := &mocks.AbstractRepositoryInterface[model]{}
	uc := NewAbstractUseCase[model](repo)

	repo.On("Delete", mock.Anything, mock.Anything).Return(nil)

	id := uuid.New()
	err := uc.Delete(nil, id)

	assert.NoError(t, err)
}

func TestAbstractUseCase_GetByID(t *testing.T) {
	repo := &mocks.AbstractRepositoryInterface[model]{}
	uc := NewAbstractUseCase[model](repo)

	repo.On("GetByID", mock.Anything, mock.Anything).Return(nil, nil)

	id := uuid.New()
	_, err := uc.GetByID(nil, id)

	assert.NoError(t, err)
}

func TestAbstractUseCase_GetByIds(t *testing.T) {
	repo := &mocks.AbstractRepositoryInterface[model]{}
	uc := NewAbstractUseCase[model](repo)

	repo.On("GetByIds", mock.Anything, mock.Anything).Return(nil, nil)

	ids := []uuid.UUID{uuid.New()}
	_, err := uc.GetByIds(nil, ids)

	assert.NoError(t, err)
}

func TestAbstractUseCase_GetAll(t *testing.T) {
	repo := &mocks.AbstractRepositoryInterface[model]{}
	uc := NewAbstractUseCase[model](repo)

	repo.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

	_, err := uc.GetAll(nil, 0, 0)

	assert.NoError(t, err)
}
