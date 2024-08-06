package job

import (
	"time"

	"github.com/gocql/gocql"

	"github.com/quanghuynguyen1902/job-schedule/internal/model"
	timeUtils "github.com/quanghuynguyen1902/job-schedule/internal/utils"
)

func (c *controller) CreateJob(job *model.Job) (*model.Job, error) {
	// Set default values if not provided
	if job.Created.IsZero() {
		job.Created = time.Now()
	}

	if job.JobID == "" {
		job.JobID = gocql.TimeUUID().String()
	}

	// Parse the execution interval
	duration, err := timeUtils.ParseISO8601Duration(job.Interval)
	if err != nil {
		return nil, err
	}

	// Calculate the next execution time
	now := time.Now()
	nextExecutionTime := now.Add(duration)

	// Convert to UNIX timestamp with minute-level granularity
	nextExecutionTimeUnix := nextExecutionTime.Unix()

	// store to job schedule
	err = c.store.Schedule.UpsertScheduleRecord(c.repo.Session(), model.JobSchedule{
		JobID:             job.JobID,
		NextExecutionTime: nextExecutionTimeUnix,
	})
	if err != nil {
		c.logger.Errorf(err, "failed to upsert schedule record")
		return nil, err
	}

	// create job
	err = c.store.Job.Create(c.repo.Session(), job)
	if err != nil {
		c.logger.Errorf(err, "failed to create job")
		return nil, err
	}

	return job, nil
}
