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
type SocialMediaHandler interface {
	GetSocialMedias(ctx *gin.Context)
	GetSocialMediaById(ctx *gin.Context)

	CreateSocialMedia(ctx *gin.Context)
	UpdateSocialMediaById(ctx *gin.Context)
	DeleteSocialMediaById(ctx *gin.Context)
}

// STRUCT
type socialMediaHandlerImpl struct {
	svc service.SocialMediaService
}

// NEW PHOTO HANDLER
func NewSocialMediaHandler(svc service.SocialMediaService) SocialMediaHandler {
	return &socialMediaHandlerImpl{
		svc: svc,
	}
}

// PHOTO HANDLER IMPL
func (u *socialMediaHandlerImpl) GetSocialMedias(ctx *gin.Context) {
	user_id_str := ctx.Query("user_id")
	var user_id uint64

	if user_id_str != "" {
		id, err := strconv.ParseUint(user_id_str, 10, 64)
		if id == 0 || err != nil {
			ctx.JSON(http.StatusBadRequest, pkg.Response{
				Status:  http.StatusBadRequest,
				Message: "Fetching ID from query param failure",
				Data:    nil,
			})
			return
		}
		user_id = id
	}

	social_medias, err := u.svc.GetSocialMedias(ctx, &user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Get social medias service failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Get social medias data successfully",
		Data:    social_medias,
	})
}

func (u *socialMediaHandlerImpl) GetSocialMediaById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	social_media, err := u.svc.GetSocialMediaById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Get Social Media by ID service failure",
			Data:    nil,
		})
		return
	} else if social_media == nil || social_media.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.Response{
			Status:  http.StatusNotFound,
			Message: "Data Not Found",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Get social media data successfully",
		Data:    social_media,
	})
}

func (u *socialMediaHandlerImpl) CreateSocialMedia(ctx *gin.Context) {
	var newSocialMedia model.SocialMedia
	newSocialMedia.User_Id = 4

	if err := ctx.Bind(&newSocialMedia); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Bind payload failure",
			Data:    nil,
		})
		return
	}

	social_media, err := u.svc.CreateSocialMedia(ctx, newSocialMedia)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Create social media service failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusCreated,
		Message: "Social media created successfully",
		Data:    social_media,
	})
}

func (u *socialMediaHandlerImpl) UpdateSocialMediaById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	var newSocialMedia model.SocialMedia
	if err := ctx.BindJSON(&newSocialMedia); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Bind payload failure",
			Data:    nil,
		})
		return
	}

	social_media, err := u.svc.UpdateSocialMediaById(ctx, newSocialMedia, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Update social media by ID failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Update social media data successfully",
		Data:    social_media,
	})
}

func (u *socialMediaHandlerImpl) DeleteSocialMediaById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Fetching ID from param failure",
			Data:    nil,
		})
		return
	}

	err = u.svc.DeleteSocialMediaById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.Response{
			Status:  http.StatusBadRequest,
			Message: "Delete social media by ID service failure",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.Response{
		Status:  http.StatusOK,
		Message: "Delete social media data successfully",
		Data:    nil,
	})
}
