package main

import (
	"log"

	"18/internal/config"
	"18/internal/handlers"
	"18/internal/middleware"
	"18/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	eventService := services.NewEventService()
	eventHandler := handlers.NewEventHandler(eventService)

	router := gin.Default()

	router.Use(middleware.LoggingMiddleware())

	api := router.Group("/api")
	{
		api.POST("/create_event", eventHandler.CreateEvent)
		api.POST("/update_event/:id", eventHandler.UpdateEvent)
		api.POST("/delete_event/:id", eventHandler.DeleteEvent)
		api.GET("/events_for_day", eventHandler.GetEventsForDay)
		api.GET("/events_for_week", eventHandler.GetEventsForWeek)
		api.GET("/events_for_month", eventHandler.GetEventsForMonth)
	}

	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
