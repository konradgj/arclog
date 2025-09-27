package cli

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
	"github.com/konradgj/arclog/internal/arclog"
	"github.com/konradgj/arclog/internal/logger"
)

const appDir = "arclog"

type Context struct {
	Debug     bool
	Verbose   bool
	CancelCtx context.Context
	*arclog.Context
}

var cli struct {
	Debug   bool `help:"Enable debug mode."`
	Verbose bool `help:"Enable verbose output."`

	Config ConfigCmd `cmd:"" help:"Manage config."`
	Watch  WatchCmd  `cmd:"" help:"Watch log directory for new logs."`
	Upload UploadCmd `cmd:"" help:"Upload pending files to dps.report."`
}

func Execute() {
	cancelCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ctx := kong.Parse(&cli)

	logger.Initlogger(cli.Verbose)
	arclogContext, err := arclog.InitContext(appDir, cli.Verbose)
	if err != nil {
		logger.Error("could not initialize context", "err", err)
	}

	cliContext := &Context{
		Verbose:   cli.Verbose,
		Debug:     cli.Debug,
		CancelCtx: cancelCtx,
		Context:   arclogContext,
	}

	err = ctx.Run(cliContext)
	ctx.FatalIfErrorf(err)
}
