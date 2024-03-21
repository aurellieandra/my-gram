package router

import (
	"github.com/aurellieandra/my-gram/internal/handler"
	"github.com/gin-gonic/gin"
)

// INTERFACE
type CommentRouter interface {
	Mount()
}

// STRUCT
type commentRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.CommentHandler
}

// NEW PHOTO ROUTER
func NewCommentRouter(v *gin.RouterGroup, handler handler.CommentHandler) CommentRouter {
	return &commentRouterImpl{v: v, handler: handler}
}

// PHOTO ROUTER IMPL
func (u *commentRouterImpl) Mount() {
	u.v.GET("/", u.handler.GetCommentsByPhotoId)
	u.v.GET("/:id", u.handler.GetCommentById)
	u.v.POST("/", u.handler.CreateComment)
	u.v.PUT("/:id", u.handler.UpdateCommentById)
	u.v.DELETE("/:id", u.handler.DeleteCommentById)
}