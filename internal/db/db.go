package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/konradgj/arclog/internal/database"
	"github.com/konradgj/arclog/internal/logger"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

type Store struct {
	*database.Queries
	db *sql.DB
}

func (s *Store) SetupDb(dbPath string, verbose bool, l *logger.Logger) {
	db, err := initDb(dbPath)
	if err != nil {
		l.Error("could not init db", "err", err)
		os.Exit(1)
	}

	if !verbose {
		goose.SetLogger(goose.NopLogger())
	}

	if err := migrateDb(db); err != nil {
		l.Error("could not migrate db", "err", err)
	}

	s.db = db
	s.Queries = database.New(db)
}

func migrateDb(db *sql.DB) error {
	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("could not set sql dialect: %w", err)
	}

	if err := goose.Up(db, "sql/schema"); err != nil {
		return fmt.Errorf("could not migrate db: %w", err)
	}

	return nil
}

func initDb(dbPath string) (*sql.DB, error) {
	var db *sql.DB

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return db, fmt.Errorf("could not open db: %w", err)
	}

	return db, nil
}
