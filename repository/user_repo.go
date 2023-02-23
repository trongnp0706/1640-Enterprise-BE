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

type IUserRepo interface {
	AddUser(ctx context.Context, imput entity.CreatUserParams) (error, entity.User)
	GetUserByID(ctx context.Context, id string) (error, entity.User)
	GetUserByEmail(ctx context.Context, email string) (error, entity.User)
	GetAllUsers(ctx context.Context, input entity.GetAllUsersParams) (error, []entity.GetAllUsersRow)
	UpdateUser(ctx context.Context, input entity.UpdateUserParams) (error, entity.User)
	DeleteUser(ctx context.Context, id string) (error, entity.User)
}

type UserRepo struct {
	sql *entity.Queries
}

func NewUserRepo(sql *entity.Queries) IUserRepo {
	return &UserRepo{
		sql: sql,
	}
}

func (u *UserRepo) AddUser(ctx context.Context, imput entity.CreatUserParams) (error, entity.User) {
	user, err := u.sql.CreatUser(ctx, imput)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errors.New("Conflict email"), entity.User{}
			}
		}
		log.Println("err: ", err)
		return errors.New("Failed to add user"), entity.User{}
	}
	return nil, user
}

func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (error, entity.User) {
	user, err := u.sql.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found"), entity.User{}
		}
		return err, entity.User{}
	}
	return nil, user
}

func (u *UserRepo) GetUserByID(ctx context.Context, email string) (error, entity.User) {
	user, err := u.sql.GetUserByID(ctx, email)
	if err != nil {
		log.Println("err", err)
		if err == sql.ErrNoRows {
			return errors.New("user not found"), entity.User{}
		}
		return err, entity.User{}
	}
	return nil, user
}

func (u *UserRepo) GetAllUsers(ctx context.Context, input entity.GetAllUsersParams) (error, []entity.GetAllUsersRow) {
	items, err := u.sql.GetAllUsers(ctx, input)
	if err != nil {
		fmt.Printf(err.Error())
		return errors.New("Cannot get all users"), []entity.GetAllUsersRow{}
	}
	return nil, items
}

func (u *UserRepo) UpdateUser(ctx context.Context, input entity.UpdateUserParams) (error, entity.User) {
	user, err := u.sql.UpdateUser(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found"), entity.User{}
		}
		return err, entity.User{}
	}
	return nil, user
}

func (u *UserRepo) DeleteUser(ctx context.Context, id string) (error, entity.User) {
	user, err := u.sql.DeleteUser(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found"), entity.User{}
		}
		return err, entity.User{}
	}
	return nil, user
}
