package handler

import (
	"net/http"
	"strconv"

	"github.com/aurellieandra/my-gram/internal/model"
	"github.com/aurellieandra/my-gram/internal/service"
	"github.com/aurellieandra/my-gram/pkg"
	"github.com/gin-gonic/gin"
)

// INTERFACE
type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)

	GetUsers(ctx *gin.Context)
	GetUserById(ctx *gin.Context)

	UpdateUserById(ctx *gin.Context)
	DeleteUserById(ctx *gin.Context)
}

// STRUCT
type userHandlerImpl struct {
	svc service.UserService
}

// NEW USER HANDLER
func NewUserHandler(svc service.UserService) UserHandler {
	return &userHandlerImpl{
		svc: svc,
	}
}

// USER HANDLER IMPL
func (u *userHandlerImpl) Register(ctx *gin.Context) {
	var newUser model.User

	if err := ctx.Bind(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Bind payload failure",
			Data:    nil,
		})
		return
	}

	user, err := u.svc.Register(ctx, newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Register service failure",
			Data:    nil,
		})
		return
	}

	token, err := u.svc.GenerateUserAccessToken(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Generate token failure",
			Data:    nil,
		})
		return
	}

	userResponse := model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Dob:       user.Dob,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	ctx.JSON(http.StatusOK, pkg.AuthResponse{
		Status:  http.StatusCreated,
		Message: "User registered successfully",
		Data:    userResponse,
		Token:   token,
	})
}

func (u *userHandlerImpl) Login(ctx *gin.Context) {
	var credentials model.User
	if err := ctx.BindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Bind credentials failure",
			Data:    nil,
		})
		return
	}

	user, err := u.svc.Login(ctx, credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid username or password",
			Data:    nil,
		})
		return
	}

	token, err := u.svc.GenerateUserAccessToken(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Generate token failure",
			Data:    nil,
		})
		return
	}

	userResponse := model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Dob:       user.Dob,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	ctx.JSON(http.StatusOK, pkg.AuthResponse{
		Status:  http.StatusOK,
		Message: "Login successfully",
		Data:    userResponse,
		Token:   token,
	})
}

func (u *userHandlerImpl) GetUsers(ctx *gin.Context) {
	users, err := u.svc.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Get users service failure",
			Data:    nil,
		})
		return
	}

	var userResponses []model.UserResponse
	for _, user := range users {
		userResponse := model.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Dob:       user.Dob,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		}

		userResponses = append(userResponses, userResponse)
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Get users data successfully",
		Data:    userResponses,
	})
}

func (u *userHandlerImpl) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	user, err := u.svc.GetUserById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Get User by ID service failure",
			Data:    nil,
		})
		return
	} else if user == nil {
		ctx.JSON(http.StatusNotFound, pkg.Response{
			Status:  http.StatusNotFound,
			Message: "Data Not Found",
			Data:    nil,
		})
		return
	}

	userResponse := model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Dob:       user.Dob,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Get user data successfully",
		Data:    userResponse,
	})
}

func (u *userHandlerImpl) UpdateUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	var newUser model.User
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Generate token failed",
			Data:    nil,
		})
		return
	}

	user, err := u.svc.UpdateUserById(ctx, newUser, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Update User by ID service failed",
			Data:    nil,
		})
		return
	}

	userResponse := model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Dob:       user.Dob,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Update user data successfully",
		Data:    userResponse,
	})
}

func (u *userHandlerImpl) DeleteUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	err = u.svc.DeleteUserById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Delete User by ID service failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Delete user data successfully",
		Data:    nil,
	})
}
