package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// MockWriter is used for testing without connecting to real Kafka
type MockWriter struct{}

func (m *MockWriter) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	// Simulate successful write
	return nil
}