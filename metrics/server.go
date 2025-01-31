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

package metrics

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartServer starts a prometheus metrics server on a given address
func (m *PrometheusMetrics) StartServer(ctx context.Context, log *slog.Logger, mux *http.ServeMux, address string) error {
	log = log.With(slog.String("server", "prometheus-metrics"))
	m.enabled = true

	// Serve default prometheus metrics
	mux.Handle("/metrics", promhttp.Handler())

	// Serve jester specific metrics
	mux.Handle("/jester/metrics", promhttp.HandlerFor(
		m.registry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	srv := &http.Server{
		Addr:     address,
		Handler:  mux,
		ErrorLog: slog.NewLogLogger(log.Handler(), slog.LevelError),
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	errChan := make(chan error, 1)
	go func() {
		log.Info("starting server", "address", address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
			close(errChan)
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("shutdown error: %w", err)
		}
		return nil
	case err := <-errChan:
		return fmt.Errorf("server error: %w", err)
	}
}
