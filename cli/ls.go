package cli

import "fmt"

type LsCmd struct {
	Uploadstatus     string `short:"s" default:"" enum:",pending,uploading,uploaded,failed,skipped" help:"Filter by upload status."`
	Relativepath     string `short:"p" help:"Filter by relative path."`
	Date             string `short:"d" help:"Filter by date (YYYY, YYYYMM or YYYYMMDD)."`
	From             string `help:"Filter from date (YYYY, YYYYMM or YYYYMMDD)."`
	To               string `help:"Filter to date (YYYY, YYYYMM or YYYYMMDD)."`
	EncounterSuccess *bool  `short:"e" help:"Filter by encounter success"`
	ChallengeMode    *bool  `short:"c" help:"Filter by challenge mode"`
}

func (cmd LsCmd) Run(ctx *Context) error {
	ctx.RunListCbtlogsByFilter(cmd.Uploadstatus, cmd.Relativepath, cmd.Date, cmd.From, cmd.To, cmd.ChallengeMode, cmd.EncounterSuccess)
	return nil
}

func (cmd LsCmd) Validate() error {
	if cmd.Date != "" && (cmd.From != "" || cmd.To != "") {
		return fmt.Errorf("cannot use -d with --from or --to")
	}

	if err := valitdateDateFormat(cmd.Date, "-d"); err != nil {
		return err
	}
	if err := valitdateDateFormat(cmd.From, "--from"); err != nil {
		return err
	}
	if err := valitdateDateFormat(cmd.To, "--to"); err != nil {
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
