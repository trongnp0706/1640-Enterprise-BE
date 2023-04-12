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
Select comments.*, users.avatar, users.username
FROM comments
         INNER JOIN users ON comments.user_id = users.id
WHERE idea_id = $1;

-- name: UpdateComment :one
UPDATE comments
SET content = $1,
    is_anonymous = $2
WHERE id = $3
    RETURNING *;

-- name: DeleteComment :one
DELETE FROM comments
WHERE id = $1
    RETURNING *;