package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

var logger *slog.Logger

const appDir = "arclog"

func Init(verbose bool) {
	logPath := getAppLogPath()

	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}

	multiWriter := io.MultiWriter(os.Stderr, f)

	logger = slog.New(slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
		Level: level,
	}))
}

// Convenience functions
func Info(msg string, args ...any)  { logger.Info(msg, args...) }
func Debug(msg string, args ...any) { logger.Debug(msg, args...) }
func Error(msg string, args ...any) { logger.Error(msg, args...) }
func Warn(msg string, args ...any)  { logger.Warn(msg, args...) }

func getAppLogPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	appDirAbs := filepath.Join(configDir, appDir)
	os.MkdirAll(appDirAbs, 0o755)

	return filepath.Join(appDirAbs, "arclog.log")
}
