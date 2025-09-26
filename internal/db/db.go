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
	Queries *database.Queries
	DB      *sql.DB
}

func (s *Store) SetupDb(dbPath string, verbose bool) {
	db, err := initDb(dbPath)
	if err != nil {
		logger.Error("could not init db", "err", err)
		os.Exit(1)
	}

	if !verbose {
		goose.SetLogger(goose.NopLogger())
	}

	if err := migrateDb(db); err != nil {
		logger.Error("could not migrate db", "err", err)
	}

	s.DB = db
	s.Queries = database.New(db)
}

func WrapNullStr(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
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
