package appconfig

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/konradgj/arclog/internal/logger"
	"github.com/spf13/viper"
)

const appDir = "arclog"

type AppConfig struct {
	LogPath   string `toml:"logpath"`
	UserToken string `toml:"usertoken"`
}

func InitConfig() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		logger.Error("Could not get user config dir", "err", err)
		os.Exit(1)
	}

	appDirAbs := filepath.Join(configDir, appDir)
	if err := os.MkdirAll(appDir, 0o755); err != nil {
		logger.Error("Could not create config dir", "err", err)
		os.Exit(1)
	}

	configFile := filepath.Join(appDirAbs, "config.toml")
	viper.AddConfigPath(appDirAbs)
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

func Unmarshal() (*AppConfig, error) {
	var cfg AppConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		return &cfg, err
	}

	return &cfg, nil
}

func printSettings(m map[string]any, indent int) {
	prefix := ""
	for range indent {
		prefix += "  " // two spaces per level
	}

	for k, v := range m {
		switch val := v.(type) {
		case map[string]any: // nested table
			fmt.Printf("%s- %s:\n", prefix, k)
			printSettings(val, indent+1)
		default:
			fmt.Printf("%s- %s = %v\n", prefix, k, val)
		}
	}
}

func Show() {
	settings := viper.AllSettings()
	printSettings(settings, 0)
}
