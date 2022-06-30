package main

import (
	"fmt"
	"mqenergy-go/app/amqp/consumer"
	"mqenergy-go/app/service/backend"
	"mqenergy-go/config"
	"mqenergy-go/pkg/lib"
)

func main() {
	config.InitConfig()
	// 实例化amqp
	amqp := lib.NewRabbitMQ("test", "", "", "", 0)
	// 启动3个消费者
	data := map[string]interface{}{
		"consumerNum": 3,
	}
	if err := consumer.New(amqp, data, backend.User.AmqpConsumerHandler).Consumer(); err != nil {
		fmt.Println("consumer failed", err.Error())
	}
}
