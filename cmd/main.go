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
	gorm := infrastructure.NewGormPostgres()

	v1 := g.Group("/api/v1")

	usersGroup := v1.Group("/users")
	userRepo := repository.NewUserQuery(gorm)
	userCmd := repository.NewUserCommand(gorm)
	userSvc := service.NewUserService(userRepo, userCmd)
	userHdl := handler.NewUserHandler(userSvc)
	userRouter := router.NewUserRouter(usersGroup, userHdl)

	photosGroup := v1.Group("/photos")
	photoRepo := repository.NewPhotoQuery(gorm)
	photoCmd := repository.NewPhotoCommand(gorm)
	photoSvc := service.NewPhotoService(photoRepo, photoCmd)
	photoHdl := handler.NewPhotoHandler(photoSvc)
	photoRouter := router.NewPhotoRouter(photosGroup, photoHdl)

	userRouter.Mount()
	photoRouter.Mount()

	g.Run(":3000")
}
