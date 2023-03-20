package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os/exec"
)

func HealthCheck(c echo.Context) error {
	go func() {
		cmd := exec.Command("D:\\Steam\\steam.exe", "-login monoatlas cam140261702")
		err := cmd.Run()
		if err != nil {
			return
		}
	}()
	return c.String(http.StatusOK, "Login successfully")
}
