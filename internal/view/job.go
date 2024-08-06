package view

import (
	"time"

	"github.com/quanghuynguyen1902/job-schedule/internal/model"
)

type Job struct {
	UserID      string    `json:"user_id"` // Partition Key
	JobID       string    `json:"job_id"`  // Sort Key (Clustering Key)
	RetryTimes  int       `json:"retry_times"`
	IsRecurring bool      `json:"is_recurring"`
	Created     time.Time `json:"created"`
	Interval    string    `json:"interval"`
}

func ToJob(job model.Job) Job {
	rs := Job{
		UserID:      job.UserID,
		JobID:       job.JobID,
		IsRecurring: job.IsRecurring,
		RetryTimes:  job.RetryTimes,
		Created:     job.Created,
		Interval:    job.Interval,
	}

	return rs
}
