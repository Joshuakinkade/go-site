package db

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/jackc/pgx/v5"
	querybuilder "github.com/joshuakinkade/go-site/db/query_builder"
	"github.com/joshuakinkade/go-site/models"
)

// Dependencies:
// db

type IPostsRepository interface {
	ListPosts(offset, limit int) ([]models.Post, error)
	GetPostBySlug(slug string) (models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
	UpdatePost(slug string, updates map[string]interface{}) error
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
		rows.Scan(&post.ID, &post.Title, &post.Slug, &post.Body, &post.CreatedAt, &post.UpdatedAt, &post.PublishedAt)
		// map row to post
		posts = append(posts, post)
	}
	return posts, nil
}

func (p PostsRepository) GetPostBySlug(slug string) (models.Post, error) {
	query := "SELECT id, title, slug, body, created_at, updated_at, published_at FROM posts WHERE slug = $1"
	row := p.db.QueryRow(context.TODO(), query, slug)
	post := models.Post{}
	err := row.Scan(&post.ID, &post.Title, &post.Slug, &post.Body, &post.CreatedAt, &post.UpdatedAt, &post.PublishedAt)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return models.Post{}, err
	}
	return post, nil
}

func (p PostsRepository) CreatePost(post models.Post) (models.Post, error) {
	sql := "INSERT INTO posts (id, title, slug, body, created_at, updated_at, published_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	row := p.db.QueryRow(context.TODO(), sql, post.ID, post.Title, post.Slug, post.Body, post.CreatedAt, post.UpdatedAt, post.PublishedAt)
	err := row.Scan(&post.ID)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (p PostsRepository) UpdatePost(slug string, updates map[string]interface{}) error {
	// This nonsense needs to be thoroughly tested to prevent security issues.
	// Maybe refactor it into a reusable function that can be used in other
	// repositories and queries.
	allowedFields := []string{"body", "published_at", "title", "updated_at"} // Needs to be in alphabetical order
	wheres, args, err := querybuilder.BuildUpdateClause(updates, allowedFields)
	q := "UPDATE posts SET " + wheres + " WHERE slug = $" + strconv.FormatInt(int64(len(args)+1), 10)
	args = append(args, slug)

	_, err = p.db.Exec(context.TODO(), q, args...)

	return err
}
