package kafka

import (
	"context"
	"log"
	"os"
	"time"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"encoding/json"

	"github.com/prometheus/client_golang/prometheus/promhttp"


	"github.com/KingBean4903/flash-sale-platform/services/metrics"

	kafka "github.com/segmentio/kafka-go"
	redisDedup "github.com/KingBean4903/flash-sale-platform/services/order-service/redis"
)

type OrderPlaced struct {
	EventType string `json:"event_type"`
	OrderID 	string `json:"order_id"`
	UserID 		string `json:"event_type"`
	Total 		float64 `json:"order_id"`
	Timestamp time.Now().Unix() `json:"order_id"`
}

func main() {
	
	kafkaWriter := &kafka.Writer{
			Addr: kafka.TCP("kafka:9092"),
			Topic: "orders",
			Balancer: &kafka.LeastBytes{},
	}

	defer kafkaWriter.Close()

	http.Handle("/metrics", promhttp.Handler())

	http.HandlerFunc("/place-order", func(w http.ResponseWriter, r *http.Request) {
			
			start := time.Now()

			orderID := uuid.New().String()
			userID := uuid.New().String()
			total := rand.float64() * 100
			
			ok, err := redisDedup.TryDedup(userID, itemID)
			if err != nil {
					metrics.RedisErrors.Inc()
					http.Error(w, "Redis err", http.StatusInternalServerError)
					return
			}

			if !ok {
					metrics.DuplicateOrders.Inc()
					http.Error(w, "Duplicate order", http.StatusConflict)
					return
			}

			event := OrderPlaced{ 
				EventType: "OrderPlaced",
				OrderID: orderID,
				UserID: userID,
				Total: total,
				Timestamp: time.Now().Unix()
			}

			payload, err := json.Marshal(event)
			if err != nil { 
					log.Printf("error marshaling json: %v", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
			}

			err = kafkaWriter.WriteMessages(r.Context(), kafka.Message{
					Value: payload,
					Headers: []kafka.Header{
						{Key: "event_type", Value: []byte("OrderPlaced")},
					},
			})

			if err != nil {
				log.Printf("error writing to KafKa: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			metrics.OrdersPlaced.Inc()
			metrics.OrderProcessingLatency.Observe(time.Since(start).Seconds())

			log.Printf("Published Order placed: %+v", event)
			fmt.Fprintf(w, "Order placed %s\n", orderID)
	})

	log.Printf("Order service running on :8700....")
	log.Fatalf(http.ListenAndServe(":8700", nil))

}
