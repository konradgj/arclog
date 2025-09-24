package appconfig

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/konradgj/arclog/internal/logger"
	"github.com/spf13/viper"
)

func InitConfig() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		logger.Error("Could not get user config dir", "err", err)
		os.Exit(1)
	}

	appDir := filepath.Join(configDir, "arclog")
	if err := os.MkdirAll(appDir, 0o755); err != nil {
		logger.Error("Could not create config dir", "err", err)
		os.Exit(1)
	}

	configFile := filepath.Join(appDir, "config.toml")
	viper.AddConfigPath(appDir)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	home, err := os.UserHomeDir()
	if err != nil {
		logger.Error("Could not get user home dir", "err", err)
		os.Exit(1)
	}

	defaultLogPath := filepath.Join("Documents", "Guild Wars 2", "addons", "arcdps.cbtlogs")
	viper.SetDefault("LogPath", filepath.Join(home, defaultLogPath))
	viper.SetDefault("UserToken", "")

	if err := viper.ReadInConfig(); err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err = viper.SafeWriteConfig(); err != nil {
				logger.Error("Could not create new config file", "err", err)
				os.Exit(1)
			}

			viper.SetConfigFile(configFile)
			logger.Debug("Created new config file", "path", viper.ConfigFileUsed())
		} else {
			logger.Error("Could not read config", "err", err)
			os.Exit(1)
		}
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
