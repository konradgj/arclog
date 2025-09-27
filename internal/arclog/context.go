package arclog

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/konradgj/arclog/internal/db"
	"github.com/konradgj/arclog/internal/dpsreport"
	"go.uber.org/zap"
)

var osUserConfigDir = os.UserConfigDir

type Context struct {
	St              *db.Store
	Config          *Config
	Logger          *zap.SugaredLogger
	Watcher         *fsnotify.Watcher
	DpsReportClient *dpsreport.Client
	RateLimiter     *RateLimiter
}

func NewContext(st *db.Store, cfg *Config, logger *zap.SugaredLogger, dpsClient *dpsreport.Client, rl *RateLimiter) *Context {
	return &Context{
		St:              st,
		Config:          cfg,
		Logger:          logger,
		DpsReportClient: dpsClient,
		RateLimiter:     rl,
	}
}

func InitContext(logger *zap.SugaredLogger, appDir string, debug bool) (*Context, error) {
	st := &db.Store{}
	dbPath, err := GetDbPath(appDir)
	if err != nil {
		return nil, fmt.Errorf("could not get db path: %w", err)
	}

	err = st.SetupDb(dbPath, debug)
	if err != nil {
		return nil, fmt.Errorf("could not setup db: %w", err)
	}

	cfg := &Config{}
	fileUsed, err := cfg.InitConfig(appDir)
	if err != nil {
		return nil, fmt.Errorf("could not initialize config: %w", err)
	}
	fmt.Printf("Using config file: %s\n", fileUsed)

	dpsClient := dpsreport.NewClient(5 * time.Minute)
	rl := NewRateLimiter(25, time.Minute)

	return NewContext(st, cfg, logger, dpsClient, rl), nil
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
