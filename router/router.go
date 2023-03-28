package router

import (
	"GDN-delivery-management/delivery/http"
	mdw "GDN-delivery-management/middleware"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo                *echo.Echo
	UserHandler         http.UserHandler
	RoleHandler         http.RoleHandler
	IdeaHandler         http.IdeaHandler
	CommentHandler      http.CommentHandler
	DepartmentHandler   http.DepartmentHandler
	CategoryHandler     http.CategoryHandler
	AcademicYearHandler http.AcademicYearHandler
	AuthMiddleware      *mdw.AuthMiddleware
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
	user.POST("/add", r.UserHandler.AddUser, r.AuthMiddleware.Authorize())
	user.GET("/:userid", r.UserHandler.UserDetails, r.AuthMiddleware.Authorize())
	user.GET("/get-me", r.UserHandler.GetMe, r.AuthMiddleware.Authorize())
	user.GET("/all", r.UserHandler.GetAllUsers, r.AuthMiddleware.Authorize())
	user.PATCH("/update", r.UserHandler.UpdateUser, r.AuthMiddleware.Authorize())
	user.PATCH("/update-avatar", r.UserHandler.UpdateAvatar, r.AuthMiddleware.Authorize())
	user.DELETE("/delete", r.UserHandler.DeleteUser, r.AuthMiddleware.Authorize())
	user.GET("/admin-check", r.UserHandler.CheckAdmin)

	role := r.Echo.Group("/role")
	role.POST("/add", r.RoleHandler.AddRole, r.AuthMiddleware.Authorize())
	role.GET("/all", r.RoleHandler.ListRoles, r.AuthMiddleware.Authorize())
	role.PATCH("/update", r.RoleHandler.UpdateRole, r.AuthMiddleware.Authorize())
	role.DELETE("/delete", r.RoleHandler.DeleteRole, r.AuthMiddleware.Authorize())

	department := r.Echo.Group("/department")
	department.POST("/add", r.DepartmentHandler.CreateDepartment, r.AuthMiddleware.Authorize())
	department.GET("/all", r.DepartmentHandler.GetAllDepartments, r.AuthMiddleware.Authorize())
	department.PATCH("/update", r.DepartmentHandler.UpdateDepartment, r.AuthMiddleware.Authorize())
	department.DELETE("/delete", r.DepartmentHandler.DeleteDepartment, r.AuthMiddleware.Authorize())

	category := r.Echo.Group("/category")
	category.POST("/add", r.CategoryHandler.CreateCategory, r.AuthMiddleware.Authorize())
	category.GET("/all", r.CategoryHandler.GetAllCategories, r.AuthMiddleware.Authorize())
	category.PATCH("/update", r.CategoryHandler.UpdateCategory, r.AuthMiddleware.Authorize())
	category.DELETE("/delete", r.CategoryHandler.DeleteCategory, r.AuthMiddleware.Authorize())

	academicYear := r.Echo.Group("/year")
	academicYear.POST("/add", r.AcademicYearHandler.CreateAcademicYear, r.AuthMiddleware.Authorize())
	academicYear.GET("/all", r.AcademicYearHandler.GetAllAcademicYears, r.AuthMiddleware.Authorize())
	academicYear.PATCH("/update", r.AcademicYearHandler.UpdateAcademicYear, r.AuthMiddleware.Authorize())
	academicYear.DELETE("/delete", r.AcademicYearHandler.DeleteAcademicYear, r.AuthMiddleware.Authorize())

	idea := r.Echo.Group("/idea")
	idea.POST("/add", r.IdeaHandler.AddIdea, r.AuthMiddleware.Authorize())
	idea.GET("/popular", r.IdeaHandler.GetMostPopularIdeas, r.AuthMiddleware.Authorize())
	idea.GET("/most-viewed", r.IdeaHandler.GetMostViewedIdeas, r.AuthMiddleware.Authorize())
	idea.GET("/latest", r.IdeaHandler.GetLatestIdeas, r.AuthMiddleware.Authorize())
	idea.GET("/by-category", r.IdeaHandler.GetIdeaByCategory, r.AuthMiddleware.Authorize())
	idea.GET("/by-year", r.IdeaHandler.GetIdeaByAcademicyear, r.AuthMiddleware.Authorize())
	idea.PATCH("/update", r.IdeaHandler.UpdateIdea, r.AuthMiddleware.Authorize())
	idea.DELETE("/delete", r.IdeaHandler.DeleteIdea, r.AuthMiddleware.Authorize())

	comment := r.Echo.Group("/comment")
	comment.POST("/add", r.CommentHandler.AddComment, r.AuthMiddleware.Authorize())
	comment.GET("/all", r.CommentHandler.GetCommentsByIdea, r.AuthMiddleware.Authorize())
	comment.GET("/latest", r.CommentHandler.GetLatestComment, r.AuthMiddleware.Authorize())
	comment.PATCH("/update", r.CommentHandler.UpdateComment, r.AuthMiddleware.Authorize())
	comment.DELETE("/delete", r.CommentHandler.DeleteComment, r.AuthMiddleware.Authorize())
}
