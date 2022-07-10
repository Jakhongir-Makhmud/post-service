package post_repo

import (
	"context"
	"post-service/internal/structs"
	"post-service/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type PostRepo interface {
	GetPost(ctx context.Context, postId int) (structs.Post, error)
	UpdatePost(ctx context.Context, post structs.Post) (structs.Post, error)
	DeletePost(ctx context.Context, postId int) error
	GetPosts(ctx context.Context, params structs.PostParams) ([]structs.Post, error)
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
