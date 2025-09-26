/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package watch

import (
	"github.com/konradgj/arclog/internal/arclog"
	"github.com/spf13/cobra"
)

// watchCmd represents the watch command
func NewWatchCmd(ctx *arclog.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "Start uploading logs",
		Long: `Start monitoring for creation of arc-dps logs.
Will upload the logs to dps.report using current configuration file.`,
		Run: func(cmd *cobra.Command, args []string) {
			arclog.RunWatch(ctx)
		},
	}
}

func init() {

}
