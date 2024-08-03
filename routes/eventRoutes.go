package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/event-management/handlers"
)

func RegisterEventRoute(server *gin.Engine) {
	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:id", handlers.GetEventById)
	server.PUT("/events/:id", handlers.UpdateEventById)
	server.DELETE("/events/:id", handlers.DeleteEventById)
	server.POST("/events", handlers.CreateEvent)
	//user logins
	server.POST("/signup", handlers.SignUp)
	server.GET("/login", handlers.LogIn)
}
