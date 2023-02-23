package repository

import (
	entity "GDN-delivery-management/db/sql"
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/lib/pq"
)

type IRoleRepo interface {
	AddRole(ctx context.Context, input entity.CreateRoleParams) (error, entity.Role)
	GetRoleByTicker(ctx context.Context, ticker string) (error, entity.Role)
	ListRoles(ctx context.Context) (error, []entity.ListRolesRow)
	UpdateRole(ctx context.Context, input entity.UpdateRoleParams) (error, entity.Role)
	DeleteRole(ctx context.Context, ticker string) (error, entity.Role)
}

type RoleRepo struct {
	sql *entity.Queries
}

func NewRoleRepo(sql *entity.Queries) IRoleRepo {
	return &RoleRepo{
		sql: sql,
	}
}

func (r *RoleRepo) ListRoles(ctx context.Context) (error, []entity.ListRolesRow) {
	roles, err := r.sql.ListRoles(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Roles are not avilable"), []entity.ListRolesRow{}
		}
		return err, []entity.ListRolesRow{}
	}
	return nil, roles
}

func (r *RoleRepo) AddRole(ctx context.Context, input entity.CreateRoleParams) (error, entity.Role) {
	role, err := r.sql.CreateRole(ctx, input)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errors.New("Conflict ticker"), entity.Role{}
			}
		}
		log.Println("err: ", err)
		return errors.New("Failed to add role"), entity.Role{}
	}
	return nil, role
}

func (r *RoleRepo) GetRoleByTicker(ctx context.Context, ticker string) (error, entity.Role) {
	role, err := r.sql.GetRoleByTicker(ctx, ticker)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("ticker not found"), entity.Role{}
		}
		return err, entity.Role{}
	}
	return nil, role
}

func (r *RoleRepo) UpdateRole(ctx context.Context, input entity.UpdateRoleParams) (error, entity.Role) {
	role, err := r.sql.UpdateRole(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("ticker not found"), entity.Role{}
		}
		return err, entity.Role{}
	}
	return nil, role
}

func (r *RoleRepo) DeleteRole(ctx context.Context, ticker string) (error, entity.Role) {
	role, err := r.sql.DeleteRole(ctx, ticker)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("ticker not found"), entity.Role{}
		}
		return err, entity.Role{}
	}
	return nil, role
}
