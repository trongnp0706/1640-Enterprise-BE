package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"net/http"
	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	CategoryRepo repo.ICategoryRepo
}

type CreateCategoryRequest struct {
	ID           string `json:"id"`
	CategoryName string `json:"category_name"`
}

func (cg *CategoryHandler) CreateCategory(c echo.Context) error {
	req := CreateCategoryRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	param := sql.CreateCategoryParams{
		ID:           req.ID,
		CategoryName: req.CategoryName,
	}
	err, category := cg.CategoryRepo.CreateCategory(c.Request().Context(), param)
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
		Data:       category,
	})
}

func (cg *CategoryHandler) GetAllCategories(c echo.Context) error {
	err, categories := cg.CategoryRepo.GetAllCategories(c.Request().Context())
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
		Data:       categories,
	})
}

type UpdateCategoryRequest struct {
	CategoryName string `json:"category_name"`
	ID           string `json:"id"`
	ID_2         string `json:"id_2"`
}

func (cg *CategoryHandler) UpdateCategory(c echo.Context) error {
	req := UpdateCategoryRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	param := sql.UpdateCategoryParams{
		CategoryName: req.CategoryName,
		ID:           req.ID,
		ID_2:         req.ID_2,
	}
	err, category := cg.CategoryRepo.UpdateCategory(c.Request().Context(), param)
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
		Data:       category,
	})
}

type DeleteCategoryRequest struct {
	ID string `json:"id"`
}

func (cg *CategoryHandler) DeleteCategory(c echo.Context) error {
	req := DeleteDepartmentRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, category := cg.CategoryRepo.DeleteCategory(c.Request().Context(), req.ID)
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
		Data:       category,
	})
}
