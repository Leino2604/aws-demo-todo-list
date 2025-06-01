package models

import "time"

type Todo struct {
    ID          string    `json:"id"`
    Description string    `json:"description"`
    EstimatedTime time.Duration `json:"estimated_time"`
    Status      string    `json:"status"` // e.g., "pending", "in_progress", "completed"
    StartTime   time.Time `json:"start_time,omitempty"` // optional, only set when the task is started
}