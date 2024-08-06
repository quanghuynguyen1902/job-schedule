package job

import (
	"errors"
	"fmt"
	"github.com/gocql/gocql"

	"github.com/quanghuynguyen1902/job-schedule/internal/model"
)

type store struct{}

func New() IStore {
	return &store{}
}

func (s store) Create(session *gocql.Session, job *model.Job) error {
	// Insert job into database
	query := `
		INSERT INTO jobs (user_id, job_id, retry_times, created, interval, is_recurring)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	if err := session.Query(query,
		job.UserID,
		job.JobID,
		job.RetryTimes,
		job.Created,
		job.Interval,
		job.IsRecurring,
	).Exec(); err != nil {
		return fmt.Errorf("failed to insert job: %w", err)
	}

	return nil
}

func (s store) GetJob(session *gocql.Session, jobId string) (*model.Job, error) {
	var job model.Job
	query := "SELECT user_id, job_id, retry_times, created, interval, is_recurring FROM jobs WHERE job_id = ?"
	iter := session.Query(query, jobId).Iter()

	if !iter.Scan(
		&job.UserID,
		&job.JobID,
		&job.RetryTimes,
		&job.Created,
		&job.Interval,
		&job.IsRecurring,
	) {
		return nil, errors.New("job not found")
	}

	return &job, nil
}
