package main

import (
	"bytes"
	"github.com/streadway/amqp"
	"helloworld/learngo/mq/mq"
	"log"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@8.130.142.45:5672/")
	mq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	mq.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil,
	)

	err = ch.Qos(
		1,
		0,
		false,
	)
	mq.FailOnError(err, "ch.Qos() failed")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	mq.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dot := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
