// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: idea.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createIdea = `-- name: CreateIdea :one
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
    CAST($6 AS VARCHAR[]),
    $7, $8, $9, $10, $11, $12, $13
)
RETURNING id, title, content, view_count, document_array, image_array, upvote_count, downvote_count, is_anonymous, user_id, category_id, academic_year, created_at
`

type CreateIdeaParams struct {
	ID            string      `json:"id"`
	Title         string      `json:"title"`
	Content       string      `json:"content"`
	ViewCount     int32       `json:"view_count"`
	Column5       interface{} `json:"column_5"`
	Column6       []string    `json:"column_6"`
	UpvoteCount   int32       `json:"upvote_count"`
	DownvoteCount int32       `json:"downvote_count"`
	IsAnonymous   bool        `json:"is_anonymous"`
	UserID        string      `json:"user_id"`
	CategoryID    string      `json:"category_id"`
	AcademicYear  string      `json:"academic_year"`
	CreatedAt     time.Time   `json:"created_at"`
}

func (q *Queries) CreateIdea(ctx context.Context, arg CreateIdeaParams) (Idea, error) {
	row := q.db.QueryRowContext(ctx, createIdea,
		arg.ID,
		arg.Title,
		arg.Content,
		arg.ViewCount,
		arg.Column5,
		pq.Array(arg.Column6),
		arg.UpvoteCount,
		arg.DownvoteCount,
		arg.IsAnonymous,
		arg.UserID,
		arg.CategoryID,
		arg.AcademicYear,
		arg.CreatedAt,
	)
	var i Idea
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.ViewCount,
		&i.DocumentArray,
		pq.Array(&i.ImageArray),
		&i.UpvoteCount,
		&i.DownvoteCount,
		&i.IsAnonymous,
		&i.UserID,
		&i.CategoryID,
		&i.AcademicYear,
		&i.CreatedAt,
	)
	return i, err
}

const decreaseDownvoteCount = `-- name: DecreaseDownvoteCount :one
UPDATE ideas
SET  downvote_count = downvote_count - 1
WHERE id = $1
    RETURNING id, title, content, view_count, document_array, image_array, upvote_count, downvote_count, is_anonymous, user_id, category_id, academic_year, created_at
`

func (q *Queries) DecreaseDownvoteCount(ctx context.Context, id string) (Idea, error) {
	row := q.db.QueryRowContext(ctx, decreaseDownvoteCount, id)
	var i Idea
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.ViewCount,
		&i.DocumentArray,
		pq.Array(&i.ImageArray),
		&i.UpvoteCount,
		&i.DownvoteCount,
		&i.IsAnonymous,
		&i.UserID,
		&i.CategoryID,
		&i.AcademicYear,
		&i.CreatedAt,
	)
	return i, err
}

const decreaseUpvoteCount = `-- name: DecreaseUpvoteCount :one
UPDATE ideas
SET  upvote_count = upvote_count - 1
WHERE id = $1
    RETURNING id, title, content, view_count, document_array, image_array, upvote_count, downvote_count, is_anonymous, user_id, category_id, academic_year, created_at
`

func (q *Queries) DecreaseUpvoteCount(ctx context.Context, id string) (Idea, error) {
	row := q.db.QueryRowContext(ctx, decreaseUpvoteCount, id)
	var i Idea
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.ViewCount,
		&i.DocumentArray,
		pq.Array(&i.ImageArray),
		&i.UpvoteCount,
		&i.DownvoteCount,
		&i.IsAnonymous,
		&i.UserID,
		&i.CategoryID,
		&i.AcademicYear,
		&i.CreatedAt,
	)
	return i, err
}

const deleteIdea = `-- name: DeleteIdea :one
DELETE FROM ideas
WHERE id = $1
    RETURNING id, title, content, view_count, document_array, image_array, upvote_count, downvote_count, is_anonymous, user_id, category_id, academic_year, created_at
`

func (q *Queries) DeleteIdea(ctx context.Context, id string) (Idea, error) {
	row := q.db.QueryRowContext(ctx, deleteIdea, id)
	var i Idea
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.ViewCount,
		&i.DocumentArray,
		pq.Array(&i.ImageArray),
		&i.UpvoteCount,
		&i.DownvoteCount,
		&i.IsAnonymous,
		&i.UserID,
		&i.CategoryID,
		&i.AcademicYear,
		&i.CreatedAt,
	)
	return i, err
}

const getDownvoteCount = `-- name: GetDownvoteCount :one
SELECT downvote_count FROM ideas WHERE id = $1
`

func (q *Queries) GetDownvoteCount(ctx context.Context, id string) (int32, error) {
	row := q.db.QueryRowContext(ctx, getDownvoteCount, id)
	var downvote_count int32
	err := row.Scan(&downvote_count)
	return downvote_count, err
}

const getIdea = `-- name: GetIdea :one
SELECT id, title, content, view_count, document_array, image_array, upvote_count, downvote_count, is_anonymous, user_id, category_id, academic_year, created_at FROM ideas WHERE id=$1
`

func (q *Queries) GetIdea(ctx context.Context, id string) (Idea, error) {
	row := q.db.QueryRowContext(ctx, getIdea, id)
	var i Idea
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.ViewCount,
		&i.DocumentArray,
		pq.Array(&i.ImageArray),
		&i.UpvoteCount,
		&i.DownvoteCount,
		&i.IsAnonymous,
		&i.UserID,
		&i.CategoryID,
		&i.AcademicYear,
		&i.CreatedAt,
	)
	return i, err
}

const getIdeaByAcademicyear = `-- name: GetIdeaByAcademicyear :many
SELECT id, title, content, view_count, document_array, image_array, upvote_count, downvote_count, is_anonymous, user_id, category_id, academic_year, created_at FROM ideas WHERE academic_year = $1
LIMIT $2
OFFSET $3
`

type GetIdeaByAcademicyearParams struct {
	AcademicYear string `json:"academic_year"`
	Limit        int32  `json:"limit"`
	Offset       int32  `json:"offset"`
}

func (q *Queries) GetIdeaByAcademicyear(ctx context.Context, arg GetIdeaByAcademicyearParams) ([]Idea, error) {
	rows, err := q.db.QueryContext(ctx, getIdeaByAcademicyear, arg.AcademicYear, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Idea
	for rows.Next() {
		var i Idea
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.ViewCount,
			&i.DocumentArray,
			pq.Array(&i.ImageArray),
			&i.UpvoteCount,
			&i.DownvoteCount,
			&i.IsAnonymous,
			&i.UserID,
			&i.CategoryID,
			&i.AcademicYear,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getIdeaByCategory = `-- name: GetIdeaByCategory :many
SELECT id, title, content, view_count, document_array, image_array, upvote_count, downvote_count, is_anonymous, user_id, category_id, academic_year, created_at FROM ideas WHERE category_id = $1
LIMIT $2
OFFSET $3
`

type GetIdeaByCategoryParams struct {
	CategoryID string `json:"category_id"`
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
}

func (q *Queries) GetIdeaByCategory(ctx context.Context, arg GetIdeaByCategoryParams) ([]Idea, error) {
	rows, err := q.db.QueryContext(ctx, getIdeaByCategory, arg.CategoryID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Idea
	for rows.Next() {
		var i Idea
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.ViewCount,
			&i.DocumentArray,
			pq.Array(&i.ImageArray),
			&i.UpvoteCount,
			&i.DownvoteCount,
			&i.IsAnonymous,
			&i.UserID,
			&i.CategoryID,
			&i.AcademicYear,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLatestIdeas = `-- name: GetLatestIdeas :many
SELECT ideas.id, ideas.title, ideas.content, ideas.view_count, ideas.document_array, ideas.image_array, ideas.upvote_count, ideas.downvote_count, ideas.is_anonymous, ideas.user_id, ideas.category_id, ideas.academic_year, ideas.created_at, users.avatar, users.username
FROM ideas
         INNER JOIN users ON ideas.user_id = users.id
ORDER BY ideas.created_at DESC
    LIMIT $1 OFFSET $2
`

type GetLatestIdeasParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetLatestIdeasRow struct {
	ID            string         `json:"id"`
	Title         string         `json:"title"`
	Content       string         `json:"content"`
	ViewCount     int32          `json:"view_count"`
	DocumentArray sql.NullString `json:"document_array"`
	ImageArray    []string       `json:"image_array"`
	UpvoteCount   int32          `json:"upvote_count"`
	DownvoteCount int32          `json:"downvote_count"`
	IsAnonymous   bool           `json:"is_anonymous"`
	UserID        string         `json:"user_id"`
	CategoryID    string         `json:"category_id"`
	AcademicYear  string         `json:"academic_year"`
	CreatedAt     time.Time      `json:"created_at"`
	Avatar        string         `json:"avatar"`
	Username      string         `json:"username"`
}

func (q *Queries) GetLatestIdeas(ctx context.Context, arg GetLatestIdeasParams) ([]GetLatestIdeasRow, error) {
	rows, err := q.db.QueryContext(ctx, getLatestIdeas, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLatestIdeasRow
	for rows.Next() {
		var i GetLatestIdeasRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.ViewCount,
			&i.DocumentArray,
			pq.Array(&i.ImageArray),
			&i.UpvoteCount,
			&i.DownvoteCount,
			&i.IsAnonymous,
			&i.UserID,
			&i.CategoryID,
			&i.AcademicYear,
			&i.CreatedAt,
			&i.Avatar,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMostPopularIdeas = `-- name: GetMostPopularIdeas :many
Select id, title, content, view_count, document_array, image_array, upvote_count, downvote_count, is_anonymous, user_id, category_id, academic_year, created_at FROM ideas ORDER BY upvote_count DESC
LIMIT $1
OFFSET $2
`

type GetMostPopularIdeasParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetMostPopularIdeas(ctx context.Context, arg GetMostPopularIdeasParams) ([]Idea, error) {
	rows, err := q.db.QueryContext(ctx, getMostPopularIdeas, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Idea
	for rows.Next() {
		var i Idea
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.ViewCount,
			&i.DocumentArray,
			pq.Array(&i.ImageArray),
			&i.UpvoteCount,
			&i.DownvoteCount,
			&i.IsAnonymous,
			&i.UserID,
			&i.CategoryID,
			&i.AcademicYear,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMostViewedIdeas = `-- name: GetMostViewedIdeas :many
Select id, title, content, view_count, document_array, image_array, upvote_count, downvote_count, is_anonymous, user_id, category_id, academic_year, created_at FROM ideas ORDER BY view_count DESC
LIMIT $1
OFFSET $2
`

type GetMostViewedIdeasParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetMostViewedIdeas(ctx context.Context, arg GetMostViewedIdeasParams) ([]Idea, error) {
	rows, err := q.db.QueryContext(ctx, getMostViewedIdeas, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Idea
	for rows.Next() {
		var i Idea
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.ViewCount,
			&i.DocumentArray,
			pq.Array(&i.ImageArray),
			&i.UpvoteCount,
			&i.DownvoteCount,
			&i.IsAnonymous,
			&i.UserID,
			&i.CategoryID,
			&i.AcademicYear,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNumberOfAllIdeas = `-- name: GetNumberOfAllIdeas :one
SELECT COUNT(*) 
FROM ideas
`

func (q *Queries) GetNumberOfAllIdeas(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getNumberOfAllIdeas)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getNumberOfIdeasByDepartment = `-- name: GetNumberOfIdeasByDepartment :one
SELECT COUNT(*) 
FROM ideas
WHERE user_id IN (
    SELECT user_id 
    FROM users 
    WHERE department_id = $1
)
`

func (q *Queries) GetNumberOfIdeasByDepartment(ctx context.Context, departmentID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getNumberOfIdeasByDepartment, departmentID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUpvoteCount = `-- name: GetUpvoteCount :one
SELECT upvote_count FROM ideas WHERE id = $1
`

func (q *Queries) GetUpvoteCount(ctx context.Context, id string) (int32, error) {
	row := q.db.QueryRowContext(ctx, getUpvoteCount, id)
	var upvote_count int32
	err := row.Scan(&upvote_count)
	return upvote_count, err
}

const increaseDownvoteCount = `-- name: IncreaseDownvoteCount :one
UPDATE ideas
SET  downvote_count = downvote_count + 1
WHERE id = $1
    RETURNING id, title, content, view_count, document_array, image_array, upvote_count, downvote_count, is_anonymous, user_id, category_id, academic_year, created_at
`

func (q *Queries) IncreaseDownvoteCount(ctx context.Context, id string) (Idea, error) {
	row := q.db.QueryRowContext(ctx, increaseDownvoteCount, id)
	var i Idea
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.ViewCount,
		&i.DocumentArray,
		pq.Array(&i.ImageArray),
		&i.UpvoteCount,
		&i.DownvoteCount,
		&i.IsAnonymous,
		&i.UserID,
		&i.CategoryID,
		&i.AcademicYear,
		&i.CreatedAt,
	)
	return i, err
}

const increaseUpvoteCount = `-- name: IncreaseUpvoteCount :one
UPDATE ideas
SET  upvote_count = upvote_count + 1
WHERE id = $1
    RETURNING id, title, content, view_count, document_array, image_array, upvote_count, downvote_count, is_anonymous, user_id, category_id, academic_year, created_at
`

func (q *Queries) IncreaseUpvoteCount(ctx context.Context, id string) (Idea, error) {
	row := q.db.QueryRowContext(ctx, increaseUpvoteCount, id)
	var i Idea
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.ViewCount,
		&i.DocumentArray,
		pq.Array(&i.ImageArray),
		&i.UpvoteCount,
		&i.DownvoteCount,
		&i.IsAnonymous,
		&i.UserID,
		&i.CategoryID,
		&i.AcademicYear,
		&i.CreatedAt,
	)
	return i, err
}

const updateIdea = `-- name: UpdateIdea :one
UPDATE ideas
SET  title = $1,
    content = $2,
    document_array = $3,
    image_array = $4,
    is_anonymous = $5,
    academic_year = $6,
    category_id = $7
WHERE id = $8
    RETURNING id, title, content, view_count, document_array, image_array, upvote_count, downvote_count, is_anonymous, user_id, category_id, academic_year, created_at
`

type UpdateIdeaParams struct {
	Title         string         `json:"title"`
	Content       string         `json:"content"`
	DocumentArray string  		 `json:"document_array"`
	ImageArray    []string       `json:"image_array"`
	IsAnonymous   bool           `json:"is_anonymous"`
	AcademicYear  string         `json:"academic_year"`
	CategoryID    string         `json:"category_id"`
	ID            string         `json:"id"`
}

func (q *Queries) UpdateIdea(ctx context.Context, arg UpdateIdeaParams) (Idea, error) {
	row := q.db.QueryRowContext(ctx, updateIdea,
		arg.Title,
		arg.Content,
		arg.DocumentArray,
		pq.Array(arg.ImageArray),
		arg.IsAnonymous,
		arg.AcademicYear,
		arg.CategoryID,
		arg.ID,
	)
	var i Idea
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.ViewCount,
		&i.DocumentArray,
		pq.Array(&i.ImageArray),
		&i.UpvoteCount,
		&i.DownvoteCount,
		&i.IsAnonymous,
		&i.UserID,
		&i.CategoryID,
		&i.AcademicYear,
		&i.CreatedAt,
	)
	return i, err
}
