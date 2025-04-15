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
	"github.com/spf13/cobra"
	"jester.noble.xyz/v2/appstate"
)

// deleteCmd represents the create command to create a Jester Siging Key
func deleteCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete Jester's signing key",

		RunE: func(cmd *cobra.Command, _ []string) error {
			k := a.Key

			if !a.KeyExists() {
				cmd.Println("No key to delete.")
				return nil
			}

			buf := bufio.NewReader(cmd.InOrStdin())

			yes, err := input.GetConfirmation("Delete Jester's signing key?", buf, cmd.ErrOrStderr())
			if err != nil {
				return fmt.Errorf("failed to read confirmation: %w", err)
			}
			if !yes {
				return nil
			}

			k.Delete(appstate.JesterSigningKey)

			cmd.Println("Jester's signing key deleted 👋")

			return nil
		},
	}

	return cmd
}
