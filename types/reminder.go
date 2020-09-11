package types

import "time"

// Reminder represents a task reminder
type Reminder struct {
	Task Task
	Time time.Time
}
