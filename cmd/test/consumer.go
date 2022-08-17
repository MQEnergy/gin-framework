package main

import (
	"github.com/MQEnergy/go-framework/app/amqp/consumer"
	"github.com/MQEnergy/go-framework/app/service/backend"
	"github.com/MQEnergy/go-framework/bootstrap"
	"github.com/MQEnergy/go-framework/config"
	"github.com/MQEnergy/go-framework/global"
	gorabbitmq "github.com/MQEnergy/go-rabbitmq"
	"time"
)

func main() {
	config.ConfEnv = "dev"
	bootstrap.BootService(bootstrap.LoggerService)
	amqpConfig := &gorabbitmq.Config{
		User:     global.Cfg.Amqp.User,
		Password: global.Cfg.Amqp.Password,
		Host:     global.Cfg.Amqp.Host,
		Port:     global.Cfg.Amqp.Port,
		Vhost:    global.Cfg.Amqp.Vhost,
	}
	// 实例化amqp
	amqp := gorabbitmq.New(amqpConfig, "test", "", "", 0, 1, true)
	time.Sleep(time.Second * 1)
	// 启动3个消费者
	data := map[string]interface{}{
		"consumerNum": 1,
	}
	if err := consumer.New(amqp, data, backend.User.AmqpConsumerHandler).Consumer(); err != nil {
	}
}
