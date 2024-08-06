package job

import (
	"github.com/gocql/gocql"
	"github.com/quanghuynguyen1902/job-schedule/internal/model"
)

type IStore interface {
	Create(session *gocql.Session, job *model.Job) error
	GetJob(session *gocql.Session, jobId string) (*model.Job, error)
}
