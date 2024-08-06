package scheduler

import (
	"time"

	"github.com/quanghuynguyen1902/job-schedule/internal/db"
	queue "github.com/quanghuynguyen1902/job-schedule/internal/kafka"
	"github.com/quanghuynguyen1902/job-schedule/internal/logger"
)

type SchedulingService struct {
	queue        *queue.Kafka
	store        *db.Store
	repo         db.DBRepo
	logger       logger.Logger
	pollInterval time.Duration
}

func NewSchedulingService(store *db.Store, repo db.DBRepo, logger logger.Logger, queue *queue.Kafka) *SchedulingService {
	return &SchedulingService{
		store:        store,
		repo:         repo,
		logger:       logger,
		queue:        queue,
		pollInterval: time.Minute,
	}
}
