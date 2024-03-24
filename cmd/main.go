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

	socialMediasGroup := v1.Group("/social-medias")
	socialMediaRepo := repository.NewSocialMediaQuery(gorm)
	socialMediaCmd := repository.NewSocialMediaCommand(gorm)
	socialMediaSvc := service.NewSocialMediaService(socialMediaRepo, socialMediaCmd)
	socialMediaHdl := handler.NewSocialMediaHandler(socialMediaSvc)
	socialMediaRouter := router.NewSocialMediaRouter(socialMediasGroup, socialMediaHdl)

	commentsGroup := v1.Group("/comments")
	commentRepo := repository.NewCommentQuery(gorm)
	commentCmd := repository.NewCommentCommand(gorm)
	commentSvc := service.NewCommentService(commentRepo, commentCmd)
	commentHdl := handler.NewCommentHandler(commentSvc)
	commentRouter := router.NewCommentRouter(commentsGroup, commentHdl)

	userRouter.Mount()
	photoRouter.Mount()
	socialMediaRouter.Mount()
	commentRouter.Mount()

	g.Run(":3000")
}
