/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"fmt"

	"github.com/konradgj/arclog/internal/logger"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show config settings",
	Long:  `Show current config settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		showConfig()
	},
}

func init() {
}

func showConfig() {
	settings := viper.AllSettings()

	data, err := toml.Marshal(settings)
	if err != nil {
		logger.Error("Failed to marshal config", "error", err)
	}

	fmt.Println("----------------------")
	fmt.Print(string(data))
	fmt.Println("----------------------")
}
