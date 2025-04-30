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
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/spf13/cobra"
	"jester.noble.xyz/v2/appstate"
)

// recoverCmd represents the recover command to recover a Jester Siging Key
func recoverCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recover",
		Short: "Recover a Jester signing key from a mnemonic",
		Long:  "Recover a Jester signing key from a mnemonic. Jester can only have one key loaded at a time.",

		RunE: func(cmd *cobra.Command, _ []string) error {
			k := a.Key

			if a.KeyExists() {
				cmd.Println("Jester's signing key already exists. Delete it before recovering a new one.")
				return nil
			}

			buf := bufio.NewReader(cmd.InOrStdin())

			mnemonicStr, err := input.GetString("Enter your bip39 mnemonic: ", buf)
			if err != nil {
				return fmt.Errorf("failed to read mnemonic: %w", err)
			}

			keyRecord, err := k.NewAccount(appstate.JesterSigningKey, mnemonicStr, "", hd.CreateHDPath(coinType, account, index).String(), algo)
			if err != nil {
				return fmt.Errorf("failed to create account: %w", err)
			}

			addr, err := keyRecord.GetAddress()
			if err != nil {
				return fmt.Errorf("failed to get address: %w", err)
			}

			cmd.Printf("\nRecovered key: %s\n", addr.String())

			return nil
		},
	}

	return cmd
}
