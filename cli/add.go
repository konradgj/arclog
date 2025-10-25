package cli

type AddCmd struct {
	Paths []string `arg:"" name:"path" type:"path" help:"Paths to add. (supports multiple paths)"`
}

func (cmd AddCmd) Run(ctx *Context) error {
	ctx.RunAddLogsToDb(cmd.Paths)
	return nil
}
