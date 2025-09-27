package arclog

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConfig_InitConfig(t *testing.T) {
	tmpDir := t.TempDir()
	cfg := &Config{}

	fileUsed, err := cfg.InitConfig(tmpDir)
	if err != nil {
		t.Fatalf("InitConfig failed: %v", err)
	}

	if _, err := os.Stat(fileUsed); os.IsNotExist(err) {
		t.Fatalf("expected config file to exist: %s", fileUsed)
	}

	if cfg.LogPath == "" {
		t.Errorf("expected default LogPath to be set")
	}

	if cfg.UserToken != "" {
		t.Errorf("expected default UserToken to be empty")
	}
}

func TestConfig_SetValues(t *testing.T) {
	tmpDir := t.TempDir()
	cfg := &Config{}

	_, err := cfg.InitConfig(tmpDir)
	if err != nil {
		t.Fatalf("InitConfig failed: %v", err)
	}

	newLogPath := filepath.Join(tmpDir, "logs")
	newToken := "mytoken123"
	if err := cfg.SetValues(tmpDir, newLogPath, newToken); err != nil {
		t.Fatalf("SetValues failed: %v", err)
	}

	if cfg.LogPath != newLogPath {
		t.Errorf("expected LogPath %q, got %q", newLogPath, cfg.LogPath)
	}

	if cfg.UserToken != newToken {
		t.Errorf("expected UserToken %q, got %q", newToken, cfg.UserToken)
	}
}

func TestConfig_GetSettingsString(t *testing.T) {
	tmpDir := t.TempDir()
	cfg := &Config{}

	_, err := cfg.InitConfig(tmpDir)
	if err != nil {
		t.Fatalf("InitConfig failed: %v", err)
	}

	cfg.SetValues(tmpDir, "somepath", "sometoken")
	output := cfg.GetSettingsString()

	if !strings.Contains(output, "logpath = somepath") || !strings.Contains(output, "usertoken = sometoken") {
		t.Errorf("expected settings string to contain keys, got: %s", output)
	}
}
