package router

import (
	"github.com/aurellieandra/my-gram/internal/handler"
	"github.com/gin-gonic/gin"
)

// INTERFACE
type PhotoRouter interface {
	Mount()
}

// STRUCT
type photoRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.PhotoHandler
}

// NEW PHOTO ROUTER
func NewPhotoRouter(v *gin.RouterGroup, handler handler.PhotoHandler) PhotoRouter {
	return &photoRouterImpl{v: v, handler: handler}
}

// PHOTO ROUTER IMPL
func (u *photoRouterImpl) Mount() {
	u.v.GET("", u.handler.GetPhotos)
	u.v.GET("/:id", u.handler.GetPhotoById)

	u.v.POST("/", u.handler.CreatePhoto)
	u.v.PUT("/:id", u.handler.UpdatePhotoById)
	u.v.DELETE("/:id", u.handler.DeletePhotoById)
}
