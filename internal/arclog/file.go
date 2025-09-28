package arclog

import (
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/konradgj/arclog/internal/database"
	"github.com/konradgj/arclog/internal/db"
)

func (cfg *Config) GetLogNameAndRelativePath(path string) (string, sql.NullString, error) {
	fileName := filepath.Base(path)

	relPath, err := filepath.Rel(cfg.LogPath, path)
	if err != nil {
		return "", sql.NullString{}, fmt.Errorf("could not get relative path: %w", err)
	}

	dir := filepath.Dir(relPath)

	var relativePath sql.NullString
	if dir != "." {
		relativePath = db.WrapNullStr(dir)
	} else {
		relativePath = sql.NullString{Valid: false}
	}

	return fileName, relativePath, nil
}

func (cfg *Config) GetLogFilePath(cbtlog database.Cbtlog) string {
	return filepath.Join(cfg.LogPath, cbtlog.RelativePath.String, cbtlog.Filename)
}
