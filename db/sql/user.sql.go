// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(
    id, username, email, password, role_ticker, department_id
) VALUES (
    $1,  $2,  $3,  $4, $5, $6
)
RETURNING id, username, email, password, role_ticker, department_id
`

type CreateUserParams struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RoleTicker   string `json:"role_ticker"`
	DepartmentID string `json:"department_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.RoleTicker,
		arg.DepartmentID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.RoleTicker,
		&i.DepartmentID,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id, username, email, password, role_ticker, department_id
`

func (q *Queries) DeleteUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.RoleTicker,
		&i.DepartmentID,
	)
	return i, err
}

const getAllUsers = `-- name: GetAllUsers :many
Select id, username, email, password, role_ticker, department_id FROM users ORDER BY created_at
LIMIT $1
OFFSET $2
`

type GetAllUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetAllUsersRow struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RoleTicker   string `json:"role_ticker"`
	DepartmentID string `json:"department_id"`
}

func (q *Queries) GetAllUsers(ctx context.Context, arg GetAllUsersParams) ([]GetAllUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllUsersRow
	for rows.Next() {
		var i GetAllUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.RoleTicker,
			&i.DepartmentID,
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

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password, role_ticker, department_id FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.RoleTicker,
		&i.DepartmentID,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, email, password, role_ticker, department_id FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.RoleTicker,
		&i.DepartmentID,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET username = $1, email = $2, password = $3, role_ticker = $4, department_id = $5
WHERE id = $6
RETURNING id, username, email, password, role_ticker, department_id
`

type UpdateUserParams struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RoleTicker   string `json:"role_ticker"`
	DepartmentID string `json:"department_id"`
	ID           string `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.RoleTicker,
		arg.DepartmentID,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.RoleTicker,
		&i.DepartmentID,
	)
	return i, err
}
