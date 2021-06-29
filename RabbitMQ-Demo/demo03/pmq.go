package pmq

const (
	TextPlain = "text/plain"
)

// 定义服务器接口
type IMQ interface {
	Start()                                                         // 启动读写协程
	Connect() error                                                 // 连接MQ
	NewWriter(msgBind *BindConfig, msgType string) (IWriter, error) // 创建生产者，必须在Connect后才可调用
	NewReader(msgBind *BindConfig, handler IReader) error           // 创建消费者，必须在Connect后才可调用
}

// 定义生产者接口
type IWriter interface {
	Publish(data []byte)
}

// 定义接收者接口
type IReader interface {
	Deliver(data []byte) error
}
