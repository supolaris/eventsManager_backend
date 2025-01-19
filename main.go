package main

import (
	"basicapis/db"
	"basicapis/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.EventRoutes(server)

	server.Run(":8080")
}
