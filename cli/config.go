package cli

import (
	"fmt"
)

type ConfigCmd struct {
	Set  ConfigSetCmd  `cmd:"" group:"config" help:"Set config values."`
	Show ConfigShowCmd `cmd:"" group:"config" help:"Show config."`
}

type ConfigSetCmd struct {
	Logpath   string `short:"p" help:"Path to log files."`
	Usertoken string `short:"t" help:"User token."`
}

func (cmd ConfigSetCmd) Run(ctx *Context) error {
	if cmd.Logpath == "" && cmd.Usertoken == "" {
		return fmt.Errorf("must provide at least one of --logpath or --usertoken")
	}

	ctx.Config.SetValues(appDir, cmd.Logpath, cmd.Usertoken)
	fmt.Println("Updated config:")
	settings := ctx.Config.GetSettingsString()
	fmt.Println(settings)
	return nil
}

type ConfigShowCmd struct {
}

func (cmd ConfigShowCmd) Run(ctx *Context) error {
	settings := ctx.Config.GetSettingsString()
	fmt.Println(settings)
	return nil
}
