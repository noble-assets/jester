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

	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"jester.noble.xyz/v2/appstate"
)

// showCmd represents the show command to show a Jester Siging Key.
func showCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show Jester's signing key",
		RunE: func(cmd *cobra.Command, _ []string) error {
			r, err := a.Key.Key(appstate.JesterSigningKey)
			if err != nil {
				cmd.Println("No key found.")
				return nil
			}

			keyOutput, err := getKeyOutput(r)
			if err != nil {
				return err
			}

			out, err := yaml.Marshal([]keys.KeyOutput{keyOutput})
			if err != nil {
				return fmt.Errorf("failed to marshal key output: %w", err)
			}

			cmd.Println(string(out))
			return nil
		},
	}

	return cmd
}

func getKeyOutput(r *keyring.Record) (keys.KeyOutput, error) {
	pk, err := r.GetPubKey()
	if err != nil {
		return keys.KeyOutput{}, fmt.Errorf("failed to get public key: %w", err)
	}

	address, err := r.GetAddress()
	if err != nil {
		return keys.KeyOutput{}, fmt.Errorf("failed to get address: %w", err)
	}

	keyOut, err := keys.NewKeyOutput(appstate.JesterSigningKey, r.GetType(), address, pk)
	if err != nil {
		return keys.KeyOutput{}, fmt.Errorf("failed to create key output: %w", err)
	}

	return keyOut, nil
}
