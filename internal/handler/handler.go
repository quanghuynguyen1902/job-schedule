package handler

import (
	"github.com/quanghuynguyen1902/job-schedule/internal/config"
	"github.com/quanghuynguyen1902/job-schedule/internal/controller"
	"github.com/quanghuynguyen1902/job-schedule/internal/db"
	"github.com/quanghuynguyen1902/job-schedule/internal/handler/job"
	"github.com/quanghuynguyen1902/job-schedule/internal/logger"
)

type Handler struct {
	Job job.IHandler
}

func New(store *db.Store, ctrl *controller.Controller, logger logger.Logger, cfg *config.Config) *Handler {
	return &Handler{
		Job: job.New(store, ctrl, logger, cfg),
	}
}
