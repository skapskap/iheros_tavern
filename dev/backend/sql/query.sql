-- name: CreateUser :one
INSERT INTO users (username, email, password)
VALUES ($1, $2, $3)
RETURNING id, username, email, created_at, updated_at;

-- name: GetUserByEmail :one
SELECT id, username, email, password, created_at, updated_at
FROM users
WHERE email = $1;
