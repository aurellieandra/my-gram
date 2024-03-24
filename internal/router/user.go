package router

import (
	"github.com/aurellieandra/my-gram/internal/handler"
	"github.com/aurellieandra/my-gram/internal/middleware"
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
	u.v.POST("/register", u.handler.Register)
	u.v.POST("/login", u.handler.Login)

	u.v.Use(middleware.CheckAuthBearer)
	u.v.GET("/", u.handler.GetUsers)
	u.v.GET("/:id", u.handler.GetUserById)
	u.v.PUT("/:id", u.handler.UpdateUserById)
	u.v.DELETE("/:id", u.handler.DeleteUserById)
	u.v.GET("/logout", u.handler.Logout)
}
