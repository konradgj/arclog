package upload

import (
	"context"

	"github.com/konradgj/arclog/internal/arclog"
	"github.com/spf13/cobra"
)

func NewUploadCmd(ctx *arclog.Context, cancelCtx context.Context) *cobra.Command {
	var watch bool
	var anonymous bool
	var detailedwvw bool

	uploadCmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload logs to arcdps",
		Long: `Upload logs marked as pending in db to arc dps.
Use -w flag to upload as files are created.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx.Config.AbortIfNoUserToken()

			if watch {
				ctx.RunWatchUploads(anonymous, detailedwvw, cancelCtx)
			} else {
				ctx.RunPendingUploads(anonymous, detailedwvw, cancelCtx)
			}
		},
	}

	uploadCmd.Flags().BoolVarP(&watch, "watch", "w", false, "Watch and upload when created")
	uploadCmd.Flags().BoolVarP(&anonymous, "anonymous", "a", false, "Upload anonymous log")
	uploadCmd.Flags().BoolVarP(&detailedwvw, "detailedwvw", "d", false, "Generate detalailedwvw log")

	return uploadCmd
}
