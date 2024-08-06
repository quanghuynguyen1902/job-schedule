package main

import (
	"fmt"

	"github.com/quanghuynguyen1902/job-schedule/internal/config"
	"github.com/quanghuynguyen1902/job-schedule/internal/db"
	"github.com/quanghuynguyen1902/job-schedule/internal/logger"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()
	log.Infof("Server starting")

	repo := db.NewCassandraStore(cfg)

	if err := db.ApplyMigrations(repo.Session(), "internal/migrations"); err != nil {
		log.Errorf(err, "failed to apply migrations")
	}

	fmt.Println("Migrations completed successfully")
}
