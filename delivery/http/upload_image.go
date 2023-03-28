package http

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func UploadImage(c echo.Context) error {
	// Get the image file from form data
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Open the file for reading
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	defer src.Close()

	// Create the destination directory if it doesn't exist
	destDir := os.Getenv("IMAGE_DIR")

	// Create the destination file
	dstPath := filepath.Join(destDir, file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	defer dst.Close()

	// Copy the file to the destination
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Return the file path as the response
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Upload Success",
		Data:       "http://localhost:3000/images/" + file.Filename,
	})
}
