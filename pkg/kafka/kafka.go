package kafka

import (
	"context"

	"github.com/greenblat17/platform-common/pkg/kafka/consumer"
)

type Consumer interface {
	Consume(ctx context.Context, topicName string, handler consumer.Handler) (err error)
	Close() error
}
