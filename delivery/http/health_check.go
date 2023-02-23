package http


import(
	"github.com/labstack/echo/v4"
	"net/http"
)


func HealthCheck(c echo.Context)error{
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Health check",
	})
} 