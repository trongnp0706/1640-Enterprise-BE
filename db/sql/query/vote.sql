-- name: CreateVote :one
INSERT INTO votes(
    id,
    user_id,
    idea_id,
    vote
) VALUES (
    $1,  $2,  $3, $4
)
RETURNING *;

-- name: GetVote :one
SELECT * FROM votes WHERE user_id = $1 AND idea_id = $2;

-- name: UpdateVote :one
UPDATE votes 
SET  vote = $1
WHERE id = $2
    RETURNING *;

-- name: DeleteVote :one
DELETE FROM votes
WHERE id = $1
    RETURNING *;