// SPDX-License-Identifier: Apache-2.0
//
// Copyright 2025 NASD Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"jester.noble.xyz/v2/appstate"
	"jester.noble.xyz/v2/cmd/config"
)

const (
	appName         = "jesterd"
	defaultLogLevel = "info"
	defaultLogStyle = "pretty"
)

func NewRootCommand() *cobra.Command {
	a := &appstate.AppState{
		Viper: viper.New(),
	}

	rootCmd := &cobra.Command{
		Use:   appName,
		Short: "Jester is a 'sidecar' meant to run alongside the nobled binary.",
		Long: `Jester is a 'sidecar' meant to run alongside the nobled binary.
	
Jester is only necessary if you are also a validator.`,
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			// Inside persistent pre-run because this takes effect after flags are
			// parsed for both root and all child commands.
			a.ConfigureViper(cmd)
			a.LoadConfig()
			a.InitLogger()
		},
	}

	rootCmd.PersistentFlags().String(appstate.FlagHome, defaultHome(), "directory for config and data \n (optional env var = JESTER_HOME)")
	// manually bind "home" instead of using viper.AutomaticEnv
	if err := a.Viper.BindEnv(appstate.FlagHome, "JESTER_HOME"); err != nil {
		panic(err)
	}
	rootCmd.PersistentFlags().String(appstate.FlagLogLevel, defaultLogLevel, "log level format (info, debug, warn, error)")
	rootCmd.PersistentFlags().String(appstate.FlagLogStyle, defaultLogStyle, "log style format (text, json, pretty)")

	rootCmd.AddCommand(
		config.ConfigCmd(a),
		startCmd(a),
		versionCmd(),
	)

	return rootCmd
}

func defaultHome() string {
	userHome, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(userHome, ".jester")
}
