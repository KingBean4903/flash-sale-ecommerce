package main

import (
	"log"
	"github.com/KingBean4903/flash-sale-platform/services/order-service/handler"
	"github.com/KingBean4903/flash-sale-platform/services/order-service/redis"
	"github.com/KingBean4903/flash-sale-platform/services/order-service/kafka"
)

func main() {

	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
			broker = "kafka-zoo:9092"
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" { 
			redisAddr = "localhost:6379"
	}

	kafka.InitProducer([]string{broker})

	redis.InitRedis(redisAddr)


	handler.StartHTTPServer()
}
