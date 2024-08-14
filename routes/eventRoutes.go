package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/event-management/handlers"
	"github.com/sarthak7509/event-management/middleware"
)

func RegisterEventRoute(server *gin.Engine) {
	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:id", handlers.GetEventById)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authentication)
	authenticated.PUT("/events/:id", handlers.UpdateEventById)
	authenticated.DELETE("/events/:id", handlers.DeleteEventById)
	authenticated.POST("/events", handlers.CreateEvent)
	authenticated.POST("/")

	//user logins
	server.POST("/signup", handlers.SignUp)
	server.GET("/login", handlers.LogIn)
}
