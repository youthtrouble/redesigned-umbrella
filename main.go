package main

import (
	"log"

	"github.com/youthtrouble/redesigned-umbrella/broker"
)

func main() {

	rabbitMQConfig, err := broker.NewRabbitMQConfig()
	if err != nil {
		log.Fatalf("Error creating RabbitMQ config: %s", err)
	}

	_, err = broker.NewQueue(rabbitMQConfig.QueueChannel, "testQueue") // not ready to be used
	if err != nil {
		log.Fatalf("Error creating main queue: %s", err)
	}
}
