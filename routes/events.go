package routes

import (
	"basicapis/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func getSingleEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("eventId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error message": "error in parsing event id",
		})
		return
	}
	event, err := models.GetSingleEvent(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error message": "internal server error get single event",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "event got from db",
		"event":   event,
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

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("eventId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error message": "error in parsing event id",
		})
		return
	}
	_, err = models.GetSingleEvent(eventId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error message": "event not found",
		})
		return
	}
	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error message": "event not found",
		})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error message": "could not update event",
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"message":      "event updated",
		"updatedEvent": updatedEvent,
	})

}
