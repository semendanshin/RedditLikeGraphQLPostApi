package graph

import (
	"GraphQLTestCase/internal/domain"
	usecaseInterfaces "GraphQLTestCase/internal/interfaces/usecases"
	"GraphQLTestCase/pkg/dataloader"
	"context"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"time"
)

const (
	userLoaderKey    = "userloader"
	commentLoaderKey = "commentloader"
	postLoaderKey    = "postloader"
)

const (
	maxBatch             = 100
	wait                 = 10 * time.Millisecond
	ttl                  = 1 * time.Second
	cacheCleanerInterval = 500 * time.Millisecond
)

type IdsLoader[TValue any] interface {
	GetByIds(ctx context.Context, ids []uuid.UUID) ([]*TValue, error)
}

func newLoader[TValue any](ctx context.Context, idsLoader IdsLoader[TValue], logger *slog.Logger) *dataloader.Loader[TValue, uuid.UUID] {
	return dataloader.NewLoader[TValue, uuid.UUID](
		maxBatch,
		wait,
		func(keys []uuid.UUID) ([]*TValue, []error) {
			const op = "DataLoaderMiddleware.FetchData"
			errors := make([]error, len(keys))
			users, err := idsLoader.GetByIds(ctx, keys)
			if err != nil {
				logger.Error(op, slog.Any("error", err.Error()))
				for i := range errors {
					errors[i] = err
				}
			}
			return users, errors
		},
		ttl,
		cacheCleanerInterval,
	)
}

func DataLoaderMiddleware(
	puc usecaseInterfaces.PostUseCase,
	cuc usecaseInterfaces.CommentUseCase,
	uuc usecaseInterfaces.UserUseCase,
	logger *slog.Logger,
) func(next http.Handler) http.Handler {
	const op = "DataLoaderMiddleware"
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userLoader := newLoader[domain.User](r.Context(), uuc, logger)
			commentLoader := newLoader[domain.Comment](r.Context(), cuc, logger)
			postLoader := newLoader[domain.Post](r.Context(), puc, logger)

			ctx := r.Context()
			ctx = context.WithValue(ctx, userLoaderKey, userLoader)
			ctx = context.WithValue(ctx, commentLoaderKey, commentLoader)
			ctx = context.WithValue(ctx, postLoaderKey, postLoader)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserLoader(ctx context.Context) *dataloader.Loader[domain.User, uuid.UUID] {
	return ctx.Value(userLoaderKey).(*dataloader.Loader[domain.User, uuid.UUID])
}

func GetCommentLoader(ctx context.Context) *dataloader.Loader[domain.Comment, uuid.UUID] {
	return ctx.Value(commentLoaderKey).(*dataloader.Loader[domain.Comment, uuid.UUID])
}

func GetPostLoader(ctx context.Context) *dataloader.Loader[domain.Post, uuid.UUID] {
	return ctx.Value(postLoaderKey).(*dataloader.Loader[domain.Post, uuid.UUID])
}
