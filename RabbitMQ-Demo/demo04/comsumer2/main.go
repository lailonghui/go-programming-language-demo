/*
@Time : 2021/6/5 9:30
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"lai.com/go_programming_language_demo/RabbitMQ-Demo/demo04/config"
)

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}
func main() {
	//exchange := "lai.test"
	//exchange := "jt808.proxy.location"
	exchange := "lai.test.1"
	conn, err := amqp.Dial(config.RMQADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	forever := make(chan bool)
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange,
		"fanout", //队列名称
		true,     //是否持久化
		false,    //是否自动删除
		false,    //是否排他,true表示是。客户端无法直接发送msg到内部交换器，只有交换器可以发送msg到内部交换器。
		false,    //是否非阻塞，阻塞：表示创建交换器的请求发送后，阻塞等待RMQ Server返回信息。非阻塞：不会阻塞等待RMQ
		nil,
	)
	failOnError(err, "Failed to declare exchange")

	q, err := ch.QueueDeclare(
		"test2", //name
		false,   //durable
		false,   //delete when usused
		true,    // exclusive
		false,   //no-wait
		nil,     // arguments
	)

	err = ch.QueueBind(
		q.Name,   //队列的名字
		"aa",     //routing key
		exchange, //所绑定的交换器
		false,
		nil,
	)
	msgs, err := ch.Consume(
		q.Name, // 引用前面的队列名
		"",     // 消费者名字，不填自动生成一个
		true,   // 自动向队列确认消息已经处理
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	// 循环处理消息
	for d := range msgs {
		//log.Printf("接收消息=%s", d.Body)
		fmt.Println(string(d.Body))
	}
	<-forever
}
