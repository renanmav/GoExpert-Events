package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/renanmav/GoExpert-Events/pkg/rabbitmq"
)

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
