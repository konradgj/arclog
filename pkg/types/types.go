package types

type Config struct {
	LogPath   string `toml:"logpath"`
	UserToken string `toml:"usertoken"`
}
