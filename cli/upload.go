package cli

import (
	"fmt"

	"github.com/konradgj/arclog/internal/db"
)

type UploadCmd struct {
	Anonymous   bool     `short:"a" help:"Upload anonymously."`
	Detailedwvw bool     `short:"d" help:"Include detailed WvW logs."`
	Watch       bool     `short:"w" help:"Monitor log dir and upload as logs are created."`
	Paths       []string `short:"p" type:"path" help:"Upload from given paths. (supports multiple paths)"`
	Status      string   `short:"s" default:"" enum:",pending,uploading,uploaded,failed,skipped" help:"Filter logs by upload status."`
}

func (cmd *UploadCmd) Validate() error {
	if cmd.Status != "" && (len(cmd.Paths) > 0 || cmd.Watch) {
		return fmt.Errorf("--status cannot be used with --path or --watch")
	}
	return nil
}

func (u *UploadCmd) Run(ctx *Context) error {
	if u.Status != "" {
		ctx.RunUploadsByStatus(u.Status, u.Anonymous, u.Detailedwvw, ctx.CancelCtx)
		return nil
	}

	if len(u.Paths) > 0 {
		ctx.RunUploadPathLogs(u.Paths, u.Anonymous, u.Detailedwvw)
		return nil
	}

	if u.Watch {
		ctx.RunWatchUploads(u.Anonymous, u.Detailedwvw, ctx.CancelCtx)
	} else {
		ctx.RunUploadsByStatus(string(db.StatusPending), u.Anonymous, u.Detailedwvw, ctx.CancelCtx)
	}

	fmt.Println()
	return nil
}
