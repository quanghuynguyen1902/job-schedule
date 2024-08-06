package history

import (
	"github.com/gocql/gocql"

	"github.com/quanghuynguyen1902/job-schedule/internal/model"
)

type store struct{}

func New() IStore {
	return &store{}
}

func (s store) CreateHistoryRecord(session *gocql.Session, history model.JobHistory) error {
	query := "INSERT INTO job_history (job_id, execution_id, status, worker_id, retry_cnt) VALUES (?, ?, ?, ?, ?)"
	return session.Query(query, history.JobID, history.ExecutionTime, history.Status, history.RetryCount, history.LastUpdateTime).Exec()
}
