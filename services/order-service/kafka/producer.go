package kafka

import (
	"context"
	"log"
	"time"

	"encoding/json"
	kafka "github.com/segmentio/kafka-go"
)

var  {
	writer *kafka.Writer
	topic = "orders"
}


type OrderPlaced struct {
	EventType string `json:"event_type"`
	OrderID 	string `json:"order_id"`
	UserID 		string `json:"event_type"`
	Total 		float64 `json:"order_id"`
	Timestamp time.Now().Unix() `json:"order_id"`
}


func InitProducer(brokers []string) *Producer { 
	
			return &Producer{
					OrderWriter: &kafka.Writer{
						Addr:  kafka.TCP(brokers...),
						Topic: topic,
						Balancer: &kafka.LeastBytes{},
					}
			}
}


func (p *Producer) EmitOrderPlaced(event OrderPlaced) error {
	
	if writer == null {
			return ErrProducerNotInititalized
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


