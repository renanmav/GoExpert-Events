package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (channel *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	channel, err = conn.Channel()
	if err != nil {
		panic(err)
	}

	return channel
}

func Consume(channel *amqp.Channel, out chan<- amqp.Delivery) error {
	messages, err := channel.Consume(
		"my-queue",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range messages {
		out <- msg
	}
	return nil
}

func Publish(channel *amqp.Channel, body string) error {
	err := channel.Publish(
		"amq.direct",
		"routing-key",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	return err
}
