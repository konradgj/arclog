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
	Log    *logger.Logger
	Config *Config
}

func NewContext() *Context {
	log := &logger.Logger{}
	log.InitLogger()

	store := &db.Store{}
	dbPath := getDbPath(log)
	store.SetupDb(dbPath, false, log)

	ctx := &Context{
		St:     store,
		Log:    log,
		Config: &Config{},
	}
	ctx.Config.InitConfig(log)

	return ctx
}

func GetAppDir(l *logger.Logger) string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		l.Error("Could not get user config dir", "err", err)
		os.Exit(1)
	}

	appDirAbs := filepath.Join(configDir, appDir)
	if err := os.MkdirAll(appDirAbs, 0o755); err != nil {
		l.Error("Could not create config dir", "err", err)
		os.Exit(1)
	}

	return appDirAbs
}

func getDbPath(l *logger.Logger) string {
	appDir := GetAppDir(l)
	return filepath.Join(appDir, "arclog.db")
}
