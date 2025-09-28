package cli

type LogCmd struct {
	Add  LogAddCmd  `cmd:"" help:"Add logs to db."`
	List LogListCmd `cmd:"" help:"List logs added to db."`
}

type LogAddCmd struct {
	Paths []string `arg:"" name:"path" type:"path" help:"Paths to add. (supports multiple paths)"`
}

func (l LogAddCmd) Run(ctx *Context) error {
	ctx.RunAddLogsToDb(l.Paths)
	return nil
}

type LogListCmd struct {
	Uploadstatus string `short:"s" default:"" enum:",pending,uploading,uploaded,failed,skipped" help:"Filter by upload status."`
	Relativepath string `short:"p" help:"Filter by relative path."`
}

func (l LogListCmd) Run(ctx *Context) error {
	ctx.RunListCbtlogsByFilter(l.Uploadstatus, l.Relativepath)
	return nil
}
