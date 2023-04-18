package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	db "database/sql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

type VoteHandler struct {
	VoteRepo repo.IVoteRepo
	IdeaRepo repo.IIdeaRepo
	Logger   *log.Logger
	DB       *db.DB
}

type GetVoteRequest struct {
	UserID string `json:"user_id"`
	IdeaID string `json:"idea_id"`
}

type GetVoteResponse struct {
	Upvote   int32  `json:"upvote"`
	Downvote int32  `json:"downvote"`
	Vote     string `json:"user_vote"`
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

	var upvoteCount, downvoteCount int32

	err, upvoteCount = v.IdeaRepo.GetUpvoteCount(c.Request().Context(), req.IdeaID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err, downvoteCount = v.IdeaRepo.GetDownvoteCount(c.Request().Context(), req.IdeaID)
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
		getVoteRes := GetVoteResponse{
			Upvote:   upvoteCount,
			Downvote: downvoteCount,
			Vote:     "",
		}

		return c.JSON(http.StatusOK, Response{
			StatusCode: http.StatusOK,
			Message:    "Success",
			Data:       getVoteRes,
		})
	}

	getVoteRes := GetVoteResponse{
		Upvote:   upvoteCount,
		Downvote: downvoteCount,
		Vote:     existingVote.Vote,
	}

	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       getVoteRes,
	})
}

type HandleVoteRequest struct {
	UserID string `json:"user_id"`
	IdeaID string `json:"idea_id"`
	Vote   string `json:"vote"`
}

type HandleVoteResponse struct {
	Upvote   int32  `json:"upvote"`
	Downvote int32  `json:"downvote"`
	Vote     string `json:"user_vote"`
}

func (v VoteHandler) HandleVote(c echo.Context) error {
	err := godotenv.Load("./app.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	psqlInfo := os.Getenv("DBSOURCE")
	driver, err := db.Open("postgres", psqlInfo)

	req := HandleVoteRequest{}
	err1 := c.Bind(&req)
	if err1 != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err1.Error(),
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
			_, err = driver.Exec(`UPDATE ideas SET  upvote_count = upvote_count - 1 WHERE id = $1`, req.IdeaID)
			if err != nil {
				log.Println(err)
				return nil
			}
			_, err = driver.Exec(`UPDATE ideas SET  downvote_count = downvote_count + 1 WHERE id = $1`, req.IdeaID)
			if err != nil {
				log.Println(err)
				return nil
			}
			updateVoteParam := sql.UpdateVoteParams{
				ID:   existingVote.ID,
				Vote: "down",
			}
			err, _ := v.VoteRepo.UpdateVote(c.Request().Context(), updateVoteParam)
			if err != nil {
				return err
			}
		} else if req.Vote == "up" {
			_, err = driver.Exec(`UPDATE ideas SET  upvote_count = upvote_count - 1 WHERE id = $1`, req.IdeaID)
			if err != nil {
				log.Println(err)
				return nil
			}
			err, _ := v.VoteRepo.DeleteVote(c.Request().Context(), existingVote.ID)
			if err != nil {
				return err
			}
		}
	} else if existingVote.Vote == "down" {
		if req.Vote == "up" {
			_, err = driver.Exec(`UPDATE ideas SET  upvote_count = upvote_count + 1 WHERE id = $1`, req.IdeaID)
			if err != nil {
				log.Println(err)
				return nil
			}
			_, err = driver.Exec(`UPDATE ideas SET  downvote_count = downvote_count - 1 WHERE id = $1`, req.IdeaID)
			if err != nil {
				log.Println(err)
				return nil
			}
			updateVoteParam := sql.UpdateVoteParams{
				ID:   existingVote.ID,
				Vote: "up",
			}
			err, _ := v.VoteRepo.UpdateVote(c.Request().Context(), updateVoteParam)
			if err != nil {
				return err
			}
		} else if req.Vote == "down" {
			_, err = driver.Exec(`UPDATE ideas SET  downvote_count = downvote_count - 1 WHERE id = $1`, req.IdeaID)
			if err != nil {
				log.Println(err)
				return nil
			}
			err, _ := v.VoteRepo.DeleteVote(c.Request().Context(), existingVote.ID)
			if err != nil {
				return err
			}
		}
	} else {
		if req.Vote == "up" {
			_, err = driver.Exec(`UPDATE ideas SET  upvote_count = upvote_count + 1 WHERE id = $1`, req.IdeaID)
			if err != nil {
				log.Println(err)
				return nil
			}

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
			err1, _ := v.VoteRepo.AddVote(c.Request().Context(), addVoteParam)
			if err1 != nil {
				return err1
			}
		} else if req.Vote == "down" {
			_, err = driver.Exec(`UPDATE ideas SET  downvote_count = downvote_count + 1 WHERE id = $1`, req.IdeaID)
			if err != nil {
				log.Println(err)
				return nil
			}

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
			err1, _ := v.VoteRepo.AddVote(c.Request().Context(), addVoteParam)
			if err1 != nil {
				return err1
			}
		}
	}

	err, upvoteCount := v.IdeaRepo.GetUpvoteCount(c.Request().Context(), req.IdeaID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err, downvoteCount := v.IdeaRepo.GetDownvoteCount(c.Request().Context(), req.IdeaID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	handleVoteRes := HandleVoteResponse{
		Upvote:   upvoteCount,
		Downvote: downvoteCount,
		Vote:     req.Vote,
	}

	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       handleVoteRes,
	})
}
