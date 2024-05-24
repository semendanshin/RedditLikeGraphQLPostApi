package mappers

import (
	"GraphQLTestCase/internal/domain"
	"GraphQLTestCase/internal/infrastructure/entities"
	"GraphQLTestCase/internal/infrastructure/graph/model"
)

func ModelToDomainPost(dto *model.Post) *domain.Post {
	return &domain.Post{
		ID:            dto.ID,
		Title:         dto.Title,
		Content:       dto.Content,
		AuthorID:      dto.AuthorID,
		AllowComments: dto.AllowComments,
		CreatedAt:     dto.CreatedAt,
		UpdatedAt:     dto.UpdatedAt,
	}
}

func DomainToModelPost(domain *domain.Post) *model.Post {
	return &model.Post{
		ID:            domain.ID,
		Title:         domain.Title,
		Content:       domain.Content,
		AuthorID:      domain.AuthorID,
		AllowComments: domain.AllowComments,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func DomainToEntityPost(domain *domain.Post) *entities.Post {
	return &entities.Post{
		ID:            domain.ID,
		Title:         domain.Title,
		Content:       domain.Content,
		AuthorID:      domain.AuthorID,
		AllowComments: domain.AllowComments,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func EntityToDomainPost(entity *entities.Post) *domain.Post {
	return &domain.Post{
		ID:            entity.ID,
		Title:         entity.Title,
		Content:       entity.Content,
		AuthorID:      entity.AuthorID,
		AllowComments: entity.AllowComments,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
	}
}

func CreateDTOToDomainPost(dto *model.NewPost) *domain.Post {
	var allowComments bool
	if dto.AllowComments != nil {
		allowComments = *dto.AllowComments
	}
	return &domain.Post{
		Title:         dto.Title,
		Content:       dto.Content,
		AuthorID:      dto.AuthorID,
		AllowComments: allowComments,
	}
}
