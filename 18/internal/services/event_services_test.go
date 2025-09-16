package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateEvent(t *testing.T) {
	service := NewEventService()

	event, err := service.CreateEvent("user1", "Meeting", "Team meeting", "2025-10-16", "10:00", "11:00")

	assert.NoError(t, err)
	assert.Equal(t, "Meeting", event.Title)
	assert.Equal(t, "user1", event.UserID)
	assert.Equal(t, "Team meeting", event.Description)
	assert.Equal(t, "10:00", event.StartTime)
	assert.Equal(t, "11:00", event.EndTime)
}

func Test_CreateEvent_InvalidDate(t *testing.T) {
	service := NewEventService()

	_, err := service.CreateEvent("user1", "Meeting", "Team meeting", "invalid-date", "10:00", "11:00")

	assert.Equal(t, ErrInvalidDate, err)
}

func Test_GetEventsForDay(t *testing.T) {
	service := NewEventService()

	_, err := service.CreateEvent("user1", "Meeting", "Team meeting", "2025-10-16", "10:00", "11:00")
	assert.NoError(t, err)

	events, err := service.GetEventsForDay("user1", "2025-10-16")

	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, "Meeting", events[0].Title)
}

func Test_GetEventsForDay_Empty(t *testing.T) {
	service := NewEventService()

	events, err := service.GetEventsForDay("user1", "2025-10-16")

	assert.NoError(t, err)
	assert.Empty(t, events)
}

func Test_GetEventsForDay_InvalidDate(t *testing.T) {
	service := NewEventService()

	_, err := service.GetEventsForDay("user1", "invalid-date")

	assert.Equal(t, ErrInvalidDate, err)
}

func Test_GetEventsForWeek(t *testing.T) {
	service := NewEventService()

	_, err := service.CreateEvent("user1", "Monday Meeting", "Meeting", "2025-10-15", "10:00", "11:00")
	assert.NoError(t, err)
	_, err = service.CreateEvent("user1", "Friday Meeting", "Meeting", "2025-10-16", "14:00", "15:00")
	assert.NoError(t, err)

	events, err := service.GetEventsForWeek("user1", "2025-10-15")

	assert.NoError(t, err)
	assert.Len(t, events, 2)
}

func Test_GetEventsForMonth(t *testing.T) {
	service := NewEventService()

	_, err := service.CreateEvent("user1", "Early", "", "2025-10-15", "10:00", "11:00")
	assert.NoError(t, err)
	_, err = service.CreateEvent("user1", "Late", "", "2025-10-16", "14:00", "15:00")
	assert.NoError(t, err)

	events, err := service.GetEventsForMonth("user1", "2025-10-16")

	assert.NoError(t, err)
	assert.Len(t, events, 2)
}

func Test_UpdateEvent(t *testing.T) {
	service := NewEventService()

	event, err := service.CreateEvent("user1", "Old Title", "Old Desc", "2025-10-16", "10:00", "11:00")
	assert.NoError(t, err)

	updated, err := service.UpdateEvent(event.ID, "user1", "New Title", "New Desc", "2025-10-16", "12:00", "13:00")

	assert.NoError(t, err)
	assert.Equal(t, "New Title", updated.Title)
	assert.Equal(t, "New Desc", updated.Description)
	assert.Equal(t, "12:00", updated.StartTime)
	assert.Equal(t, "13:00", updated.EndTime)
}

func Test_UpdateEvent_NotFound(t *testing.T) {
	service := NewEventService()

	_, err := service.UpdateEvent(123445, "user1", "New Title", "New Desc", "2025-10-16", "12:00", "13:00")

	assert.Equal(t, ErrEventNotFound, err)
}

func Test_DeleteEvent(t *testing.T) {
	service := NewEventService()

	event, err := service.CreateEvent("user1", "Meeting", "Desc", "2025-10-16", "10:00", "11:00")
	assert.NoError(t, err)

	err = service.DeleteEvent(event.ID, "user1")

	assert.NoError(t, err)

	events, err := service.GetEventsForDay("user1", "2025-10-16")
	assert.NoError(t, err)
	assert.Empty(t, events)
}

func Test_DeleteEvent_NotFound(t *testing.T) {
	service := NewEventService()

	err := service.DeleteEvent(914632, "user1")

	assert.Equal(t, ErrEventNotFound, err)
}
