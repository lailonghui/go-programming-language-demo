/*
@Time : 2021/4/15 9:56
@Author : lai
@Description :
@File : consumer
*/
package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@192.168.3.163:5672")
	if err != nil {
		fmt.Println("Fail to connect to RabbitMQ")
		return
	}
	defer conn.Close()

	forever := make(chan bool)
	ch, err := conn.Channel()
	failOnError(err, "Fail to open channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"jt808.common.writer",
		"lai-test-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for msg := range msgs {
			fmt.Println(string(msg.Body))
			fmt.Println(msg.Headers)
			msg.Ack(false) //true表示回复当前信道所有未回复的ack，用于批量确认。false表示回复当前条目
		}
	}()

	<-forever
}
