package router

import (
	"GDN-delivery-management/delivery/http"
	mdw "GDN-delivery-management/middleware"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo              *echo.Echo
	UserHandler       http.UserHandler
	RoleHandler       http.RoleHandler
	DepartmentHandler http.DepartmentHandler
	CategoryHandler   http.CategoryHandler
	AuthMiddleware    *mdw.AuthMiddleware
}

func (r *Router) SetupRouter() {
	// health check
	r.Echo.GET("/health-check", http.HealthCheck)

	admin := r.Echo.Group("/admin")
	admin.POST("/sign-up", r.UserHandler.SystemAdminSignUp)

	r.Echo.POST("/login", r.UserHandler.Login)
	r.Echo.GET("/logout", r.UserHandler.Logout)
	r.Echo.POST("otp/verify", r.UserHandler.VerifyOTP)

	// r.Echo.POST("/token/refresh-token", r.UserHandler.RenewAccessToken, r.AuthMiddleware.UserCors())

	user := r.Echo.Group("/user")
	user.POST("/add-user", r.UserHandler.AddUser, r.AuthMiddleware.Authorize())
	user.GET("/profile", r.UserHandler.UserDetails, r.AuthMiddleware.Authorize())
	user.GET("/get-me", r.UserHandler.GetMe, r.AuthMiddleware.Authorize())
	user.GET("/all-user", r.UserHandler.GetAllUsers, r.AuthMiddleware.Authorize())
	user.PATCH("/update-user", r.UserHandler.UpdateUser, r.AuthMiddleware.Authorize())
	user.DELETE("/delete-user", r.UserHandler.DeleteUser, r.AuthMiddleware.Authorize())
	user.GET("/admin-check", r.UserHandler.CheckAdmin)

	role := r.Echo.Group("/role")
	role.POST("/add-role", r.RoleHandler.AddRole, r.AuthMiddleware.Authorize())
	role.GET("/all-role", r.RoleHandler.ListRoles, r.AuthMiddleware.Authorize())
	role.PATCH("/update-role", r.RoleHandler.UpdateRole, r.AuthMiddleware.Authorize())
	role.DELETE("/delete-role", r.RoleHandler.DeleteRole, r.AuthMiddleware.Authorize())

	department := r.Echo.Group("/department")
	department.POST("/add-department", r.DepartmentHandler.CreateDepartment, r.AuthMiddleware.Authorize())
	department.GET("/all-department", r.DepartmentHandler.GetAllDepartments, r.AuthMiddleware.Authorize())
	department.PATCH("/update-department", r.DepartmentHandler.UpdateDepartment, r.AuthMiddleware.Authorize())
	department.DELETE("/delete-department", r.DepartmentHandler.DeleteDepartment, r.AuthMiddleware.Authorize())

	category := r.Echo.Group("/category")
	category.POST("/add-category", r.CategoryHandler.CreateCategory, r.AuthMiddleware.Authorize())
	category.GET("/all-category", r.CategoryHandler.GetAllCategories, r.AuthMiddleware.Authorize())
	category.PATCH("/update-category", r.CategoryHandler.UpdateCategory, r.AuthMiddleware.Authorize())
	category.DELETE("/delete-category", r.CategoryHandler.DeleteCategory, r.AuthMiddleware.Authorize())
}
