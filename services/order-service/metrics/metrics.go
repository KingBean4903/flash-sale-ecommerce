package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
		
	OrderPlaced = promauto.NewCounter(prometheus.CounterOpts{
			Name: "orders_placed_total",
			Help: "Total no of orders placed",
	})

	DuplicateOrders = promauto.NewCounter(prometheus.CounterOpts{
				Name: "orders_duplicate_total",
				Help: "Total no of duplicate orders",
	})

	RedisErrors = promauto.NewCounter(prometheus.CounterOpts{
				Name: "redis_errors_total",
				Help: "Total no of Redis errors",
	})

	OrderProcessingLatency = promauto.NewHistogram(prometheus.HistogramOpts{
				Name: "order_processing_seconds",
				Help: "Time taken to process each order",
				Buckets: prometheus.DefBuckets,
	})

)
