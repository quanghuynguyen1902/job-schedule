package db

import (
	"github.com/quanghuynguyen1902/job-schedule/internal/db/history"
	"github.com/quanghuynguyen1902/job-schedule/internal/db/job"
	"github.com/quanghuynguyen1902/job-schedule/internal/db/schedule"
)

type Store struct {
	Job      job.IStore
	History  history.IStore
	Schedule schedule.IStore
}

func New() *Store {
	return &Store{
		Job:      job.New(),
		History:  history.New(),
		Schedule: schedule.New(),
	}
}
