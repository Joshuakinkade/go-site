package db

import (
	"database/sql"

	"github.com/joshuakinkade/go-site/models"
)

// Dependencies:
// db

type Posts struct {
	db *sql.DB
}

func NewPosts(db *sql.DB) Posts {
	return Posts{db}
}

func (p Posts) ListPosts(offset, limit int) ([]models.Post, error) {
	sql := "SELECT id, title, slug, body, created_at, updated_at, published_at FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2"
	rows, err := p.db.Query(sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []models.Post{}
	for rows.Next() {
		post := models.Post{}
		// map row to post
		posts = append(posts, post)
	}
	return posts, nil
}
