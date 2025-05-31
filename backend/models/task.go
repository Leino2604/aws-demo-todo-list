package models

import "time"

type Task struct {
    ID             string    `json:"id"`
    Name           string    `json:"name"`
    EstimatedTime  int       `json:"estimated_time"` // in minutes
    Status         string    `json:"status"`         // e.g., "pending", "in_progress", "completed"
    StartTime      time.Time `json:"start_time,omitempty"`
    PomodoroCycles int       `json:"pomodoro_cycles"` // number of pomodoro cycles completed
}