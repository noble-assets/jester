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
	"errors"
	"fmt"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

// startProcessor starts the hyplane processor
func (h *Hyperlane) startProcessor(ctx context.Context, log *slog.Logger) error {
	g, ctx := errgroup.WithContext(ctx)

	// TODO: worker pool?

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case dequeued, ok := <-h.processingQueue:
				if !ok {
					return errors.New("processingQueue channel closed unexpectedly")
				}
				// TODO
				fmt.Println("Process Me!", dequeued.TxHash.Hex())

			}
		}
	})

	if err := g.Wait(); err != nil {
		log.Error("hyperlane error", "error", err)
		return err
	}

	return nil

}
