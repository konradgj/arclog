/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/konradgj/arclog/cmd/config"
	"github.com/konradgj/arclog/internal/logger"
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
	cobra.OnInitialize(func() {
		logger.Init(verbose)
	})
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Use config file (default is $HOME/.arclog.toml)")

	rootCmd.AddCommand(config.ConfigCmd)
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		logger.Error("Could not get home directory", "err", err)
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
				logger.Error("Could not create new config file", "err", err)
				os.Exit(1)
			}

			viper.SetConfigFile(filepath.Join(home, ".arclog.toml"))
			logger.Debug("Created new config file", "path", viper.ConfigFileUsed())
		} else {
			logger.Error("Could not read config", "err", err)
			os.Exit(1)
		}
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
