/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/konradgj/arclog/cmd"
	"github.com/konradgj/arclog/internal/app"
)

func main() {
	ctx := app.NewContext()

	cmd := cmd.NewRootCmd(ctx)
	if err := cmd.Execute(); err != nil {
		ctx.Log.Error("CLI execution failed", "err", err)
		os.Exit(1)
	}
}
