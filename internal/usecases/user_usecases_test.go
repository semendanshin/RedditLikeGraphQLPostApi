package usecases

import (
	"GraphQLTestCase/internal/domain"
	"GraphQLTestCase/internal/usecases/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUserUseCase_Create(t *testing.T) {
	repo := &mocks.UserRepository{}
	uc := NewUserUseCase(repo)

	repo.On("Create", mock.Anything, mock.Anything).Return(nil)

	entity := domain.User{}
	err := uc.Create(nil, entity)

	assert.NoError(t, err)
}

func TestUserUseCase_Update(t *testing.T) {
	repo := &mocks.UserRepository{}
	uc := NewUserUseCase(repo)

	repo.On("Update", mock.Anything, mock.Anything).Return(nil)

	entity := domain.User{}
	err := uc.Update(nil, entity)

	assert.NoError(t, err)
}

func TestUserUseCase_Delete(t *testing.T) {
	repo := &mocks.UserRepository{}
	uc := NewUserUseCase(repo)

	repo.On("Delete", mock.Anything, mock.Anything).Return(nil)

	id := uuid.UUID{}
	err := uc.Delete(nil, id)

	assert.NoError(t, err)
}

func TestUserUseCase_GetByID(t *testing.T) {
	repo := &mocks.UserRepository{}
	uc := NewUserUseCase(repo)

	repo.On("GetByID", mock.Anything, mock.Anything).Return(nil, nil)

	id := uuid.UUID{}
	_, err := uc.GetByID(nil, id)

	assert.NoError(t, err)
}

func TestUserUseCase_GetByIds(t *testing.T) {
	repo := &mocks.UserRepository{}
	uc := NewUserUseCase(repo)

	repo.On("GetByIds", mock.Anything, mock.Anything).Return(nil, nil)

	var ids []uuid.UUID
	_, err := uc.GetByIds(nil, ids)

	assert.NoError(t, err)
}
