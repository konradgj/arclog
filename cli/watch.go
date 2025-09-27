package cli

import "fmt"

type WatchCmd struct {
	Anonymous   bool `short:"a" help:"Upload anonymously."`
	Detailedwvw bool `short:"d" help:"Include detailed WvW logs."`
}

func (w WatchCmd) Run(ctx *Context) error {
	ctx.RunWatch(ctx.CancelCtx)
	fmt.Println()
	return nil
}
