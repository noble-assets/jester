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
	"bufio"

	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/spf13/cobra"
	"jester.noble.xyz/v2/appstate"
)

// exportCmd represents the export command to export a Jester Siging Key
func exportCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export Jester's signing key in ASCII armored format",

		RunE: func(cmd *cobra.Command, _ []string) error {
			k := a.Key

			if !a.KeyExists() {
				cmd.Println("No key to export.")
				return nil
			}

			buf := bufio.NewReader(cmd.InOrStdin())

			encryptPassword, err := input.GetPassword("Enter passphrase to encrypt the exported key:", buf)
			if err != nil {
				return err
			}

			priv, err := k.ExportPrivKeyArmor(appstate.JesterSigningKey, encryptPassword)
			if err != nil {
				return err
			}

			cmd.Println(priv)

			return nil
		},
	}

	return cmd
}
