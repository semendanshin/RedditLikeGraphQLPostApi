package usecases

import (
	"GraphQLTestCase/internal/usecases/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestPostUseCase_GetByAuthorID(t *testing.T) {
	repo := &mocks.PostRepository{}
	uc := NewPostUseCase(repo)

	repo.On("GetByAuthorID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

	id := uuid.New()
	_, err := uc.GetByAuthorID(nil, id, 0, 0)

	assert.NoError(t, err)
}
