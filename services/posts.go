package services

import "github.com/joshuakinkade/go-site/models"

// Depenenencies:
// posts repository

type PostService struct {
}

func NewPostService() PostService {
	return PostService{}
}

func (p PostService) ListPosts(offset, limit int) ([]models.Post, error) {
	posts := []models.Post{
		{
			Title:   "Hello, World!",
			Slug:    "hello-world",
			Content: "This is a test post. It's not very interesting, but it's a start.",
		},
	}
	return posts, nil
}

func (p PostService) GetPostBySlug(slug string) (models.Post, error) {
	if slug == "hello-world" {
		return models.Post{
			Title:   "Hello, World!",
			Slug:    "hello-world",
			Content: "This is a test post. It's not very interesting, but it's a start.",
		}, nil
	} else {
		return models.Post{}, nil
	}
}

func (posts PostService) CreatePost(post models.Post) (models.Post, error) {
	return models.Post{}, nil
}
