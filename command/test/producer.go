package main

import (
	"fmt"
	"mqenergy-go/app/amqp/producer"
	"mqenergy-go/app/service/backend"
	"mqenergy-go/config"
	"mqenergy-go/pkg/lib"
	"mqenergy-go/pkg/util"
	"time"
)

func main() {
	config.InitConfig()
	// 实例化amqp
	amqp := lib.NewRabbitMQ("test", "", "", "", 0)
	fmt.Println("启动生产者...")
	// 定时器 1s 执行一次生产者
	util.NewTicker(1, func() error {
		data := []byte("{\"hello\":\"world " + time.Now().Format("2006-01-02 15:04:05") + "\"}")
		if err := producer.New(amqp, data, backend.User.AmqpProducerHandler).Publish(); err != nil {
			return err
		}
		return nil
	}).Start()
}
