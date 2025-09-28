package cli

type LogCmd struct {
	Add LogAddCmd `cmd:"" help:"Add logs to db."`
}

type LogAddCmd struct {
	Paths []string `arg:"" name:"path" type:"path" help:"Paths to add."`
}

func (l LogAddCmd) Run(ctx *Context) error {
	for _, path := range l.Paths {
		ctx.AddLogsToDb(path)
	}
	return nil
}
