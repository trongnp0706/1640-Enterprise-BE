package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"GDN-delivery-management/security"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"GDN-delivery-management/mail"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo    repo.IUserRepo
	SessionRepo repo.ISessionRepo
	Email       mail.IEmailSender
}

type AddUserRequest struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Avatar       string `json:"avatar"`
	RoleTicker   string `json:"role_ticker"`
	DepartmentID string `json:"department_id"`
}

func (u *UserHandler) AddUser(c echo.Context) error {
	req := AddUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	hash := security.HashAndSalt([]byte(req.Password))
	req.Password = hash
	userId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req.ID = userId.String()
	req.Avatar = "http://localhost:3000/images/nino.jpg"
	param := sql.CreateUserParams{
		ID:           req.ID,
		Email:        req.Email,
		Username:     req.Username,
		RoleTicker:   req.RoleTicker,
		Password:     req.Password,
		Avatar:       req.Avatar,
		DepartmentID: req.DepartmentID,
	}
	err, user := u.UserRepo.AddUser(c.Request().Context(), param)
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
		Data:       user,
	})
}

type CreateAdminSystemRequest struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RoleTicker   string `json:"role_ticker"`
	DepartmentID string `json:"department_id"`
}

func (u *UserHandler) SystemAdminSignUp(c echo.Context) error {
	req := CreateAdminSystemRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	if req.Email == "" {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Email is required",
			Data:       nil,
		})
	}
	if req.Password == "" {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Password is required",
			Data:       nil,
		})
	}
	if req.Username == "" {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    "User name is required",
			Data:       nil,
		})
	}
	fmt.Println(req)
	hash := security.HashAndSalt([]byte(req.Password))
	req.Password = hash
	userId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req.ID = userId.String()
	param := sql.CreateUserParams{
		ID:           req.ID,
		Email:        req.Email,
		Username:     req.Username,
		RoleTicker:   "SAD",
		Password:     req.Password,
		DepartmentID: req.DepartmentID,
	}
	err, user := u.UserRepo.AddUser(c.Request().Context(), param)
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
		Data:       user,
	})
}

type GetUserProfilereq struct {
	ID string `json:"id"`
}

func (u *UserHandler) UserDetails(c echo.Context) error {
	req := GetUserProfilereq{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	userId := req.ID
	userId = c.Param("userid")
	err, user := u.UserRepo.GetUserByID(c.Request().Context(), userId)
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
		Data:       user,
	})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	User                  sql.User `json:"user"`
	SessionID             string   `json:"session_id"`
	AccessToken           string   `json:"access_token"`
	AccessTokenExpiresAt  int64    `json:"access_token_expires_at"`
	RefreshToken          string   `json:"refresh_token"`
	RefreshTokenExpiresAt int64    `json:"refresh_token_expires_at"`
}

func (u UserHandler) Login(c echo.Context) error {
	req := LoginRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, user := u.UserRepo.GetUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Login failed",
			Data:       nil,
		})
	}

	token, token_payload, err := security.GenToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	refresh_token, refresh_token_payload, err := security.GenRefreshtoken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	sessID, _ := uuid.NewUUID()
	sessionParam := sql.CreateSessionParams{
		ID:           sessID.String(),
		UserID:       user.ID,
		RefreshToken: refresh_token,
		UserAgent:    "agent",
		ClientIp:     "ip",
		IsBlocked:    false,
		ExpiresAt:    int64(refresh_token_payload.ExpiresAt),
		CreatedAt:    time.Now(),
	}

	err, sess := u.SessionRepo.AddSession(c.Request().Context(), sessionParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	userRes := UserLoginResponse{
		User:                  user,
		SessionID:             sess.ID,
		AccessToken:           token,
		AccessTokenExpiresAt:  token_payload.ExpiresAt,
		RefreshToken:          refresh_token,
		RefreshTokenExpiresAt: refresh_token_payload.ExpiresAt,
	}

	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       userRes,
	})
}

type RenewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
type RenewAccessTokenResponse struct {
	AccessToken string `josn:"access_token"`
	ExpiresAt   int    `json:"access_token"`
}

func (u *UserHandler) RenewAccessToken(c echo.Context) error {
	req := RenewAccessTokenRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	refresh_payload, err := security.VerifyToken(req.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err, sess := u.SessionRepo.GetSessionByID(c.Request().Context(), refresh_payload.Id)
	if err != nil {
		return c.JSON(http.StatusNotFound, Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	if sess.IsBlocked {
		err := fmt.Errorf("block session")
		return c.JSON(http.StatusUnauthorized, Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	if sess.UserID != refresh_payload.UserId {
		err := fmt.Errorf("incorrect session user")
		return c.JSON(http.StatusUnauthorized, Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	if sess.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("mismatched session token")
		return c.JSON(http.StatusUnauthorized, Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	timeT := time.Unix(sess.ExpiresAt, 0)
	if time.Now().After(timeT) {
		err := fmt.Errorf("expired session")
		return c.JSON(http.StatusUnauthorized, Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	accessToken, _, err := security.GenToken(sql.User{
		ID:    refresh_payload.UserId,
		Email: refresh_payload.Email,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	rsp := RenewAccessTokenResponse{
		AccessToken: accessToken,
	}
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       rsp,
	})
}

func (u *UserHandler) GetMe(c echo.Context) error {
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

	err, user := u.UserRepo.GetUserByID(c.Request().Context(), claim.UserId)
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
		Data:       user,
	})
}

func (u *UserHandler) Logout(c echo.Context) error {
	header := c.Request().Header
	header.Del("Authorization")

	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "logout",
		Data:       nil,
	})
}

type UpdateUserRequest struct {
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RoleTicker   string `json:"role_ticker"`
	DepartmentID string `json:"department_id"`
	Avatar       string `json:"avatar"`
	ID           string `json:"id"`
}

type UserUpdateResponse struct {
	User sql.User `json:"user"`
}

func (u *UserHandler) UpdateUser(c echo.Context) error {
	req := UpdateUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	hash := security.HashAndSalt([]byte(req.Password))
	req.Password = hash
	param := sql.UpdateUserParams{
		Username:     req.UserName,
		Email:        req.Email,
		Password:     req.Password,
		RoleTicker:   req.RoleTicker,
		DepartmentID: req.DepartmentID,
		Avatar:       req.Avatar,
		ID:           req.ID,
	}
	err, user := u.UserRepo.UpdateUser(c.Request().Context(), param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	userRes := UserUpdateResponse{
		User: user,
	}
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       userRes,
	})
}

type DeleteUserRequest struct {
	Id string `json:"id"`
}

func (u *UserHandler) DeleteUser(c echo.Context) error {
	req := DeleteUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, user := u.UserRepo.DeleteUser(c.Request().Context(), req.Id)
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
		Data:       user,
	})
}

type GetAllUserRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (u *UserHandler) GetAllUsers(c echo.Context) error {
	req := GetAllUserRequest{}
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
	err, users := u.UserRepo.GetAllUsers(c.Request().Context(), sql.GetAllUsersParams(req))
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
		Data:       users,
	})
}

type CheckAdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserHandler) CheckAdmin(c echo.Context) error {
	req := CheckAdminRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, user := u.UserRepo.GetUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Check failed",
			Data:       nil,
		})
	}

	if user.RoleTicker == "SAD" {
		return c.JSON(http.StatusOK, Response{
			StatusCode: http.StatusOK,
			Message:    "Success",
			Data:       true,
		})
	}

	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       false,
	})
}
