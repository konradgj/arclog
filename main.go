/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/konradgj/arclog/cmd"
)

func main() {
	rootCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cmd := cmd.NewRootCmd(rootCtx)
	if err := cmd.Execute(); err != nil {
		log.Fatalf("CLI execution failed: %v", err)
		os.Exit(1)
	}
}
