package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"18/internal/services"
	"18/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_CreateEvent(t *testing.T) {
	service := services.NewEventService()
	handler := NewEventHandler(service)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/create_event", handler.CreateEvent)

	reqBody := models.EventRequest{
		UserID: "test-user",
		Title:  "Test Event",
		Date:   "2023-12-31",
	}
	jsonData, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/create_event", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response models.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Result)
	assert.Empty(t, response.Error)
}

func Test_CreateEvent_InvalidRequest(t *testing.T) {
	service := services.NewEventService()
	handler := NewEventHandler(service)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/create_event", handler.CreateEvent)

	reqBody := map[string]interface{}{
		"user_id": "test-user",
	}
	jsonData, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/create_event", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

	var response models.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Error)
}

func Test_GetEventsForDay(t *testing.T) {
	service := services.NewEventService()
	handler := NewEventHandler(service)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/api/create_event", handler.CreateEvent)
	router.GET("/api/events_for_day", handler.GetEventsForDay)

	createReq := models.EventRequest{
		UserID: "test-user",
		Title:  "Test Event",
		Date:   "2023-12-31",
	}
	jsonData, _ := json.Marshal(createReq)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/create_event", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/events_for_day?user_id=test-user&date=2023-12-31", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response models.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	events, ok := response.Result.([]interface{})
	assert.True(t, ok)
	assert.Len(t, events, 1)
}

func Test_GetEventsForDay_Empty(t *testing.T) {
	service := services.NewEventService()
	handler := NewEventHandler(service)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/api/events_for_day", handler.GetEventsForDay)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/events_for_day?user_id=test-user&date=2023-12-31", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response models.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	events, ok := response.Result.([]interface{})
	assert.True(t, ok)
	assert.Empty(t, events)
}
