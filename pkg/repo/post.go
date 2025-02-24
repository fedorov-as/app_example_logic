package repo

import "github.com/fedorov-as/app_example_logic/pkg/model"

type PostsRepo interface {
	AddPost(text string, owner model.User) (model.Post, error)
	GetPostByID(id int) (model.Post, error)
	GetPostsByOwner(owner model.User) ([]model.Post, error)
}
