package pmq

import (
	"jtproto/tool"

	"github.com/streadway/amqp"
)

// 服务器配置
type Config struct {
	Server ServerConfig `yaml:"Server"`
	Writer WriterConfig `yaml:"Writer"`
	Reader ReaderConfig `yaml:"Reader"`
}

// 服务器配置
type ServerConfig struct {
	User     string `yaml:"User"`
	PassWord string `yaml:"PassWord"`
	Ip       string `yaml:"Ip"`
	Port     uint16 `yaml:"Port"`
}

// 生产者配置
type WriterConfig struct {
	Buffer int `yaml:"Buffer"`
}

// 消费者配置
type ReaderConfig struct {
	Qos int `yaml:"Qos"`
}

// 交换器与队列的绑定关系
type BindConfig struct {
	Exchange  string     `yaml:"Exchange"`
	RouteKey  string     `yaml:"RouteKey"`
	Queue     string     `yaml:"Queue"`
	QueueArgs amqp.Table `yaml:"QueueArgs"`
	Type      string     `yaml:"Type"`
}

var config = &Config{}

func init() {
	err := tool.UnmarshalConfig("mq.yaml", config)
	if err != nil {
		panic(err)
	}
}
