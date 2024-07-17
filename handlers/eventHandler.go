package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/event-management/models"
)

func CreateEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	event.ID = 1
	event.UserId = 22
	err = event.Save()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(201, event)
}

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, events)
}

func GetEventById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Could not found the event with that id",
		})
		return
	}
	ctx.JSON(200, event)
}
