package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	UserID      string    `json:"user_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" binding:"required"`
	StartTime   string    `json:"start_time"`
	EndTime     string    `json:"end_time"`
}

type EventRequest struct {
	UserID      string `json:"user_id" form:"user_id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description"`
	Date        string `json:"date" form:"date" binding:"required"`
	StartTime   string `json:"start_time" form:"start_time"`
	EndTime     string `json:"end_time" form:"end_time"`
}

type Response struct {
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}
