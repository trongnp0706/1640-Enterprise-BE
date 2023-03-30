package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type VoteHandler struct {
	VoteRepo repo.IVoteRepo
	IdeaRepo repo.IIdeaRepo
}

type GetVoteRequest struct {
	UserID string `json:"user_id"`
	IdeaID string `json:"idea_id"`
}

func (v VoteHandler) GetVote(c echo.Context) error {
	req := GetVoteRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	getVoteParam := sql.GetVoteParams{
		UserID: req.UserID,
		IdeaID: req.IdeaID,
	}
	err, existingVote := v.VoteRepo.GetVote(c.Request().Context(), getVoteParam)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"upvotes":   0,
			"downvotes": 0,
			"user_vote": "",
		})
	}

	upvoteCount := v.IdeaRepo.GetUpvoteCount(c.Request().Context(), req.IdeaID)
	downvoteCount := v.IdeaRepo.GetDownvoteCount(c.Request().Context(), req.IdeaID)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"upvotes":   upvoteCount,
		"downvotes": downvoteCount,
		"user_vote": existingVote.Vote,
	})
}

type HandleVoteRequest struct {
	UserID string `json:"user_id"`
	IdeaID string `json:"idea_id"`
	Vote   string `json:"vote"`
}

func (v VoteHandler) HandleVote(c echo.Context) error {
	req := HandleVoteRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	getVoteParam := sql.GetVoteParams{
		UserID: req.UserID,
		IdeaID: req.IdeaID,
	}
	err, existingVote := v.VoteRepo.GetVote(c.Request().Context(), getVoteParam)

	if existingVote.Vote == "up" {
		if req.Vote == "down" {
			v.IdeaRepo.DecreaseUpvoteCount(c.Request().Context(), existingVote.ID)
			v.IdeaRepo.IncreaseDownvoteCount(c.Request().Context(), existingVote.ID)
			updateVoteParam := sql.UpdateVoteParams{
				ID:   existingVote.ID,
				Vote: "down",
			}
			v.VoteRepo.UpdateVote(c.Request().Context(), updateVoteParam)
		} else if req.Vote == "" {
			v.IdeaRepo.DecreaseUpvoteCount(c.Request().Context(), existingVote.ID)
			v.VoteRepo.DeleteVote(c.Request().Context(), existingVote.ID)
		}
	} else if existingVote.Vote == "down" {
		if req.Vote == "up" {
			v.IdeaRepo.IncreaseUpvoteCount(c.Request().Context(), existingVote.ID)
			v.IdeaRepo.DecreaseDownvoteCount(c.Request().Context(), existingVote.ID)
			updateVoteParam := sql.UpdateVoteParams{
				ID:   existingVote.ID,
				Vote: "up",
			}
			v.VoteRepo.UpdateVote(c.Request().Context(), updateVoteParam)
		} else if req.Vote == "" {
			v.IdeaRepo.DecreaseDownvoteCount(c.Request().Context(), existingVote.ID)
			v.VoteRepo.DeleteVote(c.Request().Context(), existingVote.ID)
		}
	} else {
		if req.Vote == "up" {
			v.IdeaRepo.IncreaseUpvoteCount(c.Request().Context(), existingVote.ID)

			voteId, err := uuid.NewUUID()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Response{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
					Data:       nil,
				})
			}
			addVoteParam := sql.CreateVoteParams{
				ID:     voteId.String(),
				UserID: req.UserID,
				IdeaID: req.IdeaID,
				Vote:   "up",
			}
			v.VoteRepo.AddVote(c.Request().Context(), addVoteParam)
		} else if req.Vote == "down" {
			v.IdeaRepo.IncreaseDownvoteCount(c.Request().Context(), existingVote.ID)

			voteId, err := uuid.NewUUID()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Response{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
					Data:       nil,
				})
			}
			addVoteParam := sql.CreateVoteParams{
				ID:     voteId.String(),
				UserID: req.UserID,
				IdeaID: req.IdeaID,
				Vote:   "down",
			}
			v.VoteRepo.AddVote(c.Request().Context(), addVoteParam)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"upvotes":   v.IdeaRepo.GetUpvoteCount(c.Request().Context(), req.IdeaID),
		"downvotes": v.IdeaRepo.GetDownvoteCount(c.Request().Context(), req.IdeaID),
		"user_vote": req.Vote,
	})
}
