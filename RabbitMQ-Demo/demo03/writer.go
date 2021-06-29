package pmq

import (
	"jtproto/logger"

	"github.com/streadway/amqp"
)

// 定义生产者
type WriterImpl struct {
	channel *amqp.Channel
	bind    *BindConfig
	msgChan chan []byte
	msgType string
}

func newWriter(channel *amqp.Channel, bind *BindConfig, msgType string) *WriterImpl {
	return &WriterImpl{
		channel: channel,
		bind:    bind,
		msgChan: make(chan []byte, config.Writer.Buffer),
		msgType: msgType,
	}
}

func (writer *WriterImpl) Publish(data []byte) {
	writer.msgChan <- data
}

func (writer *WriterImpl) listenPublish() {
	logger.Info("MQ writer run",
		logger.FieldString("exchange", writer.bind.Exchange),
		logger.FieldString("route key", writer.bind.RouteKey))
	for data := range writer.msgChan {
		err := writer.channel.Publish(writer.bind.Exchange, writer.bind.RouteKey, false, false,
			amqp.Publishing{
				ContentType: writer.msgType,
				Body:        data,
			},
		)
		if err != nil {
			logger.Error("MQ writer publish",
				logger.FieldString("exchange", writer.bind.Exchange),
				logger.FieldString("route key", writer.bind.RouteKey),
				logger.FieldError(err))
			// 失败只记录不退出
		}
	}
}
