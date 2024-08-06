package controller

import (
	"github.com/quanghuynguyen1902/job-schedule/internal/config"
	"github.com/quanghuynguyen1902/job-schedule/internal/controller/job"
	"github.com/quanghuynguyen1902/job-schedule/internal/db"
	"github.com/quanghuynguyen1902/job-schedule/internal/logger"
)

type Controller struct {
	Job job.IController
}

func New(store *db.Store, repo db.DBRepo, logger logger.Logger, cfg *config.Config) *Controller {
	return &Controller{
		Job: job.New(store, repo, logger, cfg),
	}
}
