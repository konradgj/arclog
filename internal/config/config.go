package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	LogPath   string `json:"logpath"`
	UserToken string `json:"usertoken"`
}

const configFileName = ".arclog_config.json"

func LoadConfig() (Config, error) {
	var cfg Config
	defaultLogPath := filepath.Join("Documents", "Guild Wars 2", "addons", "arcdps.cbtlogs")

	homePath, err := os.UserHomeDir()
	if err != nil {
		return cfg, fmt.Errorf("could not get homepath: %w", err)
	}

	data, err := cfg.readOrCreate(filepath.Join(homePath, defaultLogPath))
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("error parsing config: %w", err)
	}

	return cfg, nil
}

func (cfg *Config) SetLogPath(logPath string) error {
	if fi, err := os.Stat(logPath); err != nil || !fi.IsDir() {
		return fmt.Errorf("invalid log path: %w", err)
	}

	cfg.LogPath = logPath
	if err := cfg.write(); err != nil {
		return fmt.Errorf("could not set log path: %w", err)
	}
	return nil
}

func (cfg *Config) SetUserToken(userToken string) error {
	cfg.UserToken = userToken

	if err := cfg.write(); err != nil {
		return fmt.Errorf("could not set usertoken: %w", err)
	}
	return nil
}

func (cfg *Config) readOrCreate(logPath string) ([]byte, error) {
	var data []byte

	path, err := getConfigPath()
	if err != nil {
		return data, err
	}

	data, err = os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			cfg.LogPath = logPath
			cfg.UserToken = ""

			if err := cfg.write(); err != nil {
				return data, fmt.Errorf("could not create default config: %w", err)
			}

			data, _ = json.Marshal(cfg)
			return data, nil
		}

		return data, fmt.Errorf("could not read config file: %w", err)
	}

	return data, nil
}

func (cfg *Config) write() error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	if err = os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("error writing config: %w", err)
	}
	return nil
}

func getConfigPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get home path: %w", err)
	}

	return filepath.Join(homePath, configFileName), nil
}
