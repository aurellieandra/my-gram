package handler

import (
	"net/http"
	"strconv"

	"github.com/aurellieandra/my-gram/internal/service"
	"github.com/aurellieandra/my-gram/pkg"
	"github.com/gin-gonic/gin"
)

// INTERFACE
type UserHandler interface {
	GetUsers(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	// Register(ctx *gin.Context)
	// Login(ctx *gin.Context)
	// Logout(ctx *gin.Context)
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
func (u *userHandlerImpl) GetUsers(ctx *gin.Context) {
	users, err := u.svc.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
func (u *userHandlerImpl) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	user, err := u.svc.GetUserById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
// func (u *userHandlerImpl) Register(ctx *gin.Context) {
// 	id, err := strconv.Atoi(ctx.PostForm(order))
// 	if id == 0 || err != nil {
// 		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
// 		return
// 	}
// 	user, err := u.svc.Register(ctx, uint64(id))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, user)
// }
// func (u *userHandlerImpl) Logout(ctx *gin.Context) {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if id == 0 || err != nil {
// 		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
// 		return
// 	}
// 	err = u.svc.DeleteUserById(ctx, uint64(id))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, nil)
// }
func (u *userHandlerImpl) UpdateUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	user, err := u.svc.UpdateUserById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
func (u *userHandlerImpl) DeleteUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	err = u.svc.DeleteUserById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
