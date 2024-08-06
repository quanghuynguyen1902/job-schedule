package job

import (
	"github.com/quanghuynguyen1902/job-schedule/internal/config"
	"github.com/quanghuynguyen1902/job-schedule/internal/db"
	"github.com/quanghuynguyen1902/job-schedule/internal/logger"
)

type controller struct {
	store  *db.Store
	repo   db.DBRepo
	logger logger.Logger
	config *config.Config
}

func New(store *db.Store, repo db.DBRepo, logger logger.Logger, cfg *config.Config) IController {
	return &controller{
		store:  store,
		repo:   repo,
		logger: logger,
		config: cfg,
	}
}
