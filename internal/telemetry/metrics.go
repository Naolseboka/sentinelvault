package telemetry

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	regOnce sync.Once
	reg     *prometheus.Registry
)

// Registry returns the service prometheus registry.
func Registry() *prometheus.Registry {
	regOnce.Do(func() {
		reg = prometheus.NewRegistry()
		httpRequestsTotal = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total HTTP requests processed",
			},
			[]string{"method", "path", "status"},
		)
		reg.MustRegister(httpRequestsTotal)
	})
	return reg
}

var httpRequestsTotal *prometheus.CounterVec

// ObserveHTTPRequest increments the HTTP request counter.
func ObserveHTTPRequest(method, path string, status int) {
	httpRequestsTotal.WithLabelValues(method, path, strconv.Itoa(status)).Inc()
}

// MetricsHandler exposes prometheus metrics over HTTP.
func MetricsHandler() http.Handler {
	return promhttp.HandlerFor(Registry(), promhttp.HandlerOpts{})
}
