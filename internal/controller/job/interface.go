package job

import (
	"github.com/quanghuynguyen1902/job-schedule/internal/model"
)

type IController interface {
	CreateJob(job *model.Job) (*model.Job, error)
}
