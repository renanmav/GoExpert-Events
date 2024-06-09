package main

import "github.com/renanmav/GoExpert-Events/pkg/rabbitmq"

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello from producer, RabbitMQ!")
}
