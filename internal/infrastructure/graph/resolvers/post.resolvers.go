package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"GraphQLTestCase/internal/infrastructure/graph"
	"GraphQLTestCase/internal/infrastructure/graph/model"
	"GraphQLTestCase/internal/utils/mappers"
	"context"

	"github.com/google/uuid"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	const op = "PostResolver.CreatePost"

	post := mappers.CreateDTOToDomainPost(&input)

	err := r.puc.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return mappers.DomainToModelPost(post), nil
}

// DisableComments is the resolver for the disableComments field.
func (r *mutationResolver) DisableComments(ctx context.Context, postID uuid.UUID) (*model.Post, error) {
	const op = "PostResolver.DisableComments"

	//err := r.puc.DisableComments(ctx, postID)
	//if err != nil {
	//	return nil, err
	//}
	//
	//post, err := r.puc.GetPostByID(ctx, postID, nil)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &model.Post{
	//	ID:            post.ID,
	//	Title:         post.Title,
	//	Content:       post.Content,
	//	AuthorID:      post.AuthorID,
	//	AllowComments: post.AllowComments,
	//	CreatedAt:     post.CreatedAt,
	//	UpdatedAt:     post.UpdatedAt,
	//}, nil
	return nil, nil
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *model.Post, limit *int, offset *int) ([]*model.Comment, error) {
	const op = "PostResolver.Comments"

	comments, err := r.cuc.GetByPostID(ctx, obj.ID, *limit, *offset)
	if err != nil {
		return nil, err
	}

	var modelComments []*model.Comment
	for _, comment := range comments {
		modelComments = append(modelComments, mappers.DomainToModelComment(comment))
	}

	return modelComments, nil
}

// Author is the resolver for the author field.
func (r *postResolver) Author(ctx context.Context, obj *model.Post) (*model.User, error) {
	const op = "PostResolver.Author"

	user, err := graph.GetUserLoader(ctx).Load(obj.AuthorID)
	if err != nil {
		return nil, err
	}

	return mappers.DomainToModelUser(user), nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id uuid.UUID) (*model.Post, error) {
	const op = "PostResolver.Post"

	post, err := r.puc.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mappers.DomainToModelPost(post), nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, limit *int, offset *int) ([]*model.Post, error) {
	const op = "PostResolver.Posts"

	posts, err := r.puc.GetAll(ctx, *limit, *offset)
	if err != nil {
		return nil, err
	}

	var modelPosts []*model.Post
	for _, post := range posts {
		modelPosts = append(modelPosts, mappers.DomainToModelPost(post))
	}

	return modelPosts, nil
}

// Post returns graph.PostResolver implementation.
func (r *Resolver) Post() graph.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
