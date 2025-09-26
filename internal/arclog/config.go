package arclog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	LogPath   string `toml:"logpath"`
	UserToken string `toml:"usertoken"`
}

func (cfg *Config) InitConfig(appDir string) (string, error) {
	appDirPath, err := GetAppDirPath(appDir)
	if err != nil {
		return "", fmt.Errorf("could not get app dir path: %w", err)
	}
	configFile := filepath.Join(appDirPath, "config.toml")

	viper.AddConfigPath(appDirPath)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home dir: %w", err)
	}

	defaultLogPath := filepath.Join("Documents", "Guild Wars 2", "addons", "arcdps.cbtlogs")
	viper.SetDefault("LogPath", filepath.Join(home, defaultLogPath))
	viper.SetDefault("UserToken", "")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.SafeWriteConfig(); err != nil {
				return "", fmt.Errorf("could not create new config file: %w", err)
			}
			viper.SetConfigFile(configFile)
		} else {
			return "", fmt.Errorf("could not read config: %w", err)
		}
	}

	if err := cfg.Unmarshal(); err != nil {
		return "", fmt.Errorf("error unmarshaling config: %w", err)
	}

	return viper.ConfigFileUsed(), nil
}

func (cfg *Config) Unmarshal() error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	return nil
}

func (cfg *Config) SetValues(logPath, userToken string) error {
	if logPath != "" {
		viper.Set("LogPath", logPath)
	}
	if userToken != "" {
		viper.Set("UserToken", userToken)
	}

	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("could not write config: %w", err)
	}

	err = cfg.Unmarshal()
	if err != nil {
		return fmt.Errorf("could not unmarshal config after write: %w", err)
	}

	return nil
}

func (cfg *Config) GetSettingsString() string {
	settings := viper.AllSettings()
	var sb strings.Builder
	genSettingString(settings, 0, &sb)
	return sb.String()
}

func genSettingString(m map[string]any, indent int, sb *strings.Builder) {
	prefix := strings.Repeat("  ", indent)

	for k, v := range m {
		switch val := v.(type) {
		case map[string]any:
			fmt.Fprintf(sb, "%s- %s:\n", prefix, k)
			genSettingString(val, indent+1, sb)
		default:
			fmt.Fprintf(sb, "%s- %s = %v\n", prefix, k, val)
		}
	}
}
