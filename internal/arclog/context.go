package arclog

import (
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/konradgj/arclog/internal/db"
	"github.com/konradgj/arclog/internal/dpsreport"
	"github.com/konradgj/arclog/internal/logger"
)

const appDir = "arclog"

type Context struct {
	St              *db.Store
	Config          *Config
	Watcher         *fsnotify.Watcher
	DpsReportClient *dpsreport.Client
	RateLimiter     *RateLimiter
}

func NewContext(st *db.Store, cfg *Config, dpsClient *dpsreport.Client, rl *RateLimiter) *Context {
	return &Context{
		St:              st,
		Config:          cfg,
		DpsReportClient: dpsClient,
		RateLimiter:     rl,
	}
}

func GetAppDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		logger.Error("Could not get user config dir", "err", err)
		os.Exit(1)
	}

	appDirAbs := filepath.Join(configDir, appDir)
	if err := os.MkdirAll(appDirAbs, 0o755); err != nil {
		logger.Error("Could not create config dir", "err", err)
		os.Exit(1)
	}

	return appDirAbs
}

func GetDbPath() string {
	appDir := GetAppDir()
	return filepath.Join(appDir, "arclog.db")
}
