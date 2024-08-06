package schedule

import (
	"github.com/gocql/gocql"

	"github.com/quanghuynguyen1902/job-schedule/internal/model"
)

type store struct{}

func New() IStore {
	return &store{}
}

func (s store) UpsertScheduleRecord(session *gocql.Session, schedule model.JobSchedule) error {
	query := "INSERT INTO schedules (job_id, next_execution_time) VALUES (?, ?)"
	return session.Query(query, schedule.JobID, schedule.NextExecutionTime).Exec()
}

func (s store) GetSchedules(session *gocql.Session, timeUnix int64) ([]model.JobSchedule, error) {
	query := "SELECT job_id, next_execution_time FROM schedules WHERE next_execution_time <= ? ALLOW FILTERING"
	iter := session.Query(query, timeUnix).Iter()

	var tasks []model.JobSchedule
	var task model.JobSchedule
	for iter.Scan(&task.JobID, &task.NextExecutionTime) {
		tasks = append(tasks, task)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return tasks, nil
}
