package models

import (
	"bytes"
	"time"

	"github.com/google/uuid"
	"github.com/yuin/goldmark"
)

type Post struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Body        string     `json:"body"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	PublishedAt *time.Time `json:"published_at"`
}

// Render returns the full HTML content of the post
func (post Post) Render() (string, error) {
	var buffer bytes.Buffer
	err := goldmark.Convert([]byte(post.Body), &buffer)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// Snippet returns a snippet of the post content
func (post Post) Snippet(length int) (string, error) {
	return "", nil
}
