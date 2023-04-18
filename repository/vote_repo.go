package repository

import (
	entity "GDN-delivery-management/db/sql"
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/lib/pq"
)

type IVoteRepo interface {
	AddVote(ctx context.Context, input entity.CreateVoteParams) (error, entity.Vote)
	GetVote(ctx context.Context, input entity.GetVoteParams) (error, entity.Vote)
	UpdateVote(ctx context.Context, input entity.UpdateVoteParams) (error, entity.Vote)
	DeleteVote(ctx context.Context, id string) (error, entity.Vote)
}

type VoteRepo struct {
	sql *entity.Queries
}

func NewVoteRepo(sql *entity.Queries) IVoteRepo {
	return &VoteRepo{
		sql: sql,
	}
}

func (i *VoteRepo) AddVote(ctx context.Context, input entity.CreateVoteParams) (error, entity.Vote) {
	vote, err := i.sql.CreateVote(ctx, input)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errors.New("Conflict email"), entity.Vote{}
			}
		}
		log.Println("err: ", err)
		return errors.New("Failed to add vote"), entity.Vote{}
	}
	return nil, vote
}

func (u *VoteRepo) GetVote(ctx context.Context, input entity.GetVoteParams) (error, entity.Vote) {
	vote, err := u.sql.GetVote(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("vote not found"), entity.Vote{}
		}
		return err, entity.Vote{}
	}
	return nil, vote
}

func (u *VoteRepo) UpdateVote(ctx context.Context, input entity.UpdateVoteParams) (error, entity.Vote) {
	vote, err := u.sql.UpdateVote(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("vote not found"), entity.Vote{}
		}
		return err, entity.Vote{}
	}
	return nil, vote
}

func (u *VoteRepo) DeleteVote(ctx context.Context, id string) (error, entity.Vote) {
	vote, err := u.sql.DeleteVote(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("vote not found"), entity.Vote{}
		}
		return err, entity.Vote{}
	}
	return nil, vote
}
