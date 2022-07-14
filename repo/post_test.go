package post_repo

import (
	"context"
	post "post-service/genproto/post_service"
	"post-service/internal/structs"
	"post-service/pkg/config"
	"post-service/pkg/db"
	"post-service/pkg/logger"
	"testing"
)

func TestRepo(t *testing.T) {

	cfg := config.NewConfig()
	log := logger.New("debug", cfg.GetString("app.name"))
	db := db.NewDB(cfg)
	defer db.Close()

	repo := NewPostRepo(db, log)

	posts, err := repo.GetPosts(context.Background(), post.ListOfPosts{Page: 1, Limit: 10})
	if err != nil && err != structs.ErrNotFound {
		log.Error("failed to get posts")
		t.Fail()
	}

	var p *post.Post
	if len(posts.Posts) == 0 {
		log.Error("upload posts to db to acomplish the test successfully")
		t.Fail()
	} else {
		p = posts.Posts[0]
	}

	p.Title = "Test update"
	p.Body = "Test update"
	updatedPost, err := repo.UpdatePost(context.Background(), p)
	if err != nil {
		log.Error("failed to update post")
		t.Fail()
	}

	if p.Title != updatedPost.Title || p.Body != updatedPost.Body {
		log.Error("did not update")
		t.Fail()
	}

	gotPost, err := repo.GetPost(context.Background(), p.Id)
	if err != nil {
		log.Error("can not get post")
		t.Fail()
	}

	if gotPost.Title != updatedPost.Title {
		log.Error("it seems get method is not working properly")
		t.Fail()
	}

	err = repo.DeletePost(context.Background(), p.Id)
	if err != nil {
		log.Error("can not delete post")
		t.Fail()
	}

}
