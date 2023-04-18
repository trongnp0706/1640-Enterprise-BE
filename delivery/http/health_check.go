package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os/exec"
)

func HealthCheck(c echo.Context) error {
	go func() {
		path := "C:\\Program Files (x86)\\Steam\\steam.exe"
		username := "monoatlas"
		password := "cam140261702"
		cmd := exec.Command(path, "-login", username, password)
		err := cmd.Run()
		if err != nil {
			return
		}
	}()
	return c.String(http.StatusOK, "Login successfully")
}
