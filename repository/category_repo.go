package repository

import (
	entity "GDN-delivery-management/db/sql"
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/lib/pq"
)

type ICategoryRepo interface {
	GetAllCategories(ctx context.Context) (error, []entity.Category)
	CreateCategory(ctx context.Context, input entity.CreateCategoryParams) (error, entity.Category)
	GetCategoryById(ctx context.Context, ticker string) (error, entity.Category)
	UpdateCategory(ctx context.Context, input entity.UpdateCategoryParams) (error, entity.Category)
	DeleteCategory(ctx context.Context, ticker string) (error, entity.Category)
}

type CategoryRepo struct {
	sql *entity.Queries
}

func NewCategoryRepo(sql *entity.Queries) ICategoryRepo {
	return &CategoryRepo{
		sql: sql,
	}
}

func (c *CategoryRepo) GetAllCategories(ctx context.Context) (error, []entity.Category) {
	categories, err := c.sql.GetAllCategories(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Categories are not available"), []entity.Category{}
		}
		return err, []entity.Category{}
	}
	return nil, categories
}

func (c *CategoryRepo) CreateCategory(ctx context.Context, input entity.CreateCategoryParams) (error, entity.Category) {
	category, err := c.sql.CreateCategory(ctx, input)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errors.New("Conflict id"), entity.Category{}
			}
		}
		log.Println("err: ", err)
		return errors.New("Failed to add category"), entity.Category{}
	}
	return nil, category
}

func (c *CategoryRepo) GetCategoryById(ctx context.Context, ticker string) (error, entity.Category) {
	category, err := c.sql.GetCategoryById(ctx, ticker)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("id not found"), entity.Category{}
		}
		return err, entity.Category{}
	}
	return nil, category
}

func (c *CategoryRepo) UpdateCategory(ctx context.Context, input entity.UpdateCategoryParams) (error, entity.Category) {
	category, err := c.sql.UpdateCategory(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("id not found"), entity.Category{}
		}
		return err, entity.Category{}
	}
	return nil, category
}

func (c *CategoryRepo) DeleteCategory(ctx context.Context, ticker string) (error, entity.Category) {
	category, err := c.sql.DeleteCategory(ctx, ticker)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("id not found"), entity.Category{}
		}
		return err, entity.Category{}
	}
	return nil, category
}
