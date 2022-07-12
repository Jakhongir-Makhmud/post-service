package post_repo

import (
	"context"
	"database/sql"
	"fmt"
	ps "post-service/genproto/post_service"
	"post-service/internal/structs"

	"go.uber.org/zap"
)

func columns() string {
	return `
		post_id,
		title,
		body
	`
}

func fields(f *ps.Post) []interface{} {
	return []interface{}{
		&f.Id,
		&f.Title,
		&f.Body,
	}
}

func (r *repo) GetPost(ctx context.Context, postId int64) (ps.Post, error) {

	query := fmt.Sprintf("SELECT %s FROM posts WHERE post_id = $1", columns())

	var post ps.Post

	err := r.db.QueryRowContext(ctx, query, postId).Scan(fields(&post)...)
	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Warn("nothing found", zap.Int64("post id: ", postId))
			return ps.Post{}, structs.ErrNotFound
		}
		r.logger.Error("error while selecting post by id", zap.Int64("post id: ", postId), zap.Error(err))
		return ps.Post{}, err
	}

	return post, nil
}

func (r *repo) UpdatePost(ctx context.Context, post *ps.Post) (*ps.Post, error) {

	query := fmt.Sprintf("UPDATE posts SET title = $1, body = $2 WHERE post_id = $3 RETURNING %s", columns())

	err := r.db.QueryRowContext(ctx, query, post.Title, post.Body, post.Id).Scan(fields(post)...)
	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Warn("no such post", zap.Any("post to update: ", post))
			return nil, structs.ErrBadRequest
		}
		r.logger.Error("error while updating post in repo", zap.Any("post: ", post),zap.Error(err))
		return nil, err
	}

	return post, nil
}

func (r *repo) DeletePost(ctx context.Context, postId int64) error {

	_, err := r.db.ExecContext(ctx, "DELETE FROM posts WHERE post_id = $1", postId)
	if err != nil {
		r.logger.Error("can not delete post from table", zap.Error(err))
		return err
	}

	return nil
}

func (r *repo) GetPosts(ctx context.Context, params ps.ListOfPosts) (*ps.Posts, error) {
	offset := (params.Page - 1) * params.Limit
	query := fmt.Sprintf("SELECT %s FROM posts LIMIT $1 OFFSET $2", columns())

	rows, err := r.db.QueryContext(ctx, query, params.Limit, offset)
	defer rows.Close()
	if err != nil {
		r.logger.Error("error while quering data from database", zap.Error(err))
		return &ps.Posts{}, err
	}

	var posts = &ps.Posts{}

	for rows.Next() {

		post := &ps.Post{}

		err := rows.Scan(fields(post)...)
		if err != nil {
			r.logger.Error("error while scanning values from rows", zap.Error(err))
			return &ps.Posts{}, err
		}

		posts.Posts = append(posts.Posts, post)
	}

	return posts, nil

}
