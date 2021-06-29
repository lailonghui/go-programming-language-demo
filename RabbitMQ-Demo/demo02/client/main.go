/*
@Time : 2021/1/19 9:58
@Author : lai
@Description :
@File : main
*/
package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@192.168.3.163:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"jt808.common.writer", // name
		true,                  // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "ewogICAgIkNBTlJ4VGltZSI6ICIxNTMwMDAwMDAwIiwKICAgICJDQU5BcnJheSI6IFsKICAgICAgICB7CiAgICAgICAgICAgICJJZCI6IDEsCiAgICAgICAgICAgICJEYXRhIjogImUzMD0iCiAgICAgICAgfSwgCiAgICAgICAgewogICAgICAgICAgICAiSWQiOiAyLAogICAgICAgICAgICAiRGF0YSI6ICJlMzA9IgogICAgICAgIH0KICAgIF0KfQ=="
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [lai] Sent %s", body)
}
