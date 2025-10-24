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
	Uploadstatus     string `short:"s" default:"" enum:",pending,uploading,uploaded,failed,skipped" help:"Filter by upload status."`
	Relativepath     string `short:"p" help:"Filter by relative path."`
	Date             string `short:"d" help:"Filter by date (YYYY, YYYYMM or YYYYMMDD)."`
	From             string `help:"Filter from date (YYYY, YYYYMM or YYYYMMDD)."`
	To               string `help:"Filter to date (YYYY, YYYYMM or YYYYMMDD)."`
	EncounterSuccess *bool  `short:"e" help:"Filter by encounter success"`
	ChallengeMode    *bool  `short:"c" help:"Filter by challenge mode"`
}

func (l LogListCmd) Run(ctx *Context) error {
	ctx.RunListCbtlogsByFilter(l.Uploadstatus, l.Relativepath, l.Date, l.From, l.To, l.ChallengeMode, l.EncounterSuccess)
	return nil
}

func (l LogListCmd) Validate() error {
	if l.Date != "" && (l.From != "" || l.To != "") {
		return fmt.Errorf("cannot use -d with --from or --to")
	}

	if err := valitdateDateFormat(l.Date, "-d"); err != nil {
		return err
	}
	if err := valitdateDateFormat(l.From, "--from"); err != nil {
		return err
	}
	if err := valitdateDateFormat(l.To, "--to"); err != nil {
		return err
	}
	return nil
}

func valitdateDateFormat(val, name string) error {
	if val == "" {
		return nil
	}
	if len(val) == 4 || len(val) == 6 || len(val) == 8 {
		return nil
	}
	return fmt.Errorf("invalid %s format: use YYYY, YYYYMM or YYYYMMDD", name)
}
