/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"fmt"

	"github.com/konradgj/arclog/internal/app"
	"github.com/spf13/cobra"
)

func NewSetCmd(ctx *app.Context) *cobra.Command {
	var logPath string
	var userToken string

	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Set values in config",
		Long:  `Set values for config file.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx.Config.SetValues(logPath, userToken)
			fmt.Println("Updated config:")
			ctx.Config.Show()
		},
	}
	setCmd.Flags().StringVarP(&logPath, "logpath", "l", "", "Set log path")
	setCmd.Flags().StringVarP(&userToken, "usertoken", "t", "", "Set usertoken")
	setCmd.MarkFlagsOneRequired("logpath", "usertoken")

	return setCmd
}
