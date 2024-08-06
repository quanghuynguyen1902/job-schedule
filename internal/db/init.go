package db

import (
	"time"

	"github.com/gocql/gocql"

	"github.com/quanghuynguyen1902/job-schedule/internal/config"
	"github.com/quanghuynguyen1902/job-schedule/internal/logger"
)

// NewCassandraStore initializes a Cassandra connection
func NewCassandraStore(cfg *config.Config) DBRepo {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "job_scheduler"
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.Cassandra.User,
		Password: cfg.Cassandra.Pass,
	}

	session, err := cluster.CreateSession()
	if err != nil {
		logger.L.Fatalf(err, "failed to create Cassandra session")
	}

	logger.L.Info("Cassandra database connected")

	return &repo{session: session}
}
