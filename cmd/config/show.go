/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"github.com/konradgj/arclog/internal/appconfig"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show config settings",
	Long:  `Show current config settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		appconfig.Show()
	},
}

func init() {
}
