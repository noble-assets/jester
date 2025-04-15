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
	"os"

	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/spf13/cobra"
	"jester.noble.xyz/v2/appstate"
)

// importCmd represents the import command to import a Jester Siging Key
func importCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import <keyfile>",
		Short: "Import Jester's signing key from a ASCII armored private key file",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			k := a.Key

			if a.KeyExists() {
				cmd.Println("Jester's signing key already exists. Delete it before creating a new one.")
				return nil
			}

			bz, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			buf := bufio.NewReader(cmd.InOrStdin())

			encryptPassword, err := input.GetPassword("Enter passphrase to decrypt your key:", buf)
			if err != nil {
				return err
			}

			return k.ImportPrivKey(appstate.JesterSigningKey, string(bz), encryptPassword)
		},
	}

	return cmd
}
