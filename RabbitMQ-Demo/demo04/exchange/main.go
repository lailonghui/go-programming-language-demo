/*
@Time : 2021/6/5 9:21
@Author : lai
@Description :
@File : main
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"lai.com/go_programming_language_demo/RabbitMQ-Demo/demo04/config"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}

type TestMsg struct {
	Code int
	Msg  string
}

// TODO: 广播模式
func main() {
	conn, err := amqp.Dial(config.RMQADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//err = ch.ExchangeDeclare(
	//	"lai.test.1",
	//	"fanout", //队列名称
	//	true,     //是否持久化
	//	false,    //是否自动删除
	//	false,    //是否排他,true表示是。客户端无法直接发送msg到内部交换器，只有交换器可以发送msg到内部交换器。
	//	false,    //是否非阻塞，阻塞：表示创建交换器的请求发送后，阻塞等待RMQ Server返回信息。非阻塞：不会阻塞等待RMQ
	//	nil,
	//)
	//failOnError(err, "Failed to declare exchange")
	msgBody := &TestMsg{
		Code: 200,
		Msg:  "今天天气真好",
	}
	bytes, _ := json.Marshal(msgBody)
	for i := 0; i < 10; i++ {
		//msgBody := fmt.Sprintf("今天捡到%d分钱", i)

		err = ch.Publish(
			"workflow",    //exchange，当交换器名称为空时，表示使用默认交换器
			"outage.test", //routing key
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent, //Msg set as persistent
				ContentType:  "text/plain",
				Body:         bytes,
			})

		log.Printf(" [x] Sent %s", msgBody)
		failOnError(err, "Failed to publish a message")
	}

	log.Println("All messages sent!!!!")
}
