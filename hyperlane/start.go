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

package hyperlane

import (
	"context"
	"log/slog"

	"golang.org/x/sync/errgroup"
	eth "jester.noble.xyz/v2/ethereum"
	"jester.noble.xyz/v2/metrics"
)

// Start starts Hyperlane interoperability integration.
func (h *Hyperlane) Start(
	ctx context.Context, log *slog.Logger, Eth *eth.Eth,
	metrics *metrics.PrometheusMetrics,
) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return h.startProcessor(ctx, log)
	})

	g.Go(func() error {
		return h.startHyperlaneDispatchListener(ctx, log, Eth)
	})

	if err := g.Wait(); err != nil {
		log.Error("hyperlane error", "error", err)
		return err
	}

	return nil
}
