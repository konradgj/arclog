package arclog

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	"github.com/konradgj/arclog/internal/database"
)

type MockCbtlog struct {
	database.Cbtlog
}

func TestGetLogNameAndRelativePath(t *testing.T) {
	tmpDir := t.TempDir()

	subDir := filepath.Join(tmpDir, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatal(err)
	}
	filePath := filepath.Join(subDir, "test.zevtc")
	if err := os.WriteFile(filePath, []byte("dummy"), 0644); err != nil {
		t.Fatal(err)
	}

	cfg := &Config{LogPath: tmpDir}

	name, relPath, err := cfg.GetLogNameAndRelativePath(filePath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if name != "test.zevtc" {
		t.Errorf("expected filename test.zevtc, got %s", name)
	}

	if !relPath.Valid || relPath.String != "sub" {
		t.Errorf("expected relative path 'sub', got %+v", relPath)
	}
}

func TestGetLogFilePath(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := &Config{LogPath: tmpDir}
	cbtlog := MockCbtlog{
		Cbtlog: database.Cbtlog{
			RelativePath: sql.NullString{String: "sub", Valid: true},
			Filename:     "test.zevtc",
		},
	}

	got := cfg.GetLogFilePath(cbtlog.Cbtlog)
	want := filepath.Join(tmpDir, "sub", "test.zevtc")

	if got != want {
		t.Errorf("expected %s, got %s", want, got)
	}
}

func TestGetAllFilePaths(t *testing.T) {
	tmpDir := t.TempDir()

	validFile := filepath.Join(tmpDir, "combat.zevtc")
	invalidFile := filepath.Join(tmpDir, "ignore.txt")

	if err := os.WriteFile(validFile, []byte("ok"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(invalidFile, []byte("no"), 0644); err != nil {
		t.Fatal(err)
	}

	// Case 1: Single file path (.zevtc)
	files, err := GetAllFilePaths(validFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(files) != 1 || files[0] != validFile {
		t.Errorf("expected [%s], got %v", validFile, files)
	}

	// Case 2: Single file path (not .zevtc)
	files, err = GetAllFilePaths(invalidFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(files) != 0 {
		t.Errorf("expected empty slice, got %v", files)
	}

	// Case 3: Directory containing mixed files
	files, err = GetAllFilePaths(tmpDir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(files) != 1 || files[0] != validFile {
		t.Errorf("expected [%s], got %v", validFile, files)
	}
}
