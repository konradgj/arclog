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

func NewRootCmd() *cobra.Command {
	var verbose bool
	ctx := app.NewContext()

	rootCmd := &cobra.Command{
		Use:   "arclog",
		Short: "Upload arc-dps log sto dps.report",
		Long:  `arclog is a CLI tool for uploading ardps logs to dps.report as they are generated.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctx.Init(verbose)
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	rootCmd.AddCommand(
		config.NewConfigCmd(ctx),
		watch.NewWatchCmd(ctx),
	)

	return rootCmd
}
