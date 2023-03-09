package repository

import (
	entity "GDN-delivery-management/db/sql"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type ICommentRepo interface {
	AddComment(ctx context.Context, input entity.CreateCommentParams) (error, entity.Comment)
	GetCommentsByIdea(ctx context.Context, input entity.GetCommentsByIdeaParams) (error, []entity.Comment)
	GetLatestComment(ctx context.Context, input entity.GetLatestCommentParams) (error, []entity.Comment)
	UpdateComment(ctx context.Context, input entity.UpdateCommentParams) (error, entity.Comment)
	DeleteComment(ctx context.Context, id string) (error, entity.Comment)
}

type CommentRepo struct {
	sql *entity.Queries
}

func NewCommentRepo(sql *entity.Queries) ICommentRepo {
	return &CommentRepo{
		sql: sql,
	}
}

func (i *CommentRepo) AddComment(ctx context.Context, input entity.CreateCommentParams) (error, entity.Comment) {
	comment, err := i.sql.CreateComment(ctx, input)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errors.New("Conflict id"), entity.Comment{}
			}
		}
		log.Println("err: ", err)
		return errors.New("Failed to add comment"), entity.Comment{}
	}
	return nil, comment
}

func (i *CommentRepo) GetCommentsByIdea(ctx context.Context, input entity.GetCommentsByIdeaParams) (error, []entity.Comment) {
	comment, err := i.sql.GetCommentsByIdea(ctx, input)
	if err != nil {
		log.Println("err", err)
		if err == sql.ErrNoRows {
			return errors.New("comment not found"), []entity.Comment{}
		}
		return err, []entity.Comment{}
	}
	return nil, comment
}

func (u *CommentRepo) GetLatestComment(ctx context.Context, input entity.GetLatestCommentParams) (error, []entity.Comment) {
	items, err := u.sql.GetLatestComment(ctx, input)
	if err != nil {
		fmt.Printf(err.Error())
		return errors.New("Cannot get all comments"), []entity.Comment{}
	}
	return nil, items
}

func (u *CommentRepo) UpdateComment(ctx context.Context, input entity.UpdateCommentParams) (error, entity.Comment) {
	comment, err := u.sql.UpdateComment(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("comment not found"), entity.Comment{}
		}
		return err, entity.Comment{}
	}
	return nil, comment
}

func (u *CommentRepo) DeleteComment(ctx context.Context, id string) (error, entity.Comment) {
	comment, err := u.sql.DeleteComment(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("comment not found"), entity.Comment{}
		}
		return err, entity.Comment{}
	}
	return nil, comment
}
