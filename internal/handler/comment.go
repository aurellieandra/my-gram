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
type CommentHandler interface {
	GetComments(ctx *gin.Context)
	GetCommentById(ctx *gin.Context)

	CreateComment(ctx *gin.Context)
	UpdateCommentById(ctx *gin.Context)
	DeleteCommentById(ctx *gin.Context)
}

// STRUCT
type commentHandlerImpl struct {
	svc service.CommentService
}

// NEW PHOTO HANDLER
func NewCommentHandler(svc service.CommentService) CommentHandler {
	return &commentHandlerImpl{
		svc: svc,
	}
}

// PHOTO HANDLER IMPL
func (u *commentHandlerImpl) GetComments(ctx *gin.Context) {
	photo_id_str := ctx.Query("photo_id")
	var photo_id uint64

	if photo_id_str != "" {
		id, err := strconv.ParseUint(photo_id_str, 10, 64)
		if id == 0 || err != nil {
			ctx.JSON(http.StatusBadRequest, pkg.Response{
				Status:  http.StatusBadRequest,
				Message: "Fetching ID from query param failure",
				Data:    nil,
			})
			return
		}
		photo_id = id
	}

	comments, err := u.svc.GetComments(ctx, &photo_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Get comments service failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Get comments data successfully",
		Data:    comments,
	})
}

func (u *commentHandlerImpl) GetCommentById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	comment, err := u.svc.GetCommentById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Get Social Media by ID service failure",
			Data:    nil,
		})
		return
	} else if comment == nil || comment.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.Response{
			Status:  http.StatusNotFound,
			Message: "Data Not Found",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Get comment data successfully",
		Data:    comment,
	})
}

func (u *commentHandlerImpl) CreateComment(ctx *gin.Context) {
	var newComment model.Comment
	newComment.User_Id = 4

	if err := ctx.Bind(&newComment); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Bind payload failure",
			Data:    nil,
		})
		return
	}

	comment, err := u.svc.CreateComment(ctx, newComment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Create comment service failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusCreated,
		Message: "Social media created successfully",
		Data:    comment,
	})
}

func (u *commentHandlerImpl) UpdateCommentById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	var newComment model.Comment
	if err := ctx.BindJSON(&newComment); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Bind payload failure",
			Data:    nil,
		})
		return
	}

	comment, err := u.svc.UpdateCommentById(ctx, newComment, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Update comment by ID failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Update comment data successfully",
		Data:    comment,
	})
}

func (u *commentHandlerImpl) DeleteCommentById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	err = u.svc.DeleteCommentById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Delete comment by ID service failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Delete comment data successfully",
		Data:    nil,
	})
}
