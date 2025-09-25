/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/konradgj/arclog/cmd/config"
	"github.com/konradgj/arclog/cmd/watch"
	"github.com/konradgj/arclog/internal/app"
	"github.com/spf13/cobra"
)

var (
	verbose bool
)

func NewRootCmd(ctx *app.Context) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "arclog",
		Short: "Upload arc-dps log sto dps.report",
		Long:  `arclog is a CLI tool for uploading ardps logs to dps.report as they are generated.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	rootCmd.AddCommand(
		config.NewConfigCmd(ctx),
		watch.NewWatchCmd(ctx),
	)

	return rootCmd
}
