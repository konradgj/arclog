/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package watch

import (
	"github.com/konradgj/arclog/internal/watcher"
	"github.com/spf13/cobra"
)

// watchCmd represents the watch command
var WatchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Start uploading logs",
	Long: `Start monitoring for creation of arc-dps logs.
Will upload the logs to dps.report using current configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		watcher.Watch()
	},
}

func init() {

}
