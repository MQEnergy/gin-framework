package common

import (
	"errors"
	"mqenergy-go/config"
	"mqenergy-go/entities/user"
	"mqenergy-go/global"
	"mqenergy-go/pkg/auth"
)

type UserService struct{}

var User = UserService{}

// Login 登录操作
func (s UserService) Login(requestParams user.LoginRequest) (interface{}, error) {
	var userInfo user.User
	if err := global.DB.Where("phone = ?", requestParams.Phone).First(&userInfo).Error; err != nil {
		return userInfo, errors.New("未查找到用户")
	}
	jwtToken, err := auth.GenerateJwtToken(config.Conf.Server.JwtSecret, config.Conf.Server.TokenExpire, userInfo, config.Conf.Server.TokenIssuer)
	if err != nil {
		return "", errors.New("token生成失败")
	}
	return jwtToken, nil
}
