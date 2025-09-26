package watch

import (
	"context"

	"github.com/konradgj/arclog/internal/arclog"
	"github.com/spf13/cobra"
)

func NewWatchCmd(ctx *arclog.Context, cancelCtx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "Start uploading logs",
		Long: `Start monitoring for creation of arc-dps logs.
Will upload the logs to dps.report using current configuration file.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx.RunWatch(cancelCtx)
		},
	}
}
