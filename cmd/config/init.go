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

package config

import (
	"fmt"
	"os"
	"path/filepath"

	"jester.noble.xyz/v2/appstate"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
func initCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the config file",
		Long: `Initialize the config file. Optionally use flags to 
populate values.`,

		RunE: func(_ *cobra.Command, _ []string) error {
			configPath := filepath.Join(a.Viper.GetString(appstate.FlagHome), ".jester", "config.toml")

			// programmatically overwriting config file does not treat viper keys ase expected
			// instead, notify user and have them delete config manually
			if fileExists(configPath) {
				return fmt.Errorf("file '%s' already exists", configPath)
			}

			// create config file
			dir := filepath.Dir(configPath)
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return fmt.Errorf("failed to create directories: %w", err)
			}
			file, err := os.Create(configPath)
			if err != nil {
				return err
			}
			defer file.Close()

			// ensure valid toml
			encoder := toml.NewEncoder(file)
			err = encoder.Encode(a.Config)
			if err != nil {
				return fmt.Errorf("invalid config created: %v", err)
			}
			fmt.Printf("Config created at: '%s'\n", configPath)

			return nil
		},
	}

	return cmd
}

// fileExists checks if a file exists
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
