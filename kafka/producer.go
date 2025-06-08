package kafka

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type KafkaWriter interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
}

var Writer KafkaWriter

func InitKafkaProducer() {
	brokers := os.Getenv("KAFKA_BROKER")
	if brokers == "" {
		brokers = "localhost:9092"
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "message-topic"
	}

	log.Println("🔥 InitKafkaProducer called")
	log.Println("Using broker:", brokers)
	log.Println("Using topic :", topic)

	Writer = &kafka.Writer{
		Addr:     kafka.TCP(brokers),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func PublishMessage(msg []byte) error {
	return Writer.WriteMessages(context.Background(), kafka.Message{
		Value: msg,
	})
}
