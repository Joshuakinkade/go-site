-- name: GetPosts :many
SELECT id, title, slug, body, created_at, updated_at, published_at FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: GetPostBySlug :one
SELECT id, title, slug, body, created_at, updated_at, published_at FROM posts WHERE slug = $1;