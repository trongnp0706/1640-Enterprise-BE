package repository

import (
	entity "GDN-delivery-management/db/sql"
	"context"
	"errors"
)
type ISessionRepo interface{
	AddSession(ctx context.Context,input entity.CreateSessionParams)(error, entity.Session)
	GetSessionByID(ctx context.Context, id string)(error, entity.Session)
}

type SessionRepo struct {
	sql *entity.Queries
}

func NewSessionRepo(sql *entity.Queries) ISessionRepo{
	return &SessionRepo{
		sql: sql,
	}
}

func(s *SessionRepo) AddSession(ctx context.Context,input entity.CreateSessionParams)(error, entity.Session){
	sess, err := s.sql.CreateSession(ctx, input)
	if err != nil {
		return errors.New("Failed to create session"), entity.Session{}
	}
	return nil, sess
}

func (s *SessionRepo)GetSessionByID(ctx context.Context, id string)(error, entity.Session){
	sess, err := s.sql.GetSession(ctx, id)
	if err != nil {
		return errors.New("Session not found"), entity.Session{}
	}
	return nil, sess
}