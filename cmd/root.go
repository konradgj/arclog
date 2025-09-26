/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/konradgj/arclog/cmd/config"
	"github.com/konradgj/arclog/cmd/watch"
	"github.com/konradgj/arclog/internal/arclog"
	"github.com/spf13/cobra"
)

func NewRootCmd(rootCtx context.Context) *cobra.Command {

	var verbose bool
	ctx := arclog.NewContext()

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
		watch.NewWatchCmd(ctx, rootCtx),
	)

	return rootCmd
}
