package repository

import (
	entity "GDN-delivery-management/db/sql"
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/lib/pq"
)

type IDepartmentRepo interface {
	CreateDepartment(ctx context.Context, input entity.CreateDepartmentParams) (error, entity.Department)
	GetDepartmentById(ctx context.Context, ticker string) (error, entity.Department)
	GetAllDepartments(ctx context.Context) (error, []entity.Department)
	UpdateDepartment(ctx context.Context, input entity.UpdateDepartmentParams) (error, entity.Department)
	DeleteDepartment(ctx context.Context, ticker string) (error, entity.Department)
}

type DepartmentRepo struct {
	sql *entity.Queries
}

func NewDepartmentRepo(sql *entity.Queries) IDepartmentRepo {
	return &DepartmentRepo{
		sql: sql,
	}
}

func (d *DepartmentRepo) GetAllDepartments(ctx context.Context) (error, []entity.Department) {
	departments, err := d.sql.GetAllDepartments(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Departments are not available"), []entity.Department{}
		}
		return err, []entity.Department{}
	}
	return nil, departments
}

func (d *DepartmentRepo) CreateDepartment(ctx context.Context, input entity.CreateDepartmentParams) (error, entity.Department) {
	departments, err := d.sql.CreateDepartment(ctx, input)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errors.New("Conflict id"), entity.Department{}
			}
		}
		log.Println("err: ", err)
		return errors.New("Failed to add role"), entity.Department{}
	}
	return nil, departments
}

func (d *DepartmentRepo) GetDepartmentById(ctx context.Context, ticker string) (error, entity.Department) {
	departments, err := d.sql.GetDepartmentById(ctx, ticker)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("id not found"), entity.Department{}
		}
		return err, entity.Department{}
	}
	return nil, departments
}

func (d *DepartmentRepo) UpdateDepartment(ctx context.Context, input entity.UpdateDepartmentParams) (error, entity.Department) {
	departments, err := d.sql.UpdateDepartment(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("id not found"), entity.Department{}
		}
		return err, entity.Department{}
	}
	return nil, departments
}

func (d *DepartmentRepo) DeleteDepartment(ctx context.Context, ticker string) (error, entity.Department) {
	departments, err := d.sql.DeleteDepartment(ctx, ticker)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("id not found"), entity.Department{}
		}
		return err, entity.Department{}
	}
	return nil, departments
}
