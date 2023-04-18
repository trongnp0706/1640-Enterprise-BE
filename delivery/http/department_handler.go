package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"net/http"
	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type DepartmentHandler struct {
	DepartmentRepo repo.IDepartmentRepo
}

type CreateDepartmentRequest struct {
	ID             string `json:"id"`
	DepartmentName string `json:"department_name"`
}

func (d *DepartmentHandler) CreateDepartment(c echo.Context) error {
	req := CreateDepartmentRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	param := sql.CreateDepartmentParams{
		ID:             req.ID,
		DepartmentName: req.DepartmentName,
	}
	err, department := d.DepartmentRepo.CreateDepartment(c.Request().Context(), param)
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
		Data:       department,
	})
}

func (d *DepartmentHandler) GetAllDepartments(c echo.Context) error {
	err, departments := d.DepartmentRepo.GetAllDepartments(c.Request().Context())
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
		Data:       departments,
	})
}

type UpdateDepartmentRequest struct {
	DepartmentName string `json:"department_name"`
	ID             string `json:"id"`
}

func (d *DepartmentHandler) UpdateDepartment(c echo.Context) error {
	req := UpdateDepartmentRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	param := sql.UpdateDepartmentParams{
		DepartmentName: req.DepartmentName,
		ID:             req.ID,
	}
	err, department := d.DepartmentRepo.UpdateDepartment(c.Request().Context(), param)
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
		Data:       department,
	})
}

type DeleteDepartmentRequest struct {
	ID string `json:"id"`
}

func (d *DepartmentHandler) DeleteDepartment(c echo.Context) error {
	req := DeleteDepartmentRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, department := d.DepartmentRepo.DeleteDepartment(c.Request().Context(), req.ID)
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
		Data:       department,
	})
}
