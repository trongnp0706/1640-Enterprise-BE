package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
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

	err, comments := cm.CommentRepo.GetCommentsByIdea(c.Request().Context(), req.IdeaID)
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
