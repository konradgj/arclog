package app

import (
	"os"
	"path/filepath"

	"github.com/konradgj/arclog/internal/db"
	"github.com/konradgj/arclog/internal/logger"
)

const appDir = "arclog"

type Context struct {
	St     *db.Store
	Config *Config
}

func NewContext(verbose bool) *Context {
	logger.Initlogger(verbose)

	store := &db.Store{}
	dbPath := getDbPath()
	store.SetupDb(dbPath, false)

	ctx := &Context{
		St:     store,
		Config: &Config{},
	}
	ctx.Config.InitConfig()

	return ctx
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

func getDbPath() string {
	appDir := GetAppDir()
	return filepath.Join(appDir, "arclog.db")
}
