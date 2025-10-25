package fsutil

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func GetAllFilePaths(path string) ([]string, error) {
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("log path does not exist: %s", path)
		}
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
			return errWalk
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

// File or dir
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("could not stat file %q: %w", path, err)
}

// File or dir
func RmFile(path string) error {
	exists, err := FileExists(path)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("file does not exist: %s", path)
	}

	err = os.Remove(path)
	if err != nil {
		return fmt.Errorf("could not delete file %q: %w", path, err)
	}

	return nil
}
