package cli

type LogCmd struct {
	Add LogAddCmd `cmd:"" help:"Add logs to db."`
}

type LogAddCmd struct {
	Paths []string `arg:"" name:"path" type:"path" help:"Paths to add. (supports multiple paths)"`
}

func (l LogAddCmd) Run(ctx *Context) error {
	ctx.AddLogsToDb(l.Paths)
	return nil
}
