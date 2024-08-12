package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/event-management/models"
	"github.com/sarthak7509/event-management/utils"
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

func LogIn(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = user.Validate()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, err := utils.GenerateToken(user.Id, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
