package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

type Logger struct {
	*slog.Logger
}

const appDir = "arclog"

func (l *Logger) InitLogger() {
	logPath := getAppLogPath()

	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	level := slog.LevelInfo

	multiWriter := io.MultiWriter(os.Stderr, f)

	l.Logger = slog.New(slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
		Level: level,
	}))
}

// Convenience functions
func (l *Logger) Info(msg string, args ...any)  { l.Logger.Info(msg, args...) }
func (l *Logger) Debug(msg string, args ...any) { l.Logger.Debug(msg, args...) }
func (l *Logger) Error(msg string, args ...any) { l.Logger.Error(msg, args...) }
func (l *Logger) Warn(msg string, args ...any)  { l.Logger.Warn(msg, args...) }

func getAppLogPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	appDirAbs := filepath.Join(configDir, appDir)
	if err := os.MkdirAll(appDirAbs, 0o755); err != nil {
		log.Fatalf("Could not create config dir: %v", err)
		os.Exit(1)
	}

	return filepath.Join(appDirAbs, "arclog.log")
}
