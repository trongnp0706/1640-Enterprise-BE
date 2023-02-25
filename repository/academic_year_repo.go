package repository

import (
	entity "GDN-delivery-management/db/sql"
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/lib/pq"
)

type IAcademicYearRepo interface {
	GetAllAcademicYears(ctx context.Context) (error, []entity.AcademicYear)
	CreateAcademicYear(ctx context.Context, input entity.CreateAcademicYearParams) (error, entity.AcademicYear)
	GetAcademicYear(ctx context.Context, ticker string) (error, entity.AcademicYear)
	UpdateAcademicYear(ctx context.Context, input entity.UpdateAcademicYearParams) (error, entity.AcademicYear)
	DeleteAcademicYear(ctx context.Context, ticker string) (error, entity.AcademicYear)
}

type AcademicYearRepo struct {
	sql *entity.Queries
}

func NewAcademicYearRepo(sql *entity.Queries) IAcademicYearRepo {
	return &AcademicYearRepo{
		sql: sql,
	}
}

func (a *AcademicYearRepo) GetAllAcademicYears(ctx context.Context) (error, []entity.AcademicYear) {
	academicYears, err := a.sql.GetAcademicYears(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Categories are not available"), []entity.AcademicYear{}
		}
		return err, []entity.AcademicYear{}
	}
	return nil, academicYears
}

func (a *AcademicYearRepo) CreateAcademicYear(ctx context.Context, input entity.CreateAcademicYearParams) (error, entity.AcademicYear) {
	academicYear, err := a.sql.CreateAcademicYear(ctx, input)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errors.New("Conflict id"), entity.AcademicYear{}
			}
		}
		log.Println("err: ", err)
		return errors.New("Failed to add year"), entity.AcademicYear{}
	}
	return nil, academicYear
}

func (a *AcademicYearRepo) GetAcademicYear(ctx context.Context, ticker string) (error, entity.AcademicYear) {
	academicYear, err := a.sql.GetAcademicYear(ctx, ticker)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("year not found"), entity.AcademicYear{}
		}
		return err, entity.AcademicYear{}
	}
	return nil, academicYear
}

func (a *AcademicYearRepo) UpdateAcademicYear(ctx context.Context, input entity.UpdateAcademicYearParams) (error, entity.AcademicYear) {
	academicYear, err := a.sql.UpdateAcademicYear(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("year not found"), entity.AcademicYear{}
		}
		return err, entity.AcademicYear{}
	}
	return nil, academicYear
}

func (a *AcademicYearRepo) DeleteAcademicYear(ctx context.Context, ticker string) (error, entity.AcademicYear) {
	academicYear, err := a.sql.DeleteAcademicYear(ctx, ticker)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("year not found"), entity.AcademicYear{}
		}
		return err, entity.AcademicYear{}
	}
	return nil, academicYear
}
