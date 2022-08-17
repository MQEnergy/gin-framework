package backend

import (
	"fmt"
	"github.com/MQEnergy/gin-framework/global"
	"github.com/MQEnergy/gin-framework/models"
	"github.com/MQEnergy/gin-framework/pkg/paginator"
	"github.com/MQEnergy/gin-framework/types/user"
	gorabbitmq "github.com/MQEnergy/go-rabbitmq"
	"github.com/gogf/gf/v2/util/gconv"
	"sync"
)

type UserService struct{}

var User = &UserService{}

// GetIndex 获取列表
func (s *UserService) GetIndex(requestParams user.IndexRequest) (interface{}, error) {
	var userList = make([]user.UserList, 0)
	multiFields := []paginator.SelectTableField{
		{Model: models.GinUser{}, Table: models.GinUserTbName, Field: []string{"password", "salt", "_omit"}},
		{Model: models.GinUserInfo{}, Table: models.GinUserInfoTbName, Field: []string{"id", "user_id", "role_ids"}},
	}
	pagination, err := paginator.NewBuilder().
		WithDB(global.DB).
		WithModel(models.GinUser{}).
		//WithFields(models.GinUser{}, models.GinUserTbName, []string{"password", "salt", "_omit"}).
		//WithFields(models.GinUserInfo{}, models.GinUserInfoTbName, []string{"id", "user_id", "role_ids"}).
		WithMultiFields(multiFields).
		WithJoins("left", []paginator.OnJoins{{
			LeftTableField:  paginator.JoinTableField{Table: models.GinUserTbName, Field: "id"},
			RightTableField: paginator.JoinTableField{Table: models.GinUserInfoTbName, Field: "user_id"},
		}}).
		Pagination(&userList, requestParams.Page, global.Cfg.Server.DefaultPageSize)
	return pagination, err
}

// GetList 获取列表
func (s *UserService) GetList(requestParams user.IndexRequest) (interface{}, error) {
	var userList = make([]user.GinUser, 0)
	pagination, err := paginator.NewBuilder().
		WithDB(global.DB).
		WithModel(models.GinUser{}).
		WithPreload("UserInfo").
		Pagination(&userList, requestParams.Page, global.Cfg.Server.DefaultPageSize)
	return pagination, err
}

// AmqpConsumerHandler 处理消费者方法
func (s *UserService) AmqpConsumerHandler(mq *gorabbitmq.RabbitMQ, data map[string]interface{}) error {
	var wg sync.WaitGroup
	cherrors := make(chan error)
	consumerNum := gconv.Int(data["consumerNum"])
	wg.Add(consumerNum)
	for i := 0; i < consumerNum; i++ {
		fmt.Printf("正在开启消费者：第 %d 个\n", i+1)
		go func() {
			defer wg.Done()
			deliveries, err := mq.Consume()
			if err != nil {
				cherrors <- err
			}
			for d := range deliveries {
				// 消费者逻辑 to do
				fmt.Printf("got %dbyte delivery: [%v] %s %q\n", len(d.Body), d.DeliveryTag, d.Exchange, d.Body)
				d.Ack(false)
			}
		}()
	}
	select {
	case err := <-cherrors:
		close(cherrors)
		fmt.Printf("Consumer failed: %s\n", err)
		return err
	}
	wg.Wait()
	return nil
}

// AmqpProducerHandler 处理生产者方法
func (s *UserService) AmqpProducerHandler(mq *gorabbitmq.RabbitMQ, data []byte) error {
	if err := mq.Push(data); err != nil {
		fmt.Println("Push failed: " + err.Error())
		return err
	}
	fmt.Println("Push succeeded!", string(data))
	return nil
}
