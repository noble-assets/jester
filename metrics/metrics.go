package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type PrometheusMetrics struct {
	Enabled bool

	Registry *prometheus.Registry

	EthSubInteruptionCounter *prometheus.CounterVec
}

// NewPrometheusMetrics creates a new PrometheusMetrics object
func NewPrometheusMetrics() *PrometheusMetrics {
	return &PrometheusMetrics{}
}

// Initialize initializes and configures Prometheus
func (m *PrometheusMetrics) Initialize() {
	m.Enabled = true

	m.Registry = prometheus.NewRegistry()
	factory := promauto.With(m.Registry)

	m.EthSubInteruptionCounter = factory.NewCounterVec(prometheus.CounterOpts{
		Name: "eth_sub_interruption_counter",
		Help: "The total number of ethereum subscription interruptions causing Jester to redial the websocket client",
	}, []string{})

}

// Methods to interact with metrics

func (m *PrometheusMetrics) IncEthSubInterruptionCounter() {
	m.EthSubInteruptionCounter.WithLabelValues().Inc()
}
