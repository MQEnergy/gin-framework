package app

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"mqenergy-go/entities/admin"
	"mqenergy-go/global"
	"mqenergy-go/pkg/auth"
	"strconv"
)

type AdminInfo struct {
	admin.Admin
}

type TokenPayload struct {
	UserId int64 `json:"id"`
}

// ParseUserByToken 通过token解析用户
func ParseUserByToken(token string) (TokenPayload, error) {
	user := TokenPayload{}
	if token == "" {
		return user, errors.New("token 为空")
	}
	jwtPayload, err := auth.ParseJwtToken(token, global.Cfg.Jwt.Secret)
	if err != nil {
		return user, err
	}
	byteSlice, err := json.Marshal(jwtPayload.User)
	if err != nil {
		return user, err
	}
	if err = json.Unmarshal(byteSlice, &user); err != nil {
		return user, err
	}
	if user.UserId == 0 {
		return user, errors.New("非法登录")
	}
	_, err = global.Redis.Get(context.Background(), global.Cfg.Redis.LoginPrefix+strconv.FormatInt(user.UserId, 10)).Result()
	if err != nil {
		return TokenPayload{}, errors.New("会话过期，请重新登录")
	}
	return user, nil
}

// GetAdminInfo 获取登陆者的用户信息
func GetAdminInfo(ctx *gin.Context) (AdminInfo, error) {
	info, err := ParseUserByToken(ctx.GetHeader(global.Cfg.Jwt.TokenKey))
	if err != nil {
		return AdminInfo{}, nil
	}
	Uid := info.UserId
	//	从redis查询
	result, err := global.Redis.Get(context.Background(), global.Cfg.Redis.LoginPrefix+strconv.FormatInt(Uid, 10)).Result()
	if err != nil {
		return AdminInfo{}, nil
	}
	var user admin.Admin
	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		return AdminInfo{}, nil
	}
	return AdminInfo{user}, nil
}
