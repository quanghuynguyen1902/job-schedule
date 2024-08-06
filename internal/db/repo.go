package db

import (
	"errors"

	"github.com/gocql/gocql"
)

// FinallyFunc function to finish a session
type FinallyFunc = func(error) error

// DBRepo interface for Cassandra operations
type DBRepo interface {
	Session() *gocql.Session
	NewBatch() (*gocql.Batch, FinallyFunc)
	SetNewSession(*gocql.Session)
}

// repo is implementation of repository
type repo struct {
	session *gocql.Session
}

// Session returns the Cassandra session
func (r *repo) Session() *gocql.Session {
	return r.session
}

// NewRepo creates a new repo instance
func NewRepo(session *gocql.Session) DBRepo {
	return &repo{session: session}
}

// NewBatch creates a new batch for Cassandra operations
func (r *repo) NewBatch() (*gocql.Batch, FinallyFunc) {
	batch := r.session.NewBatch(gocql.LoggedBatch)

	finallyFn := func(err error) error {
		if err != nil {
			// In Cassandra, we can't rollback a batch, so we just return the error
			return err
		}

		// Execute the batch
		if err := r.session.ExecuteBatch(batch); err != nil {
			return errors.New("failed to execute batch: " + err.Error())
		}

		return nil
	}

	return batch, finallyFn
}

// SetNewSession sets a new Cassandra session
func (r *repo) SetNewSession(session *gocql.Session) {
	r.session = session
}
