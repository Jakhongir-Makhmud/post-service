package postservice

import (
	"context"
	"fmt"
	ps "post-service/genproto/post_service"
	"post-service/internal/structs"

	"go.uber.org/zap"
)

func (s *service) GetPost(ctx context.Context, postId *ps.PostId) (*ps.Post, error) {

	post, err := s.postRepo.GetPost(ctx, postId.Id)
	if err != nil {
		if err == structs.ErrNotFound {
			return nil, err
		}
		s.logger.Error("error while getting post by id", zap.Int64("post id: ", postId.Id))
		return nil, structs.ErrInternal
	}
	return &post, nil
}

func (s *service) UpdatePost(ctx context.Context, post *ps.Post) (*ps.Post, error) {
	post, err := s.postRepo.UpdatePost(ctx, post)
	if err != nil {
		if err == structs.ErrBadRequest {
			s.logger.Warn("can not update not existing post", zap.Any("updating post:", post))
			return nil, err
		}

		s.logger.Error("can not update post, something went wrong", zap.Error(err))
		return nil, structs.ErrInternal
	}

	return post, nil
}

func (s *service) DeletePost(ctx context.Context, postId *ps.PostId) (*ps.Empty, error) {

	err := s.postRepo.DeletePost(ctx, postId.Id)
	if err != nil {
		s.logger.Error("error while deleting post",
			zap.Error(err),
			zap.Int64("post id", postId.Id))
		return nil, structs.ErrInternal
	}

	return &ps.Empty{}, nil
}

func (s *service) ListPost(ctx context.Context, params *ps.ListOfPosts) (*ps.Posts, error) {

	if params.Page <= 0{
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 10
	}
	s.logger.Info(fmt.Sprintf("params %+v",params))
	posts, err := s.postRepo.GetPosts(ctx, *params)
	if err != nil {
		if err == structs.ErrNotFound {
			s.logger.Warn("no posts found", zap.Any("params: ", params))
			return nil, structs.ErrNotFound
		}

		s.logger.Error("something went wrong", zap.Error(err))
		return nil, structs.ErrInternal
	}

	return posts, nil
}
