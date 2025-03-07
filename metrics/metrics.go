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

	EthSubInterruptionCounter   prometheus.Counter
	GetVoteExtensionCounter     prometheus.Counter
	LogMessagedPublishedCounter prometheus.Counter
	MTokenSentCounter           prometheus.Counter
	MTokenIndexSentCounter      prometheus.Counter
	VAAReceiveDuration          prometheus.Summary
	VAAFoundTotal               prometheus.Counter
	VAAFailedTotal              prometheus.Counter
	VAAFailedMaxAttemptsReached prometheus.Counter
}

// NewPrometheusMetrics creates a new PrometheusMetrics object
func NewPrometheusMetrics() *PrometheusMetrics {
	registry := prometheus.NewRegistry()
	factory := promauto.With(registry)

	return &PrometheusMetrics{
		registry: registry,
		EthSubInterruptionCounter: factory.NewCounter(prometheus.CounterOpts{
			Name: "eth_sub_interruption_counter",
			Help: "The total number of Ethereum subscription interruptions causing Jester to redial the websocket client.",
		}),
		GetVoteExtensionCounter: factory.NewCounter(prometheus.CounterOpts{
			Name: "getVoteExtension_counter",
			Help: "The number of times `getVoteExtension` is queried. If you are a validator, this should happen each time you are the proposer.",
		}),
		LogMessagedPublishedCounter: factory.NewCounter(prometheus.CounterOpts{
			Name: "logMessagePublished_counter",
			Help: "The total number of times the Ethereum event `LogMessagePublished` is observed.",
		}),
		MTokenSentCounter: factory.NewCounter(prometheus.CounterOpts{
			Name: "mTokenSent_counter",
			Help: "The total number of times the Ethereum event `MTokenSent` is observed.",
		}),

		MTokenIndexSentCounter: factory.NewCounter(prometheus.CounterOpts{
			Name: "mTokenIndexSent_counter",
			Help: "The total number of times the Ethereum event `MTokenIndexSent` is observed.",
		}),
		VAAReceiveDuration: factory.NewSummary(prometheus.SummaryOpts{
			Name:       "vaa_receive_duration_minutes",
			Help:       "Summary of the time it takes for Wormhole to pick up the VAA in minutes. This metric is not recorded if the VAA is found on the first query attempt.",
			Objectives: map[float64]float64{0.5: 0.01, 0.9: 0.01, 0.99: 0.001},
		}),
		VAAFoundTotal: factory.NewCounter(prometheus.CounterOpts{
			Name: "vaa_found_total",
			Help: "The total number of times a VAA was found.",
		}),
		VAAFailedTotal: factory.NewCounter(prometheus.CounterOpts{
			Subsystem: "vaa_failed",
			Name:      "total",
			Help:      "The total number of times fetching a VAA failed.",
		}),
		VAAFailedMaxAttemptsReached: factory.NewCounter(prometheus.CounterOpts{
			Subsystem: "vaa_failed",
			Name:      "max_attempts_reached_total",
			Help:      "The total number of times fetching a VAA failed after reaching the maximum number of attempts.",
		}),
	}
}

// Methods to interact with metrics

// IncEthSubInterruptionCounter increments the metric tracking the total number of Ethereum subscription interruptions
// causing Jester to redial the websocket client.
func (m *PrometheusMetrics) IncEthSubInterruptionCounter() {
	if m.enabled {
		m.EthSubInterruptionCounter.Inc()
	}
}

// IncGetVoteExtensionCounter increments the metric tracking the total number of times `GetVoteExtension` is queried.
func (m *PrometheusMetrics) IncGetVoteExtensionCounter() {
	if m.enabled {
		m.GetVoteExtensionCounter.Inc()
	}
}

// IncLogMessagePublishedCounter increments the metric tracking the total number of times `logMessagePublished` is observed.
func (m *PrometheusMetrics) IncLogMessagePublishedCounter() {
	if m.enabled {
		m.LogMessagedPublishedCounter.Inc()
	}
}

// AddLogMessagePublishedCounter adds to the metric tracking the total number of times `logMessagePublished` is observed.
func (m *PrometheusMetrics) AddLogMessagePublishedCounter(n int) {
	if m.enabled {
		m.LogMessagedPublishedCounter.Add(float64(n))
	}
}

// IncMTokenSentCounter increments the metric tracking the total number of times `MTokenSent` is observed.
func (m *PrometheusMetrics) IncMTokenSentCounter() {
	if m.enabled {
		m.MTokenSentCounter.Inc()
	}
}

// AddMTokenSentCounter adds to the metric tracking the total number of times `MTokenSent` is observed.
func (m *PrometheusMetrics) AddMTokenSentCounter(n int) {
	if m.enabled {
		m.MTokenSentCounter.Add(float64(n))
	}
}

// IncMTokenIndexSentCounter increments the metric tracking the total number of times `MTokenIndexSent` is observed.
func (m *PrometheusMetrics) IncMTokenIndexSentCounter() {
	if m.enabled {
		m.MTokenIndexSentCounter.Inc()
	}
}

// AddMTokenIndexSentCounter adds to the metric tracking the total number of times `MTokenIndexSent` is observed.
func (m *PrometheusMetrics) AddMTokenIndexSentCounter(n int) {
	if m.enabled {
		m.MTokenIndexSentCounter.Add(float64(n))
	}
}

// ObserveVAAReceiveDuration is a summary that tracks the time it takes wormhole to pick up the VAA in minutes.
// `duration` is in minutes.
func (m *PrometheusMetrics) ObserveVAAReceiveDuration(duration float64) {
	if m.enabled {
		m.VAAReceiveDuration.Observe(duration)
	}
}

// IncVAAFoundTotalCounter increments the metric tracking the total number of times a VAA was found.
func (m *PrometheusMetrics) IncVAAFoundTotal() {
	if m.enabled {
		m.VAAFoundTotal.Inc()
	}
}

// IncVAAFailedTotal increments the metric tracking the total number of times Jester failed to fetch a VAA.
func (m *PrometheusMetrics) IncVAAFailedTotal() {
	if m.enabled {
		m.VAAFailedTotal.Inc()
	}
}

// IncVAAFailedMaxAttemptsReached increments the metric tracking the total number of times Jester failed to
// fetch a VAA after reaching the maximum number of attempts.
func (m *PrometheusMetrics) IncVAAFailedMaxAttemptsReached() {
	if m.enabled {
		m.VAAFailedMaxAttemptsReached.Inc()
	}
}
