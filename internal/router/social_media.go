package router

import (
	"github.com/aurellieandra/my-gram/internal/handler"
	"github.com/gin-gonic/gin"
)

// INTERFACE
type SocialMediaRouter interface {
	Mount()
}

// STRUCT
type socialMediaRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.SocialMediaHandler
}

// NEW PHOTO ROUTER
func NewSocialMediaRouter(v *gin.RouterGroup, handler handler.SocialMediaHandler) SocialMediaRouter {
	return &socialMediaRouterImpl{v: v, handler: handler}
}

// PHOTO ROUTER IMPL
func (u *socialMediaRouterImpl) Mount() {
	u.v.GET("", u.handler.GetSocialMedias)
	u.v.GET("/:id", u.handler.GetSocialMediaById)

	u.v.POST("/", u.handler.CreateSocialMedia)
	u.v.PUT("/:id", u.handler.UpdateSocialMediaById)
	u.v.DELETE("/:id", u.handler.DeleteSocialMediaById)
}
