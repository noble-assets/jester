package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/noble-assets/jester/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.EnableCommandSorting = false
	cmd := cmd.NewRootCommand()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling for gracefull shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		fmt.Println("Received interrupt signal, shutting down...")
		cancel()
	}()

	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
