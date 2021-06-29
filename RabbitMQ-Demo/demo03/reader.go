package pmq

import (
	"jtproto/logger"

	"github.com/streadway/amqp"
)

// 定义消费者
type ReaderImpl struct {
	channel *amqp.Channel
	bind    *BindConfig
	handler IReader
}

func newReader(channel *amqp.Channel, bind *BindConfig, handler IReader) *ReaderImpl {
	return &ReaderImpl{
		channel: channel,
		bind:    bind,
		handler: handler,
	}
}

func (reader *ReaderImpl) listenReader() {
	// 获取消费通道
	err := reader.channel.Qos(config.Reader.Qos, 0, false)
	// 消费者名称为队列名
	deliveryChan, err := reader.channel.Consume(reader.bind.Queue, reader.bind.Queue, false, false, false, false, nil)
	if err != nil {
		logger.Error("MQ reader consume",
			logger.FieldString("queue", reader.bind.Queue),
			logger.FieldError(err))
		return
	}
	logger.Info("MQ reader run",
		logger.FieldString("queue", reader.bind.Queue))

	// 处理数据
	for msg := range deliveryChan {
		err := reader.handler.Deliver(msg.Body)
		if err != nil {
			logger.Error("MQ reader process",
				logger.FieldString("queue", reader.bind.Queue),
				logger.FieldError(err))
			// 失败只记录不退出
		}
		err = msg.Ack(false)
		if err != nil {
			logger.Error("MQ reader ack",
				logger.FieldString("queue", reader.bind.Queue),
				logger.FieldError(err))
			return
		}
	}
}
