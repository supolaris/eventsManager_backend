package routes

import "github.com/gin-gonic/gin"

func EventRoutes(server *gin.Engine) {
	server.GET("/getEvents", getEvents)
	server.POST("/addEvent", postEvents)
	server.PUT("/updateEvent/:eventId", updateEvent)
	server.GET("/getSingleEvent/:eventId", getSingleEvent)
}
