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
	userid := ctx.GetInt64("userId")
	event.UserId = userid
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

func UpdateEventById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	event, err := models.GetEventById(id)
	userid := ctx.GetInt64("userId")

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Could not found the event with that id",
		})
		return
	}
	if event.UserId != userid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You cant update this event",
		})
		return
	}
	var updateEvent models.Event
	err = ctx.ShouldBindJSON(&updateEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	updateEvent.ID = id
	err = updateEvent.Update()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Could not update the event please contact administrator",
		})
		return
	}
	ctx.JSON(201, updateEvent)
}

func DeleteEventById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	userid := ctx.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "No Event with that id",
		})
		return
	}
	if event.UserId != userid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You cant delete this event",
		})
		return
	}
	err = event.Delete()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "could not delete the event",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Event Deleted successfully",
	})
}
