-- name: CreateRole :one
INSERT INTO roles(
    ticker, role_name
)
VALUES (
           $1,  $2
       )
    RETURNING *;

-- name: ListRoles :many
SELECT role_name, ticker FROM roles;

-- name: GetRoleByTicker :one
SELECT * FROM roles WHERE ticker = $1;

-- name: UpdateRole :one
UPDATE roles
SET role_name = $1, ticker = $2
WHERE ticker = $3
    RETURNING *;

-- name: DeleteRole :one
DELETE FROM roles
WHERE ticker = $1
    RETURNING *;
