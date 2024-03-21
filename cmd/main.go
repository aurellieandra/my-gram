package main

import (
	"github.com/aurellieandra/my-gram/internal/handler"
	"github.com/aurellieandra/my-gram/internal/infrastructure"
	"github.com/aurellieandra/my-gram/internal/repository"
	"github.com/aurellieandra/my-gram/internal/router"
	"github.com/aurellieandra/my-gram/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	usersGroup := g.Group("/users")

	gorm := infrastructure.NewGormPostgres()
	userRepo := repository.NewUserQuery(gorm)
	userCmd := repository.NewUserCommand(gorm)
	userSvc := service.NewUserService(userRepo, userCmd)
	userHdl := handler.NewUserHandler(userSvc)
	userRouter := router.NewUserRouter(usersGroup, userHdl)

	userRouter.Mount()

	g.Run(":3000")
}