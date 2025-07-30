package main

import (
	"github.com/KingBean4903/flash-sale-ecommerce/services/order-service/handler"
	"os"
	"github.com/KingBean4903/flash-sale-ecommerce/services/order-service/redis"
	"github.com/KingBean4903/flash-sale-ecommerce/services/order-service/kafka"
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

	kafka.InitProducer(broker)

	redis.InitRedis(redisAddr)


	handler.StartHTTPServer()
}
