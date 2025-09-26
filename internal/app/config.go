package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/konradgj/arclog/internal/logger"
	"github.com/spf13/viper"
)

type Config struct {
	LogPath   string `toml:"logpath"`
	UserToken string `toml:"usertoken"`
}

func (cfg *Config) InitConfig() {
	appDir := GetAppDir()

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

	err = cfg.Unmarshal()
	if err != nil {
		logger.Error("error unmarshaling config", "err", err)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}

func (cfg *Config) Unmarshal() error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	return nil
}

func (cfg *Config) SetValues(logPath, userToken string) {
	if logPath != "" {
		viper.Set("LogPath", logPath)
	}
	if userToken != "" {
		viper.Set("UserToken", userToken)
	}

	err := viper.WriteConfig()
	if err != nil {
		logger.Error("could not write config", "error", err)
	}

	err = cfg.Unmarshal()
	if err != nil {
		logger.Error("error unmarshaling config", "err", err)
	}
}

func (cfg *Config) Show() {
	settings := viper.AllSettings()
	printSettings(settings, 0)
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
