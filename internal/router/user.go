package router

import (
	"github.com/aurellieandra/my-gram/internal/handler"
	"github.com/gin-gonic/gin"
)

// INTERFACE
type UserRouter interface {
	Mount()
}

// STRUCT
type userRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.UserHandler
}

// NEW USER ROUTER
func NewUserRouter(v *gin.RouterGroup, handler handler.UserHandler) UserRouter {
	return &userRouterImpl{v: v, handler: handler}
}

// USER ROUTER IMPL
func (u *userRouterImpl) Mount() {
	u.v.GET("/", u.handler.GetUsers)
	u.v.GET("/:id", u.handler.GetUserById)
	u.v.POST("/register", u.handler.Register)
	u.v.GET("/login", u.handler.Login)
	u.v.PUT("/:id", u.handler.UpdateUserById)
	u.v.DELETE("/:id", u.handler.DeleteUserById)
}
