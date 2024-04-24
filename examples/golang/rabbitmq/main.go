package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func connectRabbitMQ() *amqp.Connection {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully connected to rabbitmq")
	return conn
}

func consumer() {
	conn := connectRabbitMQ()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err :=ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("Received message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connected to rabbitmq instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}

func main() {

	conn := connectRabbitMQ()

	fmt.Println("Successfully connected to rabbitmq")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("Hello World"),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully published message to queue")

	consumer()
}