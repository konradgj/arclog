package helpers

import (
	"os"
	"path/filepath"

	"github.com/konradgj/arclog/internal/logger"
)

const appDir = "arclog"

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
