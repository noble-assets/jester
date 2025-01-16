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
