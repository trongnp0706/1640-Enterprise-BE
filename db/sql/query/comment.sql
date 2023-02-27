-- name: CreateComment :one
INSERT INTO comments(
    id, 
    content,   
    is_anonymous, 
    user_id,
    idea_id,
    created_at
) VALUES (
    $1,  $2,  $3, $4, $5, $6
)
RETURNING *;

-- name: GetCommentsByIdea :many
Select * FROM comments 
WHERE idea_id = $1
LIMIT $2
OFFSET $3;

-- name: GetLatestComment :many
Select * FROM comments 
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: UpdateComment :one
UPDATE comments
SET  title = $1,
    content = $2,
    is_anonymous = $3,
    academic_year = $4
WHERe id = $5
    RETURNING *;

-- name: DeleteComment :one
DELETE FROM comments
WHERE id = $1
    RETURNING *;