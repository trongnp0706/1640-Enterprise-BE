package http

import (
	"GDN-delivery-management/otp"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VerifyOTP struct {
	phone string `json:"phone"`
	code  string `json:"code"` 
}

func (u *UserHandler)VerifyOTP(c echo.Context) error {
	req := VerifyOTP{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err = otp.TwilioVerifyOTP(req.phone, req.code)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unauthorized OTP",
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       true,
	})
}