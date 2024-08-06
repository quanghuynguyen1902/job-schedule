package db

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

type Migration struct {
	Version   string
	Content   string
	AppliedAt time.Time
}

func ApplyMigrations(session *gocql.Session, migrationsDir string) error {
	// Ensure migrations table exists
	if err := createMigrationsTable(session); err != nil {
		return err
	}

	// Get applied migrations
	appliedMigrations, err := getAppliedMigrations(session)
	if err != nil {
		return err
	}

	// Load migration files
	migrations, err := loadMigrationFiles(migrationsDir)
	if err != nil {
		return err
	}

	// Apply new migrations
	for _, migration := range migrations {
		if !appliedMigrations[migration.Version] {
			if err := applyMigration(session, migration); err != nil {
				return err
			}
			fmt.Printf("Applied migration: %s\n", migration.Version)
		}
	}

	return nil
}

func createMigrationsTable(session *gocql.Session) error {
	query := `
    CREATE TABLE IF NOT EXISTS migrations (
        version text PRIMARY KEY,
        applied_at timestamp
    )`
	return session.Query(query).Exec()
}

func getAppliedMigrations(session *gocql.Session) (map[string]bool, error) {
	applied := make(map[string]bool)
	iter := session.Query("SELECT version FROM migrations").Iter()
	var version string
	for iter.Scan(&version) {
		applied[version] = true
	}
	return applied, iter.Close()
}

func loadMigrationFiles(dir string) ([]Migration, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var migrations []Migration
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".cql" {
			content, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			version := strings.TrimSuffix(file.Name(), ".cql")

			migrations = append(migrations, Migration{
				Version: version,
				Content: string(content),
			})
		}
	}
	return migrations, nil
}

func applyMigration(session *gocql.Session, migration Migration) error {
	// Execute migration content
	statements := strings.Split(migration.Content, ";")
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		if err := session.Query(stmt).Exec(); err != nil {
			return fmt.Errorf("error executing migration %s: %v", migration.Version, err)
		}
	}

	// Record migration as applied
	query := "INSERT INTO migrations (version, applied_at) VALUES (?, ?)"
	return session.Query(query, migration.Version, time.Now()).Exec()
}
