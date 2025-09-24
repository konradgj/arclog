/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage config file",
	Long:  `Manage config file.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("config called")
	// },
}

func init() {
	ConfigCmd.AddCommand(setCmd)
	ConfigCmd.AddCommand(showCmd)
}
