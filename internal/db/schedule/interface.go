package schedule

import (
	"github.com/gocql/gocql"
	"github.com/quanghuynguyen1902/job-schedule/internal/model"
)

type IStore interface {
	UpsertScheduleRecord(session *gocql.Session, schedule model.JobSchedule) error
	GetSchedules(session *gocql.Session, timeUnix int64) ([]model.JobSchedule, error)
}
