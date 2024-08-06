package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/quanghuynguyen1902/job-schedule/internal/config"
	"github.com/quanghuynguyen1902/job-schedule/internal/db"
	"github.com/quanghuynguyen1902/job-schedule/internal/logger"
	"github.com/quanghuynguyen1902/job-schedule/internal/routes"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()
	log.Infof("Server starting")

	s := db.New()
	repo := db.NewCassandraStore(cfg)

	_, cancel := context.WithCancel(context.Background())

	router := routes.NewRoutes(cfg, s, repo, log)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "10100"),
		Handler: router,
	}

	// serve http server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err, "failed to listen and serve")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	cancel()

	shutdownServer(srv, log)
}

func shutdownServer(srv *http.Server, l logger.Logger) {
	l.Info("Server Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		l.Error(err, "failed to shutdown server")
	}

	l.Info("Server Exit")
}
