/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"fmt"

	"github.com/konradgj/arclog/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewSetCmd(ctx *app.Context) *cobra.Command {
	var logPath string
	var userToken string

	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Set values in config",
		Long:  `Set values for config file.`,
		Run: func(cmd *cobra.Command, args []string) {
			setValues(ctx, logPath, userToken)
		},
	}
	setCmd.Flags().StringVarP(&logPath, "logpath", "l", "", "Set log path")
	setCmd.Flags().StringVarP(&userToken, "usertoken", "t", "", "Set usertoken")
	setCmd.MarkFlagsOneRequired("logpath", "usertoken")

	return setCmd
}

func setValues(ctx *app.Context, logPath, userToken string) {
	if logPath != "" {
		viper.Set("LogPath", logPath)
	}
	if userToken != "" {
		viper.Set("UserToken", userToken)
	}

	err := viper.WriteConfig()
	if err != nil {
		ctx.Log.Error("could not write config", "error", err)
	}

	fmt.Println("Updated config:")
	ctx.Config.Show()
}
