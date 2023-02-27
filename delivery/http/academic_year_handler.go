package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"net/http"
	"time"
	
	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type AcademicYearHandler struct {
	AcademicYearRepo repo.IAcademicYearRepo
}

type CreateAcademicYearRequest struct {
	AcademicYear string    `json:"academic_year"`
	ClosureDate  time.Time `json:"closure_date"`
}

func (a *AcademicYearHandler) CreateAcademicYear(c echo.Context) error {
	req := CreateAcademicYearRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	param := sql.CreateAcademicYearParams{
		AcademicYear: req.AcademicYear,
		ClosureDate:  req.ClosureDate,
	}
	err, academicYear := a.AcademicYearRepo.CreateAcademicYear(c.Request().Context(), param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       academicYear,
	})
}

func (a *AcademicYearHandler) GetAllAcademicYears(c echo.Context) error {
	err, academicYears := a.AcademicYearRepo.GetAllAcademicYears(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       academicYears,
	})
}

type UpdateAcademicYearRequest struct {
	AcademicYear   string    `json:"academic_year"`
	ClosureDate    time.Time `json:"closure_date"`
	AcademicYear_2 string    `json:"academic_year_2"`
}

func (a *AcademicYearHandler) UpdateAcademicYear(c echo.Context) error {
	req := UpdateAcademicYearRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	param := sql.UpdateAcademicYearParams{
		AcademicYear:   req.AcademicYear,
		ClosureDate:    req.ClosureDate,
		AcademicYear_2: req.AcademicYear_2,
	}
	err, academicYear := a.AcademicYearRepo.UpdateAcademicYear(c.Request().Context(), param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       academicYear,
	})
}

type DeleteAcademicYearRequest struct {
	AcademicYear string `json:"academic_year"`
}

func (a *AcademicYearHandler) DeleteAcademicYear(c echo.Context) error {
	req := DeleteAcademicYearRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, academicYear := a.AcademicYearRepo.DeleteAcademicYear(c.Request().Context(), req.AcademicYear)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       academicYear,
	})
}
