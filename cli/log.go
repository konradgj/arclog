package cli

import "fmt"

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
	Date         string `short:"d" help:"Filter by date (YYYY, YYYYMM or YYYYMMDD)."`
}

func (l LogListCmd) Run(ctx *Context) error {
	ctx.RunListCbtlogsByFilter(l.Uploadstatus, l.Relativepath, l.Date)
	return nil
}

func (l LogListCmd) Validate() error {
	if l.Date == "" {
		return nil
	}
	if len(l.Date) == 4 || len(l.Date) == 6 || len(l.Date) == 8 {
		return nil
	} else {
		return fmt.Errorf("invalid date format: use YYYY, YYYYMM or YYYYMMDD")

	}
}
