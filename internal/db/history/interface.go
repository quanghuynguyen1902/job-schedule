package history

import (
	"github.com/gocql/gocql"

	"github.com/quanghuynguyen1902/job-schedule/internal/model"
)

type IStore interface {
	CreateHistoryRecord(session *gocql.Session, history model.JobHistory) error
}
