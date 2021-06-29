/*
@Time : 2021/4/14 19:51
@Author : lai
@Description :
@File : producer
*/
package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"lai.com/go_programming_language_demo/RabbitMQ-Demo/demo04/config"
	"log"
	"sync"
)

func main() {

	conn, err := amqp.Dial(config.RMQADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(config.PRODUCERCNT)

	for routine := 0; routine < config.PRODUCERCNT; routine++ {
		go func(routineNum int) {
			ch, err := conn.Channel()
			failOnError(err, "Failed to open a channel")
			defer ch.Close()

			//q, err := ch.QueueDeclare(
			//	config.QUEUENAME, //队列名称
			//	true,             //是否持久化
			//	false,            //是否自动删除
			//	false,            //是否排他
			//	false,            //是否非阻塞，阻塞：表示创建交换器的请求发送后，阻塞等待RMQ Server返回信息。非阻塞：不会阻塞等待RMQ
			//	nil,
			//)
			//
			//failOnError(err, "Failed to declare a queue")

			for i := 0; i < 500; i++ {
				msgBody := fmt.Sprintf("Message_%d_%d", routineNum, i)

				err = ch.Publish(
					"lai.test.1", //exchange，当交换器名称为空时，表示使用默认交换器
					"aa",         //routing key
					//q.Name, //routing key
					false,
					false,
					amqp.Publishing{
						DeliveryMode: amqp.Persistent, //Msg set as persistent
						ContentType:  "text/plain",
						Body:         []byte(msgBody),
					})

				log.Printf(" [x] Sent %s", msgBody)
				failOnError(err, "Failed to publish a message")
			}

			wg.Done()
		}(routine)
	}

	wg.Wait()

	log.Println("All messages sent!!!!")
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}
