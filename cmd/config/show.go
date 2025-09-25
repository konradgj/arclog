/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"github.com/konradgj/arclog/internal/app"
	"github.com/spf13/cobra"
)

func NewShowCmd(ctx *app.Context) *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "Show config settings",
		Long:  `Show current config settings.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx.Config.Show()
		},
	}

	return showCmd
}
