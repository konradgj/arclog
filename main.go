/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"os"

	"github.com/konradgj/arclog/cmd"
)

func main() {

	cmd := cmd.NewRootCmd()
	if err := cmd.Execute(); err != nil {
		log.Fatalf("CLI execution failed: %v", err)
		os.Exit(1)
	}
}
