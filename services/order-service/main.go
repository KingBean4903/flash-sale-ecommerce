package main

import (
	"log"
	"github.com/KingBean4903/flash-sale-platform/services/handler"
	"github.com/KingBean4903/flash-sale-platform/services/order-service/redis"
)

func main() {

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" { 
			redisAddr = "localhost:6379"
	}

	redie.InitRedis(redisAddr)


	log.Println("Starting Order Service...")
	handler.StartHTTPServer()


}
