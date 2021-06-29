/*
@Time : 2021/4/15 10:54
@Author : lai
@Description :
@File : producer
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

	c, err := conn.Channel()
	ch, err := conn.Channel()
	failOnError(err, "Fail to open channel")
	defer ch.Close()
	//msgBody := `
	//{
	//   "Head":{
	//       "WorkId":33536,
	//       "ConnId":"3314589178037632",
	//       "Version":0,
	//       "PhoneNum":"10186507674",
	//       "ReqNum":0
	//   },
	//   "Body":{
	//       "TextSignal":8,
	//       "Text":"\uD5E2\uCAC7"
	//   }
	//}
	//`
	msgBody := `
	{
	   "Head":{
	       "WorkId":33028,
	       "ConnId":"3315352997404032",
	       "Version":0,
	       "PhoneNum":"10186507674",
	       "ReqNum":0
	   },
	   "Body":{}
	
	}
	`
	//q, err := ch.QueueDeclare(
	//	"test111", // name
	//	false,     // durable
	//	false,     // delete when unused
	//	false,     // exclusive
	//	false,     // no-wait
	//	nil,       // arguments
	//)
	err = c.Publish(
		"lai.test",
		"reader",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(msgBody),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

}
