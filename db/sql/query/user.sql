-- name: CreateUser :one
INSERT INTO users(
    id, username, email, password, avatar, role_ticker, department_id
) VALUES (
    $1,  $2,  $3,  $4, $5, $6, $7
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetAllUsers :many
Select * FROM users ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET username = $1, email = $2, password = $3, role_ticker = $4, department_id = $5, avatar = $6
WHERE id = $7
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;