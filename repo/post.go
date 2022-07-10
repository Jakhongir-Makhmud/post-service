package post_repo

import (
	"context"
	"database/sql"
	"fmt"
	"post-service/internal/structs"

	"go.uber.org/zap"
)

func columns() string {
	return `
		post_id,
		user_id,
		title,
		body
	`
}

func fields(f *structs.Post) []interface{} {
	return []interface{}{
		&f.Id,
		&f.UserId,
		&f.Title,
		&f.Body,
	}
}

func (r *repo) GetPost(ctx context.Context, postId int) (structs.Post, error) {

	query := fmt.Sprintf("SELECT %s FROM posts WHERE post_id = $1", columns())

	var post structs.Post

	err := r.db.QueryRowContext(ctx, query, postId).Scan(fields(&post))
	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Warn("nothing found", zap.Int("post id: ", postId))
			return structs.Post{}, structs.ErrNotFound
		}
		r.logger.Error("error while selecting post by id", zap.Int("post id: ", postId))
		return structs.Post{}, err
	}

	return post, nil
}

func (r *repo) UpdatePost(ctx context.Context, post structs.Post) (structs.Post, error) {

	query := fmt.Sprintf("UPDATE posts SET title = $1, body = $2 WHERE post_id = $3 RETURNING %s", columns())

	err := r.db.QueryRowContext(ctx, query, post.Title, post.Body, post.Id).Scan(fields(&post))
	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Warn("no such post", zap.Any("post to update: ", post))
			return structs.Post{}, structs.ErrBadRequest
		}
		r.logger.Error("error while updating post in repo", zap.Any("post: ", post))
		return structs.Post{}, err
	}

	return post, nil
}

func (r *repo) DeletePost(ctx context.Context, postId int) error {

	_, err := r.db.ExecContext(ctx, "DELETE FROM posts WHERE post_id = $1", postId)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetPosts(ctx context.Context, params structs.PostParams) ([]structs.Post, error) {
	offset := (params.Page - 1) * params.Limit
	query := fmt.Sprintf("SELECT %s FROM posts LIMIT = $1 OFFSET = $2", columns())

	rows, err := r.db.QueryContext(ctx, query, params.Limit, offset)

	if err != nil {
		return nil, err
	}

	var posts []structs.Post

	for rows.Next() {

		post := structs.Post{}

		err := rows.Scan(fields(&post))
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil

}
