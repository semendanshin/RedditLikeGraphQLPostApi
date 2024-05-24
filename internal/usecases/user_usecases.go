package usecases

import (
	"GraphQLTestCase/internal/domain"
	usecaseInterfaces "GraphQLTestCase/internal/interfaces/usecases"
)

//go:generate go run github.com/vektra/mockery/v2@v2.40.2 --name=UserRepository
type UserRepository interface {
	usecaseInterfaces.AbstractRepositoryInterface[*domain.User]
}

var _ usecaseInterfaces.UserUseCase = &UserUseCase{}

type UserUseCase struct {
	Repository UserRepository
	usecaseInterfaces.AbstractUseCase[*domain.User]
}

func NewUserUseCase(repository UserRepository) *UserUseCase {
	return &UserUseCase{
		Repository:      repository,
		AbstractUseCase: usecaseInterfaces.NewAbstractUseCase[*domain.User](repository),
	}
}
