package routes

import "github.com/gin-gonic/gin"

func EventRoutes(server *gin.Engine) {
	server.GET("/getEvents", getEvents)
	server.POST("/addEvent", postEvents)
	server.PUT("/updateEvent/:eventId", updateEvent)
	server.DELETE("/deleteEvent/:eventId", deleteEvent)
	server.GET("/getSingleEvent/:eventId", getSingleEvent)

	//users
	server.POST("/addUser", addUser)
	server.GET("/getUsers", getUsers)
	server.POST("/userLogin", loginUser)
}
