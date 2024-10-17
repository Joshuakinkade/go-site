package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/joshuakinkade/go-site/db"
	"github.com/joshuakinkade/go-site/lib/text"
	"github.com/joshuakinkade/go-site/models"
)

// Depenenencies:
// posts repository

// PostsService provides methods for working with posts.
type PostService struct {
	posts db.IPostsRepository
}

// NewPostService returns an initialized PostService.
func NewPostService(posts db.IPostsRepository) PostService {
	return PostService{
		posts: posts,
	}
}

// ListPosts returns a list of posts in reverse chronological order.
func (p PostService) ListPosts(offset, limit int) ([]models.Post, error) {
	posts, err := p.posts.ListPosts(offset, limit)
	return posts, err
}

// GetPostBySlug returns a post by its slug.
func (p PostService) GetPostBySlug(slug string) (models.Post, error) {
	post, err := p.posts.GetPostBySlug(slug)
	return post, err
}

// CreatePost saves a post and sets sensible defaults for missing data.
func (p PostService) CreatePost(post models.Post) (models.Post, error) {
	post.ID, _ = uuid.NewV6()
	if post.Slug == "" {
		post.Slug = text.Slugify(post.Title)
	}
	if post.CreatedAt.IsZero() {
		post.CreatedAt = time.Now()
	}

	post, err := p.posts.CreatePost(post)
	return post, err
}

// UpdatePost updates an existing post.
func (p PostService) UpdatePost(post models.Post) (models.Post, error) {
	post, err := p.posts.UpdatePost(post)
	return post, err
}
