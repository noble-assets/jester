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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type PrometheusMetrics struct {
	enabled  bool
	registry *prometheus.Registry

	EthSubInteruptionCounter prometheus.Counter
}

// NewPrometheusMetrics creates a new PrometheusMetrics object
func NewPrometheusMetrics() *PrometheusMetrics {
	registry := prometheus.NewRegistry()
	factory := promauto.With(registry)

	return &PrometheusMetrics{
		registry: registry,
		EthSubInteruptionCounter: factory.NewCounter(prometheus.CounterOpts{
			Name: "eth_sub_interruption_counter",
			Help: "The total number of ethereum subscription interruptions causing Jester to redial the websocket client",
		}),
	}
}

// Methods to interact with metrics

func (m *PrometheusMetrics) IncEthSubInterruptionCounter() {
	if m.enabled {
		m.EthSubInteruptionCounter.Inc()
	}
}
