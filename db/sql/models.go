// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
	"time"
)

type AcademicYear struct {
	AcademicYear string    `json:"academic_year"`
	ClosureDate  time.Time `json:"closure_date"`
}

type Category struct {
	ID           string `json:"id"`
	CategoryName string `json:"category_name"`
}

type Comment struct {
	ID          string    `json:"id"`
	Content     string    `json:"content"`
	IsAnonymous bool      `json:"is_anonymous"`
	UserID      string    `json:"user_id"`
	IdeaID      string    `json:"idea_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type Department struct {
	ID             string `json:"id"`
	DepartmentName string `json:"department_name"`
}

type Idea struct {
	ID            string         `json:"id"`
	Title         string         `json:"title"`
	Content       string         `json:"content"`
	ViewCount     int32          `json:"view_count"`
	DocumentArray sql.NullString `json:"document_array"`
	ImageArray    sql.NullString `json:"image_array"`
	UpvoteCount   int32          `json:"upvote_count"`
	DownvoteCount int32          `json:"downvote_count"`
	IsAnonymous   bool           `json:"is_anonymous"`
	UserID        string         `json:"user_id"`
	CategoryID    string         `json:"category_id"`
	AcademicYear  string         `json:"academic_year"`
	CreatedAt     time.Time      `json:"created_at"`
}

type Role struct {
	Ticker   string `json:"ticker"`
	RoleName string `json:"role_name"`
}

type Session struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    int64     `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Avatar       string `json:"avatar"`
	RoleTicker   string `json:"role_ticker"`
	DepartmentID string `json:"department_id"`
}

type Vote struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	IdeaID string `json:"idea_id"`
	Vote   bool   `json:"vote"`
}
