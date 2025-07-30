package handler

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/google/uuid"

	"time"
	"github.com/prometheus/client_golang/prometheus/promhttp"


	"github.com/KingBean4903/flash-sale-ecommerce/services/order-service/metrics"
	"github.com/KingBean4903/flash-sale-ecommerce/services/order-service/kafka"

	redisDedup "github.com/KingBean4903/flash-sale-ecommerce/services/order-service/redis"
)


func StartHTTPServer() {

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/place-order", func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		userID := r.URL.Query().Get("user_id")
		itemID := r.URL.Query().Get("item_id")

		ok, err := redisDedup.TryDedup(userID, itemID)
		if err != nil {
			metrics.RedisErrors.Inc()
			http.Error(w,  "Redis error", http.StatusInternalServerError)
			return
		}

		if !ok {
				metrics.DuplicateOrders.Inc()
				http.Error(w, "Duplicate order", http.StatusConflict)
				return
		}

		orderID := uuid.New().String()
		event := kafka.OrderPlaced{
				OrderID: orderID,
				UserID: userID,
				ItemID: itemID,
				Timestamp: time.Now().Unix(),
		}

		err = kafka.EmitOrderPlaced(event)
		if err != nil {
				http.Error(w, "Failed to place Order", http.StatusInternalServerError)
				return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
				"message": "Order placed",
				"orderID": orderID,
		})

		metrics.OrderPlaced.Inc()
		metrics.OrderProcessingLatency.Observe(time.Since(start).Seconds())

	})

	log.Printf("Order service running on :8700....")
	log.Fatal(http.ListenAndServe(":8700", nil))
}
