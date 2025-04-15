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

package keys

import (
	"github.com/spf13/cobra"
	"jester.noble.xyz/v2/appstate"
)

// KeysCmd represents the keys command
func KeysCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Keyring commands",
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	cmd.AddCommand(
		createCmd(a),
		deleteCmd(a),
		recoverCmd(a),
		exportCmd(a),
		importCmd(a),
		showCmd(a),
	)

	return cmd
}
