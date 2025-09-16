package models

import "time"

type Calendar struct {
	Year  int
	Month time.Month
	Weeks [][]CalendarDay
}

type CalendarDay struct {
	Date           time.Time
	Day            int
	IsToday        bool
	IsCurrentMonth bool
	Events         []Event
}
