package cli

type AddCmd struct {
	Paths []string `arg:"" name:"path" type:"path" help:"Paths to add. (supports multiple paths)"`
}

func (l AddCmd) Run(ctx *Context) error {
	ctx.RunAddLogsToDb(l.Paths)
	return nil
}
