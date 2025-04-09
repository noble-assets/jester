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

package wormhole

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/avast/retry-go/v4"
	"jester.noble.xyz/metrics"
)

// QueryData contains the metadata needed to query VAA's from the Wormhole API.
type queryData struct {
	sequence uint64
	txHash   string // logging purposes only
}

type WormholeResp struct {
	VaaBytes string `json:"vaaBytes"`
}

var (
	errTooManyRequests = fmt.Errorf("%d:%s", http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
	errServerError     = fmt.Errorf(">= %d:%s", http.StatusInternalServerError, "server error")
	errNotFound        = fmt.Errorf("%d:%s - waiting for wormhole to pickup tx", http.StatusNotFound, http.StatusText(http.StatusNotFound))
)

// StartWormholeWorker starts a worker to query the Wormhole API for VAA's.
// Once found, the VAA is added to the vaaList which is queryable via gRPC
func (w *Wormhole) StartWormholeWorker(
	ctx context.Context, log *slog.Logger,
	m *metrics.PrometheusMetrics,
	dequeued *queryData,
) {
	resp, err := w.fetchVaa(ctx, log, m, dequeued.sequence, dequeued.txHash)
	if err != nil {
		log.Error("wormhole VAA query failed", "error", err)
		return
	}

	log.Info("found VAA", "txHash", dequeued.txHash)

	vaa, _ := base64.StdEncoding.DecodeString(resp.VaaBytes)
	w.VaaList.Add(vaa)
}

// fetchVaa sends a GET request with retry logic to the wormhole API
//
// `txHash` is used for logging purposes only
func (w *Wormhole) fetchVaa(
	ctx context.Context, log *slog.Logger,
	m *metrics.PrometheusMetrics,
	seq uint64,
	txHash string,
) (WormholeResp, error) {
	chainIdStr := strconv.FormatUint(uint64(w.config.wormholeSrcChainId), 10)
	seqStr := strconv.FormatUint(seq, 10)
	url := fmt.Sprintf("%s/%s/%s/%s", w.config.wormholeApiUrl, chainIdStr, w.config.paddedWormholeTransceiver, seqStr)

	var (
		wormholeResp   WormholeResp
		elapsed        time.Duration
		currentAttempt uint
	)

	firstAttempt := time.Now()

	err := retry.Do(
		func() error {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return fmt.Errorf("failed to create request: %w", err)
			}

			req.Header.Set("accept", "application/json")

			client := &http.Client{Timeout: 10 * time.Second}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("request failed: %w", err)
			}
			defer resp.Body.Close()

			switch {
			case resp.StatusCode >= http.StatusInternalServerError:
				return errServerError
			case resp.StatusCode == http.StatusTooManyRequests:
				return errTooManyRequests
			case resp.StatusCode == http.StatusNotFound: // likely waiting for wormhole to pick up tx
				return errNotFound
			case resp.StatusCode != http.StatusOK:
				return retry.Unrecoverable(fmt.Errorf("non-retryable error: code %d, code: %s",
					resp.StatusCode, http.StatusText(resp.StatusCode),
				))
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read wormhole response body: %w", err)
			}

			if err := json.Unmarshal(body, &wormholeResp); err != nil {
				return fmt.Errorf("failed to unmarshall wormholeResp json: %w", err)
			}

			return nil
		},
		retry.RetryIf(func(err error) bool {
			if errors.Is(err, context.DeadlineExceeded) {
				return true
			}
			switch err {
			case errServerError, errTooManyRequests, errNotFound:
				return true
			default:
				return false
			}
		}),
		// adjust Attempts and Delay to ensure we don't give up querying
		// wormhole too soon
		retry.Attempts(w.config.fetchVAAAttempts),
		retry.Context(ctx),
		retry.Delay(30*time.Second),
		retry.DelayType(retry.FixedDelay),
		retry.OnRetry(func(attempt uint, err error) {
			elapsed = time.Since(firstAttempt).Round(time.Second)
			currentAttempt = attempt
			log.Info("retry: VAA lookup", "attempt", fmt.Sprintf(
				"%d/%d", attempt+1, w.config.fetchVAAAttempts), "seq", seq, "error", err, "since-first-attempt", elapsed, "txHash", txHash,
			)
		}),
	)
	if err != nil {
		if currentAttempt == w.config.fetchVAAAttempts-1 {
			err = fmt.Errorf("max VAA lookup attempts reached: %w", err)
			m.VAAFailedMaxAttemptsReached.Inc()
		}
		m.VAAFailedTotal.Inc()
		return WormholeResp{}, fmt.Errorf("query URL: %s: %w", url, err)
	}

	// keep metrics on how long it takes wormhole to pickup transaction
	// if current attempt is 0, we assume it was a historical query and ignore
	if currentAttempt > 0 {
		m.VAAReceiveDuration.Observe(float64(elapsed.Minutes()))
	}

	m.VAAFoundTotal.Inc()
	return wormholeResp, nil
}
