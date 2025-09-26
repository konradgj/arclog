package arclog

import (
	"os"
	"path/filepath"
	"time"

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

func NewContext() *Context {
	return &Context{}
}

func (ctx *Context) Init(verbose bool) {
	logger.Initlogger(verbose)

	ctx.St = &db.Store{}
	dbPath := getDbPath()
	ctx.St.SetupDb(dbPath, verbose)

	ctx.Config = &Config{}
	ctx.Config.InitConfig()

	ctx.DpsReportClient = dpsreport.NewClient(5 * time.Minute)
	ctx.RateLimiter = NewRateLimiter(25, time.Minute)
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
