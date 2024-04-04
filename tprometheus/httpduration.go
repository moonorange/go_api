package tprometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests in seconds.",
		},
		[]string{"handler", "method", "status"},
	)
)
