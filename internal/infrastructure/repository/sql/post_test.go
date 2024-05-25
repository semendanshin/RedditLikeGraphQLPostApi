package sql

import (
	"GraphQLTestCase/internal/domain"
	"GraphQLTestCase/pkg/logger/slogdiscard"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func setupPostSQLRepository(t *testing.T) *PostSQLRepository {
	slogger := slogdiscard.NewDiscardLogger()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatal(err)
	}
	err = db.AutoMigrate(&domain.Post{})

	postRepo := NewPostSQLRepository(db, slogger)

	return postRepo
}

func TestPostSQLRepository_GetByAuthorID(t *testing.T) {
	rep := setupPostSQLRepository(t)

	authorID := uuid.New()
	postID := uuid.New()

	err := rep.Create(context.Background(), &domain.Post{
		ID:       postID,
		AuthorID: authorID,
		Title:    "Test post",
		Content:  "Test content",
	})

	if err != nil {
		t.Fatal(err)
	}

	posts, err := rep.GetByAuthorID(context.Background(), authorID, 10, 0)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(posts))
	assert.Equal(t, postID, posts[0].ID)
}
