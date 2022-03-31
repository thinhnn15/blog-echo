package repository

import (
	"context"
	"github.com/thinhnn15/blog-echo/model"
)

type PostRepo interface {
	SavePost(context context.Context, post model.Post) (model.Post, error)
}