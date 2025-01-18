package main

import (
	"basicapis/db"
	"basicapis/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/getEvents", getEvents)
	server.POST("/addEvent", postEvents)
	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error message": "internal server error get events",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"all events": events,
	})
}

func postEvents(ctx *gin.Context) {
	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error message": "error in binding json",
		})
		return
	}

	// event.ID = 1
	event.UserID = 2

	err = event.SaveEvent()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error message": "internal server error in save event",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "event created",
		"event":   event,
	})
}
