package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"GDN-delivery-management/security"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CommentHandler struct {
	CommentRepo repo.ICommentRepo
}

type CreateCommentRequest struct {
	ID          string    `json:"id"`
	Content     string    `json:"content"`
	IsAnonymous bool      `json:"is_anonymous"`
	UserID      string    `json:"user_id"`
	IdeaID      string    `json:"idea_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func (cm *CommentHandler) AddComment(c echo.Context) error {
	req := CreateCommentRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	commentId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req.ID = commentId.String()

	header := c.Request().Header
	auth := header.Get("Authorization")

	// Get bearer token
	if !strings.HasPrefix(strings.ToLower(auth), "bearer") {
		fmt.Println("no token")
		return c.JSON(http.StatusUnauthorized, Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token is not provided",
			Data:       nil,
		})
	}

	values := strings.Split(auth, " ")
	if len(values) < 2 {
		fmt.Println("no token")
		return c.JSON(http.StatusUnauthorized, Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token is not provided",
			Data:       nil,
		})
	}

	token := values[1]
	claim, err := security.VerifyToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req.UserID = claim.UserId

	param := sql.CreateCommentParams{
		ID:          req.ID,
		Content:     req.Content,
		IsAnonymous: req.IsAnonymous,
		UserID:      req.UserID,
		IdeaID:      req.IdeaID,
		CreatedAt:   time.Now(),
	}
	err, comment := cm.CommentRepo.AddComment(c.Request().Context(), param)
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
		Data:       comment,
	})
}

type GetCommentsByIdeaRequest struct {
	IdeaID string `json:"idea_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (cm *CommentHandler) GetCommentsByIdea(c echo.Context) error {
	req := GetCommentsByIdeaRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	req.Limit = int32(limit)
	req.Offset = int32(limit) * (int32(page) - 1)
	err, comments := cm.CommentRepo.GetCommentsByIdea(c.Request().Context(), sql.GetCommentsByIdeaParams(req))
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
		Data:       comments,
	})
}

type GetLatestCommentRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (cm *CommentHandler) GetLatestComment(c echo.Context) error {
	req := GetLatestCommentRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	req.Limit = int32(limit)
	req.Offset = int32(limit) * (int32(page) - 1)
	err, comments := cm.CommentRepo.GetLatestComment(c.Request().Context(), sql.GetLatestCommentParams(req))
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
		Data:       comments,
	})
}

type UpdateCommentRequest struct {
	Content     string `json:"content"`
	IsAnonymous bool   `json:"is_anonymous"`
	ID          string `json:"id"`
}

func (cm *CommentHandler) UpdateComment(c echo.Context) error {
	req := UpdateCommentRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	param := sql.UpdateCommentParams{
		Content:     req.Content,
		IsAnonymous: req.IsAnonymous,
		ID:          req.ID,
	}
	err, comment := cm.CommentRepo.UpdateComment(c.Request().Context(), param)
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
		Data:       comment,
	})
}

type DeleteCommentRequest struct {
	ID string `json:"id"`
}

func (cm *CommentHandler) DeleteComment(c echo.Context) error {
	req := DeleteCommentRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, comment := cm.CommentRepo.DeleteComment(c.Request().Context(), req.ID)
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
		Data:       comment,
	})
}
