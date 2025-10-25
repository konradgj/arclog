package cli

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
	"github.com/konradgj/arclog/internal/arclog"
	"go.uber.org/zap"
)

const appDir = "arclog"

type Context struct {
	Debug     bool
	CancelCtx context.Context
	Logger    *zap.SugaredLogger
	*arclog.Context
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Add    AddCmd    `cmd:"" help:"Add logs to db."`
	Config ConfigCmd `cmd:"" help:"Manage config."`
	Ls     LsCmd     `cmd:"" help:"List logs added to db."`
	Rm     RmCmd     `cmd:"" help:"Remove log from db"`
	Upload UploadCmd `cmd:"" help:"Upload pending files to dps.report."`
	Watch  WatchCmd  `cmd:"" help:"Watch log directory for new logs and add to db."`
}

func Execute() {
	cancelCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ctx := kong.Parse(&cli)

	var logger *zap.Logger
	if cli.Debug {
		logger, _ = zap.NewDevelopment()
		logger.Debug("Using development level logging")
	} else {
		logger, _ = zap.NewProduction()
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	arclogContext, err := arclog.InitContext(sugar, appDir, cli.Debug)
	if err != nil {
		sugar.Error("could not initialize context", "err", err)
	}

	cliContext := &Context{
		Debug:     cli.Debug,
		CancelCtx: cancelCtx,
		Logger:    sugar,
		Context:   arclogContext,
	}

	err = ctx.Run(cliContext)
	ctx.FatalIfErrorf(err)
}
