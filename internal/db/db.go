package db

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/konradgj/arclog/internal/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

//go:embed schema/*.sql
var embedMigrations embed.FS

type Store struct {
	Queries *database.Queries
	DB      *sql.DB
}

func (s *Store) SetupDb(dbPath string, debug bool) error {
	db, err := initDb(dbPath)
	if err != nil {
		return fmt.Errorf("could not init db: %w", err)
	}

	if !debug {
		goose.SetLogger(goose.NopLogger())
	}

	if err := migrateDb(db); err != nil {
		return fmt.Errorf("could not migrate db: %w", err)
	}

	s.DB = db
	s.Queries = database.New(db)
	return nil
}

func WrapNullStr(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func PrintNullStr(s sql.NullString) string {
	if !s.Valid {
		return "-"
	}
	return s.String
}

func PrintNullBool(s sql.NullInt64) *int64 {
	if !s.Valid {
		return nil
	}
	return &s.Int64
}

func WrapNullBool(b *bool) sql.NullInt64 {
	if b == nil {
		return sql.NullInt64{Valid: false}
	}
	if *b {
		return sql.NullInt64{Int64: 1, Valid: true}
	}
	return sql.NullInt64{Int64: 0, Valid: true}
}

func migrateDb(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("could not set sql dialect: %w", err)
	}

	if err := goose.Up(db, "schema"); err != nil {
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
