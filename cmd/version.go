package cmd

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// Version defines the application version info (defined at compile time)
	Version string
	Commit  string
)

type versionInfo struct {
	Version string `json:"version" yaml:"version"`
	Commit  string `json:"commit" yaml:"commit"`
	Go      string `json:"go" yaml:"go"`
}

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Print version info",
		RunE: func(_ *cobra.Command, _ []string) error {
			verInfo := versionInfo{
				Version: Version,
				Commit:  Commit,
				Go:      fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
			}

			json, err := json.MarshalIndent(verInfo, "", "  ")
			if err != nil {
				return err
			}

			fmt.Print(string(json))
			return nil
		},
	}

	return cmd
}
