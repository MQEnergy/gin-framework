package backend

import (
	"mqenergy-go/config"
	"mqenergy-go/entities/user"
	"mqenergy-go/global"
	"mqenergy-go/models"
	"mqenergy-go/pkg/paginator"
)

type UserService struct{}

var User = UserService{}

// GetList
func (s UserService) GetList(requestParams user.IndexRequest) (interface{}, error) {
	var userList = make([]user.UserList, 0)
	fields := []string{
		models.GinUserTbName + ".user_name",
		models.GinUserInfoTbName + ".user_id",
		models.GinUserInfoTbName + ".role_ids",
	}
	pagination, err := paginator.Builder.
		WithDB(global.DB).
		WithModel(&models.GinUser{}).
		WithFields(fields).
		WithJoins("left", models.GinUserInfoTbName, paginator.OnJoins{
			LeftField:  models.GinUserTbName + ".id",
			RightField: models.GinUserInfoTbName + ".user_id",
		}).
		Pagination(userList, requestParams.Page, config.Conf.Server.DefaultPageSize)
	return pagination, err
}
