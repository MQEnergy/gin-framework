package common

import (
	"errors"
	"lyky-go/config"
	"lyky-go/entities"
	"lyky-go/global"
	"lyky-go/pkg/auth"
)

type UserService struct{}

var User = UserService{}

// Login 登录操作
func (s UserService) Login(request entities.UserLoginRequest) (string, error) {
	var user entities.User
	err := global.DB.Where(&entities.User{
		BaseUser: entities.BaseUser{
			Phone: request.Phone,
		},
	}).Take(&user).Error
	if err != nil {
		return "", errors.New("未查找到用户")
	}
	jwtToken, err := auth.GenerateJwtToken(config.Conf.Server.JwtSecret, config.Conf.Server.TokenExpire, user, config.Conf.Server.TokenIssuer)
	if err != nil {
		return "", errors.New("token生成失败")
	}
	return jwtToken, nil
}
