/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/noble-assets/jester/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.EnableCommandSorting = false
	cmd := cmd.NewRootCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
