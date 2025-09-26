/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"github.com/konradgj/arclog/internal/arclog"
	"github.com/spf13/cobra"
)

func NewConfigCmd(ctx *arclog.Context) *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Manage config file",
		Long:  `Manage config file.`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("config called")
		// },
	}
	configCmd.AddCommand(
		NewSetCmd(ctx),
		NewShowCmd(ctx),
	)
	return configCmd
}
