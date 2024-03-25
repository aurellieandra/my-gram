package middleware

import (
	"net/http"
	"strings"

	"github.com/aurellieandra/my-gram/pkg"
	"github.com/aurellieandra/my-gram/pkg/helper"
	"github.com/gin-gonic/gin"
)

const (
	CLAIM_USER_ID  = "claim_user_id"
	CLAIM_USERNAME = "claim_username"
)

func CheckAuthBearer(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	authArr := strings.Split(auth, " ")

	if len(authArr) < 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, pkg.Response{
			Status:  http.StatusUnauthorized,
			Message: "Invalid Token",
			Data:    nil,
		})
		return
	}

	if authArr[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, pkg.Response{
			Status:  http.StatusUnauthorized,
			Message: "Invalid Authorization Method",
			Data:    nil,
		})
		return
	}

	token := authArr[1]
	claims, err := helper.ValidateToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, pkg.Response{
			Status:  http.StatusUnauthorized,
			Message: "Failed to Decode",
			Data:    nil,
		})
		return
	}

	ctx.Set(CLAIM_USER_ID, claims["user_id"])
	ctx.Set(CLAIM_USERNAME, claims["username"])
	ctx.Next()
}
