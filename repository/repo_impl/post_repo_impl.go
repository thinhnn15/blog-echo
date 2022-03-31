package repo_impl

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	"github.com/thinhnn15/blog-echo/db"
	"github.com/thinhnn15/blog-echo/model"
	"github.com/thinhnn15/blog-echo/repository"
	"time"
)

type PostRepoImpl struct {
	sql *db.Sql
}

func NewPostRepo(sql *db.Sql) repository.PostRepo{
	return &PostRepoImpl{
		sql: sql,
	}
}

func (p *PostRepoImpl) SavePost(context context.Context, post model.Post) (model.Post, error) {
	statement := `
		INSERT INTO posts(title, content, created_at)
		VALUES(:title, :content, :created_at)
	`
	post.CreateAt = time.Now()
	_, err := p.sql.Db.NamedExecContext(context, statement, post)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return post, nil		// TNN-TODO
			}
		}
		return post, nil				// TNN-TODO
	}
	return post, nil					// TNN-TODO
}

