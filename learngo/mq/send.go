package main

import (
	"github.com/streadway/amqp"
	"helloworld/learngo/mq/mq"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@8.130.142.45:5672/")
	mq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	mq.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	mq.FailOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	mq.FailOnError(err, "Failed to publish a message")
}
