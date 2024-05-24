package in_memory

import (
	"GraphQLTestCase/internal/domain"
	"GraphQLTestCase/lib/logger/slogdiscard"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryCommentRepository_Create(t *testing.T) {
	logger := slogdiscard.NewDiscardLogger()
	// Initialize repository with in-memory storage
	commentRepo := NewCommentInMemoryRepository(logger)

	// Define test data
	ID := uuid.New()
	comment := &domain.Comment{
		ID:      ID,
		Content: "Test comment",
	}

	// Call the repository method being tested
	err := commentRepo.Create(context.Background(), comment)

	// Assert the result
	assert.NoError(t, err)

	// Verify that the comment was added to the repository
	createdComment, err := commentRepo.GetByID(context.Background(), ID)

	assert.NoError(t, err)
	assert.NotNil(t, createdComment)
	assert.Equal(t, "Test comment", createdComment.Content)
}

func TestInMemoryCommentRepository_Create_AlreadyExists(t *testing.T) {
	logger := slogdiscard.NewDiscardLogger()
	commentRepo := NewCommentInMemoryRepository(logger)

	ID := uuid.New()
	comment := &domain.Comment{
		ID:      ID,
		Content: "Test comment",
	}

	err := commentRepo.Create(context.Background(), comment)
	assert.NoError(t, err)

	err = commentRepo.Create(context.Background(), comment)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrAlreadyExists, err)
}

func TestInMemoryCommentRepository_Update(t *testing.T) {
	logger := slogdiscard.NewDiscardLogger()
	commentRepo := NewCommentInMemoryRepository(logger)

	ID := uuid.New()
	comment := &domain.Comment{
		ID:      ID,
		Content: "Test comment",
	}

	err := commentRepo.Create(context.Background(), comment)
	assert.NoError(t, err)

	comment.Content = "Updated comment"
	err = commentRepo.Update(context.Background(), comment)
	assert.NoError(t, err)

	updatedComment, err := commentRepo.GetByID(context.Background(), ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedComment)
	assert.Equal(t, "Updated comment", updatedComment.Content)
}

func TestInMemoryCommentRepository_Update_NotFound(t *testing.T) {
	logger := slogdiscard.NewDiscardLogger()
	commentRepo := NewCommentInMemoryRepository(logger)

	ID := uuid.New()
	comment := &domain.Comment{
		ID:      ID,
		Content: "Test comment",
	}

	err := commentRepo.Update(context.Background(), comment)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrNotFound, err)

}

func TestInMemoryCommentRepository_Delete(t *testing.T) {
	logger := slogdiscard.NewDiscardLogger()
	commentRepo := NewCommentInMemoryRepository(logger)

	ID := uuid.New()
	comment := &domain.Comment{
		ID:      ID,
		Content: "Test comment",
	}

	err := commentRepo.Create(context.Background(), comment)
	assert.NoError(t, err)

	err = commentRepo.Delete(context.Background(), ID)
	assert.NoError(t, err)

	deletedComment, err := commentRepo.GetByID(context.Background(), ID)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrNotFound, err)
	assert.Nil(t, deletedComment)
}

func TestInMemoryCommentRepository_Delete_NotFound(t *testing.T) {
	logger := slogdiscard.NewDiscardLogger()
	commentRepo := NewCommentInMemoryRepository(logger)

	ID := uuid.New()
	err := commentRepo.Delete(context.Background(), ID)
	assert.NoError(t, err)
}

func TestInMemoryCommentRepository_GetByID(t *testing.T) {
	logger := slogdiscard.NewDiscardLogger()
	commentRepo := NewCommentInMemoryRepository(logger)

	ID := uuid.New()
	comment := &domain.Comment{
		ID:      ID,
		Content: "Test comment",
	}

	err := commentRepo.Create(context.Background(), comment)
	assert.NoError(t, err)

	foundComment, err := commentRepo.GetByID(context.Background(), ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundComment)
	assert.Equal(t, "Test comment", foundComment.Content)
}

func TestInMemoryCommentRepository_GetByID_NotFound(t *testing.T) {
	logger := slogdiscard.NewDiscardLogger()
	commentRepo := NewCommentInMemoryRepository(logger)

	ID := uuid.New()
	comment, err := commentRepo.GetByID(context.Background(), ID)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrNotFound, err)
	assert.Nil(t, comment)
}

func TestInMemoryCommentRepository_GetAll(t *testing.T) {
	logger := slogdiscard.NewDiscardLogger()
	commentRepo := NewCommentInMemoryRepository(logger)

	ID1 := uuid.New()
	comment1 := &domain.Comment{
		ID:      ID1,
		Content: "Test comment 1",
	}

	ID2 := uuid.New()
	comment2 := &domain.Comment{
		ID:      ID2,
		Content: "Test comment 2",
	}

	err := commentRepo.Create(context.Background(), comment1)
	assert.NoError(t, err)

	err = commentRepo.Create(context.Background(), comment2)
	assert.NoError(t, err)

	comments, err := commentRepo.GetAll(context.Background(), 10, 0)
	assert.NoError(t, err)
	assert.Len(t, comments, 2)
	assert.Equal(t, "Test comment 1", comments[0].Content)
	assert.Equal(t, "Test comment 2", comments[1].Content)
}
