package arclog

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetAppDirPath(t *testing.T) {
	tmp := t.TempDir()
	dirName := "testapp"

	old := osUserConfigDir
	defer func() { osUserConfigDir = old }()
	osUserConfigDir = func() (string, error) { return tmp, nil }

	appDirPath, err := GetAppDirPath(dirName)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := filepath.Join(tmp, dirName)
	if appDirPath != expected {
		t.Errorf("expected %q, got %q", expected, appDirPath)
	}

	info, err := os.Stat(appDirPath)
	if err != nil {
		t.Fatalf("could not stat path: %v", err)
	}
	if !info.IsDir() {
		t.Errorf("expected %q to be a directory", appDirPath)
	}
}

func TestGetDbPath(t *testing.T) {
	tmp := t.TempDir()
	dirName := "testapp"

	old := osUserConfigDir
	defer func() { osUserConfigDir = old }()
	osUserConfigDir = func() (string, error) { return tmp, nil }

	dbPath, err := GetDbPath(dirName)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := filepath.Join(tmp, dirName, "arclog.db")
	if dbPath != expected {
		t.Errorf("expected %q, got %q", expected, dbPath)
	}
}
