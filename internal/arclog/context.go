package arclog

import (
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/konradgj/arclog/internal/db"
	"github.com/konradgj/arclog/internal/dpsreport"
)

var osUserConfigDir = os.UserConfigDir

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

func GetAppDirPath(appDir string) (string, error) {
	configDir, err := osUserConfigDir()
	if err != nil {
		return "", err
	}

	appDirPath := filepath.Join(configDir, appDir)
	if err := os.MkdirAll(appDirPath, 0o755); err != nil {
		return "", err
	}

	return appDirPath, nil
}

func GetDbPath(appDir string) (string, error) {
	appDirPath, err := GetAppDirPath(appDir)
	if err != nil {
		return "", err
	}

	return filepath.Join(appDirPath, "arclog.db"), nil
}
