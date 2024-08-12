package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/event-management/utils"
)

func Authentication(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "No Auth token added",
		})
		return

	}
	userid, err := utils.VerifyToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Set("userId", userid)

	ctx.Next()

}
