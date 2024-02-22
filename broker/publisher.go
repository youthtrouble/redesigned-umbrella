package broker

import (
	"fmt"

	"github.com/streadway/amqp"
)

//dynamic queue config
type PublisherQueueConfig struct {
	amqpQueue    amqp.Queue
	QueueChannel *amqp.Channel
}

// Re-usable function to create a new queue
func NewQueue(rabbitMQChannel *amqp.Channel, name string) (*PublisherQueueConfig, error) {
	amqpQueue, err := rabbitMQChannel.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &PublisherQueueConfig{
		amqpQueue:    amqpQueue,
		QueueChannel: rabbitMQChannel,
	}, nil
}

// Re-usable function to publish a message to a queue
func (mq *PublisherQueueConfig) Publish(message string) error {

	err := mq.QueueChannel.Publish(
		"",
		mq.amqpQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		return fmt.Errorf("error publishing message to %s queue : %s", mq.amqpQueue.Name, err)
	}

	return nil
}
