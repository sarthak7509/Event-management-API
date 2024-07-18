package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/event-management/handlers"
)

func RegisterEventRoute(server *gin.Engine) {
	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:id", handlers.GetEventById)
	server.PUT("/events/:id", handlers.UpdateEventById)
	server.POST("/events", handlers.CreateEvent)

}
