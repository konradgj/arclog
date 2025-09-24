package helpers

import "log/slog"

var Logger *slog.Logger

type Config struct {
	LogPath   string `toml:"logpath"`
	UserToken string `toml:"usertoken"`
}
