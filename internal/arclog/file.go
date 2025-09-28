package arclog

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

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

func GetAllFilePaths(path string) ([]string, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("could not get path stats: %w", err)
	}

	if !stat.IsDir() {
		if strings.HasSuffix(path, ".zevtc") {
			return []string{path}, nil
		}
		return nil, nil
	}

	var filePaths []string
	err = filepath.WalkDir(path, func(p string, d fs.DirEntry, errWalk error) error {
		if errWalk != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(p, ".zevtc") {
			filePaths = append(filePaths, p)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not walk dir: %w", err)
	}

	return filePaths, nil
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("could not stat file: %w", err)
}
