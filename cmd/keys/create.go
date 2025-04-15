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
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/go-bip39"
	"github.com/spf13/cobra"
	"jester.noble.xyz/v2/appstate"
)

const (
	coinType            = uint32(118)
	account             = uint32(0)
	index               = uint32(0)
	mnemonicEntropySize = 256
)

var algo keyring.SignatureAlgo = hd.Secp256k1

// createCmd represents the create command to create a Jester Siging Key
func createCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Jester signing key",
		Long:  "Create a new Jester signing key. Jester can only have one key loaded at a time.",

		RunE: func(cmd *cobra.Command, _ []string) error {
			k := a.Key

			if a.KeyExists() {
				cmd.Println("Jester's signing key already exists. Delete it before creating a new one.")
				return nil
			}

			mnemonicStr, err := createMnemonic()
			if err != nil {
				return fmt.Errorf("failed to create mnemonic: %w", err)
			}

			keyRecord, err := k.NewAccount(appstate.JesterSigningKey, mnemonicStr, "", hd.CreateHDPath(coinType, account, index).String(), algo)
			if err != nil {
				return fmt.Errorf("failed to create account: %w", err)
			}

			addr, err := keyRecord.GetAddress()
			if err != nil {
				return fmt.Errorf("failed to get address: %w", err)
			}

			cmd.Printf("Generated new key: %s\n", addr.String())
			cmd.Printf("\n\n**MNEMONIC**\n\n%s\n", mnemonicStr)

			return nil
		},
	}

	return cmd
}

// CreateMnemonic generates a new mnemonic.
func createMnemonic() (string, error) {
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropySeed)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}
