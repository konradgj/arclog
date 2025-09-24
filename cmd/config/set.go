/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"github.com/konradgj/arclog/pkg/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setCmd represents the set command
var (
	logPath   string
	userToken string

	setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set values in config",
		Long:  `Set values for config file.`,
		Run: func(cmd *cobra.Command, args []string) {
			setValues()
		},
	}
)

func init() {
	setCmd.Flags().StringVarP(&logPath, "logpath", "l", "", "Set log path")
	setCmd.Flags().StringVarP(&userToken, "usertoken", "t", "", "Set usertoken")
	setCmd.MarkFlagsOneRequired("logpath", "usertoken")
}

func setValues() {
	if logPath != "" {
		viper.Set("LogPath", logPath)
	}
	if userToken != "" {
		viper.Set("UserToken", userToken)
	}

	err := viper.WriteConfig()
	if err != nil {
		helpers.Logger.Error("could not write config", "error", err)
	}
}
