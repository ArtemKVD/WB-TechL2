package handlers

import (
	"net/http"
	"strconv"

	"18/internal/services"
	"18/pkg/models"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service *services.EventService
}

func NewEventHandler(service *services.EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var req models.EventRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	event, err := h.service.CreateEvent(
		req.UserID,
		req.Title,
		req.Description,
		req.Date,
		req.StartTime,
		req.EndTime,
	)

	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.Response{Result: event})
}

func (h *EventHandler) UpdateEvent(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	var req models.EventRequest
	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	event, err := h.service.UpdateEvent(
		eventID,
		req.UserID,
		req.Title,
		req.Description,
		req.Date,
		req.StartTime,
		req.EndTime,
	)

	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.Response{Result: event})
}

func (h *EventHandler) DeleteEvent(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: "Invalid ID"})
		return
	}

	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, models.Response{Error: "user_id parameter empty"})
		return
	}

	err = h.service.DeleteEvent(eventID, userID)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.Response{Result: "Event deleted"})
}

func (h *EventHandler) GetEventsForDay(c *gin.Context) {
	userID := c.Query("user_id")
	date := c.Query("date")
	events, err := h.service.GetEventsForDay(userID, date)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.Response{Result: events})
}

func (h *EventHandler) GetEventsForWeek(c *gin.Context) {
	userID := c.Query("user_id")
	date := c.Query("date")
	events, err := h.service.GetEventsForWeek(userID, date)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.Response{Result: events})
}

func (h *EventHandler) GetEventsForMonth(c *gin.Context) {
	userID := c.Query("user_id")
	date := c.Query("date")
	events, err := h.service.GetEventsForMonth(userID, date)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.Response{Result: events})
}

func (h *EventHandler) handleServiceError(c *gin.Context, err error) {
	switch err {
	case services.ErrEventNotFound:
		c.JSON(http.StatusServiceUnavailable, models.Response{Error: err.Error()})
	case services.ErrInvalidDate:
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, models.Response{Error: "Internal server error"})
	}
}
