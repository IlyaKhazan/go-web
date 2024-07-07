-- name: CreateUser :one
INSERT INTO users (id, firstname, lastname, email, age)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY $1
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET firstname = $2, lastname = $3, email = $4, age = $5
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
