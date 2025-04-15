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

package appstate

import (
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const JesterSigningKey = "jesterSigningKey"

func (a *AppState) NewKeyring() (err error) {
	configureSDK()

	a.Key, err = keyring.New("jester", keyring.BackendTest, a.Home, os.Stdin, a.cdc)
	if err != nil {
		return err
	}

	return nil
}

// configureSDK sets up the Cosmos SDK with the jester bech32 prefix
func configureSDK() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("jester", "jesterpub")
	config.Seal()
}

// KeyExists returns true if a Jester signing key exists in the keystore, it returns false otherwise.
func (a *AppState) KeyExists() bool {
	r, err := a.Key.Key(JesterSigningKey)
	if err != nil {
		return false
	}

	return r.Name == JesterSigningKey
}
