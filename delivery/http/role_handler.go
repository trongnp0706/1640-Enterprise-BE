package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"net/http"
	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	RoleRepo repo.IRoleRepo
}

type RequestCreateRole struct {
	RoleName string `json:"role_name"`
	Ticker   string `json:"ticker"`
}

func (r *RoleHandler) AddRole(c echo.Context) error {
	req := RequestCreateRole{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	param := sql.CreateRoleParams{
		RoleName: req.RoleName,
		Ticker:   req.Ticker,
	}
	err, role := r.RoleRepo.AddRole(c.Request().Context(), param)
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
		Data:       role,
	})
}

func (r *RoleHandler) ListRoles(c echo.Context) error {
	err, roles := r.RoleRepo.ListRoles(c.Request().Context())
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
		Data:       roles,
	})
}

type UpdateRoleRequest struct {
	RoleName string `json:"role_name"`
	Ticker   string `json:"ticker"`
	Ticker_2 string `json:"ticker_2"`
}

func (r *RoleHandler) UpdateRole(c echo.Context) error {
	req := UpdateRoleRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	param := sql.UpdateRoleParams{
		RoleName: req.RoleName,
		Ticker:   req.Ticker,
		Ticker_2: req.Ticker_2,
	}
	err, role := r.RoleRepo.UpdateRole(c.Request().Context(), param)
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
		Data:       role,
	})
}

type DeleteRoleRequest struct {
	Ticker string `json:"ticker"`
}

func (r *RoleHandler) DeleteRole(c echo.Context) error {
	req := DeleteRoleRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, role := r.RoleRepo.DeleteRole(c.Request().Context(), req.Ticker)
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
		Data:       role,
	})
}
