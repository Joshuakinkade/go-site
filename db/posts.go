package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/joshuakinkade/go-site/models"
)

// Dependencies:
// db

type IPostsRepository interface {
	ListPosts(offset, limit int) ([]models.Post, error)
	GetPostBySlug(slug string) (models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
	UpdatePost(post models.Post) (models.Post, error)
}

type PostsRepository struct {
	db *pgx.Conn
}

func NewPosts(db *pgx.Conn) PostsRepository {
	return PostsRepository{db}
}

func (p PostsRepository) ListPosts(offset, limit int) ([]models.Post, error) {
	sql := "SELECT id, title, slug, body, created_at, updated_at, published_at FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2"
	rows, err := p.db.Query(context.TODO(), sql, limit, offset)
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

func (p PostsRepository) GetPostBySlug(slug string) (models.Post, error) {
	query := "SELECT id, title, slug, body, created_at, updated_at, published_at FROM posts WHERE slug = $1"
	row := p.db.QueryRow(context.TODO(), query, slug)
	post := models.Post{}
	err := row.Scan(&post.ID, &post.Title, &post.Slug, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.PublishedAt)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return models.Post{}, err
	}
	return post, nil
}

func (p PostsRepository) CreatePost(post models.Post) (models.Post, error) {
	sql := "INSERT INTO posts (id, title, slug, body, created_at, updated_at, published_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	row := p.db.QueryRow(context.TODO(), sql, post.ID, post.Title, post.Slug, post.Content, post.CreatedAt, post.UpdatedAt, post.PublishedAt)
	err := row.Scan(&post.ID)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (p PostsRepository) UpdatePost(post models.Post) (models.Post, error) {
	sql := "UPDATE posts SET title = $1, slug = $2, body = $3, created_at = $4, updated_at = $5, published_at = $6 WHERE id = $7"
	_, err := p.db.Exec(context.TODO(), sql, post.Title, post.Slug, post.Content, post.CreatedAt, post.UpdatedAt, post.PublishedAt, post.ID)
	if err != nil {
		return models.Post{}, err
	}
	return models.Post{}, nil
}
