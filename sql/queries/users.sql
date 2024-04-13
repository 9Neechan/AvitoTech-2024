-- name: CreateUser :one
INSERT INTO users (id, is_admin, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;
