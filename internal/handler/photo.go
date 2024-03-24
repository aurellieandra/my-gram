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
type PhotoHandler interface {
	GetPhotos(ctx *gin.Context)
	GetPhotosByUserId(ctx *gin.Context)
	GetPhotoById(ctx *gin.Context)

	CreatePhoto(ctx *gin.Context)
	UpdatePhotoById(ctx *gin.Context)
	DeletePhotoById(ctx *gin.Context)
}

// STRUCT
type photoHandlerImpl struct {
	svc service.PhotoService
}

// NEW PHOTO HANDLER
func NewPhotoHandler(svc service.PhotoService) PhotoHandler {
	return &photoHandlerImpl{
		svc: svc,
	}
}

// PHOTO HANDLER IMPL
func (u *photoHandlerImpl) GetPhotos(ctx *gin.Context) {
	photos, err := u.svc.GetPhotos(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Get photos service failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Get photos data successfully",
		Data:    photos,
	})
}

func (u *photoHandlerImpl) GetPhotosByUserId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	photos, err := u.svc.GetPhotosByUserId(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Get photos by photo id service failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Get photos data successfully",
		Data:    photos,
	})
}

func (u *photoHandlerImpl) GetPhotoById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	photo, err := u.svc.GetPhotoById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Get Photo by ID service failure",
			Data:    nil,
		})
		return
	} else if photo == nil {
		ctx.JSON(http.StatusNotFound, pkg.Response{
			Status:  http.StatusNotFound,
			Message: "Data Not Found",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Get photo data successfully",
		Data:    photo,
	})
}

func (u *photoHandlerImpl) CreatePhoto(ctx *gin.Context) {
	var newPhoto model.Photo
	newPhoto.User_Id = 4

	if err := ctx.Bind(&newPhoto); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Bind payload failure",
			Data:    nil,
		})
		return
	}

	photo, err := u.svc.CreatePhoto(ctx, newPhoto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Create photo service failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusCreated,
		Message: "Photo created successfully",
		Data:    photo,
	})
}

func (u *photoHandlerImpl) UpdatePhotoById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	var newPhoto model.Photo
	if err := ctx.BindJSON(&newPhoto); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Bind payload failure",
			Data:    nil,
		})
		return
	}

	photo, err := u.svc.UpdatePhotoById(ctx, newPhoto, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Update photo by ID failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Update photo data successfully",
		Data:    photo,
	})
}

func (u *photoHandlerImpl) DeletePhotoById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	err = u.svc.DeletePhotoById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Delete photo by ID service failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Delete photo data successfully",
		Data:    nil,
	})
}
