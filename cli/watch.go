package cli

import "fmt"

type WatchCmd struct {
}

func (w WatchCmd) Run(ctx *Context) error {
	ctx.RunWatch(ctx.CancelCtx)
	fmt.Println()
	return nil
}
