package mappers

import (
	"GraphQLTestCase/internal/domain"
	"GraphQLTestCase/internal/infrastructure/entities"
	"GraphQLTestCase/internal/infrastructure/graph/model"
)

func ModelToDomainUser(dto *model.User) *domain.User {
	return &domain.User{
		ID:   dto.ID,
		Name: dto.Name,
	}
}

func DomainToModelUser(domain *domain.User) *model.User {
	return &model.User{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

func DomainToEntityUser(domain *domain.User) *entities.User {
	return &entities.User{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

func EntityToDomainUser(entity *entities.User) *domain.User {
	return &domain.User{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func CreateDTOToDomainUser(dto *model.NewUser) *domain.User {
	return &domain.User{
		Name: dto.Name,
	}
}
