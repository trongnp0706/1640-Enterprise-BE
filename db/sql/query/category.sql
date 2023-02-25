-- name: CreateCategory :one
INSERT INTO categories(
    id, category_name
)
VALUES (
           $1,  $2
       )
    RETURNING *;

-- name: GetAllCategories :many
SELECT id, category_name FROM categories;

-- name: GetCategoryById :one
SELECT * FROM categories WHERE id = $1;

-- name: UpdateCategory :one
UPDATE categories
SET category_name = $1,
    id = $2
WHERE id = $3
    RETURNING *;

-- name: DeleteCategory :one
DELETE FROM categories
WHERE id = $1
    RETURNING *;
