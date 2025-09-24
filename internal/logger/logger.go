package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func Init(verbose bool) {
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))
}

// Convenience functions
func Info(msg string, args ...any)  { logger.Info(msg, args...) }
func Debug(msg string, args ...any) { logger.Debug(msg, args...) }
func Error(msg string, args ...any) { logger.Error(msg, args...) }
func Warn(msg string, args ...any)  { logger.Warn(msg, args...) }
