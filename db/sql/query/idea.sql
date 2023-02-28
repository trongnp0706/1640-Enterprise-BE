-- name: CreateIdea :one
INSERT INTO ideas(
    id, 
    title, 
    content, 
    view_count, 
    document_array, 
    image_array, 
    upvote_count, 
    downvote_count, 
    is_anonymous, 
    user_id, 
    category_id, 
    academic_year, 
    created_at
) VALUES (
    $1,  $2,  $3,  $4, 
    CASE WHEN $5 = '' THEN 'null' ELSE $5 END, 
    CASE WHEN $6 = '' THEN 'null' ELSE $6 END,
    $7, $8, $9, $10, $11, $12, $13
)
RETURNING *;

-- name: GetMostPopularIdeas :many
Select * FROM ideas ORDER BY upvote_count DESC
LIMIT $1
OFFSET $2;

-- name: GetMostViewedIdeas :many
Select * FROM ideas ORDER BY view_count DESC
LIMIT $1
OFFSET $2;

-- name: GetLatestIdeas :many
Select * FROM ideas ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetNumberOfAllIdeas :one
SELECT COUNT(*) 
FROM ideas;

-- name: GetNumberOfIdeasByDepartment :one
SELECT COUNT(*) 
FROM ideas
WHERE user_id IN (
    SELECT user_id 
    FROM users 
    WHERE department_id = $1
);

-- name: GetIdea :one
SELECT * FROM ideas WHERE id=$1;

-- name: GetIdeaByCategory :many
SELECT * FROM ideas WHERE category_id = $1
LIMIT $2
OFFSET $3;

-- name: GetIdeaByAcademicyear :many
SELECT * FROM ideas WHERE academic_year = $1
LIMIT $2
OFFSET $3;

-- name: UpdateIdea :one
UPDATE ideas
SET  title = $1,
    content = $2,
    document_array = $3,
    image_array = $4,
    is_anonymous = $5,
    academic_year =  $6
WHERE id = $7
    RETURNING *;

-- name: DeleteIdea :one
DELETE FROM ideas
WHERE id = $1
    RETURNING *;

