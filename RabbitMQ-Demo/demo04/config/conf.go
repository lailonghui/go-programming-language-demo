/*
@Time : 2021/4/14 19:50
@Author : lai
@Description :
@File : conf
*/
package config

const (
	//RMQADDR     = "amqp://admin:admin@192.168.3.130:5672/"
	//RMQADDR     = "amqp://admin:admin@192.168.3.163:5672/"
	RMQADDR     = "amqp://admin:admin@120.37.177.122:18002/"
	QUEUENAME   = "msgQueueWithPersist"
	PRODUCERCNT = 5
	CONSUMERCNT = 20
)
