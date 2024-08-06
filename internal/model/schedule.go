package model

type JobSchedule struct {
	NextExecutionTime int64  `json:"next_execution_time"` // Partition Key
	JobID             string `json:"job_id"`              // Sort Key (Clustering Key)
}
