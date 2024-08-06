package model

import "time"

type Job struct {
	UserID      string    `json:"user_id"` // Partition Key
	JobID       string    `json:"job_id"`  // Sort Key (Clustering Key)
	IsRecurring bool      `json:"is_recurring"`
	RetryTimes  int       `json:"retry_times"`
	Created     time.Time `json:"created"`
	Interval    string    `json:"interval"`
}
