/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/konradgj/arclog/cmd/config"
	"github.com/konradgj/arclog/cmd/watch"
	"github.com/konradgj/arclog/internal/appconfig"
	"github.com/konradgj/arclog/internal/logger"
	"github.com/spf13/cobra"
)

var (
	verbose bool

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
	cobra.OnInitialize(appconfig.InitConfig)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(watch.WatchCmd)
}
