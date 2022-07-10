package post_repo

import (
	"context"
	ps "post-service/genproto/post_service"
	"post-service/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type PostRepo interface {
	GetPost(ctx context.Context, postId int64) (ps.Post, error)
	UpdatePost(ctx context.Context, post *ps.Post) (*ps.Post, error)
	DeletePost(ctx context.Context, postId int64) error
	GetPosts(ctx context.Context, params ps.ListOfPosts) (*ps.Posts, error)
}

type repo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewPostRepo(db *sqlx.DB, logger logger.Logger) PostRepo {
	return &repo{
		db:     db,
		logger: logger,
	}
}
