package postservice

import (
	"post-service/pkg/logger"
	post_repo "post-service/repo"
)

type service struct {
	postRepo post_repo.PostRepo
	logger   logger.Logger
}

// type PostService interface {
// 	GetPost(ctx context.Context, postId *ps.PostId) (*ps.Post, error)
// 	UpdatePost(ctx context.Context, post *ps.Post) (*ps.Post, error)
// 	DeletePost(ctx context.Context, postId *ps.PostId) (*ps.Empty, error)
// 	GetPosts(ctx context.Context, params *ps.ListOfPosts) (*[]ps.Post, error)
// }

func NewPostService(repo post_repo.PostRepo, logger logger.Logger) *service {

	return &service{
		postRepo: repo,
		logger:   logger,
	}
}
