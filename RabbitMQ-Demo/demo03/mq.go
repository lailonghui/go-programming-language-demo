package pmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// 定义RabbitMQ对象
type RabbitMQ struct {
	conn    *amqp.Connection
	writers []*WriterImpl
	readers []*ReaderImpl
}

func NewRabbitMQ() IMQ {
	return &RabbitMQ{
		conn: nil,
	}
}
func (mq *RabbitMQ) Start() {
	for _, writer := range mq.writers {
		go writer.listenPublish()
	}
	for _, reader := range mq.readers {
		go reader.listenReader()
	}
}

// 连接RabbitMQ
func (mq *RabbitMQ) Connect() error {
	err := error(nil)
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Server.User, config.Server.PassWord, config.Server.Ip, config.Server.Port)
	mq.conn, err = amqp.Dial(url)
	if err != nil {
		return fmt.Errorf("connect %s failed, %s", url, err)
	}
	return nil
}

// 创建生产者，必须在Start后才可调用
func (mq *RabbitMQ) NewWriter(bind *BindConfig, msgType string) (IWriter, error) {
	channel, err := mq.newChannel(bind)
	if err != nil {
		return nil, err
	}

	writer := newWriter(channel, bind, msgType)
	mq.writers = append(mq.writers, writer)
	return writer, nil
}

// 创建消费者，必须在Start后才可调用
func (mq *RabbitMQ) NewReader(bind *BindConfig, handler IReader) error {
	channel, err := mq.newChannel(bind)
	if err != nil {
		return err
	}

	reader := newReader(channel, bind, handler)
	mq.readers = append(mq.readers, reader)
	return nil
}

// 创建channel
func (mq *RabbitMQ) newChannel(bind *BindConfig) (*amqp.Channel, error) {
	if mq.conn == nil || mq.conn.IsClosed() {
		return nil, fmt.Errorf("mq conn is nil or closed")
	}
	channel, err := mq.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("create channel failed, %s", err)
	}
	err = mq.applyBind(channel, bind)
	if err != nil {
		_ = channel.Close()
		return nil, err
	}
	return channel, nil
}

// 创建交换器、队列，并将交换器与队列绑定
func (mq *RabbitMQ) applyBind(channel *amqp.Channel, bind *BindConfig) error {
	// 创建交换器，已经存在不会重复创建
	err := channel.ExchangeDeclare(bind.Exchange, bind.Type, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("declare exchange %s failed, %s", bind.Exchange, err)
	}

	// 创建队列，已经存在不会重复创建
	_, err = channel.QueueDeclare(bind.Queue, true, false, false, false, bind.QueueArgs)
	if err != nil {
		return fmt.Errorf("declare queue %s failed, %s", bind.Queue, err)
	}

	// 队列绑定
	err = channel.QueueBind(bind.Queue, bind.RouteKey, bind.Exchange, false, nil)
	if err != nil {
		return fmt.Errorf("bind queue %s with exchange %s failed, %s", bind.Queue, bind.Exchange, err)
	}

	return nil
}
