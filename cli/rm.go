package cli

type RmCmd struct {
	FileName string `arg:"" name:"filename" help:"File to drop from db."`
	Delete   bool   `short:"d" help:"Delete file."`
}

func (cmd RmCmd) Run(ctx *Context) error {
	ctx.RunRmCmd(cmd.FileName, cmd.Delete)
	return nil
}
