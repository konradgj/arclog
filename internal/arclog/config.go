package arclog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	LogPath   string `toml:"logpath"`
	UserToken string `toml:"usertoken"`
}

func (cfg *Config) InitConfig(appDir string) (string, error) {
	configFile, err := GetConfigFilePath(appDir)
	if err != nil {
		return "", fmt.Errorf("could not get config file path: %w", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home dir: %w", err)
	}

	if cfg.LogPath == "" {
		cfg.LogPath = filepath.Join(home, "Documents", "Guild Wars 2", "addons", "arcdps.cbtlogs")
	}
	if cfg.UserToken == "" {
		cfg.UserToken = ""
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		if err := cfg.Save(configFile); err != nil {
			return "", fmt.Errorf("could not create new config file: %w", err)
		}
		return configFile, nil
	} else if err != nil {
		return "", fmt.Errorf("could not stat config file: %w", err)
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return "", fmt.Errorf("could not read config file: %w", err)
	}

	if err := toml.Unmarshal(data, cfg); err != nil {
		return "", fmt.Errorf("could not decode config file: %w", err)
	}

	return configFile, nil
}

func (cfg *Config) Save(path string) error {
	data, err := toml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("could not marshal config: %w", err)
	}

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}

	return nil
}

func (cfg *Config) SetValues(appDir, path, userToken string) error {
	if path != "" {
		cfg.LogPath = path
	}
	if userToken != "" {
		cfg.UserToken = userToken
	}

	configFile, err := GetConfigFilePath(appDir)
	if err != nil {
		return err
	}

	return cfg.Save(configFile)
}

func (cfg *Config) AbortIfNoUserToken() {
	if cfg.UserToken != "" {
		return
	}

	const msg = `
	Missing UserToken!
	Get your UserToken at: https://dps.report/getUserToken
	Set your token with: arclog config set -t <your_token>
	`

	fmt.Fprint(os.Stderr, msg)
	os.Exit(1)
}

func (cfg *Config) GetSettingsString() string {
	var sb strings.Builder

	type field struct {
		name  string
		value any
	}

	fields := []field{
		{"LogPath", cfg.LogPath},
		{"UserToken", cfg.UserToken},
	}

	for _, f := range fields {
		fmt.Fprintf(&sb, "  - %s = %v\n", f.name, f.value)
	}

	return sb.String()
}

func GetConfigFilePath(appDir string) (string, error) {
	appDirPath, err := GetAppDirPath(appDir)
	if err != nil {
		return "", fmt.Errorf("could not get app dir path: %w", err)
	}

	return filepath.Join(appDirPath, "config.toml"), nil
}
