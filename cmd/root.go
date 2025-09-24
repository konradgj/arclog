/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/konradgj/arclog/cmd/config"
	"github.com/konradgj/arclog/pkg/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	verbose bool
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "arclog",
		Short: "Upload arc-dps log sto dps.report",
		Long:  `arclog is a CLI tool for uploading ardps logs to dps.report as they are generated.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initLogger)
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Use config file (default is $HOME/.arclog.toml)")

	rootCmd.AddCommand(config.ConfigCmd)
}

func initLogger() {
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}

	helpers.Logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		helpers.Logger.Error("Could not get home directory", "err", err)
		os.Exit(1)
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".arclog")

	}

	viper.AutomaticEnv()

	defaultLogPath := filepath.Join("Documents", "Guild Wars 2", "addons", "arcdps.cbtlogs")
	viper.SetDefault("LogPath", filepath.Join(home, defaultLogPath))
	viper.SetDefault("UserToken", "")

	if err := viper.ReadInConfig(); err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err = viper.SafeWriteConfig(); err != nil {
				helpers.Logger.Error("Could not create new config file", "err", err)
				os.Exit(1)
			}

			viper.SetConfigFile(filepath.Join(home, ".arclog.toml"))
			helpers.Logger.Debug("Created new config file", "path", viper.ConfigFileUsed())
		} else {
			helpers.Logger.Error("Could not read config", "err", err)
			os.Exit(1)
		}
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
