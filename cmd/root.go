/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"

	"github.com/konradgj/arclog/cmd/config"
	"github.com/konradgj/arclog/cmd/upload"
	"github.com/konradgj/arclog/cmd/watch"
	"github.com/konradgj/arclog/internal/arclog"
	"github.com/konradgj/arclog/internal/logger"
	"github.com/spf13/cobra"
)

const appDir = "arclog"

func NewRootCmd(rootCtx context.Context) *cobra.Command {
	var ctx *arclog.Context
	var verbose bool

	rootCmd := &cobra.Command{
		Use:   "arclog",
		Short: "Upload arc-dps log sto dps.report",
		Long:  `arclog is a CLI tool for uploading ardps logs to dps.report as they are generated.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 	logger.Initlogger(verbose)

		// },
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	cobra.OnInitialize(func() { logger.Initlogger(verbose) })

	ctx, err := arclog.InitContext(appDir, verbose)
	if err != nil {
		logger.Error("could not init config", "err", err)
		os.Exit(1)
	}

	rootCmd.AddCommand(
		config.NewConfigCmd(ctx),
		watch.NewWatchCmd(ctx, rootCtx),
		upload.NewUploadCmd(ctx, rootCtx),
	)

	return rootCmd
}
