package model

import (
	"time"
)

const (
	JobHistoryStatusCompleted = "COMPLETED"
	JobHistoryStatusFailed    = "FAILED"
	JobHistoryStatusRunning   = "RUNNING"
	JobHistoryStatusPending   = "PENDING"
	jobHistoryStatusScheduled = "SCHEDULED"
)

type JobHistory struct {
	JobID          string    `json:"job_id"`         // Partition Key
	ExecutionTime  time.Time `json:"execution_time"` // Sort Key (Clustering Key)
	Status         string    `json:"status"`
	RetryCount     int       `json:"retry_count"`
	LastUpdateTime time.Time `json:"last_update_time"`
}
