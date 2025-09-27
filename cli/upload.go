package cli

import "fmt"

type UploadCmd struct {
	Anonymous   bool `short:"a" help:"Upload anonymously."`
	Detailedwvw bool `short:"d" help:"Include detailed WvW logs."`
	Watch       bool `short:"w" help:"Monitor log dir and upload as logs are created."`
}

func (u *UploadCmd) Run(ctx *Context) error {
	if u.Watch {
		ctx.RunWatchUploads(u.Anonymous, u.Detailedwvw, ctx.CancelCtx)
	} else {
		ctx.RunPendingUploads(u.Anonymous, u.Detailedwvw, ctx.CancelCtx)
	}

	fmt.Println()
	return nil
}
