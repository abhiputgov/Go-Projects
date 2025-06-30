package main

import (
	models "go/restapi/models"

	http "net/http"

	gin "github.com/gin-gonic/gin"

	db "go/restapi/db"
)

func main() {
	db.IniDB()
	engine := gin.Default()
	engine.GET("/", home)
	engine.GET("/events", listEvents)
	engine.POST("/events", createEvent)
	engine.Run(":8080") //localhost:8080 for dev
}

func home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"destination": "home"})
}
func listEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent(context *gin.Context) {
	var event models.Event
	error := context.ShouldBindJSON(&event)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error_message": "Could not parse JSON object!", "error": error})
		return
	}

	event.EventID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event is created", "event": event})
}
