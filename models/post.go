package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Body        string     `json:"content"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	PublishedAt *time.Time `json:"published_at"`
}

// Render returns the full HTML content of the post
func (post Post) Render() (string, error) {
	return "", nil
}

// Snippet returns a snippet of the post content
func (post Post) Snippet(length int) (string, error) {
	return "", nil
}
