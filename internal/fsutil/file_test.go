package fsutil

import (
	"os"
	"path/filepath"
	"testing"
)

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
