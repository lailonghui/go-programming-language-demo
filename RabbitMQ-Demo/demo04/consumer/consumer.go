/*
@Time : 2021/4/14 19:51
@Author : lai
@Description :
@File : consumer
*/
package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"lai.com/go_programming_language_demo/RabbitMQ-Demo/demo04/config"
	"log"
)

func main() {

	conn, err := amqp.Dial(config.RMQADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	forever := make(chan bool)

	for routine := 0; routine < config.CONSUMERCNT; routine++ {
		go func(routineNum int) {
			ch, err := conn.Channel()
			failOnError(err, "Failed to open a channel")
			defer ch.Close()

			q, err := ch.QueueDeclare(
				config.QUEUENAME,
				true, //durable
				false,
				false,
				false,
				nil,
			)

			failOnError(err, "Failed to declare a queue")

			msgs, err := ch.Consume(
				q.Name,            //队列名称
				"MsgWorkConsumer", //消费者标签，用于区分不同的消费者。
				false,             //是否自动回复ACK，回复ACK表示高速服务器我收到消息了。建议为false，手动回复，这样可控性强。
				false,             //设置是否排他，排他表示当前队列只能给一个消费者使用。
				false,             //如果为true，表示生产者和消费者不能是同一个connect。
				false,             //是否非阻塞，true表示是。阻塞：表示创建交换器的请求发送后，阻塞等待RMQ Server返回信息。非阻塞：不会阻塞等待RMQ Server的返回信息，而RMQ Server也不会返回信息。（不推荐使用）
				nil,
			)

			if err != nil {
				log.Fatal(err)
			}

			for msg := range msgs {
				log.Printf("In %d consume a message: %s\n", 0, msg.Body)
				log.Printf("Done")
				msg.Ack(false) //Ack
			}

		}(routine)
	}

	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}
