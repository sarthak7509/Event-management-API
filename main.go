package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/event-management/db"
	"github.com/sarthak7509/event-management/models"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", func(ctx *gin.Context) {
		events, err := models.GetAllEvents()
		if err != nil {
			log.Fatal(err)
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, events)
	})
	server.POST("/events", func(ctx *gin.Context) {
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
	})
	server.Run()
}
