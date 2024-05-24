package usecases

import (
	"GraphQLTestCase/internal/usecases/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCommentUseCase_GetByPostID(t *testing.T) {
	repo := &mocks.CommentRepository{}
	uc := NewCommentUseCase(repo)

	repo.On("GetByPostID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

	id := uuid.New()
	_, err := uc.GetByPostID(nil, id, 0, 0)

	assert.NoError(t, err)
}

func TestCommentUseCase_GetChildren(t *testing.T) {
	repo := &mocks.CommentRepository{}
	uc := NewCommentUseCase(repo)

	repo.On("GetChildren", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

	id := uuid.New()
	_, err := uc.GetChildren(nil, id, 0, 0)

	assert.NoError(t, err)
}
