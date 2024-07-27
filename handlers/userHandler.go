package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/event-management/models"
)

func SignUp(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	err = user.Save()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(201, user)
}
