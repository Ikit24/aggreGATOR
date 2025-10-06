-- name: CreatePosts :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
