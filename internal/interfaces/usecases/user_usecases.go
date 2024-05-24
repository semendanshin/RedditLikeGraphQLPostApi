package usecases

import (
	"GraphQLTestCase/internal/domain"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=UserUseCase
type UserUseCase interface {
	AbstractUseCaseInterface[*domain.User]
}
