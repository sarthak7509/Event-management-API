package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sarthak7509/event-management/db"
	route "github.com/sarthak7509/event-management/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	route.RegisterEventRoute(server)
	server.Run()
}
