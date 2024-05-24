package sql

import (
	"GraphQLTestCase/internal/domain"
	"GraphQLTestCase/internal/infrastructure/entities"
	"GraphQLTestCase/internal/usecases"
	"GraphQLTestCase/internal/utils/mappers"
	"gorm.io/gorm"
	"log/slog"
)

var _ usecases.UserRepository = &UserSQLRepository{}

type UserSQLRepository struct {
	AbstractSQLRepository[*domain.User, entities.User]
}

func NewUserSQLRepository(db *gorm.DB, logger *slog.Logger) *UserSQLRepository {
	return &UserSQLRepository{
		AbstractSQLRepository: NewAbstractSQLRepository[*domain.User, entities.User](
			db, logger, mappers.DomainToEntityUser, mappers.EntityToDomainUser,
		),
	}
}
