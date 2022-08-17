package main

import (
	gorabbitmq "github.com/MQEnergy/go-rabbitmq"
	"mqenergy-go/app/amqp/consumer"
	"mqenergy-go/app/service/backend"
	"mqenergy-go/bootstrap"
	"mqenergy-go/config"
	"mqenergy-go/global"
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
