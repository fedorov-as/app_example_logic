package repo

import (
	"fmt"

	"github.com/fedorov-as/app_example_logic/pkg/model"
)

type PostsMemoryRepo struct {
	data   map[int]model.Post
	nextID int
}

var _ PostsRepo = &PostsMemoryRepo{}

func NewPostsMemoryRepo() *PostsMemoryRepo {
	return &PostsMemoryRepo{
		data:   make(map[int]model.Post),
		nextID: 0,
	}
}

func (repo *PostsMemoryRepo) AddPost(text string, owner model.User) (model.Post, error) {
	post := model.NewPost(repo.nextID, text, owner)
	repo.data[repo.nextID] = post
	repo.nextID++
	return post, nil
}

func (repo PostsMemoryRepo) GetPostByID(id int) (model.Post, error) {
	if post, ok := repo.data[id]; ok {
		return post, nil
	}

	return model.Post{}, fmt.Errorf("no post with id %d", id)
}

func (repo PostsMemoryRepo) GetPostsByOwner(owner model.User) ([]model.Post, error) {
	posts := make([]model.Post, 0)

	for _, post := range repo.data {
		posts = append(posts, post)
	}

	return posts, nil
}

func (repo *PostsMemoryRepo) DeletePost(id int) error {
	if _, ok := repo.data[id]; ok {
		delete(repo.data, id)
		return nil
	}

	return fmt.Errorf("no post with id %d", id)
}
