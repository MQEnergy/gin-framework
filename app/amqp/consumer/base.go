package consumer

import (
	gorabbitmq "github.com/MQEnergy/go-rabbitmq"
	"time"
)

type (
	BaseConfig struct {
		Amqp     *gorabbitmq.RabbitMQ
		Data     map[string]interface{}
		CallBack Fn
	}
	Fn func(mq *gorabbitmq.RabbitMQ, Data map[string]interface{}) error
)

// New 实例化
func New(mq *gorabbitmq.RabbitMQ, data map[string]interface{}, f Fn) *BaseConfig {
	return &BaseConfig{
		Amqp:     mq,
		Data:     data,
		CallBack: f,
	}
}

// Consumer 消费者
func (c *BaseConfig) Consumer() error {
	// 防止消费过快 加一秒时间
	time.Sleep(time.Second * 1)
	if err := c.CallBack(c.Amqp, c.Data); err != nil {
		return err
	}
	return nil
}
