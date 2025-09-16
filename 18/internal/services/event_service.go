package services

import (
	"18/pkg/models"
	"errors"
	"time"
)

var (
	ErrEventNotFound = errors.New("event not found")
	ErrInvalidDate   = errors.New("invalid date")
)

type EventService struct {
	events         map[string][]models.Event
	eventIDCounter int
}

func NewEventService() *EventService {
	return &EventService{
		events:         make(map[string][]models.Event),
		eventIDCounter: 1,
	}
}

func (s *EventService) CreateEvent(userID, title, description, dateStr, startTime, endTime string) (models.Event, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return models.Event{}, ErrInvalidDate
	}

	event := models.Event{
		ID:          s.eventIDCounter,
		UserID:      userID,
		Title:       title,
		Description: description,
		Date:        date,
		StartTime:   startTime,
		EndTime:     endTime,
	}

	s.events[userID] = append(s.events[userID], event)
	s.eventIDCounter++

	return event, nil
}

func (s *EventService) UpdateEvent(eventID int, userID, title, description, dateStr, startTime, endTime string) (models.Event, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return models.Event{}, ErrInvalidDate
	}

	events, exists := s.events[userID]
	if !exists {
		return models.Event{}, ErrEventNotFound
	}

	for i, event := range events {
		if event.ID == eventID {
			updatedEvent := models.Event{
				ID:          eventID,
				UserID:      userID,
				Title:       title,
				Description: description,
				Date:        date,
				StartTime:   startTime,
				EndTime:     endTime,
			}
			s.events[userID][i] = updatedEvent
			return updatedEvent, nil
		}
	}

	return models.Event{}, ErrEventNotFound
}

func (s *EventService) DeleteEvent(eventID int, userID string) error {

	events, exists := s.events[userID]
	if !exists {
		return ErrEventNotFound
	}

	for i, event := range events {
		if event.ID == eventID {
			s.events[userID] = append(events[:i], events[i+1:]...)
			return nil
		}
	}

	return ErrEventNotFound
}

func (s *EventService) GetEventsForDay(userID, dateStr string) ([]models.Event, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, ErrInvalidDate
	}

	events, exists := s.events[userID]
	if !exists {
		return []models.Event{}, nil
	}

	var result []models.Event
	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.Date.Day() == date.Day() {
			result = append(result, event)
		}
	}

	return result, nil
}

func (s *EventService) GetEventsForWeek(userID, dateStr string) ([]models.Event, error) {

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, ErrInvalidDate
	}

	startOfWeek := date
	for startOfWeek.Weekday() != time.Monday {
		startOfWeek = startOfWeek.AddDate(0, 0, -1)
	}

	endOfWeek := startOfWeek.AddDate(0, 0, 6)

	events, exists := s.events[userID]
	if !exists {
		return []models.Event{}, nil
	}

	var result []models.Event
	for _, event := range events {
		if (event.Date.Equal(startOfWeek) || event.Date.After(startOfWeek)) &&
			(event.Date.Equal(endOfWeek) || event.Date.Before(endOfWeek)) {
			result = append(result, event)
		}
	}

	return result, nil
}

func (s *EventService) GetEventsForMonth(userID, dateStr string) ([]models.Event, error) {

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, ErrInvalidDate
	}

	startOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	events, exists := s.events[userID]
	if !exists {
		return []models.Event{}, nil
	}

	var result []models.Event
	for _, event := range events {
		if (event.Date.Equal(startOfMonth) || event.Date.After(startOfMonth)) &&
			(event.Date.Equal(endOfMonth) || event.Date.Before(endOfMonth)) {
			result = append(result, event)
		}
	}

	return result, nil
}
