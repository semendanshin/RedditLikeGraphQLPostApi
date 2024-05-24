package mappers

import (
	"GraphQLTestCase/internal/domain"
	"GraphQLTestCase/internal/infrastructure/entities"
	"GraphQLTestCase/internal/infrastructure/graph/model"
)

func ModelToDomainComment(dto *model.Comment) *domain.Comment {
	return &domain.Comment{
		ID:        dto.ID,
		PostID:    dto.PostID,
		ParentID:  dto.ParentID,
		Content:   dto.Content,
		AuthorID:  dto.AuthorID,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}

func DomainToModelComment(domain *domain.Comment) *model.Comment {
	return &model.Comment{
		ID:        domain.ID,
		PostID:    domain.PostID,
		ParentID:  domain.ParentID,
		Content:   domain.Content,
		AuthorID:  domain.AuthorID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func DomainToEntityComment(domain *domain.Comment) *entities.Comment {
	return &entities.Comment{
		ID:        domain.ID,
		PostID:    domain.PostID,
		ParentID:  domain.ParentID,
		Content:   domain.Content,
		AuthorID:  domain.AuthorID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func EntityToDomainComment(entity *entities.Comment) *domain.Comment {
	return &domain.Comment{
		ID:        entity.ID,
		PostID:    entity.PostID,
		ParentID:  entity.ParentID,
		Content:   entity.Content,
		AuthorID:  entity.AuthorID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func CreateDTOToDomainComment(dto *model.NewComment) *domain.Comment {
	return &domain.Comment{
		PostID:   dto.PostID,
		ParentID: dto.ParentID,
		Content:  dto.Content,
		AuthorID: dto.AuthorID,
	}
}
