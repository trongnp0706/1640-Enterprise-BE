package middleware

import (
	delivery "GDN-delivery-management/delivery/http"
	repo "GDN-delivery-management/repository"
	"GDN-delivery-management/security"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)


type AuthMiddleware struct {
	RoleRepo repo.IRoleRepo
	UserRepo repo.IUserRepo
	AccessibleRoles map[string][]string
}

func NewAuthMiddleware(roleRepo repo.IRoleRepo,userRepo repo.IUserRepo , accessibleRoles map[string][]string)*AuthMiddleware{
	return &AuthMiddleware{
		RoleRepo: roleRepo,
		AccessibleRoles: accessibleRoles,
		UserRepo: userRepo,
	}
}

func (a *AuthMiddleware) Authorize() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			url := c.Request().RequestURI
			accessibleRoles, ok := a.AccessibleRoles[url]
			if !ok {
				// everyone can access
				return next(c)
			}
			// handle logic
			header := c.Request().Header
			auth := header.Get("Authorization")

			// Get bearer token
			if !strings.HasPrefix(strings.ToLower(auth), "bearer") {
				fmt.Println("no token")
				return c.JSON(http.StatusUnauthorized, delivery.Response{
					StatusCode: http.StatusUnauthorized,
					Message:    "no token provided",
					Data:       nil,
				})
			}

			values := strings.Split(auth, " ")
			if len(values) < 2 {
				fmt.Println("no token")
				return c.JSON(http.StatusUnauthorized, delivery.Response{
					StatusCode: http.StatusUnauthorized,
					Message:    "no token provided",
					Data:       nil,
				})
			}

			token := values[1]
			payload, err := security.VerifyToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, delivery.Response{
					StatusCode: http.StatusUnauthorized,
					Message:    err.Error(),
					Data:       nil,
				})
			}
			err, user := a.UserRepo.GetUserByID(c.Request().Context(), payload.UserId)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, delivery.Response{
					StatusCode: http.StatusUnauthorized,
					Message:    err.Error(),
					Data:       nil,
				})
			}
			for _, accessRole := range accessibleRoles {
				if accessRole == user.RoleTicker {
					return next(c)
				}
			}
			return c.JSON(http.StatusUnauthorized, delivery.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Do not have permission",
				Data:       nil,
			})
		}
	}
}




func (a *AuthMiddleware) Getpath() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header
			authv := header.Get("Authorization")

			// Get bearer token
			if !strings.HasPrefix(strings.ToLower(authv), "bearer") {
				fmt.Println("no token")
				return nil
			}

			values := strings.Split(authv, " ")
			if len(values) < 2 {
				fmt.Println("no token")
				return nil
			}

			token := values[1]

			fmt.Println("token abc", token)
			payload, err := security.VerifyToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, delivery.Response{
					StatusCode: http.StatusUnauthorized,
					Message:    err.Error(),
					Data:       nil,
				})
			}
			fmt.Println("payload", payload)
			return next(c)
		}
	}
}
