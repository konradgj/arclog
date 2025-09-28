package cli

import "fmt"

type UploadCmd struct {
	Anonymous   bool     `short:"a" help:"Upload anonymously."`
	Detailedwvw bool     `short:"d" help:"Include detailed WvW logs."`
	Watch       bool     `short:"w" xor:"X" help:"Monitor log dir and upload as logs are created."`
	Paths       []string `short:"p" xor:"X" type:"path" help:"Upload from given paths. (supports multiple paths)"`
}

func (u *UploadCmd) Run(ctx *Context) error {
	if len(u.Paths) > 0 {
		ctx.RunUploadLogs(u.Paths, u.Anonymous, u.Detailedwvw)
		return nil
	}

	if u.Watch {
		ctx.RunWatchUploads(u.Anonymous, u.Detailedwvw, ctx.CancelCtx)
	} else {
		ctx.RunPendingUploads(u.Anonymous, u.Detailedwvw, ctx.CancelCtx)
	}

	fmt.Println()
	return nil
}
