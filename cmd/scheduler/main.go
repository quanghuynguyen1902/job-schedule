package main

import (
	"github.com/quanghuynguyen1902/job-schedule/internal/config"
	"github.com/quanghuynguyen1902/job-schedule/internal/db"
	"github.com/quanghuynguyen1902/job-schedule/internal/kafka"
	"github.com/quanghuynguyen1902/job-schedule/internal/logger"
	"github.com/quanghuynguyen1902/job-schedule/internal/scheduler"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()
	log.Infof("Server starting")

	repo := db.NewCassandraStore(cfg)

	s := db.New()
	queue := kafka.New(cfg.Kafka.Broker)

	errCh := make(chan error)
	go func(ch chan error) {
		err := queue.RunProducer()
		if err != nil {
			errCh <- err
		}
	}(errCh)

	// Khởi tạo và chạy SchedulingService
	service := scheduler.NewSchedulingService(s, repo, log, queue)
	service.Run()
}
