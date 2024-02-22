package broker

import (
	"os"

	"github.com/streadway/amqp"
)

func NewRabbitMQConfig() (*PublisherQueueConfig, error) {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	rabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		return nil, err
	}

	defer rabbitMQ.Close()
	rabbitMQChannel, err := rabbitMQ.Channel()
	if err != nil {
		return nil, err
	}
	defer rabbitMQChannel.Close()

	return &PublisherQueueConfig{
		QueueChannel:  rabbitMQChannel,
	}, nil
}
