package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type IdeaHandler struct {
	IdeaRepo repo.IIdeaRepo
}

type CreateIdeaRequest struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	ViewCount     int32     `json:"view_count"`
	DocumentArray string    `json:"document_array"`
	ImageArray    []string  `json:"image_array"`
	UpvoteCount   int32     `json:"upvote_count"`
	DownvoteCount int32     `json:"downvote_count"`
	IsAnonymous   bool      `json:"is_anonymous"`
	UserID        string    `json:"user_id"`
	CategoryID    string    `json:"category_id"`
	AcademicYear  string    `json:"academic_year"`
	CreatedAt     time.Time `json:"created_at"`
}

func (i *IdeaHandler) AddIdea(c echo.Context) error {
	req := CreateIdeaRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	ideaId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req.ID = ideaId.String()

	param := sql.CreateIdeaParams{
		ID:            req.ID,
		Title:         req.Title,
		Content:       req.Content,
		ViewCount:     0,
		Column5:       req.DocumentArray,
		Column6:       req.ImageArray,
		UpvoteCount:   0,
		DownvoteCount: 0,
		IsAnonymous:   req.IsAnonymous,
		UserID:        req.UserID,
		CategoryID:    req.CategoryID,
		AcademicYear:  req.AcademicYear,
		CreatedAt:     time.Now(),
	}
	err, idea := i.IdeaRepo.AddIdea(c.Request().Context(), param)
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
		Data:       idea,
	})
}

type GetMostPopularIdeasRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (i *IdeaHandler) GetMostPopularIdeas(c echo.Context) error {
	req := GetMostPopularIdeasRequest{}
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
	err, ideas := i.IdeaRepo.GetMostPopularIdeas(c.Request().Context(), sql.GetMostPopularIdeasParams(req))
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
		Data:       ideas,
	})
}

type GetMostViewedIdeasRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (i *IdeaHandler) GetMostViewedIdeas(c echo.Context) error {
	req := GetMostViewedIdeasRequest{}
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
	err, ideas := i.IdeaRepo.GetMostViewedIdeas(c.Request().Context(), sql.GetMostViewedIdeasParams(req))
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
		Data:       ideas,
	})
}

type GetLatestIdeasRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (i *IdeaHandler) GetLatestIdeas(c echo.Context) error {
	req := GetLatestIdeasRequest{}
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
	err, ideas := i.IdeaRepo.GetLatestIdeas(c.Request().Context(), sql.GetLatestIdeasParams(req))
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
		Data:       ideas,
	})
}

type GetIdeaByCategoryRequest struct {
	CategoryID string `json:"category_id"`
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
}

func (i *IdeaHandler) GetIdeaByCategory(c echo.Context) error {
	req := GetIdeaByCategoryRequest{}
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
	param := sql.GetIdeaByCategoryParams{
		CategoryID: req.CategoryID,
		Limit:      req.Limit,
		Offset:     req.Offset,
	}
	err, ideas := i.IdeaRepo.GetIdeaByCategory(c.Request().Context(), param)
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
		Data:       ideas,
	})
}

type GetIdeaByAcademicyearRequest struct {
	AcademicYear string `json:"academic_year"`
	Limit        int32  `json:"limit"`
	Offset       int32  `json:"offset"`
}

func (i *IdeaHandler) GetIdeaByAcademicyear(c echo.Context) error {
	req := GetIdeaByAcademicyearRequest{}
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
	param := sql.GetIdeaByAcademicyearParams{
		AcademicYear: req.AcademicYear,
		Limit:        req.Limit,
		Offset:       req.Offset,
	}
	err, ideas := i.IdeaRepo.GetIdeaByAcademicyear(c.Request().Context(), param)
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
		Data:       ideas,
	})
}

type UpdateIdeaRequest struct {
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	DocumentArray string   `json:"document_array"`
	ImageArray    []string `json:"image_array"`
	IsAnonymous   bool     `json:"is_anonymous"`
	AcademicYear  string   `json:"academic_year"`
	CategoryID    string   `json:"category_id"`
	ID            string   `json:"id"`
}

func (i *IdeaHandler) UpdateIdea(c echo.Context) error {
	req := UpdateIdeaRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	param := sql.UpdateIdeaParams{
		Title:         req.Title,
		Content:       req.Content,
		DocumentArray: req.DocumentArray,
		ImageArray:    req.ImageArray,
		IsAnonymous:   req.IsAnonymous,
		CategoryID:    req.CategoryID,
		AcademicYear:  req.AcademicYear,
		ID:            req.ID,
	}
	err, idea := i.IdeaRepo.UpdateIdea(c.Request().Context(), param)
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
		Data:       idea,
	})
}

type DeleteIdeaRequest struct {
	ID string `json:"id"`
}

func (i *IdeaHandler) DeleteIdea(c echo.Context) error {
	req := DeleteIdeaRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, idea := i.IdeaRepo.DeleteIdea(c.Request().Context(), req.ID)
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
		Data:       idea,
	})
}
