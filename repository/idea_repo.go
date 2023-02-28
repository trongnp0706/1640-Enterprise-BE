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

type IIdeaRepo interface {
	AddIdea(ctx context.Context, imput entity.CreateIdeaParams) (error, entity.Idea)
	GetNumberOfAllIdeas(ctx context.Context) (error, int64)
	GetIdeaByCategory(ctx context.Context, input entity.GetIdeaByCategoryParams) (error, []entity.Idea)
	GetNumberOfIdeasByDepartment(ctx context.Context, department_id string) (error, int64)
	GetIdeaByAcademicyear(ctx context.Context, input entity.GetIdeaByAcademicyearParams) (error, []entity.Idea)
	GetIdea(ctx context.Context, id string) (error, entity.Idea)
	GetMostPopularIdeas(ctx context.Context, input entity.GetMostPopularIdeasParams) (error, []entity.Idea)
	GetMostViewedIdeas(ctx context.Context, input entity.GetMostViewedIdeasParams) (error, []entity.Idea)
	GetLatestIdeas(ctx context.Context, input entity.GetLatestIdeasParams) (error, []entity.Idea)
	UpdateIdea(ctx context.Context, input entity.UpdateIdeaParams) (error, entity.Idea)
	DeleteIdea(ctx context.Context, id string) (error, entity.Idea)
}

type IdeaRepo struct {
	sql *entity.Queries
}

func NewIdeaRepo(sql *entity.Queries) IIdeaRepo {
	return &IdeaRepo{
		sql: sql,
	}
}

func (i *IdeaRepo) AddIdea(ctx context.Context, imput entity.CreateIdeaParams) (error, entity.Idea) {
	idea, err := i.sql.CreateIdea(ctx, imput)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errors.New("Conflict email"), entity.Idea{}
			}
		}
		log.Println("err: ", err)
		return errors.New("Failed to add idea"), entity.Idea{}
	}
	return nil, idea
}

func (i *IdeaRepo) GetNumberOfAllIdeas(ctx context.Context) (error, int64) {
	number, err := i.sql.GetNumberOfAllIdeas(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("idea not found"), number
		}
		return err, number
	}
	return nil, number
}

func (i *IdeaRepo)  GetNumberOfIdeasByDepartment(ctx context.Context, department_id string) (error, int64) {
	number, err := i.sql. GetNumberOfIdeasByDepartment(ctx, department_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("idea not found"), number
		}
		return err, number
	}
	return nil, number
}

func (i *IdeaRepo) GetIdea(ctx context.Context, id string) (error, entity.Idea) {
	idea, err := i.sql.GetIdea(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("idea not found"), entity.Idea{}
		}
		return err, entity.Idea{}
	}
	return nil, idea
}

func (i *IdeaRepo) GetIdeaByCategory(ctx context.Context, input entity.GetIdeaByCategoryParams) (error, []entity.Idea) {
	idea, err := i.sql.GetIdeaByCategory(ctx, input)
	if err != nil {
		log.Println("err", err)
		if err == sql.ErrNoRows {
			return errors.New("idea not found"), []entity.Idea{}
		}
		return err, []entity.Idea{}
	}
	return nil, idea
}

func (i *IdeaRepo) GetIdeaByAcademicyear(ctx context.Context, input entity.GetIdeaByAcademicyearParams) (error, []entity.Idea) {
	idea, err := i.sql.GetIdeaByAcademicyear(ctx, input)
	if err != nil {
		log.Println("err", err)
		if err == sql.ErrNoRows {
			return errors.New("idea not found"), []entity.Idea{}
		}
		return err, []entity.Idea{}
	}
	return nil, idea
}

func (u *IdeaRepo) GetMostPopularIdeas(ctx context.Context, input entity.GetMostPopularIdeasParams) (error, []entity.Idea) {
	items, err := u.sql.GetMostPopularIdeas(ctx, input)
	if err != nil {
		fmt.Printf(err.Error())
		return errors.New("Cannot get all ideas"), []entity.Idea{}
	}
	return nil, items
}

func (u *IdeaRepo) GetMostViewedIdeas(ctx context.Context, input entity.GetMostViewedIdeasParams) (error, []entity.Idea) {
	items, err := u.sql.GetMostViewedIdeas(ctx, input)
	if err != nil {
		fmt.Printf(err.Error())
		return errors.New("Cannot get all ideas"), []entity.Idea{}
	}
	return nil, items
}


func (u *IdeaRepo) GetLatestIdeas(ctx context.Context, input entity.GetLatestIdeasParams) (error, []entity.Idea) {
	items, err := u.sql.GetLatestIdeas(ctx, input)
	if err != nil {
		fmt.Printf(err.Error())
		return errors.New("Cannot get all ideas"), []entity.Idea{}
	}
	return nil, items
}

func (u *IdeaRepo) UpdateIdea(ctx context.Context, input entity.UpdateIdeaParams) (error, entity.Idea) {
	idea, err := u.sql.UpdateIdea(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("idea not found"), entity.Idea{}
		}
		return err, entity.Idea{}
	}
	return nil, idea
}

func (u *IdeaRepo) DeleteIdea(ctx context.Context, id string) (error, entity.Idea) {
	idea, err := u.sql.DeleteIdea(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("idea not found"), entity.Idea{}
		}
		return err, entity.Idea{}
	}
	return nil, idea
}
