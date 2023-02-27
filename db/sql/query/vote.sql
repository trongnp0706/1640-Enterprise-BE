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

-- name: UpdateVoteUP :one
UPDATE votes 
SET  vote = TRUE
WHERE id = $1
    RETURNING *;

-- name: UpdateVoteDOWN :one
UPDATE votes 
SET  vote = FALSE
WHERE id = $1
    RETURNING *;

-- name: DeleteVote :one
DELETE FROM votes
WHERE id = $1
    RETURNING *;