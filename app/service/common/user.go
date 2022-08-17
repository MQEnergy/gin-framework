package common

import (
	"errors"
	"mqenergy-go/global"
	"mqenergy-go/pkg/auth"
	"mqenergy-go/types/user"
)

type UserService struct{}

var User = UserService{}

// Login 登录操作
func (s UserService) Login(requestParams user.LoginRequest) (interface{}, error) {
	var userInfo user.User
	if err := global.DB.Where("phone = ?", requestParams.Phone).First(&userInfo).Error; err != nil {
		return userInfo, errors.New("未查找到用户")
	}
	jwtToken, err := auth.GenerateJwtToken(global.Cfg.Jwt.Secret, global.Cfg.Jwt.TokenExpire, userInfo, global.Cfg.Jwt.TokenIssuer)
	if err != nil {
		return "", errors.New("token生成失败")
	}
	return jwtToken, nil
}
