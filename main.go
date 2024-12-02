package main

import (
	"context"
	"os"

	"github.com/noble-assets/jester/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.EnableCommandSorting = false
	cmd := cmd.NewRootCommand()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
