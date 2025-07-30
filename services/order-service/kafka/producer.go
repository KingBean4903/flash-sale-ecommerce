package kafka

import (
	"context"
	"log"
	"time"

	"encoding/json"
	kafka "github.com/segmentio/kafka-go"
)

var ( 
	writer *kafka.Writer
	topic = "orders"
)


type OrderPlaced struct {
	EventType string  `json:"event_type"`
	OrderID 	string  `json:"order_id"`
	ItemID    string  `json:"item_id"`
	UserID 		string  `json:"event_type"`
	Total 		float64 `json:"order_id"`
	Timestamp int64   `json:"order_id"`
}


func InitProducer(brokerAddress string) { 
	
			writer = kafka.NewWriter(kafka.WriterConfig{
						Brokers:  []string{brokerAddress},
						Topic: topic,
						Balancer: &kafka.LeastBytes{},
					})
				log.Println("Producer initialized")
			
}


func EmitOrderPlaced(event OrderPlaced) error {
	
	if writer == nil {
			return ErrProducerNotInitialized
	}

	data, err := json.Marshal(event)
	if err != nil {
			return err
	}

	msg := kafka.Message{
			Key: []byte(event.OrderID),
			Value: data,
			Time: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err = writer.WriteMessages(ctx, msg)
	if err != nil {
			log.Printf("Failed to write message: %v\n", err)
			return err
	}

	log.Printf("Sent order.placed event for orderID: %s\n",event.OrderID)
	return nil
}

var ErrProducerNotInitialized = &ProducerError{"KafKa producer not initialized"}

type ProducerError struct {
	msg string
}

func (e *ProducerError) Error() string {
		return e.msg
}


