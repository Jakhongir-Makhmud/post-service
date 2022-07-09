package post_repo

import (
	"context"
	"post-service/internal/structs"
)

type PostRepo interface {
	GetPost(ctx context.Context, postId int) (structs.Post, error)
	UpdatePost(ctx context.Context, post structs.Post) (structs.Post, error)
	DeletePost(ctx context.Context, postId int) error
	GetPosts(ctx context.Context) ([]structs.Post, error)
}

func NewPostRepo()