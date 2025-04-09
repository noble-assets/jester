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

package ethereum

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

// FilterLogs wraps the RPC client's FilterLogs method with retry logic.
func (e *Eth) FilterLogs(ctx context.Context, log *slog.Logger, q ethereum.FilterQuery) (ethLogs []ethTypes.Log, err error) {
	err = retry.Do(
		func() error {
			ethLogs, err = e.RPCClient.FilterLogs(ctx, q)
			if err != nil {
				return fmt.Errorf("filterlogs failed: %w", err)
			}
			return nil
		},
		retry.Context(ctx),
		retry.OnRetry(func(attempt uint, err error) {
			log.Warn("retrying FilterLogs", "attempt", fmt.Sprintf("%d/%d", attempt+1, 10), "error", err)
		}),
	)
	if err != nil {
		return nil, err
	}

	return ethLogs, nil
}

// headerByNumber wraps the RPC client's HeaderByNumber method with retry logic.
func (e *Eth) headerByNumber(ctx context.Context, log *slog.Logger, block *big.Int) (header *ethTypes.Header, err error) {
	err = retry.Do(
		func() error {
			header, err = e.RPCClient.HeaderByNumber(ctx, block)
			if err != nil {
				return fmt.Errorf("headerByNumber failed: %w", err)
			}
			return nil
		},
		retry.Context(ctx),
		retry.OnRetry(func(attempt uint, err error) {
			log.Warn("retrying HeaderByNumber", "attempt", fmt.Sprintf("%d/%d", attempt+1, 10), "error", err)
		}),
	)
	if err != nil {
		return nil, err
	}

	return header, nil
}
